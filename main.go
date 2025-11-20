package main

import (
	"fmt"
	"os"

	"github.com/Polqt/doc-drift-detector/internal/parser"

)

func main() {
	symbols := parser.ParseGoCode("./test")
	docSymbols := parser.ParseDocs("./docs")
	for _, sym := range docSymbols {
		fmt.Println("Doc symbol: ", sym.Name, "-", sym.Kind)
	}
	fmt.Println(symbols)

	os.Exit(0)
}