package analyzer

import (
	"io/fs"
	"path/filepath"

	"github.com/Polqt/doc-drift-detector/internal/parser"
)

type UnifiedSymbol struct {
	Name 	string
	Kind 	string
}

type AnalyzerResult struct {
	CodeSymbols 	[]UnifiedSymbol
	DocSymbols  	[]UnifiedSymbol
}

func Analyze(srcPath string, docsPath string) AnalyzerResult {
	codeSymbols := []UnifiedSymbol{}
	docSymbols := []UnifiedSymbol{}


	goSymbols := parser.ParseGoCode(srcPath)
	for _, sym := range goSymbols {
		codeSymbols = append(codeSymbols, UnifiedSymbol{
			Name: sym.Name,
			Kind: sym.Kind,
		})
	}

	var tsSymbols []parser.TSSymbol

	for _, sym := range tsSymbols {
		codeSymbols = append(codeSymbols, UnifiedSymbol{
			Name: sym.Name,
			Kind: sym.Kind,
		})
	}

	filepath.Walk(docsPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".md" {
			headings := parser.ParseDocs(path)
			for _, h := range headings {
				docSymbols = append(docSymbols, UnifiedSymbol{
					Name: Normalize(h),
					Kind: "doc-heading",
				})
			}
		}

		return nil
	})

	
	return AnalyzerResult{
		CodeSymbols: codeSymbols,
		DocSymbols: docSymbols,
	}
}