package secondaryHeaderProcessor

import (
	"reflect"
	"testing"
	"uploader/pkg/domains/users"
)

func TestFormatHeader(t *testing.T) {
	headerFirstEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone"}
	headerSecondEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone", "test"}
	headerFirstOutput := []string{"FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	configHeader := users.ConfigurationHeader{
		FullName:   []string{"name", "fullname"},
		FirstName:  []string{"firstname", "first", "fname"},
		MiddleName: []string{"middlename", "middle"},
		LastName:   []string{"lastname", "last", "lname"},
		Email:      []string{"email", "mail", "airmail", "electronicmail", "junkmail", "mail", "postalcard", "postcard"},
		Salary:     []string{"emolument", "hire", "packet", "pay", "payenvelope", "paycheck", "payment", "stipend", "wage", "salary", "rate"},
		Identifier: []string{"id", "key", "identify", "uid", "hash", "hashid", "idhash", "number", "seq", "sequence", "employeenumber", "empid"},
		Phone:      []string{"phone", "call", "dial", "ring", "telephone"},
		Mobile:     []string{"mobile"},
	}
	type args struct {
		header           []string
		configHeader     users.ConfigurationHeader
		sizeStructHeader int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test_Header_Identification",
			args{headerFirstEntry, configHeader, 9},
			headerFirstOutput,
		},
		{
			"Test_Unmapped_Column",
			args{headerSecondEntry, configHeader, 9},
			headerFirstOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatHeader(tt.args.header, tt.args.configHeader, tt.args.sizeStructHeader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCSVExport(t *testing.T) {
	inputUser := []string{"Matthew", "Vargas", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeader := []string{"FirstName", "MiddleName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	inputUserFullName := []string{"Matthew Doe", "Matthew", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeaderFullName := []string{"FullName", "FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	validUserExit := users.ConfigurationHeaderExport{
		Name:       "Matthew Vargas Doe",
		Email:      "matthew.doe@test.com",
		Salary:     "2,451.45",
		Identifier: "RT6",
	}
	validUserExitFullName := users.ConfigurationHeaderExport{
		Name:       "Matthew Doe",
		Email:      "matthew.doe@test.com",
		Salary:     "2,451.45",
		Identifier: "RT6",
	}

	type args struct {
		row    []string
		header []string
	}
	tests := []struct {
		name string
		args args
		want users.ConfigurationHeaderExport
	}{
		{
			"Test_To_Format_String_Array_In_Struct_Header_FirstName_MiddleName_And_LastName",
			args{inputUser, inputHeader},
			validUserExit,
		},
		{
			"Test_To_Format_String_Array_In_Struct_Header_FullName",
			args{inputUserFullName, inputHeaderFullName},
			validUserExitFullName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCSVExport(tt.args.row, tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatCSVExport() = %v, want %v", got, tt.want)
			}
		})
	}
}
