package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TSSymbol struct {
	Name 	string
	Kind 	string
}

func ParseTsFile(path string) []TSSymbol {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening TS File: ", err)
		return nil
	}

	defer file.Close()

	symbols := []TSSymbol{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Detect function patterns
		if strings.HasPrefix(line, "function ") {
			parts := strings.Split(strings.TrimPrefix(line, "function "), "(")[0]
			symbols = append(symbols, TSSymbol{
				Name: parts,
				Kind: "function",
			})
		} 

		if strings.HasPrefix(line, "export function ") {
			parts := strings.Split(strings.TrimPrefix(line, "export function "), "(")[0]
			symbols = append(symbols, TSSymbol{
				Name: parts,
				Kind: "export function",
			})
		}

		if strings.HasPrefix(line, "const ") && strings.Contains(line, "=") {
			beforeEqual := strings.Split(line, "=")[0]
			parts := strings.Split(strings.TrimPrefix(beforeEqual, "const "), " ")
			name := parts[0]

			symbols = append(symbols, TSSymbol{
				Name: name,
				Kind: "arrow function",
			})
		}

		// Detect interfaces patterns
		if strings.HasPrefix(line, "interface ") {
			parts := strings.Split(strings.TrimPrefix(line, "interface "), " ")[0]
			symbols = append(symbols, TSSymbol{
				Name: parts,
				Kind: "interface",
			})
		}

		// Detect class patterns
		if strings.HasPrefix(line, "class ") {
			parts := strings.Split(strings.TrimPrefix(line, "class "), " ")[0]
			symbols = append(symbols, TSSymbol{
				Name: parts,
				Kind: "class",
			})
		}

		if strings.HasSuffix(p, ".ts") {
			tsSymbols := ParseTsFile(p)
			
			for _, sym := range tsSymbols {
				symbols = append(symbols, TSSymbol{
					Name: sym.Name,
					Kind: sym.Kind,
				})
			}
		}
	}

	return symbols

}