module Testdata.Test5 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Testdata.EmbeddedStruct

type alias Test5 = {aStringMap : (Dict.Dict String String), aIntMap : (Dict.Dict String Int), aFloatMap : (Dict.Dict String Float), aBoolMap : (Dict.Dict String Bool), embeddedStruct : Testdata.EmbeddedStruct.EmbeddedStruct}

decode : Json.Decode.Decoder Test5
decode = (Json.Decode.object5 Test5 ("aStringMap" := (Go2Elm.decodeMap Json.Decode.string)) ("aIntMap" := (Go2Elm.decodeMap Json.Decode.int)) ("aFloatMap" := (Go2Elm.decodeMap Json.Decode.float)) ("aBoolMap" := (Go2Elm.decodeMap Json.Decode.bool)) ("EmbeddedStruct" := Testdata.EmbeddedStruct.decode))

empty : Test5
empty = {aStringMap=Dict.empty, aIntMap=Dict.empty, aFloatMap=Dict.empty, aBoolMap=Dict.empty, embeddedStruct=Testdata.EmbeddedStruct.empty}

encode : Test5 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aStringMap", (Go2Elm.encodeMap Json.Encode.string) x.aStringMap), ("aIntMap", (Go2Elm.encodeMap Json.Encode.int) x.aIntMap), ("aFloatMap", (Go2Elm.encodeMap Json.Encode.float) x.aFloatMap), ("aBoolMap", (Go2Elm.encodeMap Json.Encode.bool) x.aBoolMap), ("EmbeddedStruct", Testdata.EmbeddedStruct.encode x.embeddedStruct) ])

