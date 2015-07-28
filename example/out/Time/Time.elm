module Time.Time where

import Date
import Json.Decode exposing ((:=))
import Json.Encode
import Result

import Rfc3339
import Go2Elm

type alias T = Rfc3339.Date

fromString : String -> T
fromString string =
    case Rfc3339.decode string of
      Result.Ok time -> time

decode : Json.Decode.Decoder T
decode = Json.Decode.map fromString Json.Decode.string

empty : T
empty = Rfc3339.empty

encode : T -> Json.Encode.Value
encode = Json.Encode.string << Rfc3339.encode
