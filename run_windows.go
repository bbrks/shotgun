package main

import (
	"fmt"
	"os/exec"
)

// Run a command from a directory.
func run(dir, cmd string) *exec.Cmd {
	return exec.Command("cmd", "/C", fmt.Sprintf("cd %s & %s", dir, cmd))
}
