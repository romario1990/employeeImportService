package secondaryReadFileCSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"uploader/constants"
	secondaryFile "uploader/internal/adapters/secondary/fileHelp/file"
	secondaryHeaderProcessor "uploader/internal/adapters/secondary/headerProcessor"
	secondaryUserValidator "uploader/internal/adapters/secondary/userValidator"
	"uploader/pkg/domains/users"
)

//TODO remover daqui e quabra em menas funções
func ReadFile(f *os.File, hasHeader bool, configHeader users.ConfigurationHeader,
	sizeStructHeader int) ([]users.ConfigurationHeaderExport, []users.ConfigurationHeaderExport, error) {
	if f == nil {
		return []users.ConfigurationHeaderExport{}, []users.ConfigurationHeaderExport{}, fmt.Errorf("no files provided")
	}
	rows := csv.NewReader(f)
	var header []string
	if hasHeader {
		row, _ := rows.Read()
		header = secondaryHeaderProcessor.FormatHeader(row, configHeader, sizeStructHeader)
	}
	var usersValid, usersInvalid []users.ConfigurationHeaderExport
	for {
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return usersValid, usersInvalid, err
		}
		newUser := secondaryHeaderProcessor.FormatCSVExport(row, header)
		var oldValues [][]string
		if _, err := os.Stat("./" + constants.SUCCESSPATHNAME); err == nil {
			oldValues, err = GetDataCSV("./" + constants.SUCCESSPATHNAME)
			if err != nil {
				return []users.ConfigurationHeaderExport{}, []users.ConfigurationHeaderExport{}, fmt.Errorf("error reading existing values")
			}
		}
		userValid, err := secondaryUserValidator.CheckUserValid(newUser, usersValid, oldValues)
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

//TODO remover daqui e quabra em menas funções pkg
func GetDataCSV(filename string) ([][]string, error) {
	file, err := secondaryFile.Read(filename)
	if err != nil {
		return [][]string{}, fmt.Errorf("error trying to edit file")
	}
	defer file.Close()
	var data [][]string
	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return [][]string{}, nil
		}
		data = append(data, rec)
	}
	return data, nil
}
