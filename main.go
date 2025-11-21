package main

import (
	"fmt"
	"os"

	"github.com/Polqt/doc-drift-detector/internal/parser"
	"github.com/Polqt/doc-drift-detector/internal/report"
	"github.com/Polqt/doc-drift-detector/internal/comparer"
)

func main() {
	codeSymbols := parser.ParseGoCode("./test")
	docSymbols := parser.ParseDocs("./docs")

	codeNames := []string{}
	for _, sym := range codeSymbols {
		codeNames = append(codeNames, sym.Name)
	}

	docNames := []string{}
	for _, sym := range docSymbols {
		docNames = append(docNames, sym.Name)
	}

	drift := comparer.Compare(codeNames, docNames)

	report.GenerateReport(drift, "report.json")
	report.GenerateMarkdownReport(drift, "report.md")

	fmt.Println("Done.")
	os.Exit(0)
}