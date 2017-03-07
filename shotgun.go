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
	f    *regexp.Regexp
	c    int
	v    bool
	dry  bool
	dirs []os.FileInfo
	cmd  string
	sync.WaitGroup
}

// TODO: Use channels to return command output and errors
func (p *params) shotgun() error {
	if p == nil {
		return errors.New("invalid command")
	}

	t := make(chan struct{}, p.c)

	for _, d := range p.dirs {
		p.Add(1)
		t <- struct{}{}

		go func(d os.FileInfo) {
			defer p.Done()
			defer func() { <-t }()

			if d.IsDir() && p.f.MatchString(d.Name()) {
				if p.dry {
					fmt.Printf("Would run '%s' in '%s'\n", p.cmd, d.Name())
					return
				}

				command := run(d.Name(), p.cmd)

				if p.v {
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
