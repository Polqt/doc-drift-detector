package parser

import (
	"bufio"
	"os"
	"strings"
)

func ParseDocs(path string) []string {
	headings := []string{}

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Detect if line is a header 
		if strings.HasPrefix(line, "# ") {
			headings = append(headings, strings.TrimPrefix(line, "# "))
		}

		if strings.HasPrefix(line, "## ") {
			headings = append(headings, strings.TrimPrefix(line, "## "))
		}
	}

	return headings
}