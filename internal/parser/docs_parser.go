package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseDocs(path string) []CodeSymbol {
	symbols := []CodeSymbol{}

	fmt.Println("Parsing docs from: ", path)

	file, err := os.Open(path + "/example.md")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return symbols
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Detect if line is a header 
		if strings.HasPrefix(line, "##") {
			headerName := strings.TrimSpace(line[3:])
			symbols = append(symbols, CodeSymbol{
				Name: headerName,
				Kind: "doc",
			})
		}
	}

	return symbols
}