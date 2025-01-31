# LinkGo
A link checker written in Go.

_**Note, this tool is in development**_

## Overview

I built this tool for two reasons. First, because many link-checkers that currently exist are designed to operate on webpages. Second, I simply wanted to write a tool in Go.

## Usage

This tool takes a very specific input format: a CSV, without a header, containing one URL per line. It also produces a very specific output: a JSON lines file.

Here is an example of how to run the tool:

```
linkgo -i input.csv -o output.jsonl
```

## TODO

I have several features and basic functionality planned for development, including outputting only 'dead' links and printing to standard out.