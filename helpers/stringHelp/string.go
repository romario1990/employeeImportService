package secondaryStringHelp

import (
	"regexp"
	"strings"
)

func StringInSlice(columnName string, list []string) bool {
	for _, configColumnName := range list {
		if configColumnName == columnName {
			return true
		}
	}
	return false
}

func RemoveSpecialChar(word string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(word, "")
	return processedString, nil
}

func StandardizeColumn(column string) (string, error) {
	word, err := RemoveSpecialChar(column)
	if err != nil {
		return "", err
	}
	return strings.ToLower(strings.TrimSpace(word)), nil
}

func RemoveIndex(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}
