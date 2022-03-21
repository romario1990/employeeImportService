package check

import "uploader/internal/core/domains/usecases/userUsecase"

func Exec(filename string, hasH bool, fileType string) error {
	var userExternalService userUsecase.UserExternalService
	userExternalService = userUsecase.NewUserExternalService(filename, hasH, fileType)
	return userExternalService.Exec()
}
