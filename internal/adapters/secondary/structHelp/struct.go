package secondaryStructHelp

import (
	"fmt"
	"reflect"
	"uploader/constants"
	secondaryArrayHelp "uploader/internal/adapters/secondary/arrayHelp"
	"uploader/pkg/domains/users"
)

func GetField(v users.ConfigurationHeader, field string) ([]string, error) {
	if !secondaryArrayHelp.Contains(constants.HEADERCONFIGURATION, field) {
		return []string{}, fmt.Errorf("field not configured")
	}
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string), nil
}
