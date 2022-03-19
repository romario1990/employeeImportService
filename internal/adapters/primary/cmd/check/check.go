package primaryCheck

import secondaryFileProcessor "uploader/internal/adapters/secondary/fileProcessor"

func Exec(filename string, hasH bool) error {
	return secondaryFileProcessor.Exec(filename, hasH)
}
