package report

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Polqt/doc-drift-detector/internal/comparer"
)

func GenerateReport(drift comparer.Drift, path string) {
	// Convert Drift struct to a JSON
	data, err := json.MarshalIndent(drift, "", " ")
	if err != nil {
		fmt.Println("Error generating JSON: ", err)
		return
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

	fmt.Println("JSON report generated at: ", path)
}

func GenerateMarkdownReport(drift comparer.Drift, path string) {
	content := "# Documentation Drift Report\n\n"
	
	content += "## Missing in Documentation\n"
	if len(drift.MissingInDocs) == 0 {
		content += "None\n"
	} else {
		for _, item := range drift.MissingInDocs {
			content += "- " + item + "\n"
		}
	}

	content += "\n## Missing in Code\n"
	if len(drift.MissingInCode) == 0 {
		content += "None\n"
	} else {
		for _, item := range drift.MissingInCode {
			content += "- " + item + "\n"
		}
	}

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing Markdown file: ", err)
		return
	}

	fmt.Println("Markdown report generated at: ", path)
}