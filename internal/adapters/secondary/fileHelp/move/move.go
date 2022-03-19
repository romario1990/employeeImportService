package secondaryMoveFile

import (
	"io"
	"os"
	"regexp"
	"uploader/constants"
)

func MoveFile(source, destination string) (err error) {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()
	fi, err := src.Stat()
	if err != nil {
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		dst.Close()
		os.Remove(destination)
		return err
	}
	err = dst.Close()
	if err != nil {
		return err
	}
	err = src.Close()
	if err != nil {
		return err
	}
	err = os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}

func MoveFileProcessed(filename string, defaultPah string) error {
	if defaultPah == "" {
		defaultPah = constants.PATHPROCESSED
	}
	var re = regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	newPath := re.FindStringSubmatch(filename)
	err := MoveFile(filename, defaultPah+newPath[2]+".csv")
	return err
}

func MoveFileProcessedError(filename string, defaultPah string) error {
	if defaultPah == "" {
		defaultPah = constants.PATHPROCESSEDERROR
	}
	var re = regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	newPath := re.FindStringSubmatch(filename)
	err := MoveFile(filename, defaultPah+newPath[2]+".csv")
	return err
}
