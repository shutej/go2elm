module Testdata.Test3 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias Test3 = {aStringSlice : (Array.Array String), aIntSlice : (Array.Array Int), aFloatSlice : (Array.Array Float), aBoolSlice : (Array.Array Bool)}

decode : Json.Decode.Decoder Test3
decode = (Json.Decode.object4 Test3 ("aStringSlice" := (Go2Elm.decodeSlice Json.Decode.string)) ("aIntSlice" := (Go2Elm.decodeSlice Json.Decode.int)) ("aFloatSlice" := (Go2Elm.decodeSlice Json.Decode.float)) ("aBoolSlice" := (Go2Elm.decodeSlice Json.Decode.bool)))

empty : Test3
empty = {aStringSlice=Array.empty, aIntSlice=Array.empty, aFloatSlice=Array.empty, aBoolSlice=Array.empty}

encode : Test3 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aStringSlice", (Go2Elm.encodeSlice Json.Encode.string) x.aStringSlice), ("aIntSlice", (Go2Elm.encodeSlice Json.Encode.int) x.aIntSlice), ("aFloatSlice", (Go2Elm.encodeSlice Json.Encode.float) x.aFloatSlice), ("aBoolSlice", (Go2Elm.encodeSlice Json.Encode.bool) x.aBoolSlice) ])

