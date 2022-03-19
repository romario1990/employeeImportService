package primaryAll

import secondaryAllFileProcessor "uploader/internal/adapters/secondary/allFileProcessor"

func Exec(hasH bool) error {
	return secondaryAllFileProcessor.Exec(hasH)
}
