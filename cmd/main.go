package main

//TODO: maybe use sp13/afero for filesystem abstractions.
//TODO: testing, ... uf.

import (
	"log"
	"os"

	"github.com/gbrls/Gorganizer/pkg/org"

	"github.com/gbrls/Gorganizer/cmd/cfg"
)

func main() {
	f, err := os.Open("config.txt")
	if err != nil {
		log.Fatalf("config.txt not found (%s)\n", err)
	}

	config := cfg.NewConfig(f)
	org.Org(config)
}
