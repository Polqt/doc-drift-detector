package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Polqt/doc-drift-detector/internal/parser"
	"github.com/Polqt/doc-drift-detector/internal/comparer"
	"github.com/Polqt/doc-drift-detector/internal/report"

)

func main() {
	codePath := flag.String("code", "./src", "Path to the source code")
	docsPath := flag.String("docs", "./docs", "Path to the markdown documentation")
	outPath := flag.String("output", "./report", "Path to the report")
	flag.Parse();

	fmt.Println("Scanning code: " *codePath)
	fmt.Println("Scanning documentation: " *docsPath)

	codeSymbols := parser.ParseGoCode(*codePath)


	docSymbols := parser.ParseDocs(*docsPath)

	drift := comparer.Compare(codeSymbols, docSymbols)
	
	report.GenerateReport(drift, *outPath)

	if len(drift.MissingInDocs) > 0 || len(drift.MissingInCode) > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}