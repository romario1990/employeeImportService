package secondaryAllFileProcessor

import (
	"uploader/constants"
	secondaryFileProcessor "uploader/internal/adapters/secondary/fileProcessor"
	"uploader/src/services/ioUsefulService"
)

func Exec(hasH bool) error {
	filesNames, err := ioUsefulService.ReadAllNameFilesPath(constants.PATHPPENDINGROCESSED)
	if err != nil {
		return err
	}
	for _, name := range filesNames {
		secondaryFileProcessor.Exec(constants.PATHPPENDINGROCESSED+name, hasH)
	}
	return nil
}
