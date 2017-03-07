package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

const shotgunVersion = "v1.0"

func main() {
	params, err := parseFlags()
	if err != nil {
		log.Println(errors.WithMessage(err, "Error parsing flags"))
		usage()
		os.Exit(1)
	}

	if err := params.shotgun(); err != nil {
		log.Fatal(errors.WithMessage(err, "Unable to run shotgun"))
	}
}
