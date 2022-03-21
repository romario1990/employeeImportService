package all

import (
	"fmt"
	"os"
	"uploader/constants"
	"uploader/helpers/fileHelp"
	"uploader/internal/core/domains/usecases/userUsecase"
)

func ExecAll(hasH bool, fileType string) error {
	filesNames, err := fileHelp.ReadAllNameFilesPath(constants.PATHPPENDINGROCESSED)
	if err != nil {
		return err
	}
	if _, err := os.Stat("./headerConfiguration"); err != nil {
		return fmt.Errorf("configure headerConfiguration file")
	}
	var userExternalService userUsecase.UserExternalService
	for _, name := range filesNames {
		userExternalService = userUsecase.NewUserExternalService(constants.PATHPPENDINGROCESSED+name, hasH, fileType)
		err = userExternalService.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}
