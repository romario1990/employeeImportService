package csv

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/repositories/ports"
)

func setUp() {
	defaultPath := "./../../../../../mock/transfer/pending/roster3.csv"
	file1, _ := os.Create(defaultPath)
	w := csv.NewWriter(file1)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file1.Close()
}

func tearDown() {
	os.Remove("./../../../../../mock/transfer/processedError/roster3.csv")
}

func Test_fileRepo_CreateDefault(t *testing.T) {
	fileName := "./../../../../../mock/transfer/pending/roster.csv"
	var filePending ports.FileRepository
	filePending = NewFileCSVRepo(fileName, "./../../../../../mock/transfer/processed/")
	type fields struct {
		repo           *ports.FileRepository
		filename       string
		destinationPah string
	}
	type args struct {
		header [][]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Test_Create_default_Files",
			fields{&filePending, fileName, "./../../../../../mock/transfer/processed/"},
			args{[][]string{[]string{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fileRepo{
				repo:           tt.fields.repo,
				filename:       tt.fields.filename,
				destinationPah: tt.fields.destinationPah,
			}
			repo.CreateDefault(tt.args.header)
			if _, err := os.Stat(fileName); err != nil {
				t.Errorf("CreateDefault() error, error create file = %v", fileName)
			}
			os.Remove(fileName)
		})
	}
}

func Test_fileRepo_GetData(t *testing.T) {
	fileName := "./../../../../../mock/transfer/pending/roster2.csv"
	validUsersT := [][]string{
		[]string{"John Doe"},
	}
	var filePending ports.FileRepository
	filePending = NewFileCSVRepo(fileName, "./../../../../../mock/transfer/processed/")
	type fields struct {
		repo           *ports.FileRepository
		filename       string
		destinationPah string
	}
	tests := []struct {
		name    string
		fields  fields
		want    [][]string
		wantErr bool
	}{
		{
			"Test_Get_Data_File",
			fields{&filePending, fileName, "./../../../../../mock/transfer/processed/"},
			validUsersT,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fileRepo{
				repo:           tt.fields.repo,
				filename:       tt.fields.filename,
				destinationPah: tt.fields.destinationPah,
			}
			got, err := repo.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileRepo_MoveFileProcessed(t *testing.T) {
	setUp()
	fileName := "./../../../../../mock/transfer/pending/roster3.csv"
	var filePending ports.FileRepository
	filePending = NewFileCSVRepo(fileName, "./../../../../../mock/transfer/processed/")
	type fields struct {
		repo           *ports.FileRepository
		filename       string
		destinationPah string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Test_Move_File_Processed",
			fields{&filePending, fileName, "./../../../../../mock/transfer/processedError/"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fileRepo{
				repo:           tt.fields.repo,
				filename:       tt.fields.filename,
				destinationPah: tt.fields.destinationPah,
			}
			if err := repo.MoveFileProcessed(); (err != nil) != tt.wantErr {
				t.Errorf("MoveFileProcessed() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat("./../../../../../mock/transfer/processed/roster3.csv"); err == nil {
				t.Errorf("MoveFileProcessed() error, error moving processed error file = roster1.csv")
			}
			if _, err := os.Stat("./../../../../../mock/transfer/processedError/roster3.csv"); err != nil {
				t.Errorf("MoveFileProcessed() error, error moving processed file = roster1.csv")
			}
			tearDown()
		})
	}
}

func Test_fileRepo_SaveInFile(t *testing.T) {
	defaultPath := "./../../../../../mock/transfer/success/employeesucess.csv"
	file1, _ := os.Create(defaultPath)
	w := csv.NewWriter(file1)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file1.Close()
	var filePending ports.FileRepository
	filePending = NewFileCSVRepo(defaultPath, "./../../../../../mock/transfer/processed/")
	validUsers := []userModel.ConfigurationHeaderExport{
		{Name: "Jane Doe", Email: "doe@test.com", Salary: "$8.45", Identifier: "5", Phone: "", Mobile: ""},
	}
	validUsersData := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
	}
	valid := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	type fields struct {
		repo           *ports.FileRepository
		filename       string
		destinationPah string
	}
	type args struct {
		oldValues [][]string
		users     []userModel.ConfigurationHeaderExport
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Test_Save_In_File_Processed",
			fields{&filePending, defaultPath, "./../../../../../mock/transfer/processed/"},
			args{validUsersData, validUsers},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fileRepo{
				repo:           tt.fields.repo,
				filename:       tt.fields.filename,
				destinationPah: tt.fields.destinationPah,
			}
			if err := repo.SaveInFile(tt.args.oldValues, tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("SaveInFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(defaultPath); err != nil {
				t.Errorf("SaveInFile() error, file created = %v", defaultPath)
			}
			got, err := repo.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, valid) {
				t.Errorf("GetData() got = %v, want %v", got, valid)
			}
			os.Remove(defaultPath)
		})
	}
}
