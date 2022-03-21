package fileHelp

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CreateFileCSV(filePathName string, initialValue [][]string) error {
	file, err := os.Create(filePathName)
	if err != nil {
		return fmt.Errorf("unable to create output file " + filePathName)
	}
	WriteCSV(file, initialValue)
	file.Close()
	return nil
}

func WriteCSV(file *os.File, values [][]string) {
	w := csv.NewWriter(file)
	for _, row := range values {
		_ = w.Write(row)
	}
	w.Flush()
}

func MoveFile(source, destination string) (err error) {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()
	fi, err := src.Stat()
	if err != nil {
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		dst.Close()
		os.Remove(destination)
		return err
	}
	err = dst.Close()
	if err != nil {
		return err
	}
	err = src.Close()
	if err != nil {
		return err
	}
	err = os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}

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
