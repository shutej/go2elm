module Anonymous.T2 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T2 = {x : Float}

decode : Json.Decode.Decoder T2
decode = (Json.Decode.object1 T2 ("X" := Json.Decode.float))

empty : T2
empty = {x=0.0}

encode : T2 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.float x.x) ])

