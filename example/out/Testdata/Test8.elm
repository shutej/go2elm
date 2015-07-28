module Testdata.Test8 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.Test2

type alias Test8 = (Array.Array (Maybe.Maybe Testdata.Test2.T))

type T = T Test8

t : T -> Test8
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Go2Elm.decodeSlice (Go2Elm.decodePtr Testdata.Test2.decode))

empty : T
empty = T Array.empty

encode : T -> Json.Encode.Value
encode = (Go2Elm.encodeSlice (Go2Elm.encodePtr Testdata.Test2.encode)) << t

