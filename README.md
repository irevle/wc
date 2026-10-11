# wc (word counter)

A small `wc` command clone I'm writing in Go to learn the language. It's simple enough that it could probably survive real use once I finish it.

## Usage

```sh
go run main.go [-l | -w | -c | -m] <filename>
```

With no flags it prints lines, words, and bytes similar to the real `wc`:

```sh
$ go run main.go test.txt
  3  9 67 test.txt
```

Numbers are right-justified to the width of the widest count, so the columns stay aligned without padding smaller numbers more than they need. Column order is always lines → words → bytes → chars.

Each count has a short and a long flag, and the short ones can be combined, so `-lw` works:

- `-l` / `--lines` — newline count
- `-w` / `--words` — word count
- `-c` / `--bytes` — byte count
- `-m` / `--chars` — character (rune) count

```sh
$ go run main.go -l test.txt
 3 test.txt
$ go run main.go -lw test.txt
 3 9 test.txt
$ go run main.go -m test.txt
 67 test.txt
```

`-c` and `-m` only differ when the file has multi-byte text. A file containing `café\n` is 6 bytes but 5 characters, because `é` is two bytes and one rune:

```sh
$ go run main.go -c cafe.txt
 6 cafe.txt
$ go run main.go -m cafe.txt
 5 cafe.txt
```

Unknown short flags get the same treatment as GNU:

```sh
$ go run main.go -f test.txt
wc: invalid option -- 'f'
```

## How it counts

- Lines are counted per `bufio.Scanner` iteration.
- Words come from `strings.Fields`, which splits on runs of whitespace — so `"double  space"` in `test.txt` doesn't need special handling.
- Bytes are `len(line) + 1` (`+1` is the newline).
- Characters are `utf8.RuneCountInString(line) + 1`.
- Bundled flags like `-lw` are expanded into `-l -w` by `expandFlags` before the `flag` package sees them.

## Building

```sh
go build -o wc .
./wc test.txt
```

Requires Go 1.27.1+ (whatever's pinned in `go.mod`).

## Known problems

Honest list of things I know are off and want to fix:

- **Assumes every line ends in a newline.** The `+1` per line means a file without a trailing `\n` reports one extra count.
- **Errors are panics.** `check()` panics on failure, so a typo'd filename dumps a stack trace instead of a clean error message.
- **One file at a time.** No multiple files, no stdin, no `total` row.
- **Flags must come before the filename.** The `flag` package stops parsing at the first positional argument, so `wc test.txt -l` silently ignores the `-l`.

## To do

- [x] Implement `-l`, `-w`, `-c`, `-m` (and allow combining them, e.g. `-lw`)
- [ ] Accept multiple files and print a `total` row
- [ ] Read from stdin when no file is given
- [x] Count characters as runes instead of bytes
- [ ] Replace panics with proper error messages and exit codes
- [ ] Handle files without a trailing newline
- [ ] Accept flags after the filename
- [ ] Tests

`test.txt` is the sample file I've been running it against.

## License

MIT — see [LICENSE](LICENSE). GNU's `wc` is GPLv3, but this is a fresh implementation that borrows no coreutils code, so it didn't have to be.
