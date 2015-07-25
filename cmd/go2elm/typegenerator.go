package main

import (
	"github.com/shutej/go2elm/model"
)

type TypeGenerator struct {
	*Generator
}

func (self *TypeGenerator) withType(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		typ := UpperCamelcase(Name(name))
		self.printf("type alias %s = ", typ)
		resume()
		self.printf("\n\n")
	})
}

func (self *TypeGenerator) VisitString(name string, resume func()) {
	self.withType(name, func() {
		self.printf("String")
	})
}

func (self *TypeGenerator) VisitInt(name string, resume func()) {
	self.withType(name, func() {
		self.printf("Int")
	})
}

func (self *TypeGenerator) VisitFloat(name string, resume func()) {
	self.withType(name, func() {
		self.printf("Float")
	})
}

func (self *TypeGenerator) VisitBool(name string, resume func()) {
	self.withType(name, func() {
		self.printf("Bool")
	})
}

func (self *TypeGenerator) VisitPtr(name string, resume func()) {
	self.withType(name, func() {
		self.printf("(Maybe.Maybe ")
		resume()
		self.printf(")")
	})
}

func (self *TypeGenerator) VisitSlice(name string, resume func()) {
	self.withType(name, func() {
		self.printf("(Array.Array ")
		resume()
		self.printf(")")
	})
}

func (self *TypeGenerator) VisitStruct(name string, _ []model.Field, resume func()) {
	self.withType(name, func() {
		self.printf("{")
		resume()
		self.printf("}")
	})
}

func (self *TypeGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("%s : ", LowerCamelcase(field.Name))
	resume()
}

func (self *TypeGenerator) VisitMap(name string, resume func()) {
	self.withType(name, func() {
		self.printf("(Dict.Dict String ")
		resume()
		self.printf(")")
	})
}

func (self *TypeGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *TypeGenerator) VisitReference(name string, resume func()) {
	pkg, ref := Package(name), UpperCamelcase(Name(name))
	imports := self.imports()

	if self.stack.top() == pkg {
		self.printf("%s", ref)
	} else {
		imports[pkg] = struct{}{}
		self.printf("%s.%s", pkg, ref)
	}
}
