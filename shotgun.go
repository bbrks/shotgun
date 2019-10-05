package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

type params struct {
	filter        *regexp.Regexp
	maxConcurrent int
	verbose       bool
	dry           bool
	dirs          []os.FileInfo
	cmd           string
	sync.WaitGroup
}

// TODO: Use channels to return command output and errors
func (p *params) shotgun() error {
	if p == nil {
		return errors.New("invalid command")
	}

	throttle := make(chan struct{}, p.maxConcurrent)

	for _, d := range p.dirs {
		p.Add(1)
		throttle <- struct{}{}

		go func(d os.FileInfo) {
			defer p.Done()
			defer func() { <-throttle }()

			if d.IsDir() && p.filter.MatchString(d.Name()) {
				if p.dry {
					fmt.Printf("Would run '%s' in '%s'\n", p.cmd, d.Name())
					return
				}

				command := run(d.Name(), p.cmd)

				if p.verbose {
					fmt.Printf("Running '%s' in '%s'\n", p.cmd, d.Name())
					command.Stdout = os.Stdout
					command.Stderr = os.Stderr
				}

				if err := command.Run(); err != nil {
					log.Printf("ERROR: %s", err)
				}
			}

		}(d)

	}

	p.Wait()

	return nil
}
