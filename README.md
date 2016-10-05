# shotgun

Take the shotgun approach to your terminal.

A tool for running commands in parallel on a set of directories.

## Installation

```
go get -u github.com/bbrks/shotgun
```

## Usage

```
Shotgun is a tool for running commands in parallel on a set of directories.

Usage:
  shotgun [options] command

Options:
  -d string
    	Working directory - search from here (default ".")
  -dry-run
    	Print what would be run where, without actually doing it
  -f string
    	Filter directories by RE2 regexp (default ".*")
  -version
    	Print version and exit

Examples:
  shotgun git pull
    	Run a command in each child directory of the current
  shotgun -f '^a.*$' git pull
    	Run a command in each directory beginning with the letter 'a'
  shotgun -dir $GOPATH/src/github.com/bbrks git status --short
    	Run a command for each directory in '$GOPATH/src/github.com/bbrks'
  shotgun 'git checkout -- .; git checkout develop; git fetch; git pull'
    	Wrap commands in quotes and separate by semicolons to chain sequentially
  shotgun -dry-run 'rm LICENSE.md'
    	Print what would be run where
```
