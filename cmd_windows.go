package main

import (
	"fmt"
	"os/exec"
)

// Run a command from a directory.
func cmd(d, c string) *exec.Cmd {
	return exec.Command("cmd", "/C", fmt.Sprintf("dir %s & %s", d, c))
}
