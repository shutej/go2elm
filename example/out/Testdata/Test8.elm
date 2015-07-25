module Testdata.Test8 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.Test2

type alias Test8 = (Array.Array (Maybe.Maybe Testdata.Test2.Test2))

decode : Json.Decode.Decoder Test8
decode = (Go2Elm.decodeSlice (Go2Elm.decodePtr Testdata.Test2.decode))

empty : Test8
empty = Array.empty

encode : Test8 -> Json.Encode.Value
encode = (Go2Elm.encodeSlice (Go2Elm.encodePtr Testdata.Test2.encode))

