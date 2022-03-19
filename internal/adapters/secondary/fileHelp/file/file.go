package secondaryFile

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"uploader/constants"
	secondaryWriteFile "uploader/internal/adapters/secondary/fileHelp/write/csv"
)

func createCSV(path string, filename string, extension string, initialValue [][]string) error {
	file, err := os.Create(path + filename + extension)
	if err != nil {
		return fmt.Errorf("unable to create output file " + filename)
	}
	secondaryWriteFile.WriteCSV(file, initialValue)
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
