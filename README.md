# go-iec62056

This library is intended to provide a robust parser for IEC1107/IEC62056
telegrams written in pure Go. It was born out of my
[iec62056 library for Python](https://github.com/jonkerj/iec62056), but I 
decided to rewrite it in Go.

My philosophy is to actually parse the incoming data based on a formal grammar,
as opposed to dissecting it with regexes.

In its core, it uses [participle](https://github.com/alecthomas/participle).

# How to use it

For now, read the source. I promise to write docs at some point.

```bash
$ go test -v ./pkg/ast
$ go run main.go
```