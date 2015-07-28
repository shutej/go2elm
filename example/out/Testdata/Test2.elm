module Testdata.Test2 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm


type alias Test2 = {aStringPtr : (Maybe.Maybe String), aIntPtr : (Maybe.Maybe Int), aFloatPtr : (Maybe.Maybe Float), aBoolPtr : (Maybe.Maybe Bool)}

type T = T Test2

t : T -> Test2
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object4 Test2 ("aStringPtr" := (Go2Elm.decodePtr Json.Decode.string)) ("aIntPtr" := (Go2Elm.decodePtr Json.Decode.int)) ("aFloatPtr" := (Go2Elm.decodePtr Json.Decode.float)) ("aBoolPtr" := (Go2Elm.decodePtr Json.Decode.bool)))

empty : T
empty = T {aStringPtr=Maybe.Nothing, aIntPtr=Maybe.Nothing, aFloatPtr=Maybe.Nothing, aBoolPtr=Maybe.Nothing}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aStringPtr", (Go2Elm.encodePtr Json.Encode.string) x.aStringPtr), ("aIntPtr", (Go2Elm.encodePtr Json.Encode.int) x.aIntPtr), ("aFloatPtr", (Go2Elm.encodePtr Json.Encode.float) x.aFloatPtr), ("aBoolPtr", (Go2Elm.encodePtr Json.Encode.bool) x.aBoolPtr) ]) << t

