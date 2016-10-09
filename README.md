[![Shotgun Logo](https://cdn.rawgit.com/bbrks/shotgun/master/logo.svg)](https://github.com/bbrks/shotgun)

A tool for running commands in parallel on a set of directories.

[![GitHub Tag](https://img.shields.io/github/tag/bbrks/shotgun.svg?maxAge=3600)](https://github.com/bbrks/shotgun/tags)
[![Build Status](https://img.shields.io/travis/bbrks/shotgun.svg?maxAge=3600)](https://travis-ci.org/bbrks/shotgun)
[![License](https://img.shields.io/github/license/bbrks/shotgun.svg?maxAge=2592000)](https://github.com/bbrks/shotgun/blob/master/LICENSE)

Take the shotgun approach to your terminal. :boom:

## Installation/Updating

```
go get -u github.com/bbrks/shotgun
```

## Why?

Imagine you're working on a system with 20, 50, 100, or more microservices, all in separate repositories.
How do you update them all at once? Easy.

```
shotgun git pull
```

What about removing all local changes and syncing with the remote master branches?

```
shotgun 'git fetch origin; git reset --hard origin/master; git clean -f'
```

See how fast running commands in parallel vs. sequentially actually is:
<a href="https://asciinema.org/a/d3kj4vdi47orpl5tleqn0c9rx" target="_blank"><img src="https://asciinema.org/a/d3kj4vdi47orpl5tleqn0c9rx.png" width="250px"/></a>
<a href="https://asciinema.org/a/b0d16ry57hsn1vfmq2ez7u1an" target="_blank"><img src="https://asciinema.org/a/b0d16ry57hsn1vfmq2ez7u1an.png" width="250px"/></a>

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
  shotgun -f '^a' git pull
    	Run a command in each directory beginning with the letter 'a'
  shotgun -dir $GOPATH/src/github.com/bbrks git status --short
    	Run a command for each directory in '$GOPATH/src/github.com/bbrks'
  shotgun 'git checkout -- .; git checkout develop; git fetch; git pull'
    	Wrap commands in quotes and separate by semicolons to chain sequentially
  shotgun -dry-run 'rm .travis.yml'
    	Print what would be run where
```

## License
This project is licensed under the [MIT License](LICENSE).
