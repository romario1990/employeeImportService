package help

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"uploader/constants"
	"uploader/src/entities"
)

func StringInSlice(columnName string, list []string) bool {
	for _, configColumnName := range list {
		if configColumnName == columnName {
			return true
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func GetField(v entities.ConfigurationHeader, field string) ([]string, error) {
	if !contains(constants.HEADERCONFIGURATION, field) {
		return []string{}, fmt.Errorf("field not configured")
	}
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string), nil
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
