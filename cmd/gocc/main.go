package main

import (
	"github.com/knsh14/gocc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gocc.Analyzer)
}
