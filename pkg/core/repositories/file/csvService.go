package file

import (
	"encoding/csv"
	"fmt"
	"io"
	secondaryFile "uploader/helpers/fileHelp"
)

type fileRepo struct {
	repo *FileRepository
}

func NewFileCSVRepo() FileRepository {
	return &fileRepo{}
}

func (repo *fileRepo) GetData(filename string) ([][]string, error) {
	file, err := secondaryFile.Read(filename)
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
