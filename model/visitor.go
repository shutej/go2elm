package model

type Interface interface {
	VisitString(name string, resume func())
	VisitInt(name string, resume func())
	VisitFloat(name string, resume func())
	VisitBool(name string, resume func())
	VisitPtr(name string, resume func())
	VisitSlice(name string, resume func())
	VisitStruct(name string, fields []Field, resume func())
	VisitStructField(field Field, resume func())
	VisitMap(name string, resume func())
	VisitCustom(name string, resume func())
	VisitReference(name string, resume func())
}

type Visitor struct{}

func (self *Visitor) VisitString(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitInt(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitFloat(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitBool(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitPtr(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitSlice(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitStruct(name string, fields []Field, resume func()) {
	resume()
}

func (self *Visitor) VisitStructField(field Field, resume func()) {
	resume()
}

func (self *Visitor) VisitMap(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitCustom(name string, resume func()) {
	resume()
}

func (self *Visitor) VisitReference(name string, resume func()) {
	resume()
}
