package complexity

import "go/ast"

func Count(funcNode ast.Node) int {
	count := 1
	ast.Inspect(funcNode, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt:
			count++
		}
		return true
	})
	return count
}
