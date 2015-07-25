# go2elm

This is a simple tool that generates Elm types, empty value initializers, and
JSON encoders and decoders from Go types.  This eliminates some of the drudgery
of keeping your frontend code consistent with an API implemented in Go.

## Installation

```sh
go get -u github.com/shutej/go2elm/...
go install github.com/shutej/go2elm/cmd/go2elm
```

## Usage

The tool looks for a list of packages and types to convert from a YAML file, by
default looking for `go2elm.yml` in the current working directory.  This
package [includes an example](example/go2elm.yml) of just such a file.

You may consider making a synthetic package so `go generate` will compile your
Elm static assets for you:

```go
//go:generate go2elm --yml go2elm.yml --out out
```
