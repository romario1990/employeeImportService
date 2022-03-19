package secondaryAllFileProcessor

import (
	"uploader/constants"
	secondaryReadFile "uploader/internal/adapters/secondary/fileHelp/read"
	secondaryFileProcessor "uploader/internal/adapters/secondary/fileProcessor"
)

func Exec(hasH bool) error {
	filesNames, err := secondaryReadFile.ReadAllNameFilesPath(constants.PATHPPENDINGROCESSED)
	if err != nil {
		return err
	}
	for _, name := range filesNames {
		secondaryFileProcessor.Exec(constants.PATHPPENDINGROCESSED+name, hasH)
	}
	return nil
}
