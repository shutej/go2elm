module Testdata.Test6 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.EmbeddedStruct

type alias Test6 = {embeddedStruct : Testdata.EmbeddedStruct.T}

type T = T Test6

t : T -> Test6
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object1 Test6 ("EmbeddedStruct" := Testdata.EmbeddedStruct.decode))

empty : T
empty = T {embeddedStruct=Testdata.EmbeddedStruct.empty}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("EmbeddedStruct", Testdata.EmbeddedStruct.encode x.embeddedStruct) ]) << t

