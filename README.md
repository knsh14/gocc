gocc
---

gocc is analysis tool to check cyclomatic complexity of go codes.
this package provides Analyzer for [analysis]( https://godoc.org/golang.org/x/tools/go/analysis ).
you can use this analyzer in your analysis tool.

# Install

```
$ go get -u github.com/knsh14/gocc/cmd/gocc
```

# Usage
```
$ gocc [-flag] [package]
```

## example
```
$ gocc -over 10 ./src
```
