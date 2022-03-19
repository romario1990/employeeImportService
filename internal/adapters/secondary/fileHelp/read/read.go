package secondaryReadFile

import (
	"io/ioutil"
	"os"
)

func Read(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ReadAllNameFilesPath(pathName string) ([]string, error) {
	var filesName []string
	files, err := ioutil.ReadDir(pathName)
	if err != nil {
		return []string{}, nil
	}
	for _, file := range files {
		filesName = append(filesName, file.Name())
	}
	return filesName, nil
}
