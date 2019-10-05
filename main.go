package main

import (
	"log"
	"os"
)

const shotgunVersion = "v1.0"

func main() {
	params, err := parseFlags()
	if err != nil {
		log.Printf("Error parsing flags: %v\n", err)
		usage()
		os.Exit(1)
	}

	if err := params.shotgun(); err != nil {
		log.Printf("Unable to run shotgun: %v\n", err)
	}
}
