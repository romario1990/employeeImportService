package file

type FileRepository interface {
	GetData(filename string) ([][]string, error)
}
