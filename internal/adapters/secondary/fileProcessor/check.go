package secondaryFileProcessor

import (
	"fmt"
	"uploader/config"
	"uploader/constants"
	secondaryFile "uploader/internal/adapters/secondary/fileHelp/file"
	secondaryReadFileCSV "uploader/internal/adapters/secondary/fileHelp/read/csv"
	secondaryWriteFile "uploader/internal/adapters/secondary/fileHelp/write/csv"
	"uploader/pkg/domains/users"
)

func Exec(filename string, hasH bool) error {
	fmt.Println("############################ STARTING ############################")
	fmt.Println("---------------------------- Start of processing information ----------------------------")
	file, err := secondaryFile.Read(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	secondaryFile.InitExec(nil)
	var usersValid, userInvalid []users.ConfigurationHeaderExport
	confHeader, err := config.LoadConfigHeader()
	if err != nil {
		file.Close()
		secondaryFile.MoveFileProcessedError(filename, "")
		return fmt.Errorf("headerconfiguration.json file not configured")
	}
	configHeader := users.ConfigurationHeader{
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
	usersValid, userInvalid, err = secondaryReadFileCSV.ReadFile(file, hasH, configHeader, 9)
	if err != nil {
		file.Close()
		secondaryFile.MoveFileProcessedError(filename, "")
		return err
	}
	oldValues, err := secondaryReadFileCSV.GetDataCSV(constants.SUCCESSPATHNAME)
	fmt.Println("AQUIIIIIIIII 1 ", usersValid)
	fmt.Println("AQUIIIIIIIII 2 ", userInvalid)
	fmt.Println("AQUIIIIIIIII 3 ", oldValues)
	if err != nil {
		return fmt.Errorf("error reading existing values")
	}
	err = secondaryWriteFile.SaveUsers(oldValues, usersValid, "", true)
	err = secondaryWriteFile.SaveUsers(oldValues, userInvalid, "", false)
	fmt.Println("AQUIIIIIIIII 4 ")
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
