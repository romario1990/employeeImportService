package users

type UserRepository interface {
	FindAll() ([]ConfigurationHeaderExport, error)
	Save([]ConfigurationHeaderExport) error
	ConvertDataToHeaderExport(dataFile [][]string, header []string) ([]ConfigurationHeaderExport, error)
	GetField(v ConfigurationHeader, field string) ([]string, error)
	FormatHeader(header []string, configHeader ConfigurationHeader, sizeStructHeader int) []string
	CheckUserValid(user ConfigurationHeaderExport, users []ConfigurationHeaderExport, oldValues [][]string) (bool, error)
}
