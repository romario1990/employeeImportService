package secondaryFileProcessor

import (
	"fmt"
	"os"
	"strings"
	"uploader/config"
	"uploader/constants"
	secondaryFile "uploader/helpers/fileHelp"
	secondaryWriteFile "uploader/helpers/fileHelp/write/csv"
	"uploader/helpers/stringHelp"
	users2 "uploader/pkg/core/domains/users"
	file3 "uploader/pkg/core/repositories/file"
)

func Exec(filename string, hasH bool, fileType string) error {
	if strings.ToLower(fileType) != constants.CSV {
		return fmt.Errorf("unsupported file format %s", fileType)
	}
	fmt.Println("############################ STARTING ############################")
	var userRepo users2.UserRepository
	var fileRepo file3.FileRepository
	if fileType == constants.CSV {
		userRepo = users2.NewUserRepo()
		fileRepo = file3.NewFileCSVRepo()
		userRepo.FindAll()
		userRepo.Save([]users2.ConfigurationHeaderExport{})
	}

	fmt.Println("---------------------------- Start of processing information ----------------------------")
	file, err := secondaryFile.Read(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	secondaryFile.InitExec(nil)
	var usersValid, userInvalid []users2.ConfigurationHeaderExport
	confHeader, err := config.LoadConfigHeader()
	if err != nil {
		file.Close()
		secondaryFile.MoveFileProcessedError(filename, "")
		return fmt.Errorf("headerconfiguration.json file not configured")
	}
	configHeader := users2.ConfigurationHeader{
		FullName:   confHeader.GetStringSlice(constants.FULLNAME),
		FirstName:  confHeader.GetStringSlice(constants.FIRSTNAME),
		MiddleName: confHeader.GetStringSlice(constants.MIDDLENAME),
		LastName:   confHeader.GetStringSlice(constants.LASTNAME),
		Email:      confHeader.GetStringSlice(constants.EMAIL),
		Salary:     confHeader.GetStringSlice(constants.SALARY),
		Identifier: confHeader.GetStringSlice(constants.IDENTIFIER),
		Phone:      confHeader.GetStringSlice(constants.PHONE),
		Mobile:     confHeader.GetStringSlice(constants.MOBILE),
	}
	fmt.Printf("### The \"%v\" file is being processed\n", filename)
	fmt.Printf("### Header de configuration file = %v\n", configHeader)
	oldValues3, err := fileRepo.GetData(filename)
	if err != nil {
		return fmt.Errorf("error reading existing values")
	}
	var header []string
	if hasH {
		header = userRepo.FormatHeader(oldValues3[0], configHeader, 9)
		oldValues3 = secondaryStringHelp.RemoveIndex(oldValues3, 0)
	}
	users2, err := userRepo.ConvertDataToHeaderExport(oldValues3, header)
	var oldValues2 [][]string
	for _, newUser := range users2 {
		if _, err := os.Stat("./" + constants.SUCCESSPATHNAME); err == nil {
			oldValues2, err = fileRepo.GetData("./" + constants.SUCCESSPATHNAME)
			if err != nil {
				return fmt.Errorf("error reading existing values")
			}
		}
		userValid, err := userRepo.CheckUserValid(newUser, usersValid, oldValues2)
		if err != nil {
			return err
		}
		if userValid {
			usersValid = append(usersValid, newUser)
		} else {
			userInvalid = append(userInvalid, newUser)
		}
	}

	if err != nil {
		file.Close()
		secondaryFile.MoveFileProcessedError(filename, "")
		return err
	}
	oldValues, err := fileRepo.GetData(constants.SUCCESSPATHNAME)
	if err != nil {
		return fmt.Errorf("error reading existing values")
	}
	err = secondaryWriteFile.SaveUsers(oldValues, usersValid, "", true)
	err = secondaryWriteFile.SaveUsers(oldValues, userInvalid, "", false)
	if err != nil {
		file.Close()
		secondaryFile.MoveFileProcessedError(filename, "")
		return err
	}
	file.Close()
	err = secondaryFile.MoveFileProcessed(filename, "")
	if err != nil {
		return err
	}
	fmt.Println("---------------------------- End of Processing Information ----------------------------")
	fmt.Println("---------------------------- Start of summary ----------------------------")
	fmt.Println("Valid users: ", usersValid)
	fmt.Println("Invalid users: ", userInvalid)
	fmt.Println("Success file is located: ", "./"+constants.SUCCESSPATHNAME)
	fmt.Println("Fault file is located: ", "./"+constants.ERRORPATHNAME)
	fmt.Println("Processed files are located: ", constants.PATHPROCESSED)
	fmt.Println("Processed with error file is located: ", constants.PATHPROCESSEDERROR)
	fmt.Println("---------------------------- End of summary ----------------------------")
	fmt.Println("############################ FINISHED ############################")
	return nil
}
