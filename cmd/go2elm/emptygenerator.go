package main

import (
	"github.com/shutej/go2elm/model"
)

type EmptyGenerator struct {
	*Generator
}

func (self *EmptyGenerator) withEmpty(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("empty : T\n")
		self.printf("empty = T ")
		resume()
		self.printf("\n\n")
	})
}

func (self *EmptyGenerator) VisitString(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("%q", "")
	})
}

func (self *EmptyGenerator) VisitInt(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("0")
	})
}

func (self *EmptyGenerator) VisitFloat(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("0.0")
	})
}

func (self *EmptyGenerator) VisitBool(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("False")
	})
}

func (self *EmptyGenerator) VisitPtr(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("Maybe.Nothing")
	})
}

func (self *EmptyGenerator) VisitSlice(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("Array.empty")
	})
}

func (self *EmptyGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withEmpty(name, func() {
		self.printf("{")
		resume()
		self.printf("}")
	})
}

func (self *EmptyGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("%s=", LowerCamelcase(field.Name))
	resume()
}

func (self *EmptyGenerator) VisitMap(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("Dict.empty")
	})
}

func (self *EmptyGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *EmptyGenerator) VisitReference(name string, resume func()) {
	pkg := Package(name)
	imports := self.imports()

	if self.stack.top() == pkg {
		self.printf("empty")
	} else {
		imports[pkg] = struct{}{}
		self.printf("%s.empty", pkg)
	}
}
