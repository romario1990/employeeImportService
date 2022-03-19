package secondaryAllFileProcessor

import (
	"uploader/constants"
	secondaryFile "uploader/internal/adapters/secondary/fileHelp/file"
	secondaryFileProcessor "uploader/internal/adapters/secondary/fileProcessor"
)

func Exec(hasH bool) error {
	filesNames, err := secondaryFile.ReadAllNameFilesPath(constants.PATHPPENDINGROCESSED)
	if err != nil {
		return err
	}
	for _, name := range filesNames {
		secondaryFileProcessor.Exec(constants.PATHPPENDINGROCESSED+name, hasH)
	}
	return nil
}
