module Time.Time where

import Date
import Json.Decode exposing ((:=))
import Json.Encode
import Result

import Rfc3339
import Go2Elm

type alias Time = Rfc3339.Date

fromString : String -> Time
fromString string =
    case Rfc3339.decode string of
      Result.Ok time -> time

decode : Json.Decode.Decoder Time
decode = Json.Decode.map fromString Json.Decode.string

empty : Time
empty = Rfc3339.empty

encode : Time -> Json.Encode.Value
encode = Json.Encode.string << Rfc3339.encode
