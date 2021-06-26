# Camelot

A CLI tool for searching the web and bookmarking sites.

## TECH STACK

- Go
- Cobra (CLI pkg for go)

## Install

```bash

git clone https://github.com/Arkaraj/Camelot.git

or

go get -v github.com/Arkaraj/Camelot

cd Camelot

go build .

go install Camelot

```

< ## Go Setup

https://github.com/spf13/cobra

```bash

# Get cobra package
$ go get -u github.com/spf13/cobra/cobra

# initialize repo with cobra file struct
$ cobra init --pkg-name camelot

# Add new command
$ cobra add cmd_name

```

>

## Bookmark DB

- save in db config disk (no fancy sql, nosql dbs)
- json format (generate auto increment id)
- bookmark.json

- -l list Gets all the bookmarked website
- -a url add Add to bookmarks
- -r id Remove from bookmark

## Providers

- Google (default)
- Amazon
- bing
- codepen
- dribbble
- duckduckgo
- facebook
- flipkart
- giphy
- github
- hackernews
- imdb
- instagram
- linkedin
- medium
- netflix
- npm, npmsearch
- reddit
- spotify
- stackoverflow
- twitter
- wikipedia
- youtube

## Commands

```bash

# Search in google by default, opens in default web browser
$ Camelot search <query> -flags

# flags
## -l lists out the providers
$ Camelot search -l

## -p specify the provider
$ Camelot search -p amazon t-shirt

## -o to get Link of the website
$ Camelot search -p imdb fight club -o

## -i open chrome in incognito mode (only works for chrome browsers)
$ Camelot search -i -p duckduckgo secret stuff

## Raw url search
$ Camelot search -u https://www.google.com/

# Bookmark (Not implemented yet)
$ Camelot bookmark https://www.google.com/

### In Camelot Directory

## View all Bookmarks
$ Camelot bookmark --list

## Add Url to Bookmarks
$ Camelot bookmark -a https://rxjs.dev/

## Remove Bookmark
$ Camelot bookmark -r 4

### Globally,  By Entering Directory path
$ camelot bookmark -l -d /Users/arkarajghosh/go/src/Camelot

```

## Requirements

- Go
