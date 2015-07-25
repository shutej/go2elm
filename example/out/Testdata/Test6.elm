module Testdata.Test6 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.EmbeddedStruct

type alias Test6 = {embeddedStruct : Testdata.EmbeddedStruct.EmbeddedStruct}

decode : Json.Decode.Decoder Test6
decode = (Json.Decode.object1 Test6 ("EmbeddedStruct" := Testdata.EmbeddedStruct.decode))

empty : Test6
empty = {embeddedStruct=Testdata.EmbeddedStruct.empty}

encode : Test6 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("EmbeddedStruct", Testdata.EmbeddedStruct.encode x.embeddedStruct) ])

