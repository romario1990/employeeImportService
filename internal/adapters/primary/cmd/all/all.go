package primaryAll

import secondaryAllFileProcessor "uploader/internal/adapters/secondary/allFileProcessor"

func ExecAll(hasH bool) error {
	return secondaryAllFileProcessor.Exec(hasH)
}
