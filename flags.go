package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strings"
)

func parseFlags() (*params, error) {
	var (
		filter        = flag.String("f", ".*", "Filter directories by RE2 `regexp`")
		workingDir    = flag.String("d", ".", "Working `directory` - search from here")
		maxConcurrent = flag.Int("c", runtime.GOMAXPROCS(0)*8, "Maximum `number` of concurrent commands")
		verbose       = flag.Bool("v", false, "Print all lines of command output")
		dry           = flag.Bool("dry", false, "Print what would be run where, without actually doing it")
		version       = flag.Bool("version", false, "Print version and exit")
	)

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("shotgun %s\n", shotgunVersion)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		return nil, errors.New("invalid command")
	}

	fRegexp, err := regexp.Compile(*filter)
	if err != nil {
		return nil, fmt.Errorf("invalid directory filter: %w", err)
	}

	dirs, err := ioutil.ReadDir(*workingDir)
	if err != nil {
		return nil, fmt.Errorf("invalid working directory: %w", err)
	}

	return &params{
		filter:        fRegexp,
		maxConcurrent: *maxConcurrent,
		verbose:       *verbose,
		dry:           *dry,
		dirs:          dirs,
		cmd:           strings.Join(flag.Args(), " "),
	}, nil
}

var usage = func() {
	fmt.Printf("Shotgun is a tool for running commands in parallel on a set of directories.\n")

	fmt.Printf("\nUsage:\n")
	fmt.Printf("  shotgun [options] command\n")

	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("\nExamples:\n")
	fmt.Printf("  %s\n    \t%s\n",
		"shotgun git pull",
		"Run a command in each child directory of the current")
	fmt.Printf("  %s\n    \t%s\n",
		"shotgun -f '^a' git pull",
		"Run a command in each directory beginning with the letter 'a'")
	fmt.Printf("  %s\n    \t%s\n",
		"shotgun -d $GOPATH/src/github.com/bbrks git status --short",
		"Run a command for each directory in '$GOPATH/src/github.com/bbrks'")
	fmt.Printf("  %s\n    \t%s\n",
		"shotgun 'git checkout -- .; git checkout develop; git fetch; git pull'",
		"Wrap commands in quotes and separate by semicolons to chain sequentially")
	fmt.Printf("  %s\n    \t%s\n",
		"shotgun -dry 'rm .travis.yml'",
		"Print what would be run where")
}
