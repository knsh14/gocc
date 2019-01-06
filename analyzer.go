package gocc

import (
	"go/ast"

	"github.com/knsh14/gocc/complexity"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	threshold int
)

// Analyzer does check about cyclomatic complexity
var Analyzer = &analysis.Analyzer{
	Name:     "gocc",
	Doc:      "checks cyclomatic complexity",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func init() {
	Analyzer.Flags.IntVar(&threshold, "over", 10, "minimum complexity to report")
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		count := complexity.Count(n)
		if count >= threshold {
			fd := n.(*ast.FuncDecl)
			name := getFuncName(fd)
			pass.Reportf(n.Pos(), "func %s complexity=%d", name, count)
		}
	})
	return nil, nil
}
