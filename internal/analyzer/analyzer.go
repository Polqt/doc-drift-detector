package analyzer

import (
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
	goSymbols := parser.ParseGoCode(srcPath)

	var tsSymbols []parser.TSSymbol
	
	var docHeadings []string

	unifiedCode := []UnifiedSymbol{}
	unifiedDocs := []UnifiedSymbol{}

	for _, sym := range goSymbols {
		unifiedCode = append(unifiedCode, UnifiedSymbol{
			Name: sym.Name,
			Kind: sym.Kind,
		})
	}

	for _, sym := range tsSymbols {
		unifiedCode = append(unifiedCode, UnifiedSymbol{
			Name: sym.Name,
			Kind: sym.Kind,
		})
	}

	for _, head := range docHeadings {
		unifiedDocs = append(unifiedDocs, UnifiedSymbol{
			Name: head,
			Kind: "doc-heading",
		})
	}

	return AnalyzerResult{
		CodeSymbols: unifiedCode,
		DocSymbols: unifiedDocs,
	}
}