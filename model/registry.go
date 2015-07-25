package model

import "reflect"

type Registry struct {
	s map[string]*Type
	c int
}

func NewRegistry() *Registry {
	return &Registry{
		s: map[string]*Type{},
	}
}

func (self *Registry) Register(t reflect.Type) *Type {
	return getTypes(t, self.s, &self.c)
}

type Types []*Type

func (self Types) Visit(visitor Interface) {
	for _, typ := range self {
		typ.Visit(visitor)
	}
}

func (self *Registry) Types() Types {
	retval := Types{}
	for _, typ := range self.s {
		retval = append(retval, typ)
	}
	return retval
}
