/*

Shotgun is a tool for running commands in parallel on a set of directories.

Usage:
	shotgun [options] command

Options:
	-d directory
		Working directory - search from here (default ".")
	-f regexp
		Filter directories by RE2 regexp (default ".*")
	-c number
		Maximum number of concurrent commands (default 64)
	-v
		Print all lines of command output
	-dry
		Print what would be run where, without actually doing it
	-version
		Print version and exit

Examples:
	shotgun git pull
		Run a command in each child directory of the current

	shotgun -f '^a' git pull
		Run a command in each directory beginning with the letter 'a'

	shotgun -d $GOPATH/src/github.com/bbrks git status --short
		Run a command for each directory in '$GOPATH/src/github.com/bbrks'

	shotgun 'git checkout -- .; git checkout develop; git fetch; git pull'
		Wrap commands in quotes and separate by semicolons to chain sequentially

	shotgun -dry 'rm .travis.yml'
		Print what would be run where

*/
package main
