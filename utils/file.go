package utils

import (
	"fmt"
	"strings"
)

func FindFunctionCalls(lines []string, funcNames []string) map[string]bool {
	called := make(map[string]bool)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

	}

	return called
}

func GetUnusedFunctions(defined map[string]bool, called map[string]bool) []string {
	unused := []string{}

	return unused
}

