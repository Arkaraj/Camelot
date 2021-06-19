# Camelot

A CLI tool for controlling Searches in Google Chrome

### Work on Progress

🚧🚧 Work on Progress 🚧🚧

## TECH STACK

- Go
- Cobra (CLI pkg for go)

## Go Setup

https://github.com/spf13/cobra

```bash

# Get cobra package
$ go get -u github.com/spf13/cobra/cobra

# initialize repo with cobra file struct
cobra init --pkg-name camelot

# Add new command
cobra add cmd_name

```

### AIM

To create a CLI tool that allows users to web search in google chrome inside terminal.

## Commands (Till Now)

```bash

# search in google by default
$ Camelot search <query> -flags

```
