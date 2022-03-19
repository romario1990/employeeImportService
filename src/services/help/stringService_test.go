package help

import (
	"reflect"
	"testing"
	"uploader/src/entities"
)

func TestRemoveSpecialChar(t *testing.T) {
	input := "!@#$%¨&*()_+``{`}^:?>< ¹²³£¢¬§ªºf. name"
	output := "fname"
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test_Removes_Special_Characters_And_Whitespace",
			args{input},
			output,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveSpecialChar(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveSpecialChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveSpecialChar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardizeColumn(t *testing.T) {
	input := "F.NamE"
	output := "fname"
	type args struct {
		column string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test_Defaults_To_All_Lowercase",
			args{input},
			output,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StandardizeColumn(tt.args.column)
			if (err != nil) != tt.wantErr {
				t.Errorf("StandardizeColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StandardizeColumn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	inputStringList := []string{"emolument", "hire", "packet", "pay", "payenvelope", "paycheck", "payment", "stipend", "wage", "salary", "rate"}
	field := "wage"
	type args struct {
		columnName string
		list       []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_Existing_String_In_String_Array",
			args{field, inputStringList},
			true,
		},
		{
			"Test_Non-Existent_String_In_String_Array",
			args{"", inputStringList},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.columnName, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
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
			"Test_Search_Field_Configuration",
			args{configHeader, inputField},
			output,
			false,
		},
		{
			"Test_Search_Field_Configuration",
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
