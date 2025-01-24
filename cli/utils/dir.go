package utils

import (
	"fmt"
	"slices"
	"strings"
)

var operationTypes = []string{"d", "md", "mc", "mcd", "l"}

func GetOperationType(fileName string) (string, error) {
	rawType := strings.Split(fileName, "-")[0]
	if !slices.Contains(operationTypes, rawType) {
		return "", fmt.Errorf("invalid operation type")
	}
	return rawType, nil
}

func WindowsMountIssue(inputString string) string {
	firstIndex := (strings.Index(inputString, ":")) + 1
	beforeMount := inputString[:firstIndex]
	afterMount := inputString[firstIndex:]
	afterMount = strings.ReplaceAll(afterMount, ":", "_")
	return beforeMount + afterMount
}
