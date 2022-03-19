package secondaryUserValidator

import (
	"os"
	"uploader/constants"
	secondaryEmailValidator "uploader/internal/adapters/secondary/emailValidator"
	"uploader/pkg/domains/users"
)

func ValidateFieldsAlreadyRegistered(newUser users.ConfigurationHeaderExport, user []string) bool {
	return newUser.Email == user[0] || newUser.Identifier == user[1]
}

func CheckUserValid(user users.ConfigurationHeaderExport, users []users.ConfigurationHeaderExport, oldValues [][]string) (bool, error) {
	valid := true
	if user.Email == "" || user.Identifier == "" || !secondaryEmailValidator.ValidateEmail(user.Email) {
		valid = false
	}
	if _, err := os.Stat("./" + constants.SUCCESSPATHNAME); err == nil {
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
