package gocc

import (
	"bytes"
	"go/ast"
)

func getFuncName(node *ast.FuncDecl) string {
	var buf bytes.Buffer
	if node.Recv != nil {
		buf.WriteString("(")
		for _, r := range node.Recv.List {
			if len(r.Names) > 0 {
				buf.WriteString(r.Names[0].Name)
				buf.WriteString(" ")
			}
			if t, ok := r.Type.(*ast.StarExpr); ok {
				buf.WriteString("*")
				if ident, ok := t.X.(*ast.Ident); ok {
					buf.WriteString(ident.Name)
				}
			}
			if t, ok := r.Type.(*ast.Ident); ok {
				buf.WriteString(t.Name)
			}
		}
		buf.WriteString(") ")
	}
	buf.WriteString(node.Name.Name)
	return buf.String()
}
