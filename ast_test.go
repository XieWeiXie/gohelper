package gohelper

import (
	"fmt"
	"go/parser"
	"testing"
)

func TestNewAstFile(t *testing.T) {
	ast := NewAstFile("ast.go", parser.ParseComments)
	if ast != nil {
		fmt.Println(ast.GetComment())
	}

	ast1 := NewAstFile("ast.go", parser.PackageClauseOnly)
	if ast1 != nil {
		fmt.Println(ast1.GetPackages())
	}

	ast2 := NewAstFile("ast.go", parser.ImportsOnly)
	if ast2 != nil {
		fmt.Println(ast2.GetImports())
	}
}
