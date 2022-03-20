package primaryAll

import secondaryAllFileProcessor "uploader/internal/adapters/secondary/allFileProcessor"

func ExecAll(hasH bool, fileType string) error {
	return secondaryAllFileProcessor.Exec(hasH, fileType)
}
