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
$ cobra init --pkg-name camelot

# Add new command
$ cobra add cmd_name

```

## Bookmark DB

- save in db config disk (no fancy sql, nosql dbs)
- json format (generate random id)
- bookmark.json 

- GET list Gets all the bookmarked website
- POST add Add to bookmarks
- DELETE rmv Remove from bookmark

#### Providers

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
- instagram
- linkedin
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

## -i open chrome in incognito mode (only works for chrome browsers)
$ Camelot search -i -p duckduckgo secret stuff

# Bookmark (Not implemented yet)
$ Camelot bookmark www.google.com

## View all Bookmarks
$ Camelot bookmark --list

```
