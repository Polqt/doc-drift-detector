package main

import (
	"fmt"
	"os"

	"github.com/Polqt/doc-drift-detector/internal/parser"
	"github.com/Polqt/doc-drift-detector/internal/report"
)

func main() {
	symbols := parser.ParseGoCode("./test")
	docSymbols := parser.ParseDocs("./docs")
	report.GenerateReport(drift, "report.json")
	report.GenerateMarkdownReport(drift, "report.md")
		

	for _, sym := range docSymbols {
		fmt.Println("Doc symbol: ", sym.Name, "-", sym.Kind)
	}
	fmt.Println(symbols)

	os.Exit(0)
}