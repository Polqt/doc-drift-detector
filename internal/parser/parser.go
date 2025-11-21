package parser

import (
	"io/fs"
	"path/filepath"
)

func ParseCode(path string) []CodeSymbol {
	symbols := []CodeSymbol{}

	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(d.Name())

		if ext == ".go" {
			goSym := ParseGoCodeFile(p)
			symbols = append(symbols, goSym...)
		}

		if ext == ".ts" || ext == ".tsx" {
			tsSym := ParseTsFile(p)
			for _, ts := range tsSym {
				symbols = append(symbols, CodeSymbol{
					Name: ts.Name,
					Kind: ts.Kind,
				})
			}
		}

		return nil
	})
	return symbols
}