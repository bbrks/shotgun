package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func parseFlags() (*params, error) {
	var (
		f   = flag.String("f", ".*", "Filter directories by RE2 `regexp`")
		d   = flag.String("d", ".", "Working `directory` - search from here")
		c   = flag.Int("c", runtime.GOMAXPROCS(0)*8, "Maximum `number` of concurrent commands")
		v   = flag.Bool("v", false, "Print all lines of command output")
		dry = flag.Bool("dry", false, "Print what would be run where, without actually doing it")
		ver = flag.Bool("version", false, "Print version and exit")
	)

	flag.Usage = usage
	flag.Parse()

	if *ver {
		fmt.Printf("shotgun %s\n", shotgunVersion)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		return nil, errors.New("invalid command")
	}

	fRegexp, err := regexp.Compile(*f)
	if err != nil {
		return nil, errors.Wrap(err, "invalid directory filter")
	}

	dirs, err := ioutil.ReadDir(*d)
	if err != nil {
		return nil, errors.Wrap(err, "invalid working directory")
	}

	return &params{
		f:    fRegexp,
		c:    *c,
		v:    *v,
		dry:  *dry,
		dirs: dirs,
		cmd:  strings.Join(flag.Args(), " "),
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
