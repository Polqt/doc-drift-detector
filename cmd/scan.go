package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var srcPath string
var docsPath string

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&srcPath, "src", "s", "", "Path to the source code directory")
	scanCmd.Flags().StringVarP(&docsPath, "docs", "d", "", "Path to the documentation directory")

	scanCmd.MarkFlagRequired("src")
	scanCmd.MarkFlagRequired("docs")
}

var scanCmd = &cobra.Command{
	Use: "scan",
	Short: "Scan code and documentation for drift.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scanning...")
		fmt.Println("Source: ", srcPath)
		fmt.Println("Documentation: ", docsPath)
	},

	// codeSymbols := parser.ParseGoCode("./test")
	// docSymbols := parser.ParseDocs("./docs")

	// codeNames := []string{}
	// for _, sym := range codeSymbols {
	// 	codeNames = append(codeNames, sym.Name)
	// }

	// docNames := []string{}
	// for _, sym := range docSymbols {
	// 	docNames = append(docNames, sym.Name)
	// }

	// drift := comparer.Compare(codeNames, docNames)

	// report.GenerateReport(drift, "report.json")
	// report.GenerateMarkdownReport(drift, "report.md")

	// fmt.Println("Done.")
	// os.Exit(0)
}