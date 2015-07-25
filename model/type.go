package model

import (
	"fmt"
	"log"
	"reflect"
)

type Kind int

const (
	String = iota + 1
	Int
	Float
	Bool
	Ptr
	Slice
	Struct
	Map
	Custom
	Reference
)

type Field struct {
	Name      string
	Type      *Type
	Index     int
	OmitEmpty bool
}

type Type struct {
	Kind      Kind    `json:",omitempty"`
	Name      string  `json:",omitempty"`
	PkgPath   string  `json:",omitempty"`
	Elem      *Type   `json:",omitempty"` // Ptr/Slice/Map
	Fields    []Field `json:",omitempty"` // Struct
	OmitEmpty bool    `json:",omitempty"` // Struct fields may be omitted!
}

func (self *Type) Visit(visitor Interface) {
	resume := func() {}

	resumeElem := func() {
		self.Elem.Visit(visitor)
	}

	switch self.Kind {
	case String:
		visitor.VisitString(self.Name, resume)
	case Int:
		visitor.VisitInt(self.Name, resume)
	case Float:
		visitor.VisitFloat(self.Name, resume)
	case Bool:
		visitor.VisitBool(self.Name, resume)
	case Ptr:
		visitor.VisitPtr(self.Name, resumeElem)
	case Slice:
		visitor.VisitSlice(self.Name, resumeElem)
	case Struct:
		visitor.VisitStruct(self.Name, self.Fields, func() {
			for _, field := range self.Fields {
				visitor.VisitStructField(field, func() {
					field.Type.Visit(visitor)
				})
			}
		})
	case Map:
		visitor.VisitMap(self.Name, resumeElem)
	case Custom:
		visitor.VisitCustom(self.Name, resume)
	case Reference:
		visitor.VisitReference(self.Name, resume)
	}
}

func anonTypeName(c *int) string {
	tmp := fmt.Sprintf("anonymous.T%d", *c)
	*c++
	return tmp
}

type supportingTypes map[string]*Type

func (self supportingTypes) set(n, p string, f func() *Type) *Type {
	// This type is a reference type.
	if n != "" && p != "" {
		_, ok := self[n]
		if ok {
			// We have already marked this type, return.
			return reference(n, p)
		}

		// Mark this type to prevent infinite recursion.
		self[n] = nil

		tmp := f()
		tmp.Name = n
		tmp.PkgPath = p
		self[n] = tmp
		return reference(n, p)
	}

	// This is an anonymous type.
	return f()
}

func getTypes(t reflect.Type, s supportingTypes, c *int) *Type {
	n, p := t.String(), t.PkgPath()

	if _, ok := t.MethodByName("MarshalJSON"); ok {
		return custom(n, p)
	}

	switch t.Kind() {
	case reflect.String:
		return simple(String)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return simple(Int)
	case reflect.Float32, reflect.Float64:
		return simple(Float)
	case reflect.Bool:
		return simple(Bool)
	case reflect.Ptr:
		return s.set(n, p, func() *Type {
			return ptrOf(getTypes(t.Elem(), s, c))
		})
	case reflect.Array, reflect.Slice:
		return s.set(n, p, func() *Type {
			return arrayOf(getTypes(t.Elem(), s, c))
		})
	case reflect.Struct:
		if p == "" {
			// gets an anonymous type name and fake the package
			n = anonTypeName(c)
			p = "github.com/shutej/go2elm/anonymous"
		}
		return s.set(n, p, func() *Type {
			return structOf(getFields(t, s, c))
		})
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			log.Fatalf("unexpected map key type: %v", t.Key())
		}
		return s.set(n, p, func() *Type {
			return mapOf(getTypes(t.Elem(), s, c))
		})
	}
	log.Fatalf("unknown kind for type: %v", t)
	return nil
}

func getFields(t reflect.Type, s supportingTypes, c *int) []Field {
	fields := []Field{}
	// TODO(shutej): Anonymous fields?
	n := t.NumField()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if name == "" {
			name = field.Name
		}
		tmp := getTypes(field.Type, s, c)
		fields = append(fields, Field{
			Name:      name,
			Type:      tmp,
			Index:     len(fields),
			OmitEmpty: opts.Contains("omitempty"),
		})
	}
	return fields
}

func arrayOf(t *Type) *Type {
	return &Type{
		Kind: Slice,
		Elem: t,
	}
}

func mapOf(t *Type) *Type {
	return &Type{
		Kind: Map,
		Elem: t,
	}
}

func ptrOf(t *Type) *Type {
	return &Type{
		Kind: Ptr,
		Elem: t,
	}
}

func structOf(fields []Field) *Type {
	return &Type{
		Kind:   Struct,
		Fields: fields,
	}
}

func simple(k Kind) *Type {
	return &Type{Kind: k}
}

func custom(n, p string) *Type {
	return &Type{
		Kind:    Custom,
		Name:    n,
		PkgPath: p,
	}
}

func reference(n, p string) *Type {
	return &Type{
		Kind:    Reference,
		Name:    n,
		PkgPath: p,
	}
}
