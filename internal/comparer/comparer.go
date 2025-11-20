package comparer

type Drift struct {
	MissingInDocs []string
	MissingInCode []string
}

func Compare(codeSymbols, docSymbols []string) Drift {
	drift := Drift{}

	return drift
}