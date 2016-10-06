package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync"
)

const version = "v0.2.1"

var (
	f      = flag.String("f", ".*", "Filter directories by RE2 regexp")
	dir    = flag.String("d", ".", "Working directory - search from here")
	dryRun = flag.Bool("dry-run", false, "Print what would be run where, without actually doing it")
	ver    = flag.Bool("version", false, "Print version and exit")
)

func main() {
	flag.Parse()

	if *ver {
		fmt.Printf("shotgun %s\n", version)
		os.Exit(0)
	}

	// no arguments passed or invalid directory filter
	fRegexp, err := regexp.Compile(*f)
	if len(flag.Args()) < 1 || err != nil {
		usage()
		os.Exit(1)
	}

	// invalid working directory
	f, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	var (
		c  = strings.Join(flag.Args(), " ")
		wg = sync.WaitGroup{}
	)

	for _, d := range f {
		wg.Add(1)
		go func(d os.FileInfo) {
			defer wg.Done()
			if d.IsDir() && fRegexp.MatchString(d.Name()) {
				if *dryRun {
					fmt.Printf("Would run '%s' in '%s'\n", c, d.Name())
					return
				}
				cmd(d.Name(), c).Run()
				fmt.Printf("Running '%s' in '%s'\n", c, d.Name())
			}
		}(d)
	}
	wg.Wait()
}

var usage = func() {
	fmt.Fprintf(os.Stderr, "Shotgun is a tool for running commands in parallel on a set of directories.\n")
	fmt.Fprintf(os.Stderr, "\nUsage:\n")
	fmt.Fprintf(os.Stderr, "  shotgun [options] command\n")
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  %s\n    \t%s\n",
		"shotgun git pull",
		"Run a command in each child directory of the current")
	fmt.Fprintf(os.Stderr, "  %s\n    \t%s\n",
		"shotgun -f '^a' git pull",
		"Run a command in each directory beginning with the letter 'a'")
	fmt.Fprintf(os.Stderr, "  %s\n    \t%s\n",
		"shotgun -dir $GOPATH/src/github.com/bbrks git status --short",
		"Run a command for each directory in '$GOPATH/src/github.com/bbrks'")
	fmt.Fprintf(os.Stderr, "  %s\n    \t%s\n",
		"shotgun 'git checkout -- .; git checkout develop; git fetch; git pull'",
		"Wrap commands in quotes and separate by semicolons to chain sequentially")
	fmt.Fprintf(os.Stderr, "  %s\n    \t%s\n",
		"shotgun -dry-run 'rm .travis.yml'",
		"Print what would be run where")

}
