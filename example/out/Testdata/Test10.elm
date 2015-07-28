module Testdata.Test10 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias Test10 = (Maybe.Maybe T)

type T = T Test10

t : T -> Test10
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Go2Elm.decodePtr decode)

empty : T
empty = T Maybe.Nothing

encode : T -> Json.Encode.Value
encode = (Go2Elm.encodePtr encode) << t

