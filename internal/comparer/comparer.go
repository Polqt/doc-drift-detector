package comparer

type Drift struct {
	MissingInDocs []string
	MissingInCode []string
}

func Compare(codeSymbols, docSymbols []string) Drift {
	drift := Drift{}

	docMap := map[string]bool{}
	for _, s := range docSymbols {
		docMap[s.Name] = true 
	}
	

	return drift
}