package gocc

import (
	"go/ast"

	"github.com/knsh14/gocc/complexity"
	"golang.org/x/tools/go/analysis"
)

var (
	threshold int
)

// Analyzer does check about cyclomatic complexity
var Analyzer = &analysis.Analyzer{
	Name:     "gocc",
	Doc:      "checks cyclomatic complexity",
	Run:      run,
	Requires: []*analysis.Analyzer{complexity.Analyzer},
}

func init() {
	Analyzer.Flags.IntVar(&threshold, "over", 10, "minimum complexity to report")
}

func run(pass *analysis.Pass) (interface{}, error) {
	complexities := pass.ResultOf[complexity.Analyzer].(map[ast.Node]int)
	for n, count := range complexities {
		if count >= threshold {
			fd := n.(*ast.FuncDecl)
			name := getFuncName(fd)
			pass.Reportf(n.Pos(), "func %s complexity=%d", name, count)
		}
	}
	return nil, nil
}
