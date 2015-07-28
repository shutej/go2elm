package main

import (
	"github.com/shutej/go2elm/model"
)

type DecodeGenerator struct {
	*Generator
}

func (self *DecodeGenerator) withDecode(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("decode : Json.Decode.Decoder T\n")
		self.printf("decode = Json.Decode.map T ")
		resume()
		self.printf("\n\n")
	})
}

func (self *DecodeGenerator) VisitString(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("Json.Decode.string")
	})
}

func (self *DecodeGenerator) VisitInt(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("Json.Decode.int")
	})
}

func (self *DecodeGenerator) VisitFloat(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("Json.Decode.float")
	})
}

func (self *DecodeGenerator) VisitBool(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("Json.Decode.bool")
	})
}

func (self *DecodeGenerator) VisitPtr(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("(Go2Elm.decodePtr ")
		resume()
		self.printf(")")
	})
}

func (self *DecodeGenerator) VisitSlice(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("(Go2Elm.decodeSlice ")
		resume()
		self.printf(")")
	})
}

func (self *DecodeGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withDecode(name, func() {
		typ := UpperCamelcase(Name(name))
		self.printf("(Json.Decode.object%d %s ", len(fields), typ)
		resume()
		self.printf(")")
	})
}

func (self *DecodeGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(" ")
	}
	if field.OmitEmpty {
		self.printf("(Json.Decode.oneOf [ ")
	}
	self.printf("(%q := ", field.Name)
	resume()
	self.printf(")")
	if field.OmitEmpty {
		self.printf(", Json.Decode.succeed ")
		field.Type.Visit(&EmptyGenerator{self.Generator})
		self.printf(" ])")

	}
}

func (self *DecodeGenerator) VisitMap(name string, resume func()) {
	self.withDecode(name, func() {
		self.printf("(Go2Elm.decodeMap ")
		resume()
		self.printf(")")
	})
}

func (self *DecodeGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *DecodeGenerator) VisitReference(name string, resume func()) {
	pkg := Package(name)
	imports := self.imports()

	if self.stack.top() == pkg {
		self.printf("decode")
	} else {
		imports[pkg] = struct{}{}
		self.printf("%s.decode", pkg)
	}
}
