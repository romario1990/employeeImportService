package secondaryReadFileCSV

import (
	"os"
	"reflect"
	"testing"
	"uploader/pkg/domains/users"
)

func TestReadFile(t *testing.T) {
	filenameWithHeader := "../../../../../../transfer/mock/transfer/pending/roster1.csv"
	filenameNonExistent := "../../../../../../transfer/mock/transfer/pending/non_existent_file.csv"
	testCaseWithHeader, _ := os.Open(filenameWithHeader)
	testCaseNonExistent, _ := os.Open(filenameNonExistent)
	configHeader := users.ConfigurationHeader{
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
		configHeader     users.ConfigurationHeader
		sizeStructHeader int
	}
	tests := []struct {
		name    string
		args    args
		want    []users.ConfigurationHeaderExport
		want1   []users.ConfigurationHeaderExport
		wantErr bool
	}{
		{
			"Read csv File With Header Test Case",
			args{testCaseWithHeader, true, configHeader, 9},
			[]users.ConfigurationHeaderExport{
				{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"},
				{Name: "Mary Jane", Email: "mary@tes.com", Salary: "$15", Identifier: "2"},
				{Name: "Max Topperson", Email: "max@test.com", Salary: "$11", Identifier: "3"},
			},
			[]users.ConfigurationHeaderExport{
				{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "4"},
				{Name: "Jane Doe", Email: "doe@test.com", Salary: "$8.45", Identifier: "5"},
			},
			false,
		},
		{
			"Non Existent csv File With Header Test Case",
			args{testCaseNonExistent, true, configHeader, 9},
			[]users.ConfigurationHeaderExport{},
			[]users.ConfigurationHeaderExport{},
			true,
		},
		{
			"Non Existent csv File Without Header Test Case",
			args{testCaseNonExistent, false, configHeader, 9},
			[]users.ConfigurationHeaderExport{},
			[]users.ConfigurationHeaderExport{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReadFile(tt.args.f, tt.args.hasHeader, tt.args.configHeader, tt.args.sizeStructHeader)
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
