package secondaryFileProcessor

import (
	"fmt"
	"uploader/config"
	"uploader/constants"
	"uploader/src/entities"
	"uploader/src/services/ioUsefulService"
	"uploader/src/services/parseService"
)

func Exec(filename string, hasH bool) error {
	fmt.Println("############################ STARTING ############################")
	fmt.Println("---------------------------- Start of processing information ----------------------------")
	file, err := ioUsefulService.Read(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	ioUsefulService.InitExec(nil)
	var usersValid, userInvalid []entities.ConfigurationHeaderExport
	confHeader, err := config.LoadConfigHeader()
	if err != nil {
		file.Close()
		ioUsefulService.MoveFileProcessedError(filename, "")
		return fmt.Errorf("headerconfiguration.json file not configured")
	}
	configHeader := entities.ConfigurationHeader{
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
	fmt.Printf("### The %v file is being processed", filename)
	fmt.Printf("### Header de configuration file = %v", configHeader)
	usersValid, userInvalid, err = parseService.ReadCSV(file, hasH, configHeader, 9)
	if err != nil {
		file.Close()
		ioUsefulService.MoveFileProcessedError(filename, "")
		return err
	}
	err = ioUsefulService.SaveUsers(usersValid, userInvalid, "", "")
	if err != nil {
		file.Close()
		ioUsefulService.MoveFileProcessedError(filename, "")
		return err
	}
	file.Close()
	err = ioUsefulService.MoveFileProcessed(filename, "")
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
