package userRepo

import (
	"uploader/internal/core/domains/model/userModel"
)

type UserRepository interface {
	Save([]userModel.ConfigurationHeaderExport) error
	GetList() ([][]string, error)
	MoveFileProcessed() error
}
