package main

import (
	"bytes"
	"fmt"
	"strings"

	camelcase "github.com/segmentio/go-camelcase"
	"github.com/shutej/go2elm/model"
)

const Go2Elm = `
module Go2Elm where

import Array
import Dict
import Json.Decode
import Json.Encode
import List
import Maybe

decodePtr : Json.Decode.Decoder a -> Json.Decode.Decoder (Maybe.Maybe a)
decodePtr decoder =
    Json.Decode.oneOf [ Json.Decode.null Maybe.Nothing, Json.Decode.map Maybe.Just decoder ]

decodeSlice : Json.Decode.Decoder a -> Json.Decode.Decoder (Array.Array a)
decodeSlice decoder =
    Json.Decode.oneOf [ Json.Decode.null Array.empty, Json.Decode.array decoder ]

decodeMap : Json.Decode.Decoder a -> Json.Decode.Decoder (Dict.Dict String a)
decodeMap decoder =
    Json.Decode.oneOf [ Json.Decode.null Dict.empty, Json.Decode.dict decoder ]

encodeMaybe : (a -> Json.Encode.Value) -> Maybe.Maybe a -> Json.Encode.Value
encodeMaybe encode maybe =
  case maybe of
    Maybe.Just value -> encode value
    Maybe.Nothing    -> Json.Encode.null

encodeSlice : (a -> Json.Encode.Value) -> Array.Array a -> Json.Encode.Value
encodeSlice encode slice =
  Array.map encode slice |> Json.Encode.array

encodeMap : (a -> Json.Encode.Value) -> Dict.Dict String a -> Json.Encode.Value
encodeMap encoder map =
  Dict.map (\ignore -> encoder) map |> Dict.toList |> Json.Encode.object
`

func UpperCamelcase(s string) string {
	s = camelcase.Camelcase(s)
	return strings.ToUpper(s[0:1]) + s[1:len(s)]
}

func LowerCamelcase(s string) string {
	s = camelcase.Camelcase(s)
	return strings.ToLower(s[0:1]) + s[1:len(s)]
}

func Package(s string) string {
	tmp := strings.SplitN(s, ".", 2)
	return fmt.Sprintf("%s.%s", UpperCamelcase(tmp[0]), UpperCamelcase(tmp[1]))
}

func Name(s string) string {
	tmp := strings.SplitN(s, ".", 2)
	return tmp[1]
}

type Generator struct {
	buffers  map[string]*bytes.Buffer
	imports_ map[string]map[string]struct{}
	stack    StringStack
}

func (self *Generator) Visit(types model.Types) {
	types.Visit(&TypeGenerator{self})
	types.Visit(&DecodeGenerator{self})
	types.Visit(&EmptyGenerator{self})
	types.Visit(&EncodeGenerator{self})
}

const prelude = `
import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

`

func (self *Generator) Buffers() map[string]*bytes.Buffer {
	retval := map[string]*bytes.Buffer{}
	for pkg, epilogue := range self.buffers {
		buffer := bytes.NewBuffer([]byte{})
		fmt.Fprintf(buffer, "module %s where\n", pkg)
		fmt.Fprintf(buffer, "%s", prelude)
		imports := self.imports_[pkg]
		for import_ := range imports {
			fmt.Fprintf(buffer, "import %s\n", import_)
		}
		fmt.Fprintf(buffer, "\n")
		epilogue.WriteTo(buffer)
		retval[pkg] = buffer
	}
	return retval
}

func (self *Generator) imports() map[string]struct{} {
	if self.imports_ == nil {
		self.imports_ = map[string]map[string]struct{}{}
	}
	pkg := self.stack.top()
	imports, ok := self.imports_[pkg]
	if !ok {
		imports = map[string]struct{}{}
		self.imports_[pkg] = imports
	}
	return imports
}

func (self *Generator) buffer() *bytes.Buffer {
	if self.buffers == nil {
		self.buffers = map[string]*bytes.Buffer{}
	}

	pkg := self.stack.top()
	buffer, ok := self.buffers[pkg]
	if !ok {
		buffer = bytes.NewBuffer([]byte{})
		self.buffers[pkg] = buffer
	}
	return buffer
}

func (self *Generator) printf(format string, args ...interface{}) {
	fmt.Fprintf(self.buffer(), format, args...)
}

func (self *Generator) withPackage(name string, fn func()) {
	if name == "" {
		fn()
		return
	}
	self.stack.with(Package(name), fn)
}
