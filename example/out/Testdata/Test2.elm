module Testdata.Test2 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias Test2 = {aStringPtr : (Maybe.Maybe String), aIntPtr : (Maybe.Maybe Int), aFloatPtr : (Maybe.Maybe Float), aBoolPtr : (Maybe.Maybe Bool)}

decode : Json.Decode.Decoder Test2
decode = (Json.Decode.object4 Test2 ("aStringPtr" := (Go2Elm.decodePtr Json.Decode.string)) ("aIntPtr" := (Go2Elm.decodePtr Json.Decode.int)) ("aFloatPtr" := (Go2Elm.decodePtr Json.Decode.float)) ("aBoolPtr" := (Go2Elm.decodePtr Json.Decode.bool)))

empty : Test2
empty = {aStringPtr=Maybe.Nothing, aIntPtr=Maybe.Nothing, aFloatPtr=Maybe.Nothing, aBoolPtr=Maybe.Nothing}

encode : Test2 -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aStringPtr", (Go2Elm.encodeMaybe Json.Encode.string) x.aStringPtr), ("aIntPtr", (Go2Elm.encodeMaybe Json.Encode.int) x.aIntPtr), ("aFloatPtr", (Go2Elm.encodeMaybe Json.Encode.float) x.aFloatPtr), ("aBoolPtr", (Go2Elm.encodeMaybe Json.Encode.bool) x.aBoolPtr) ])

