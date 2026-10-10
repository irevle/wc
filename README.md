# wc (word counter)

A small `wc` command clone I'm writing in Go to learn the language. It's simple enough that it could probably survive real use once I finish it, but it's not there yet. The flags (`-l`, `-w`, `-c`, `-m`) haven't been implemented.

## What it does

Point it at a file and it prints the line count, word count, character count, and filename:

```sh
$ go run main.go test.txt
---- FINAL OUTPUT ----
 3  9 67 test.txt
```

How each count works:

- **Lines** are counted per `bufio.Scanner` iteration.
- **Words** come from `strings.Fields`, which splits on runs of whitespace — so `"double  space"` in `test.txt` is handled fine, no special casing needed.
- **Characters** are counted as `len(line) + 1`, where the `+1` is the newline. Note that this is really a _byte_ count, not a character count. See the problems section.

## Building

```sh
go build -o wc .
./wc test.txt
```

Requires Go 1.27.1+ (whatever's pinned in `go.mod`).

## Known problems

Honest list of things I know are off and want to fix:

- **Byte count isn't a character count.** `len()` counts bytes, so UTF-8 text gets overcounted. Real character counting needs `utf8.RuneCountInString`.
- **Assumes every line ends in a newline.** The `+1` per line means a file without a trailing `\n` reports one extra "character".
- **Errors are panics.** `check()` panics on failure, so a typo'd filename dumps a stack trace instead of a clean `no such file: foo.txt`. Fine for learning, bad for production.
- **One file at a time.** No multiple files, no stdin, no `total` row.

## To do

- [ ] Implement `-l`, `-w`, `-c`, `-m` (and allow combining them, e.g. `-lw`)
- [ ] Accept multiple files and print a `total` row
- [ ] Read from stdin when no file is given
- [ ] Count characters as runes instead of bytes
- [ ] Replace panics with proper error messages and exit codes
- [ ] Add tests

`test.txt` is the sample file I've been running it against.

## License

MIT (see [LICENSE](LICENSE)). GNU's `wc` is GPLv3, but this is a fresh implementation that borrows no coreutils code, so it didn't have to be.
