module Anonymous.T3 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T3 = {x : Bool}

decode : Json.Decode.Decoder T3
decode = (Json.Decode.object1 T3 ("X" := Json.Decode.bool))

empty : T3
empty = {x=False}

encode : T3 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.bool x.x) ])

