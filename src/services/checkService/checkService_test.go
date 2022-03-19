package checkService

import (
	"testing"
	"uploader/src/entities"
)

func TestCheckUserValid(t *testing.T) {
	user := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	users := []entities.ConfigurationHeaderExport{
		{Name: "Mary Jane", Email: "Mary@tes.com", Salary: "$15", Identifier: "2"},
		{Name: "Max Topperson", Email: "mary@tes.com", Salary: "$11", Identifier: "3"},
		{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "3"},
		{Name: "Jane Doe", Email: "jane_doe@test.com", Salary: "$8.45", Identifier: "5"},
	}
	userEmailInvalid := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: "1"}
	userIdentifierInvalid := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doet@est.com", Salary: "$10.00", Identifier: "2"}
	userWithoutEmail := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "", Salary: "$10.00", Identifier: "1"}
	userWithoutIdentifier := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: ""}
	type args struct {
		user  entities.ConfigurationHeaderExport
		users []entities.ConfigurationHeaderExport
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Test_Only_Valid_Email_And_Identifier",
			args{user, users},
			true,
			false,
		},
		{
			"Test_Only_Valid_Email",
			args{users[0], users},
			false,
			false,
		},
		{
			"Test_Only_Valid_Identifier",
			args{users[1], users},
			false,
			false,
		},
		{
			"Test_Invalid_Email",
			args{userEmailInvalid, users},
			false,
			false,
		},
		{
			"Test_Invalid_Identifier",
			args{userIdentifierInvalid, users},
			false,
			false,
		},
		{
			"Test_And_Email_Required",
			args{userWithoutEmail, users},
			false,
			false,
		},
		{
			"Test_And_Identifier_Required",
			args{userWithoutIdentifier, users},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUserValid(tt.args.user, tt.args.users)
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

func TestValidateFieldsAlreadyRegistered(t *testing.T) {
	userWithAlreadyRegisteredEmail := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userWithAlreadyRegisteredIdentifier := entities.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "2"}
	values := []string{"doe@test.com", "2"}
	type args struct {
		newUser entities.ConfigurationHeaderExport
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
			if got := ValidateFieldsAlreadyRegistered(tt.args.newUser, tt.args.user); got != tt.want {
				t.Errorf("ValidateFieldsAlreadyRegistered() = %v, want %v", got, tt.want)
			}
		})
	}
}
