module Anonymous.T3 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T3 = {x : Bool}

type T = T T3

t : T -> T3
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 T3 ("X" := Json.Decode.bool))

empty : T
empty = T {x=False}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.bool x.x) ]) << t

