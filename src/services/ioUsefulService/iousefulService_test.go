package ioUsefulService

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
	"uploader/src/entities"
)

func setUp() {
	defaultPath := "./../../../transfer/mock/"
	file1, _ := os.Create(defaultPath + constants.SUCCESSPATHNAME)
	w := csv.NewWriter(file1)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file1.Close()
	file2, _ := os.Create(defaultPath + constants.ERRORPATHNAME)
	w = csv.NewWriter(file2)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file2.Close()
	file3, _ := os.Create(defaultPath + "transfer/processed/roster1.csv")
	file3.Close()
}

func tearDown() {
	defaultPath := "./../../../transfer/mock/"
	os.Remove(defaultPath + constants.SUCCESSPATHNAME)
	os.Remove(defaultPath + constants.ERRORPATHNAME)
	os.Remove(defaultPath + "transfer/processed/roster1.csv")
	os.Remove(defaultPath + "transfer/processedError/roster1.csv")
}

func TestSaveInFile(t *testing.T) {
	setUp()
	defaultPath := "./../../../transfer/mock/"
	validUsers := []entities.ConfigurationHeaderExport{
		{"John Doe", "doe@test.com", "$10.00", "1", "", ""},
		{"Mary Jane", "mary@tes.com", "$15", "2", "", ""},
		{"Max Topperson", "max@test.com", "$11", "3", "", ""},
	}
	validUsersT := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"John Doe", "doe@test.com", "$10.00", "1", "", ""},
		[]string{"Mary Jane", "mary@tes.com", "$15", "2", "", ""},
		[]string{"Max Topperson", "max@test.com", "$11", "3", "", ""},
	}
	type args struct {
		filename string
		users    []entities.ConfigurationHeaderExport
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Valid_And_Invalid_Users",
			args{defaultPath + constants.SUCCESSPATHNAME, validUsers},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveInFile(tt.args.filename, tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("SaveInFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			file, _ := os.Open(defaultPath + constants.SUCCESSPATHNAME)
			var validUsersTest [][]string
			csvReader := csv.NewReader(file)
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				validUsersTest = append(validUsersTest, rec)
			}
			if !reflect.DeepEqual(validUsersT, validUsersTest) {
				t.Errorf("TestSaveInFile() got = %v, want %v", validUsersT, validUsersTest)
			}
			file.Close()
			tearDown()
		})
	}
}

func TestSaveUsers(t *testing.T) {
	setUp()
	defaultPath := "./../../../transfer/mock/"
	validUsers := []entities.ConfigurationHeaderExport{
		{"John Doe", "doe@test.com", "$10.00", "1", "", ""},
		{"Mary Jane", "mary@tes.com", "$15", "2", "", ""},
		{"Max Topperson", "max@test.com", "$11", "3", "", ""},
	}
	inValidUsers := []entities.ConfigurationHeaderExport{
		{"Alfred Donald", "", "$11.5", "4", "", ""},
		{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	validUsersT := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"John Doe", "doe@test.com", "$10.00", "1", "", ""},
		[]string{"Mary Jane", "mary@tes.com", "$15", "2", "", ""},
		[]string{"Max Topperson", "max@test.com", "$11", "3", "", ""},
	}
	inValidUsersT := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"Alfred Donald", "", "$11.5", "4", "", ""},
		[]string{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	type args struct {
		validUsers   []entities.ConfigurationHeaderExport
		inValidUsers []entities.ConfigurationHeaderExport
		successPath  string
		errorPath    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Valid_And_Invalid_Users",
			args{validUsers, inValidUsers, defaultPath + constants.SUCCESSPATHNAME, defaultPath + constants.ERRORPATHNAME},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveUsers(tt.args.validUsers, tt.args.inValidUsers, tt.args.successPath, tt.args.errorPath); (err != nil) != tt.wantErr {
				t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			file, _ := os.Open(defaultPath + constants.SUCCESSPATHNAME)
			var validUsersTest [][]string
			csvReader := csv.NewReader(file)
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				validUsersTest = append(validUsersTest, rec)
			}
			if !reflect.DeepEqual(validUsersT, validUsersTest) {
				t.Errorf("TestSaveUsers() got = %v, want %v", validUsersT, validUsersTest)
			}
			file.Close()
			file2, _ := os.Open(defaultPath + constants.ERRORPATHNAME)
			var inValidUsersTest [][]string
			csvReader2 := csv.NewReader(file2)
			for {
				rec, err := csvReader2.Read()
				if err == io.EOF {
					break
				}
				inValidUsersTest = append(inValidUsersTest, rec)
			}
			file2.Close()
			if !reflect.DeepEqual(inValidUsersT, inValidUsersTest) {
				t.Errorf("TestSaveUsers() got = %v, want %v", inValidUsersT, inValidUsersTest)
			}
			tearDown()
		})
	}
}

func TestCreateDefaultFiles(t *testing.T) {
	defaultPath := "./../../../transfer/mock/"
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

func TestGetDataCSV(t *testing.T) {
	setUp()
	defaultPathAndName := "./../../../transfer/mock/transfer/success/employeesucess.csv"
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			"Test_Header_Export",
			args{defaultPathAndName},
			constants.HEADER,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataCSV(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataCSV() got = %v, want %v", got, tt.want)
			}
			tearDown()
		})
	}
}

func TestMoveFileProcessed(t *testing.T) {
	pathProcessedSuccess := "./../../../transfer/mock/transfer/processed/"
	fileName := "./../../../transfer/mock/transfer/pending/roster1.csv"
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
			if _, err := os.Stat("./../../../transfer/mock/transfer/processed/roster1.csv"); err != nil {
				t.Errorf("TestMoveFileProcessed() error, error moving processed file = roster1.csv")
			}
			if _, err := os.Stat("./../../../transfer/mock/transfer/pending/roster1.csv"); err == nil {
				t.Errorf("TestMoveFileProcessed() error, error moving pending file = roster1.csv")
			}
			source := "./../../../transfer/mock/transfer/processed/roster1.csv"
			destination := "./../../../transfer/mock/transfer/pending/roster1.csv"
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
	pathProcessedError := "./../../../transfer/mock/transfer/processedError/"
	fileName := "./../../../transfer/mock/transfer/processed/roster1.csv"
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
			if _, err := os.Stat("./../../../transfer/mock/transfer/processed/roster1.csv"); err == nil {
				t.Errorf("TestMoveFileProcessedError() error, error moving processed error file = roster1.csv")
			}
			if _, err := os.Stat("./../../../transfer/mock/transfer/processedError/roster1.csv"); err != nil {
				t.Errorf("TestMoveFileProcessedError() error, error moving processed file = roster1.csv")
			}
			tearDown()
		})
	}
}

func TestReadAllNameFilesPath(t *testing.T) {
	pathTestMock := "./../../../transfer/mock/transfer/pending/"
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
			"Test_Filename_Reads",
			args{pathTestMock},
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
