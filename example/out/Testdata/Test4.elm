module Testdata.Test4 where

import Array
import Dict
import Json.Decode exposing ((:=))
import Json.Encode
import List
import Maybe
import Go2Elm

import Anonymous.T0
import Anonymous.T1
import Anonymous.T2
import Anonymous.T3

type alias Test4 = {aStringObject : Anonymous.T0.T, aIntObject : Anonymous.T1.T, aFloatObject : Anonymous.T2.T, aBoolObject : Anonymous.T3.T}

type T = T Test4

t : T -> Test4
t x = case x of T y -> y

decode : Json.Decode.Decoder T
decode = Json.Decode.map T (Json.Decode.object4 Test4 ("aStringObject" := Anonymous.T0.decode) ("aIntObject" := Anonymous.T1.decode) ("aFloatObject" := Anonymous.T2.decode) ("aBoolObject" := Anonymous.T3.decode))

empty : T
empty = T {aStringObject=Anonymous.T0.empty, aIntObject=Anonymous.T1.empty, aFloatObject=Anonymous.T2.empty, aBoolObject=Anonymous.T3.empty}

encode : T -> Json.Encode.Value
encode = (\x -> Json.Encode.object [ ("aStringObject", Anonymous.T0.encode x.aStringObject), ("aIntObject", Anonymous.T1.encode x.aIntObject), ("aFloatObject", Anonymous.T2.encode x.aFloatObject), ("aBoolObject", Anonymous.T3.encode x.aBoolObject) ]) << t

