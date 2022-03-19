package ioUsefulService

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"uploader/constants"
	"uploader/src/entities"
)

// Read ##IMPORTANT## Remember to close the file read in the function call (example defer file.Close())
func Read(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func writeCSV(file *os.File, values [][]string) {
	w := csv.NewWriter(file)
	for _, row := range values {
		_ = w.Write(row)
	}
	w.Flush()
}

func createCSV(path string, filename string, extension string, initialValue [][]string) error {
	file, err := os.Create(path + filename + extension)
	if err != nil {
		return fmt.Errorf("unable to create output file " + filename)
	}
	writeCSV(file, initialValue)
	file.Close()
	return nil
}

func moveFile(source, destination string) (err error) {
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

func CreateDefaultFiles(header [][]string, defaultPath string) error {
	if defaultPath == "" {
		defaultPath = "./"
	}
	if _, err := os.Stat(defaultPath + constants.SUCCESSPATHNAME); err != nil {
		fmt.Println("----------- Create valid data output file -----------")
		err = createCSV(defaultPath+constants.SUCCESSPATH, constants.SUCCESSNAMEFILE, constants.EXTENSIONCSV, header)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(defaultPath + constants.ERRORPATHNAME); err != nil {
		fmt.Println("----------- Create invalid data output file -----------")
		err = createCSV(defaultPath+constants.ERRORPATH, constants.ERRORNAMEFILE, constants.EXTENSIONCSV, header)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDataCSV(filename string) ([][]string, error) {
	file, err := Read(filename)
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

func SaveInFile(filename string, users []entities.ConfigurationHeaderExport) error {
	oldValues, err := GetDataCSV(filename)
	if err != nil {
		return fmt.Errorf("error reading existing values")
	}
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	defer file.Close()
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
	writeCSV(file, data)
	return nil
}

func MoveFileProcessed(filename string, defaultPah string) error {
	if defaultPah == "" {
		defaultPah = constants.PATHPROCESSED
	}
	var re = regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	newPath := re.FindStringSubmatch(filename)
	err := moveFile(filename, defaultPah+newPath[2]+".csv")
	return err
}

func MoveFileProcessedError(filename string, defaultPah string) error {
	if defaultPah == "" {
		defaultPah = constants.PATHPROCESSEDERROR
	}
	var re = regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	newPath := re.FindStringSubmatch(filename)
	err := moveFile(filename, defaultPah+newPath[2]+".csv")
	return err
}

func InitExec(header [][]string) (err error) {
	if header == nil {
		header = constants.HEADER
	}
	err = CreateDefaultFiles(header, "")
	if err != nil {
		return err
	}
	return nil
}

func SaveUsers(validUsers []entities.ConfigurationHeaderExport, inValidUsers []entities.ConfigurationHeaderExport,
	successPath string, errorPath string) (err error) {
	if successPath == "" {
		successPath = "./" + constants.SUCCESSPATHNAME
	}
	if errorPath == "" {
		errorPath = "./" + constants.ERRORPATHNAME
	}
	err = SaveInFile(successPath, validUsers)
	if err != nil {
		return err
	}
	err = SaveInFile(errorPath, inValidUsers)
	if err != nil {
		return err
	}
	return nil
}
