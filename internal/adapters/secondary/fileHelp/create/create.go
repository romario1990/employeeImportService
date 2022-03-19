package secondaryCreateFile

import (
	"fmt"
	"os"
	"uploader/constants"
	secondaryWriteFile "uploader/internal/adapters/secondary/fileHelp/write"
)

// TODO testar
func createCSV(path string, filename string, extension string, initialValue [][]string) error {
	file, err := os.Create(path + filename + extension)
	if err != nil {
		return fmt.Errorf("unable to create output file " + filename)
	}
	secondaryWriteFile.WriteCSV(file, initialValue)
	file.Close()
	return nil
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

// TODO verificar para remover
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
