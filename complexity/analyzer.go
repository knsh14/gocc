package complexity

import (
	"go/ast"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:       "complexity",
	Doc:        "pass result of complexity",
	ResultType: reflect.TypeOf(map[ast.Node]int{}),
	Requires:   []*analysis.Analyzer{inspect.Analyzer},
	Run:        run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	res := map[ast.Node]int{}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	var err error
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		count := Count(n)
		if _, hasKey := res[n]; hasKey {
			err = errors.New("Node is duplicate")
			return
		}
		res[n] = count
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to calc complexity")
	}
	return res, nil
}
