package utils

import (
	"fmt"
	"strings"
)

func GetOperationType(fileName string) (string, error) {
	rawType := strings.Split(fileName, "-")[0]
	if rawType != "d" && rawType != "md" && rawType != "mc" {
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
