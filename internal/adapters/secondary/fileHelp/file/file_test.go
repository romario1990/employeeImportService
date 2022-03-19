package secondaryFile

import (
	"io"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
)

func setUp() {
	defaultPath := "./../../../../../transfer/mock/"
	file3, _ := os.Create(defaultPath + "transfer/processed/roster1.csv")
	file3.Close()
}

func tearDown() {
	defaultPath := "./../../../../../transfer/mock/"
	os.Remove(defaultPath + "transfer/processed/roster1.csv")
	os.Remove(defaultPath + "transfer/processedError/roster1.csv")
}

func TestMoveFileProcessed(t *testing.T) {
	pathProcessedSuccess := "./../../../../../transfer/mock/transfer/processed/"
	fileName := "./../../../../../transfer/mock/transfer/pending/roster1.csv"
	type args struct {
		filename   string
		defaultPah string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test_Move_File_Processed",
			args{fileName, pathProcessedSuccess},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MoveFileProcessed(tt.args.filename, tt.args.defaultPah); (err != nil) != tt.wantErr {
				t.Errorf("MoveFileProcessed() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat("./../../../../../transfer/mock/transfer/processed/roster1.csv"); err != nil {
				t.Errorf("TestMoveFileProcessed() error, error moving processed file = roster1.csv")
			}
			if _, err := os.Stat("./../../../../../transfer/mock/transfer/pending/roster1.csv"); err == nil {
				t.Errorf("TestMoveFileProcessed() error, error moving pending file = roster1.csv")
			}
			source := "./../../../../../transfer/mock/transfer/processed/roster1.csv"
			destination := "./../../../../../transfer/mock/transfer/pending/roster1.csv"
			src, _ := os.Open(source)
			defer src.Close()
			fi, _ := src.Stat()
			flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			perm := fi.Mode() & os.ModePerm
			dst, _ := os.OpenFile(destination, flag, perm)
			defer dst.Close()
			_, _ = io.Copy(dst, src)
			dst.Close()
			src.Close()
			os.Remove(source)
		})
	}
}

func TestMoveFileProcessedError(t *testing.T) {
	setUp()
	pathProcessedError := "./../../../../../transfer/mock/transfer/processedError/"
	fileName := "./../../../../../transfer/mock/transfer/processed/roster1.csv"
	type args struct {
		filename   string
		defaultPah string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test_Move_File_Processed_Error",
			args{fileName, pathProcessedError},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MoveFileProcessedError(tt.args.filename, tt.args.defaultPah); (err != nil) != tt.wantErr {
				t.Errorf("MoveFileProcessedError() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat("./../../../../../transfer/mock/transfer/processed/roster1.csv"); err == nil {
				t.Errorf("TestMoveFileProcessedError() error, error moving processed error file = roster1.csv")
			}
			if _, err := os.Stat("./../../../../../transfer/mock/transfer/processedError/roster1.csv"); err != nil {
				t.Errorf("TestMoveFileProcessedError() error, error moving processed file = roster1.csv")
			}
			tearDown()
		})
	}
}

func TestCreateDefaultFiles(t *testing.T) {
	defaultPath := "./../../../../../transfer/mock/"
	type args struct {
		header      [][]string
		defaultPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Creation_Of_Standard_Files_Of_Success_And_Error",
			args{constants.HEADER, defaultPath},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDefaultFiles(tt.args.header, tt.args.defaultPath); (err != nil) != tt.wantErr {
				t.Errorf("CreateDefaultFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(defaultPath + constants.SUCCESSPATHNAME); err != nil {
				t.Errorf("CreateDefaultFiles() error, success file not created = %v", constants.SUCCESSNAMEFILE)
			}
			if _, err := os.Stat(defaultPath + constants.ERRORPATHNAME); err != nil {
				t.Errorf("CreateDefaultFiles() error, error file not created = %v", constants.ERRORNAMEFILE)
			}
			e := os.Remove(defaultPath + constants.SUCCESSPATHNAME)
			if e != nil {
				t.Errorf("CreateDefaultFiles() error, success file not deleted = %v", constants.ERRORNAMEFILE)
			}
			e = os.Remove(defaultPath + constants.ERRORPATHNAME)
			if e != nil {
				t.Errorf("CreateDefaultFiles() error, error file not deleted = %v", constants.ERRORNAMEFILE)
			}
		})
	}
}

func TestReadAllNameFilesPath(t *testing.T) {
	type args struct {
		pathName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAllNameFilesPath(tt.args.pathName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAllNameFilesPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAllNameFilesPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
