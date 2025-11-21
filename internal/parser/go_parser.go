package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
)

type CodeSymbol struct {
	Name 	string
	Kind 	string
}

// Traverse Go files, extract functions, structs, constants, and etc.
// Return a map where keys are symbol names and values indicate presence.

func ParseGoCode(path string) []CodeSymbol {
	fmt.Printf("Parsing Go code at: %s\n", path)
	symbols := []CodeSymbol{}

	// Walk through the folder
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		// Skip Directories
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		// Check if file extension is .go
		if filepath.Ext(d.Name()) == ".go" {
			fileSymbols := ParseGoCodeFile(p)
			symbols = append(symbols, fileSymbols...)
		}

		return nil
	})

	for _, sym := range symbols {
		fmt.Println("Found symbol: ", sym.Name, "=", sym.Kind)
	}

	return symbols
}

func ParseGoCodeFile(filePath string) []CodeSymbol {
	symbols := []CodeSymbol{}


	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return symbols
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if fd, ok := n.(*ast.FuncDecl); ok {
			symbols = append(symbols, CodeSymbol{
				Name: fd.Name.Name,
				Kind: "func",
			})
		}

		return true
	})
	return symbols
}

