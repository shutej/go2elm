module Testdata.Test9 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Time.Time

type alias Test9 = {x : Time.Time.Time}

decode : Json.Decode.Decoder Test9
decode = (Json.Decode.object1 Test9 ("X" := Time.Time.decode))

empty : Test9
empty = {x=Time.Time.empty}

encode : Test9 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Time.Time.encode x.x) ])

