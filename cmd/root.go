package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "doc-drift-detector",
	Short: "Detects documentation drift between code and its documentation.",
	Long: "A tool to help developers identify and manage documentation drift in their projects. Moreover, it scans Go/TS codebases and compares them to documentation to detect missing or outdated docs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `docdrift scan` to check documentation drift.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}