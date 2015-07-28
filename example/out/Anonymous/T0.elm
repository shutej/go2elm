module Anonymous.T0 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias T0 = {x : String}

type T = T T0

t : T -> T0
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 T0 ("X" := Json.Decode.string))

empty : T
empty = T {x=""}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Json.Encode.string x.x) ]) << t

