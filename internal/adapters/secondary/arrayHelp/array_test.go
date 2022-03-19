package secondaryArrayHelp

import (
	"reflect"
	"testing"
	"uploader/entities"
)

func TestContains(t *testing.T) {
	array := []string{"email", "mail", "airmail", "electronicmail", "junkmail", "mail", "postalcard", "postcard"}
	existingFieldArray := "electronicmail"
	noFieldArray := "golang"
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_Existing_Field_In_Array",
			args{array, existingFieldArray},
			true,
		},
		{
			"Test_No_Field_In_Array",
			args{array, noFieldArray},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetField(t *testing.T) {
	configHeader := entities.ConfigurationHeader{
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
	inputField := "FullName"
	inputFieldNonExistent := "Instagram"
	output := []string{"name", "fullname"}
	type args struct {
		v     entities.ConfigurationHeader
		field string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Test_Existing_Field_Configuration_In_Lookup",
			args{configHeader, inputField},
			output,
			false,
		},
		{
			"Test_Non-existent_Field_Setting_In_Search",
			args{configHeader, inputFieldNonExistent},
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetField(tt.args.v, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetField() got = %v, want %v", got, tt.want)
			}
		})
	}
}
