package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"uploader/helpers/fileHelp"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/repositories/ports"
)

type fileRepo struct {
	repo           *ports.FileRepository
	filename       string
	destinationPah string
}

func NewFileCSVRepo(filename string, destinationPah string) ports.FileRepository {
	return &fileRepo{nil, filename, destinationPah}
}

func (repo *fileRepo) CreateDefault(header [][]string) {
	if _, err := os.Stat(repo.filename); err != nil {
		fileHelp.CreateFileCSV(repo.filename, header)
	}
}

func (repo *fileRepo) GetData() ([][]string, error) {
	if _, err := os.Stat(repo.filename); err == nil {
		file, err := fileHelp.Read(repo.filename)
		if err != nil {
			return [][]string{}, fmt.Errorf("error trying to edit file")
		}
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
		file.Close()
		return data, nil
	}
	return [][]string{[]string{}}, nil
}

func (repo *fileRepo) SaveInFile(oldValues [][]string, users []userModel.ConfigurationHeaderExport) error {
	file, err := os.OpenFile(repo.filename, os.O_RDWR, 0644)
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
	fileHelp.WriteCSV(file, data)
	file.Close()
	return nil
}

func (repo *fileRepo) MoveFileProcessed() error {
	var re = regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	newPath := re.FindStringSubmatch(repo.filename)
	err := fileHelp.MoveFile(repo.filename, repo.destinationPah+newPath[2]+".csv")
	return err
}
