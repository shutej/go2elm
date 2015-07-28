module Anonymous.T1 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T1 = {x : Int}

type T = T T1

t : T -> T1
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 T1 ("X" := Json.Decode.int))

empty : T
empty = T {x=0}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.int x.x) ]) << t

