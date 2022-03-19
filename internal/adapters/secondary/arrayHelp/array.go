package secondaryArrayHelp

import (
	"fmt"
	"reflect"
	"uploader/constants"
	"uploader/entities"
)

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func GetField(v entities.ConfigurationHeader, field string) ([]string, error) {
	if !Contains(constants.HEADERCONFIGURATION, field) {
		return []string{}, fmt.Errorf("field not configured")
	}
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string), nil
}
