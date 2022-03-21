package ports

import "uploader/internal/core/domains/model/userModel"

type FileRepository interface {
	CreateDefault(header [][]string)
	GetData() ([][]string, error)
	SaveInFile(oldValues [][]string, users []userModel.ConfigurationHeaderExport) error
	MoveFileProcessed() error
}
