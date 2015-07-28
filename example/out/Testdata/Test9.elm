module Testdata.Test9 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Time.Time

type alias Test9 = {x : Time.Time.T}

type T = T Test9

t : T -> Test9
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 Test9 ("X" := Time.Time.decode))

empty : T
empty = T {x=Time.Time.empty}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("X", Time.Time.encode x.x) ]) << t

