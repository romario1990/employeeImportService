package formatService

import (
	"fmt"
	"reflect"
	"strings"
	"uploader/constants"
	"uploader/src/entities"
	"uploader/src/services/help"
)

func FormatHeader(header []string, configHeader entities.ConfigurationHeader, sizeStructHeader int) []string {
	var parseHeader []string
	for i := 0; i < len(header); i++ {
		word, _ := help.StandardizeColumn(header[i])
		for j := 0; j < sizeStructHeader; j++ {
			val := reflect.Indirect(reflect.ValueOf(configHeader))
			nameColumn := val.Type().Field(j).Name
			field, _ := help.GetField(configHeader, nameColumn)
			contain := help.StringInSlice(word, field)
			if contain {
				parseHeader = append(parseHeader, nameColumn)
				break
			}
		}
	}
	return parseHeader
}

func FormatCSVExport(row []string, header []string) entities.ConfigurationHeaderExport {
	fullName := ""
	firstName := ""
	middleName := ""
	lastName := ""
	email := ""
	salary := ""
	identifier := ""
	phone := ""
	mobile := ""
	for i := 0; i < len(row); i++ {
		switch header[i] {
		case constants.FULLNAME:
			fullName = row[i]
		case constants.FIRSTNAME:
			firstName = row[i]
		case constants.MIDDLENAME:
			middleName = " " + row[i]
		case constants.LASTNAME:
			lastName = " " + row[i]
		case constants.EMAIL:
			email = strings.Trim(strings.ToLower(row[i]), " ")
		case constants.SALARY:
			salary = row[i]
		case constants.IDENTIFIER:
			identifier = row[i]
		case constants.PHONE:
			phone = row[i]
		case constants.MOBILE:
			mobile = row[i]
		default:
			fmt.Printf("Type column not found")
		}
	}
	name := fullName
	if name == constants.INVALIDNAME {
		name = strings.Join([]string{firstName, middleName, lastName}, "")
	}
	user := entities.ConfigurationHeaderExport{
		Name:       name,
		Email:      email,
		Salary:     salary,
		Identifier: identifier,
		Phone:      phone,
		Mobile:     mobile,
	}
	return user
}
