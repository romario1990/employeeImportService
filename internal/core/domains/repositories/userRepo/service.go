package userRepo

import (
	"fmt"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/repositories/ports"
)

type userRepo struct {
	userRepo *UserRepository
	fileRepo ports.FileRepository
}

func NewUserRepo(header [][]string, fileRepo ports.FileRepository) UserRepository {
	fileRepo.CreateDefault(header)
	return &userRepo{nil, fileRepo}
}

func (repo *userRepo) Save(usersList []userModel.ConfigurationHeaderExport) error {
	oldValues, err := repo.fileRepo.GetData()
	if err != nil {
		fmt.Errorf("error querying recorded value")
	}

	return repo.fileRepo.SaveInFile(oldValues, usersList)
}

func (repo *userRepo) GetList() ([][]string, error) {
	data, err := repo.fileRepo.GetData()
	if err != nil {
		return [][]string{[]string{}}, nil
	}
	return data, nil
}

func (repo *userRepo) MoveFileProcessed() error {
	return repo.fileRepo.MoveFileProcessed()
}
