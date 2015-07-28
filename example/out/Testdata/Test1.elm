module Testdata.Test1 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias Test1 = {aString : String, aInt : Int, aFloat : Float, aBool : Bool}

type T = T Test1

t : T -> Test1
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object4 Test1 ("aString" := Json.Decode.string) ("aInt" := Json.Decode.int) ("aFloat" := Json.Decode.float) ("aBool" := Json.Decode.bool))

empty : T
empty = T {aString="", aInt=0, aFloat=0.0, aBool=False}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aString", Json.Encode.string x.aString), ("aInt", Json.Encode.int x.aInt), ("aFloat", Json.Encode.float x.aFloat), ("aBool", Json.Encode.bool x.aBool) ]) << t

