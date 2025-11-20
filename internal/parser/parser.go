package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

type CodeSymbol struct {
	Name 		string
	Kind 		string
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

		// Check if file extension is .go
		if filepath.Ext(d.Name()) != ".go" {
			return nil
		}

		// Call function to parse this file
		if !d.IsDir() && strings.HasSuffix(p, ".go") {
			parseGoFile(p, &symbols)
		}

		return nil
	})

	for _, sym := range symbols {
		fmt.Println("Found symbol: ", sym.Name, "=", sym.Kind)
	}

	return symbols
}

func parseGoFile(filePath string, result *[]CodeSymbol) {
	fmt.Println("Parsing Go file: ", filePath)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		fmt.Println("Error parsing: ", filePath, err)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if fd, ok := n.(*ast.FuncDecl); ok {
			*result = append(*result, CodeSymbol{
				Name: fd.Name.Name,
				Kind: "func",
			})
		}

		return true
	})
}

