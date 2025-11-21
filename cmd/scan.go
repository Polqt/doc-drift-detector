package cmd

import (
	"fmt"

	"github.com/Polqt/doc-drift-detector/internal/analyzer"
	"github.com/Polqt/doc-drift-detector/internal/comparer"
	"github.com/Polqt/doc-drift-detector/internal/report"
	"github.com/spf13/cobra"
)

var (
	srcPath 	string
	docsPath 	string
	jsonPath 	string
	mdPath		string
)

var scanCmd = &cobra.Command{
	Use: "scan",
	Short: "Scan code and documentation for drift.",
	Run: func(cmd *cobra.Command, args []string) {
		if srcPath == "" || docsPath == "" {
			fmt.Println("Error: --src and --docs are required.")
			return
		}

		fmt.Println("Scanning...")
		results := analyzer.Analyze(srcPath, docsPath)

		codeNames := []string{}
		for _, sym := range results.CodeSymbols {
			codeNames = append(codeNames, sym.Name)
		}

		docNames := []string{}
		for _, sym := range results.DocSymbols {
			docNames = append(docNames, sym.Name)
		}

		missing := comparer.Compare(codeNames, docNames)

		fmt.Println("Code symbols: ", results.CodeSymbols)
		fmt.Println("Doc symbols: ", results.DocSymbols)

		fmt.Println("Missing in Docs: ", missing.MissingInDocs)
		fmt.Println("Missing in Code: ", missing.MissingInCode)

		if jsonPath != "" {
			report.GenerateReport(missing, jsonPath)
		}

		if mdPath != "" {
			report.GenerateMarkdownReport(missing, mdPath)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&srcPath, "src", "s", "", "Path to source code directory")
	scanCmd.Flags().StringVarP(&docsPath, "docs", "d", "", "Path to documentation directory")

	scanCmd.Flags().StringVarP(&jsonPath, "json", "j", "", "Path to output JSON report")
	scanCmd.Flags().StringVarP(&mdPath, "md", "m", "", "Path to output Markdown report")

	scanCmd.MarkFlagRequired("src")
	scanCmd.MarkFlagRequired("docs")

	rootCmd.AddCommand(scanCmd)
}