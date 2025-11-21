package comparer

type Drift struct {
	MissingInDocs []string
	MissingInCode []string
}

func Compare(codeSymbols, docSymbols []string) Drift {
	drift := Drift{
		MissingInDocs: []string{},
		MissingInCode: []string{},
	}

	docMap := map[string]bool{}
	for _, docSym := range docSymbols {
		docMap[docSym] = true
	}

	codeMap := map[string]bool{}
	for _, codeSym := range codeSymbols {
		codeMap[codeSym] = true
	}

	// Code -> Docs
	for _, code := range codeSymbols {
		if !docMap[code] {
			drift.MissingInDocs = append(drift.MissingInDocs, code)
		}
	}

	// Docs -> Code
	for _, doc := range docSymbols {
		if !codeMap[doc] {
			drift.MissingInCode = append(drift.MissingInCode, doc)
		}
	}

	return drift
}