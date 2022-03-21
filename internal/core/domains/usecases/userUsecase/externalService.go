package userUsecase

import (
	"fmt"
	"os"
	"uploader/config"
	"uploader/constants"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/core/domains/repositories/userRepo"
	"uploader/internal/repositories/adapters/file/csv"
	"uploader/internal/repositories/ports"
)

type userExternalService struct {
	useCases *UserExternalService
	filename string
	hasH     bool
	fileType string
}

func NewUserExternalService(filename string, hasH bool, fileType string) UserExternalService {
	return &userExternalService{nil, filename, hasH, fileType}
}

func (u userExternalService) Exec() error {
	if u.fileType != constants.CSV {
		return fmt.Errorf("unsupported file format %s", u.fileType)
	}
	fmt.Println("############################ STARTING ############################")
	var userS UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	if u.fileType == constants.CSV {
		filePending = csv.NewFileCSVRepo(u.filename, constants.PATHPROCESSED)
		var fileRepoE ports.FileRepository
		fileRepoE = csv.NewFileCSVRepo("./"+constants.ERRORPATHNAME, constants.PATHPROCESSEDERROR)
		userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
		var fileRepoS ports.FileRepository
		fileRepoS = csv.NewFileCSVRepo("./"+constants.SUCCESSPATHNAME, constants.PATHPROCESSED)
		userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	}
	if _, err := os.Stat("./headerConfiguration"); err != nil {
		fmt.Println("configure headerConfiguration file")
		return fmt.Errorf("configure headerConfiguration file")
	}
	confHeader, err := config.LoadConfigHeader()
	configHeader := userModel.ConfigurationHeader{
		FullName:   confHeader.GetStringSlice(constants.FULLNAME),
		FirstName:  confHeader.GetStringSlice(constants.FIRSTNAME),
		MiddleName: confHeader.GetStringSlice(constants.MIDDLENAME),
		LastName:   confHeader.GetStringSlice(constants.LASTNAME),
		Email:      confHeader.GetStringSlice(constants.EMAIL),
		Salary:     confHeader.GetStringSlice(constants.SALARY),
		Identifier: confHeader.GetStringSlice(constants.IDENTIFIER),
		Phone:      confHeader.GetStringSlice(constants.PHONE),
		Mobile:     confHeader.GetStringSlice(constants.MOBILE),
	}

	userS, err = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, u.hasH)
	if err != nil {
		return err
	}
	err = userS.Exec()
	fmt.Println("############################ FINISHED ############################")
	return err
}
