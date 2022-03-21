package all

import (
	"uploader/constants"
	"uploader/handlers/cmd/check"
	secondaryFile "uploader/helpers/fileHelp"
)

func ExecAll(hasH bool, fileType string) error {
	filesNames, err := secondaryFile.ReadAllNameFilesPath(constants.PATHPPENDINGROCESSED)
	if err != nil {
		return err
	}
	for _, name := range filesNames {
		check.Exec(constants.PATHPPENDINGROCESSED+name, hasH, fileType)
	}
	return nil
}
