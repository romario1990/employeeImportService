package users

import (
	"reflect"
	"testing"
)

func Test_userRepo_GetField(t *testing.T) {
	configHeader := ConfigurationHeader{
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
	var userRepo UserRepository
	userRepo = NewUserRepo()
	type fields struct {
		repo UserRepository
	}
	type args struct {
		v     ConfigurationHeader
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Test_Existing_Field_Configuration_In_Lookup",
			fields{userRepo},
			args{configHeader, inputField},
			output,
			false,
		},
		{
			"Test_Non-existent_Field_Setting_In_Search",
			fields{userRepo},
			args{configHeader, inputFieldNonExistent},
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.repo.GetField(tt.args.v, tt.args.field)
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

func TestFormatExport(t *testing.T) {
	inputUser := []string{"Matthew", "Vargas", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeader := []string{"FirstName", "MiddleName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	inputUserFullName := []string{"Matthew Doe", "Matthew", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeaderFullName := []string{"FullName", "FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	validUserExit := ConfigurationHeaderExport{
		Name:       "Matthew Vargas Doe",
		Email:      "matthew.doe@test.com",
		Salary:     "2,451.45",
		Identifier: "RT6",
	}
	validUserExitFullName := ConfigurationHeaderExport{
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
		want ConfigurationHeaderExport
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
			if got := FormatExport(tt.args.row, tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatExport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_FormatHeader(t *testing.T) {
	headerFirstEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone"}
	headerSecondEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone", "test"}
	headerFirstOutput := []string{"FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	configHeader := ConfigurationHeader{
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
	var userRepo UserRepository
	userRepo = NewUserRepo()
	type fields struct {
		repo UserRepository
	}
	type args struct {
		header           []string
		configHeader     ConfigurationHeader
		sizeStructHeader int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			"Test_Header_Identification",
			fields{userRepo},
			args{headerFirstEntry, configHeader, 9},
			headerFirstOutput,
		},
		{
			"Test_Unmapped_Column",
			fields{userRepo},
			args{headerSecondEntry, configHeader, 9},
			headerFirstOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.repo.FormatHeader(tt.args.header, tt.args.configHeader, tt.args.sizeStructHeader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatHeader123213() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFieldsAlreadyRegistered(t *testing.T) {
	userWithAlreadyRegisteredEmail := ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userWithAlreadyRegisteredIdentifier := ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "2"}
	values := []string{"doe@test.com", "2"}
	type args struct {
		newUser ConfigurationHeaderExport
		user    []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_Email_Already_Registered",
			args{userWithAlreadyRegisteredEmail, values},
			true,
		},
		{
			"Test_Identifier_Already_Registered",
			args{userWithAlreadyRegisteredIdentifier, values},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFieldsAlreadyRegistered(tt.args.newUser, tt.args.user); got != tt.want {
				t.Errorf("validateFieldsAlreadyRegistered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_CheckUserValid(t *testing.T) {
	user := ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userEmailInvalid := ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: "1"}
	userIdentifierInvalid := ConfigurationHeaderExport{Name: "John Doe", Email: "doet@est.com", Salary: "$10.00", Identifier: "2"}
	userWithoutEmail := ConfigurationHeaderExport{Name: "John Doe", Email: "", Salary: "$10.00", Identifier: "1"}
	userWithoutIdentifier := ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: ""}
	users := []ConfigurationHeaderExport{
		{Name: "Mary Jane", Email: "Mary@tes.com", Salary: "$15", Identifier: "2"},
		{Name: "Max Topperson", Email: "mary@tes.com", Salary: "$11", Identifier: "3"},
		{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "3"},
		{Name: "Jane Doe", Email: "jane_doe@test.com", Salary: "$8.45", Identifier: "5"},
	}
	var userRepo UserRepository
	userRepo = NewUserRepo()
	type fields struct {
		repo UserRepository
	}
	type args struct {
		user      ConfigurationHeaderExport
		users     []ConfigurationHeaderExport
		oldValues [][]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Test_Only_Valid_Email_And_Identifier",
			fields{userRepo},
			args{user, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			true,
			false,
		},
		{
			"Test_Only_Valid_Email",
			fields{userRepo},
			args{users[0], users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Only_Valid_Identifier",
			fields{userRepo},
			args{users[1], users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Invalid_Email",
			fields{userRepo},
			args{userEmailInvalid, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Invalid_Identifier",
			fields{userRepo},
			args{userIdentifierInvalid, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_And_Email_Required",
			fields{userRepo},
			args{userWithoutEmail, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_And_Identifier_Required",
			fields{userRepo},
			args{userWithoutIdentifier, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.repo.CheckUserValid(tt.args.user, tt.args.users, tt.args.oldValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckUserValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
