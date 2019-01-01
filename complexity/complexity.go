package complexity

import (
	"go/ast"
	"go/token"
)

func Count(funcNode ast.Node) int {
	count := 1
	ast.Inspect(funcNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.IfStmt:
			count++
		case *ast.ForStmt:
			count++
		case *ast.CaseClause:
			if n.List == nil {
				count++
				break
			}
			count += len(n.List)
		case *ast.BinaryExpr:
			if n.Op == token.LAND || n.Op == token.LOR {
				count++
			}
		}
		return true
	})
	return count
}
