
module Go2Elm where

import Array
import Dict
import Json.Decode
import Json.Encode
import List
import Maybe

decodePtr : Json.Decode.Decoder a -> Json.Decode.Decoder (Maybe.Maybe a)
decodePtr decoder =
    Json.Decode.oneOf [ Json.Decode.null Maybe.Nothing, Json.Decode.map Maybe.Just decoder ]

decodeSlice : Json.Decode.Decoder a -> Json.Decode.Decoder (Array.Array a)
decodeSlice decoder =
    Json.Decode.oneOf [ Json.Decode.null Array.empty, Json.Decode.array decoder ]

decodeMap : Json.Decode.Decoder a -> Json.Decode.Decoder (Dict.Dict String a)
decodeMap decoder =
    Json.Decode.oneOf [ Json.Decode.null Dict.empty, Json.Decode.dict decoder ]

encodeMaybe : (a -> Json.Encode.Value) -> Maybe.Maybe a -> Json.Encode.Value
encodeMaybe encode maybe =
  case maybe of
    Maybe.Just value -> encode value
    Maybe.Nothing    -> Json.Encode.null

encodeSlice : (a -> Json.Encode.Value) -> Array.Array a -> Json.Encode.Value
encodeSlice encode slice =
  Array.map encode slice |> Json.Encode.array

encodeMap : (a -> Json.Encode.Value) -> Dict.Dict String a -> Json.Encode.Value
encodeMap encoder map =
  Dict.map (\ignore -> encoder) map |> Dict.toList |> Json.Encode.object
