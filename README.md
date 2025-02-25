# LinkGo
A link checker written in Go.

_**Note, this tool is in development**_

## Overview

I built this tool for two reasons. First, because many link-checkers that currently exist are designed to operate on webpages. Second, I simply wanted to write a tool in Go.

## Usage

This tool takes a very specific input format: a CSV, without a header, containing one URL per line. It also produces a very specific output: a JSON lines file.

There are no 'commands' for this tool, since it's doing only and exactly one thing, though it does have several flags.
- `-i` indicates the input .csv file. (**required**)
- `-o` indicates the output .jsonl file. (**required**)
- `-a` is an optional flag that returns all (`true`) or only broken (`false`) links. The default behavior is false, only returning broken links.



Here is an example of how to run the tool:

```
linkgo -i input.csv -o output.jsonl -a false
```

### Example

Using the tool on the file provided in the example directory looks like this:
```
linkgo -i Links.csv -o Result.jsonl -a true
```
Where `Links.csv` looks like:
```
https://www.id.uscourts.gov/glossary.htm
https://www.merriam-webster.com/dictionary/anthropogenic
```
Will produce `Result.jsonl` that looks like:
```
{"url":"https://www.id.uscourts.gov/glossary.htm","response_status_code":404,"response_status":"Not Found","time_of_validation":"2025-02-01T13:32:20Z"}

{"url":"https://www.merriam-webster.com/dictionary/anthropogenic","response_status_code":200,"response_status":"OK","time_of_validation":"2025-02-01T13:32:20Z"}

```


## TODO

I have several features and basic functionality planned for development, including:
- ✅ Outputting only 'dead' links
- ⬜️ Printing to standard out.
- ⬜️ Using GoRoutines to speed up link checking
- ⬜️ Handle different output formats, like .csv or .tsv.