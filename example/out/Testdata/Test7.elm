module Testdata.Test7 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.Test2

type alias Test7 = (Maybe.Maybe Testdata.Test2.Test2)

decode : Json.Decode.Decoder Test7
decode = (Go2Elm.decodePtr Testdata.Test2.decode)

empty : Test7
empty = Maybe.Nothing

encode : Test7 -> Json.Encode.Value
encode = (Go2Elm.encodeMaybe Testdata.Test2.encode)

