// Copyright 2012 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inception

// This file contains the model construction by reflection.

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/shutej/go2elm/model"
)

var (
	progOnly = flag.Bool("prog_only", false, "(reflect mode) Only generate the reflection program; write it to stdout.")
	execOnly = flag.String("exec_only", "", "(reflect mode) If set, execute this reflection program.")
)

type Configs []Config

type Config struct {
	Import string
	Types  []string
}

func Inception(configs Configs) (model.Types, error) {
	// TODO: sanity check arguments

	progPath := *execOnly
	if *execOnly == "" {
		// We use TempDir instead of TempFile so we can control the filename.
		tmpDir, err := ioutil.TempDir("", "gomock_reflect_")
		if err != nil {
			return nil, err
		}
		defer func() { os.RemoveAll(tmpDir) }()
		const progSource = "prog.go"
		var progBinary = "prog.bin"
		if runtime.GOOS == "windows" {
			// Windows won't execute a program unless it has a ".exe" suffix.
			progBinary += ".exe"
		}

		// Generate program.
		var program bytes.Buffer
		if err := reflectProgram.Execute(&program, configs); err != nil {
			return nil, err
		}
		if *progOnly {
			io.Copy(os.Stdout, &program)
			os.Exit(0)
		}
		if err := ioutil.WriteFile(filepath.Join(tmpDir, progSource), program.Bytes(), 0600); err != nil {
			return nil, err
		}

		// Build the program.
		cmd := exec.Command("go", "build", "-o", progBinary, progSource)
		cmd.Dir = tmpDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return nil, err
		}
		progPath = filepath.Join(tmpDir, progBinary)
	}

	// Run it.
	cmd := exec.Command(progPath)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Process output.
	types := model.Types{}
	if err := json.NewDecoder(&stdout).Decode(&types); err != nil {
		return nil, err
	}
	return types, nil
}

// This program reflects on an interface value, and prints the
// JSON encoding of a model.Package to standard output.
// JSON doesn't work because of the model.Type interface.
var reflectProgram = template.Must(template.New("program").Parse(`
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/shutej/go2elm/model"

	{{range $index, $config := .}}
	pkg_{{$index}}_ {{printf "%q" $config.Import}}
	{{end}}
)

func main() {
	registry := model.NewRegistry()
	{{range $index, $config := .}}
	{{range $config.Types}}
	registry.Register(reflect.TypeOf((*pkg_{{$index}}_.{{.}})(nil)).Elem())
	{{end}}
	{{end}}

	if err := json.NewEncoder(os.Stdout).Encode(registry.Types()); err != nil {
		fmt.Fprintf(os.Stderr, "json encode: %v\n", err)
		os.Exit(1)
	}
}
`))
