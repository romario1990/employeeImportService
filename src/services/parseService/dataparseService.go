package parseService

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"uploader/src/entities"
	"uploader/src/services/checkService"
	"uploader/src/services/formatService"
)

func ReadCSV(f *os.File, hasHeader bool, configHeader entities.ConfigurationHeader,
	sizeStructHeader int) ([]entities.ConfigurationHeaderExport, []entities.ConfigurationHeaderExport, error) {
	if f == nil {
		return []entities.ConfigurationHeaderExport{}, []entities.ConfigurationHeaderExport{}, fmt.Errorf("no files provided")
	}
	rows := csv.NewReader(f)
	var header []string
	if hasHeader {
		row, _ := rows.Read()
		header = formatService.FormatHeader(row, configHeader, sizeStructHeader)
	}
	var usersValid, usersInvalid []entities.ConfigurationHeaderExport
	for {
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return usersValid, usersInvalid, err
		}
		newUser := formatService.FormatCSVExport(row, header)
		userValid, err := checkService.CheckUserValid(newUser, usersValid)
		if err != nil {
			return usersValid, usersInvalid, err
		}
		if userValid {
			usersValid = append(usersValid, newUser)
		} else {
			usersInvalid = append(usersInvalid, newUser)
		}

	}
	return usersValid, usersInvalid, nil
}
