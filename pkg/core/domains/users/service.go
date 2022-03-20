package users

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"uploader/constants"
	"uploader/helpers/arrayHelp"
	"uploader/helpers/stringHelp"
	"uploader/helpers/validatorHelp/emailValidator"
)

type userRepo struct {
	repo *UserRepository
}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func validateFieldsAlreadyRegistered(newUser ConfigurationHeaderExport, user []string) bool {
	return newUser.Email == user[0] || newUser.Identifier == user[1]
}

func (repo *userRepo) FindAll() ([]ConfigurationHeaderExport, error) {
	fmt.Println("FindAll")

	return []ConfigurationHeaderExport{}, nil
}

func (repo *userRepo) Save(usersList []ConfigurationHeaderExport) error {
	fmt.Println("Save", usersList)

	return nil
}

func FormatExport(row []string, header []string) ConfigurationHeaderExport {
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
	user := ConfigurationHeaderExport{
		Name:       name,
		Email:      email,
		Salary:     salary,
		Identifier: identifier,
		Phone:      phone,
		Mobile:     mobile,
	}
	return user
}

func (repo *userRepo) GetField(v ConfigurationHeader, field string) ([]string, error) {
	if !arrayHelp.Contains(constants.HEADERCONFIGURATION, field) {
		return []string{}, fmt.Errorf("field not configured")
	}
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string), nil
}

func (repo *userRepo) ConvertDataToHeaderExport(data [][]string, header []string) ([]ConfigurationHeaderExport, error) {
	var users []ConfigurationHeaderExport
	for _, row := range data {
		newUser := FormatExport(row, header)
		users = append(users, newUser)
	}
	return users, nil
}

func (repo *userRepo) FormatHeader(header []string, configHeader ConfigurationHeader, sizeStructHeader int) []string {
	var parseHeader []string
	for i := 0; i < len(header); i++ {
		word, _ := secondaryStringHelp.StandardizeColumn(header[i])
		for j := 0; j < sizeStructHeader; j++ {
			val := reflect.Indirect(reflect.ValueOf(configHeader))
			nameColumn := val.Type().Field(j).Name
			field, _ := repo.GetField(configHeader, nameColumn)
			contain := secondaryStringHelp.StringInSlice(word, field)
			if contain {
				parseHeader = append(parseHeader, nameColumn)
				break
			}
		}
	}
	return parseHeader
}

func (repo *userRepo) CheckUserValid(user ConfigurationHeaderExport, users []ConfigurationHeaderExport, oldValues [][]string) (bool, error) {
	valid := true
	if user.Email == "" || user.Identifier == "" || !secondaryEmailValidator.ValidateEmail(user.Email) {
		valid = false
	}
	if _, err := os.Stat("./" + constants.SUCCESSPATHNAME); err == nil {
		if valid {
			for _, record := range oldValues {
				if validateFieldsAlreadyRegistered(user, []string{record[constants.POSITIONHEADEREMAIL], record[constants.POSITIONHEADERIDENTIFIER]}) {
					valid = false
					break
				}
			}
		}
	}

	if valid {
		for _, record := range users {
			if validateFieldsAlreadyRegistered(user, []string{record.Email, record.Identifier}) {
				valid = false
				break
			}
		}
	}
	return valid, nil
}
