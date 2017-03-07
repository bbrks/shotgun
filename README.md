# [![Shotgun Logo](https://cdn.rawgit.com/bbrks/shotgun/master/logo.svg)](https://github.com/bbrks/shotgun)

A tool for running commands in parallel on a set of directories.

[![Build Status](https://travis-ci.org/bbrks/shotgun.svg)](https://travis-ci.org/bbrks/shotgun)
[![GoDoc](https://godoc.org/github.com/bbrks/shotgun?status.svg)](https://godoc.org/github.com/bbrks/shotgun)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbrks/shotgun)](https://goreportcard.com/report/github.com/bbrks/shotgun)
[![GitHub tag](https://img.shields.io/github/tag/bbrks/shotgun.svg)](https://github.com/bbrks/shotgun/releases)
[![license](https://img.shields.io/github/license/bbrks/shotgun.svg)](https://github.com/bbrks/shotgun/blob/master/LICENSE)

Take the shotgun approach to your terminal. :boom:

## Installation/Updating

```sh
$ go get -u github.com/bbrks/shotgun
```

## Why?

Imagine you're working on a system with 20, 50, 100, or more microservices, all in separate repositories.
How do you update them all at once? Easy.

```sh
$ shotgun git pull
```

What about removing all local changes and syncing with the remote master branches?

```sh
$ shotgun 'git fetch origin; git reset --hard origin/master; git clean -f'
```

Not convinced yet? See just how fast running commands in parallel vs. sequentially actually is:

<a href="https://asciinema.org/a/d3kj4vdi47orpl5tleqn0c9rx" target="_blank"><img src="http://i.imgur.com/7xqA67x.gif" width="250px"/></a>
<a href="https://asciinema.org/a/b0d16ry57hsn1vfmq2ez7u1an" target="_blank"><img src="http://i.imgur.com/e9T6YY0.gif" width="250px"/></a>

## Usage

[embedmd]:# (doc.go text /\*\n/ /\n\*/)
```text
*

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

*
```

## License
This project is licensed under the [MIT License](LICENSE).
