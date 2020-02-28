package gohelper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// AstFile struct
type AstFile struct {
	path string
	file *ast.File
}

// NewAstFile ...
func NewAstFile(path string, mode parser.Mode) *AstFile {
	f, e := os.Open(path)
	if e != nil {
		return nil
	}
	set := token.NewFileSet()
	fs, e := parser.ParseFile(set, f.Name(), nil, mode)
	if e != nil {
		return nil
	}
	return &AstFile{
		path: path,
		file: fs,
	}

}

// GetComment...
func (A *AstFile) GetComment() []string {
	if len(A.file.Comments) == 0 {
		return nil
	}
	var results []string
	for _, i := range A.file.Comments {
		results = append(results, strings.Trim(i.Text(), "\n"))

	}
	return results
}

func (A *AstFile) GetPackages() []string {
	var results []string
	results = append(results, A.file.Name.String())
	return results
}

func (A *AstFile) GetImports() []string {
	if len(A.file.Imports) == 0 {
		return nil
	}
	var results []string
	for _, i := range A.file.Imports {
		results = append(results, i.Path.Value)
	}
	return results
}
