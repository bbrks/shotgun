// +build !windows

package main

import (
	"fmt"
	"os/exec"
)

// Run a command from a directory.
func cmd(d, c string) *exec.Cmd {
	return exec.Command("sh", "-c", fmt.Sprintf("cd %s; %s", d, c))
}
