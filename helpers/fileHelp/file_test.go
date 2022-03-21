package fileHelp

import (
	"os"
	"reflect"
	"testing"
)

func setUp() {
	defaultPath := "./../../transfer/mock/"
	file3, _ := os.Create(defaultPath + "transfer/processed/roster1.csv")
	file3.Close()
}

func tearDown() {
	defaultPath := "./../../transfer/mock/"
	os.Remove(defaultPath + "transfer/processedError/roster1.csv")
}

func TestReadAllNameFilesPath(t *testing.T) {
	fileName := "./../../transfer/mock/transfer/pending/"
	type args struct {
		pathName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Test_Read_All_Files_In_A_Directory",
			args{fileName},
			[]string{"roster1.csv", "roster_no_header.csv"},
			false,
		},
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

func TestCreateFileCSV(t *testing.T) {
	fileName := "../../transfer/mock/test_create.csv"
	type args struct {
		filePathName string
		initialValue [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_File_Creation",
			args{"../../transfer/mock/test_create.csv", [][]string{[]string{"teste"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFileCSV(tt.args.filePathName, tt.args.initialValue); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(fileName); err != nil {
				t.Errorf("CreateFileCSV() error = file was not created")
			}
			os.Remove(fileName)
		})
	}
}

func TestMoveFile(t *testing.T) {
	setUp()
	pathProcessedSuccess := "./../../transfer/mock/transfer/processedError/roster1.csv"
	fileName := "./../../transfer/mock/transfer/processed/roster1.csv"
	type args struct {
		source      string
		destination string
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
			if err := MoveFile(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("MoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat("./../../transfer/mock/transfer/processedError/roster1.csv"); err != nil {
				t.Errorf("MoveFile() error, error moving processedError file = roster1.csv")
			}
			if _, err := os.Stat("./../../transfer/mock/transfer/processed/roster1.csv"); err == nil {
				t.Errorf("MoveFile() error, error moving processed file = roster1.csv")
			}
			tearDown()
		})
	}
}
