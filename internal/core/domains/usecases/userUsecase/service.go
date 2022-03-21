package userUsecase

import (
	"fmt"
	"reflect"
	"strings"
	"uploader/constants"
	"uploader/helpers/arrayHelp"
	"uploader/helpers/stringHelp"
	"uploader/helpers/validatorHelp/emailValidator"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/core/domains/repositories/userRepo"
	"uploader/internal/repositories/ports"
)

type userService struct {
	useCases        *UserService
	userSuccessRepo userRepo.UserRepository
	userErrorRepo   userRepo.UserRepository
	filePending     ports.FileRepository
	configHeader    userModel.ConfigurationHeader
	hasH            bool
}

func NewUserService(filePending ports.FileRepository, userSuccessRepo userRepo.UserRepository,
	userErrorRepo userRepo.UserRepository, configHeader userModel.ConfigurationHeader, hasH bool) (UserService, error) {
	return &userService{nil, userSuccessRepo, userErrorRepo,
		filePending, configHeader, hasH}, nil
}

func validateFieldsAlreadyRegistered(newUser userModel.ConfigurationHeaderExport, user []string) bool {
	return newUser.Email == user[0] || newUser.Identifier == user[1]
}

func FormatExport(row []string, header []string) userModel.ConfigurationHeaderExport {
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
	user := userModel.ConfigurationHeaderExport{
		Name:       name,
		Email:      email,
		Salary:     salary,
		Identifier: identifier,
		Phone:      phone,
		Mobile:     mobile,
	}
	return user
}

func (userService *userService) Exec() error {
	var usersValid, userInvalid []userModel.ConfigurationHeaderExport
	fmt.Printf("### Header de configuration file = %v\n", userService.configHeader)
	data, err := userService.filePending.GetData()
	if err != nil {
		return fmt.Errorf("error reading existing values")
	}
	var header []string
	if userService.hasH {
		header = userService.FormatHeader(data[0], userService.configHeader, 9)
		data = secondaryStringHelp.RemoveIndex(data, 0)
	}
	users, err := userService.ConvertDataToHeaderExport(data, header)
	for _, newUser := range users {
		userValid := userService.CheckUserValid(newUser, usersValid)
		if userValid {
			usersValid = append(usersValid, newUser)
		} else {
			userInvalid = append(userInvalid, newUser)
		}
	}
	err = userService.SaveUsers(usersValid, true)
	if err != nil {
		userService.userErrorRepo.MoveFileProcessed()
		return err
	}
	err = userService.SaveUsers(userInvalid, false)
	if err != nil {
		userService.userErrorRepo.MoveFileProcessed()
		return err
	}
	userService.filePending.MoveFileProcessed()
	//if err != nil {
	//	return err
	//}
	fmt.Println("---------------------------- End of Processing Information ----------------------------")
	fmt.Println("---------------------------- Start of summary ----------------------------")
	fmt.Println("Valid userRepos: ", usersValid)
	fmt.Println("Invalid userRepos: ", userInvalid)
	fmt.Println("Success file is located: ", "./"+constants.SUCCESSPATHNAME)
	fmt.Println("Fault file is located: ", "./"+constants.ERRORPATHNAME)
	fmt.Println("Processed files are located: ", constants.PATHPROCESSED)
	fmt.Println("Processed with error file is located: ", constants.PATHPROCESSEDERROR)
	fmt.Println("---------------------------- End of summary ----------------------------")
	fmt.Println("############################ FINISHED ############################")
	return nil
}

func (userService *userService) GetField(v userModel.ConfigurationHeader, field string) ([]string, error) {
	if !arrayHelp.Contains(constants.HEADERCONFIGURATION, field) {
		return []string{}, fmt.Errorf("field not configured")
	}
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string), nil
}

func (userService *userService) ConvertDataToHeaderExport(data [][]string, header []string) ([]userModel.ConfigurationHeaderExport, error) {
	var users []userModel.ConfigurationHeaderExport
	for _, row := range data {
		newUser := FormatExport(row, header)
		users = append(users, newUser)
	}
	return users, nil
}

func (userService *userService) FormatHeader(header []string, configHeader userModel.ConfigurationHeader, sizeStructHeader int) []string {
	var parseHeader []string
	for i := 0; i < len(header); i++ {
		word, _ := secondaryStringHelp.StandardizeColumn(header[i])
		for j := 0; j < sizeStructHeader; j++ {
			val := reflect.Indirect(reflect.ValueOf(configHeader))
			nameColumn := val.Type().Field(j).Name
			field, _ := userService.GetField(configHeader, nameColumn)
			contain := secondaryStringHelp.StringInSlice(word, field)
			if contain {
				parseHeader = append(parseHeader, nameColumn)
				break
			}
		}
	}
	return parseHeader
}

func (userService *userService) CheckUserValid(user userModel.ConfigurationHeaderExport, users []userModel.ConfigurationHeaderExport) bool {
	valid := true
	if user.Email == "" || !secondaryEmailValidator.ValidateEmail(user.Email) {
		return false
	}
	if user.Identifier == "" {
		return false
	}
	oldValues, _ := userService.userSuccessRepo.GetList()
	if valid {
		for _, record := range oldValues {
			if len(record) > 0 {
				if validateFieldsAlreadyRegistered(user, []string{record[constants.POSITIONHEADEREMAIL], record[constants.POSITIONHEADERIDENTIFIER]}) {
					return false
				}
			}
		}
	}
	if valid {
		for _, record := range users {
			if validateFieldsAlreadyRegistered(user, []string{record.Email, record.Identifier}) {
				return false
			}
		}
	}
	return valid
}

func (userService *userService) SaveUsers(users []userModel.ConfigurationHeaderExport, userValid bool) (err error) {
	if userValid {
		userService.userSuccessRepo.Save(users)
	} else {
		userService.userErrorRepo.Save(users)
	}
	return nil
}
