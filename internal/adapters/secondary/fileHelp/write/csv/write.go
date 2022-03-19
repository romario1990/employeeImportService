package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"uploader/constants"
	"uploader/pkg/domains/users"
)

func WriteCSV(file *os.File, values [][]string) {
	w := csv.NewWriter(file)
	for _, row := range values {
		_ = w.Write(row)
	}
	w.Flush()
}

//TODO remover pkg
func SaveInFile(oldValues [][]string, filename string, users []users.ConfigurationHeaderExport) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error trying to edit file")
	}
	var data [][]string
	for _, record := range oldValues {
		row := []string{record[0], record[1], record[2], record[3], record[4], record[5]}
		data = append(data, row)
	}
	for _, record := range users {
		row := []string{record.Name, record.Email, record.Salary, record.Identifier, record.Phone, record.Mobile}
		data = append(data, row)
	}
	WriteCSV(file, data)
	file.Close()
	return nil
}

//TODO remover pkg
func SaveUsers(oldValues [][]string, validUsers []users.ConfigurationHeaderExport, path string, userValid bool) (err error) {
	if path == "" && userValid {
		path = "./" + constants.SUCCESSPATHNAME
	} else if path == "" && !userValid {
		path = "./" + constants.ERRORPATHNAME
	}
	err = SaveInFile(oldValues, path, validUsers)
	if err != nil {
		return err
	}
	return nil
}
