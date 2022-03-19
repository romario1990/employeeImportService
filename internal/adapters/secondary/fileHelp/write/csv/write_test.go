package csv

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
	"uploader/pkg/domains/users"
)

func setUp() {
	defaultPath := "./../../../../../../transfer/mock/"
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
}

func tearDown() {
	defaultPath := "./../../../../../../transfer/mock/"
	os.Remove(defaultPath + constants.SUCCESSPATHNAME)
	os.Remove(defaultPath + constants.ERRORPATHNAME)
}

func TestSaveUsersValid(t *testing.T) {
	setUp()
	defaultPath := "./../../../../../../transfer/mock/"
	validUsers := []users.ConfigurationHeaderExport{
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
		oldValues  [][]string
		validUsers []users.ConfigurationHeaderExport
		path       string
		userValid  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Valid_Users",
			args{[][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}, validUsers, defaultPath + constants.SUCCESSPATHNAME, true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveUsers(tt.args.oldValues, tt.args.validUsers, tt.args.path, tt.args.userValid); (err != nil) != tt.wantErr {
				t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			file, _ := os.Open(tt.args.path)
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
			tearDown()
		})
	}
}

func TestSaveUsersInValid(t *testing.T) {
	setUp()
	defaultPath := "./../../../../../../transfer/mock/"
	inValidUsers := []users.ConfigurationHeaderExport{
		{"Alfred Donald", "", "$11.5", "4", "", ""},
		{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	invalidUsersT := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"Alfred Donald", "", "$11.5", "4", "", ""},
		[]string{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	type args struct {
		oldValues  [][]string
		validUsers []users.ConfigurationHeaderExport
		path       string
		userValid  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Invalid_Users",
			args{[][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}, inValidUsers, defaultPath + constants.ERRORPATHNAME, false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveUsers(tt.args.oldValues, tt.args.validUsers, tt.args.path, tt.args.userValid); (err != nil) != tt.wantErr {
				t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			file, _ := os.Open(tt.args.path)
			var validUsersTest [][]string
			csvReader := csv.NewReader(file)
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				validUsersTest = append(validUsersTest, rec)
			}
			if !reflect.DeepEqual(invalidUsersT, validUsersTest) {
				t.Errorf("TestSaveUsers() got = %v, want %v", invalidUsersT, validUsersTest)
			}
			file.Close()
			tearDown()
		})
	}
}

func TestSaveInFile(t *testing.T) {
	setUp()
	defaultPath := "./../../../../../../transfer/mock/"
	validUsers := []users.ConfigurationHeaderExport{
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
		oldValues [][]string
		filename  string
		users     []users.ConfigurationHeaderExport
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Valid_And_Invalid_Users",
			args{[][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}, defaultPath + constants.SUCCESSPATHNAME, validUsers},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveInFile(tt.args.oldValues, tt.args.filename, tt.args.users); (err != nil) != tt.wantErr {
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
