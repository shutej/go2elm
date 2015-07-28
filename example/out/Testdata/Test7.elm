module Testdata.Test7 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.Test2

type alias Test7 = (Maybe.Maybe Testdata.Test2.T)

type T = T Test7

t : T -> Test7
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Go2Elm.decodePtr Testdata.Test2.decode)

empty : T
empty = T Maybe.Nothing

encode : T -> Json.Encode.Value
encode = (Go2Elm.encodePtr Testdata.Test2.encode) << t

