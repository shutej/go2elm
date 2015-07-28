module Anonymous.T2 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T2 = {x : Float}

type T = T T2

t : T -> T2
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 T2 ("X" := Json.Decode.float))

empty : T
empty = T {x=0.0}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.float x.x) ]) << t

