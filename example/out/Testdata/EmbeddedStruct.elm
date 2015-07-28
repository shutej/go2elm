module Testdata.EmbeddedStruct where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias EmbeddedStruct = {aString : String, aInt : Int, aFloat : Float, aBool : Bool}

type T = T EmbeddedStruct

t : T -> EmbeddedStruct
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object4 EmbeddedStruct (Json.Decode.oneOf [ ("aString" := Json.Decode.string), Json.Decode.succeed "" ]) (Json.Decode.oneOf [ ("aInt" := Json.Decode.int), Json.Decode.succeed 0 ]) (Json.Decode.oneOf [ ("aFloat" := Json.Decode.float), Json.Decode.succeed 0.0 ]) (Json.Decode.oneOf [ ("aBool" := Json.Decode.bool), Json.Decode.succeed False ]))

empty : T
empty = T {aString="", aInt=0, aFloat=0.0, aBool=False}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aString", Json.Encode.string x.aString), ("aInt", Json.Encode.int x.aInt), ("aFloat", Json.Encode.float x.aFloat), ("aBool", Json.Encode.bool x.aBool) ]) << t

