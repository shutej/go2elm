package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shutej/go2elm/inception"
	yaml "gopkg.in/yaml.v1"
)

var (
	yml = flag.String("yml", "go2elm.yml", "configuration file")
	out = flag.String("out", ".", "directory to write output to")
)

func main() {
	flag.Parse()

	configs := &inception.Configs{}
	input, err := ioutil.ReadFile(*yml)
	if err != nil {
		log.Fatalf("Opening config failed: %v", err)
	}
	if err := yaml.Unmarshal(input, configs); err != nil {
		log.Fatalf("Parsing config failed: %v", err)
	}

	types, err := inception.Inception(*configs)
	if err != nil {
		log.Fatalf("Loading input failed: %v", err)
	}

	if err := os.Chdir(*out); err != nil {
		log.Fatalf("Changing to output directory %q failed: %v", *out, err)
	}

	if err := ioutil.WriteFile("Go2Elm.elm", []byte(Go2Elm), 0640); err != nil {
		log.Fatalf("Writing runtime failed: %v", err)
	}

	visitor := &Generator{}
	visitor.Visit(types)
	for pkg, buffer := range visitor.Buffers() {
		parts := strings.Split(pkg, ".")
		os.MkdirAll(filepath.Join(parts[0:len(parts)-1]...), 0750)
		if err := ioutil.WriteFile(filepath.Join(parts...)+".elm", buffer.Bytes(), 0640); err != nil {
			log.Fatalf("Writing output failed: %v", err)
		}
	}
}
