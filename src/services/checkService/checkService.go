package checkService

import (
	"fmt"
	"net/mail"
	"os"
	"uploader/constants"
	"uploader/src/entities"
	"uploader/src/services/ioUsefulService"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateFieldsAlreadyRegistered(newUser entities.ConfigurationHeaderExport, user []string) bool {
	return newUser.Email == user[0] || newUser.Identifier == user[1]
}

func CheckUserValid(user entities.ConfigurationHeaderExport, users []entities.ConfigurationHeaderExport) (bool, error) {
	valid := true
	if user.Email == "" || user.Identifier == "" || !validateEmail(user.Email) {
		valid = false
	}
	if _, err := os.Stat("./" + constants.SUCCESSPATHNAME); err == nil {
		oldValues, err := ioUsefulService.GetDataCSV("./" + constants.SUCCESSPATHNAME)
		if err != nil {
			return false, fmt.Errorf("error reading existing values")
		}
		if valid {
			for _, record := range oldValues {
				if ValidateFieldsAlreadyRegistered(user, []string{record[constants.POSITIONHEADEREMAIL], record[constants.POSITIONHEADERIDENTIFIER]}) {
					valid = false
					break
				}
			}
		}
	}

	if valid {
		for _, record := range users {
			if ValidateFieldsAlreadyRegistered(user, []string{record.Email, record.Identifier}) {
				valid = false
				break
			}
		}
	}
	return valid, nil
}
