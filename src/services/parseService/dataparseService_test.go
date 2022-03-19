package parseService

import (
	"os"
	"reflect"
	"testing"
	"uploader/src/entities"
)

func TestReadCSV(t *testing.T) {
	filenameWithHeader := "../../../transfer/mock/pending/roster1.csv"
	filenameNonExistent := "../../../transfer/mock/pending/non_existent_file.csv"
	testCaseWithHeader, _ := os.Open(filenameWithHeader)
	testCaseNonExistent, _ := os.Open(filenameNonExistent)
	configHeader := entities.ConfigurationHeader{
		FullName:   []string{"name", "fullname", "full_name"},
		FirstName:  []string{"firstname", "first", "fname"},
		MiddleName: []string{"middlename", "middle"},
		LastName:   []string{"lastname", "last", "lname"},
		Email:      []string{"email", "mail", "airmail", "electronicmail", "junkmail", "junk_mail", "mail", "postalcard", "postcard"},
		Salary:     []string{"emolument", "hire", "packet", "pay", "payenvelope", "paycheck", "payment", "stipend", "wage", "salary", "rate"},
		Identifier: []string{"id", "key", "identify", "uid", "hash", "hashid", "idhash", "number", "seq", "sequence", "employeenumber", "empid"},
		Phone:      []string{"phone", "call", "dial", "ring", "telephone"},
		Mobile:     []string{"mobile"},
	}
	type args struct {
		f                *os.File
		hasHeader        bool
		configHeader     entities.ConfigurationHeader
		sizeStructHeader int
	}
	tests := []struct {
		name    string
		args    args
		want    []entities.ConfigurationHeaderExport
		want1   []entities.ConfigurationHeaderExport
		wantErr bool
	}{
		{
			"Read CSV File With Header Test Case",
			args{testCaseWithHeader, true, configHeader, 9},
			[]entities.ConfigurationHeaderExport{
				{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"},
				{Name: "Mary Jane", Email: "mary@tes.com", Salary: "$15", Identifier: "2"},
				{Name: "Max Topperson", Email: "max@test.com", Salary: "$11", Identifier: "3"},
			},
			[]entities.ConfigurationHeaderExport{
				{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "4"},
				{Name: "Jane Doe", Email: "doe@test.com", Salary: "$8.45", Identifier: "5"},
			},
			false,
		},
		{
			"Non Existent CSV File With Header Test Case",
			args{testCaseNonExistent, true, configHeader, 9},
			[]entities.ConfigurationHeaderExport{},
			[]entities.ConfigurationHeaderExport{},
			true,
		},
		{
			"Non Existent CSV File Without Header Test Case",
			args{testCaseNonExistent, false, configHeader, 9},
			[]entities.ConfigurationHeaderExport{},
			[]entities.ConfigurationHeaderExport{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReadCSV(tt.args.f, tt.args.hasHeader, tt.args.configHeader, tt.args.sizeStructHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadCSV() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
