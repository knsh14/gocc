package complexity

import (
	"go/ast"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

var testdata = analysistest.TestData()

func TestAnalyzerSimple(t *testing.T) {
	expects := map[string]int{
		"main": 1,
	}
	testdata := analysistest.TestData()
	res := analysistest.Run(t, testdata, Analyzer, "simple")[0]
	v, ok := res.Result.(map[ast.Node]int)
	if !ok {
		t.Fatalf("type of result must be map[ast.Node]int: %T", res.Result)
	}
	if len(v) != len(expects) {
		t.Fatalf("unexpected complexity num. want=%d, got=%d", len(expects), len(v))
	}

	for k := range v {
		switch n := k.(type) {
		case *ast.FuncDecl:
			expect, ok := expects[n.Name.Name]
			if !ok {
				t.Fatalf("%s is not in expected function list", n.Name.Name)
			}
			if v[k] != expect {
				t.Fatalf("unexpected complexity. want=%d, got=%d", expect, v[k])
			}
		}
	}
}
