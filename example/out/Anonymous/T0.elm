module Anonymous.T0 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T0 = {x : String}

decode : Json.Decode.Decoder T0
decode = (Json.Decode.object1 T0 ("X" := Json.Decode.string))

empty : T0
empty = {x=""}

encode : T0 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.string x.x) ])

