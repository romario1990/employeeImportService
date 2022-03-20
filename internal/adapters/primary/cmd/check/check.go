package primaryCheck

import (
	secondaryFileProcessor "uploader/internal/adapters/secondary/fileProcessor"
)

func Exec(filename string, hasH bool, fileType string) error {
	return secondaryFileProcessor.Exec(filename, hasH, fileType)
}
