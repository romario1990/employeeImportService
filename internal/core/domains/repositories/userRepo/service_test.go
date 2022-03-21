package userRepo

import (
	csv2 "encoding/csv"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/repositories/adapters/file/csv"
	"uploader/internal/repositories/ports"
)

func setUp() {
	defaultPath := "./../../../../../mock/transfer/pending/roster3.csv"
	file1, _ := os.Create(defaultPath)
	w := csv2.NewWriter(file1)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file1.Close()
}

func tearDown() {
	os.Remove("./../../../../../mock/transfer/pending/roster3.csv")
	os.Remove("./../../../../../mock/transfer/processed/roster3.csv")
	os.Remove("./../../../../../mock/transfer/processedError/roster3.csv")
}

func Test_userRepo_GetList(t *testing.T) {
	var userR UserRepository
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/transfer/pending/roster2.csv", "./../../../../../mock//transfer/processedError/")
	userR = NewUserRepo(constants.HEADER, fileRepoE)

	type fields struct {
		userRepo *UserRepository
		fileRepo ports.FileRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    [][]string
		wantErr bool
	}{
		{
			"Test_Get_List_User",
			fields{&userR, fileRepoE},
			[][]string{[]string{"John Doe"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				userRepo: tt.fields.userRepo,
				fileRepo: tt.fields.fileRepo,
			}
			got, err := repo.GetList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_MoveFileProcessed(t *testing.T) {
	setUp()
	var userR UserRepository
	fileName := "./../../../../../mock/transfer/processed/roster3.csv"
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo(fileName, "./../../../../../mock//transfer/processedError/")
	userR = NewUserRepo(constants.HEADER, fileRepoE)
	type fields struct {
		userRepo *UserRepository
		fileRepo ports.FileRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Test_Get_List_User",
			fields{&userR, fileRepoE},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				userRepo: tt.fields.userRepo,
				fileRepo: tt.fields.fileRepo,
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

func Test_userRepo_Save(t *testing.T) {
	defaultPath := "./../../../../../mock/transfer/success/employeesucess.csv"
	file1, _ := os.Create(defaultPath)
	w := csv2.NewWriter(file1)
	for _, row := range constants.HEADER {
		_ = w.Write(row)
	}
	w.Flush()
	file1.Close()
	var userR UserRepository
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo(defaultPath, "./../../../../../mock//transfer/processed/")
	userR = NewUserRepo(constants.HEADER, fileRepoS)
	validUsers := []userModel.ConfigurationHeaderExport{
		{Name: "Jane Doe", Email: "doe@test.com", Salary: "$8.45", Identifier: "5", Phone: "", Mobile: ""},
	}
	valid := [][]string{
		[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		[]string{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	type fields struct {
		userRepo *UserRepository
		fileRepo ports.FileRepository
	}
	type args struct {
		usersList []userModel.ConfigurationHeaderExport
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Test_Save_Users",
			fields{&userR, fileRepoS},
			args{validUsers},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				userRepo: tt.fields.userRepo,
				fileRepo: tt.fields.fileRepo,
			}
			if err := repo.Save(tt.args.usersList); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(defaultPath); err != nil {
				t.Errorf("SaveInFile() error, file created = %v", defaultPath)
			}
			got, err := repo.fileRepo.GetData()
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
