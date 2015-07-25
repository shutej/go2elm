package main

import (
	"github.com/shutej/go2elm/model"
)

type EncodeGenerator struct {
	*Generator
}

func (self *EncodeGenerator) withEncode(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		typ := UpperCamelcase(Name(name))
		self.printf("encode : %s -> Json.Encode.Value\n", typ)
		self.printf("encode = ")
		resume()
		self.printf("\n\n")
	})
}

func (self *EncodeGenerator) VisitString(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("Json.Encode.string")
	})
}

func (self *EncodeGenerator) VisitInt(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("Json.Encode.int")
	})
}

func (self *EncodeGenerator) VisitFloat(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("Json.Encode.float")
	})
}

func (self *EncodeGenerator) VisitBool(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("Json.Encode.bool")
	})
}

func (self *EncodeGenerator) VisitPtr(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("(Go2Elm.encodeMaybe ")
		resume()
		self.printf(")")
	})
}

func (self *EncodeGenerator) VisitSlice(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("(Go2Elm.encodeSlice ")
		resume()
		self.printf(")")
	})
}

func (self *EncodeGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withEncode(name, func() {
		self.printf("(\\x -> Json.Encode.object [ ")
		resume()
		self.printf(" ])")
	})
}

func (self *EncodeGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("(%q, ", field.Name)
	resume()
	self.printf(" x.%s)", LowerCamelcase(field.Name))
}

func (self *EncodeGenerator) VisitMap(name string, resume func()) {
	self.withEncode(name, func() {
		self.printf("(Go2Elm.encodeMap ")
		resume()
		self.printf(")")
	})
}

func (self *EncodeGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *EncodeGenerator) VisitReference(name string, resume func()) {
	pkg := Package(name)
	imports := self.imports()

	if self.stack.top() == pkg {
		self.printf("encode")
	} else {
		imports[pkg] = struct{}{}
		self.printf("%s.encode", pkg)
	}
}
