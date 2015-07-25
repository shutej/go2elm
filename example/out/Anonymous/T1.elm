module Anonymous.T1 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T1 = {x : Int}

decode : Json.Decode.Decoder T1
decode = (Json.Decode.object1 T1 ("X" := Json.Decode.int))

empty : T1
empty = {x=0}

encode : T1 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.int x.x) ])

