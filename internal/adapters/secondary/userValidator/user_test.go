package secondaryUserValidator

import (
	"testing"
	"uploader/pkg/domains/users"
)

func TestValidateFieldsAlreadyRegistered(t *testing.T) {
	userWithAlreadyRegisteredEmail := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userWithAlreadyRegisteredIdentifier := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "2"}
	values := []string{"doe@test.com", "2"}
	type args struct {
		newUser users.ConfigurationHeaderExport
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

func TestCheckUserValid(t *testing.T) {
	user := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userEmailInvalid := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: "1"}
	userIdentifierInvalid := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doet@est.com", Salary: "$10.00", Identifier: "2"}
	userWithoutEmail := users.ConfigurationHeaderExport{Name: "John Doe", Email: "", Salary: "$10.00", Identifier: "1"}
	userWithoutIdentifier := users.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: ""}
	type args struct {
		user      users.ConfigurationHeaderExport
		users     []users.ConfigurationHeaderExport
		oldValues [][]string
	}
	users := []users.ConfigurationHeaderExport{
		{Name: "Mary Jane", Email: "Mary@tes.com", Salary: "$15", Identifier: "2"},
		{Name: "Max Topperson", Email: "mary@tes.com", Salary: "$11", Identifier: "3"},
		{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "3"},
		{Name: "Jane Doe", Email: "jane_doe@test.com", Salary: "$8.45", Identifier: "5"},
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Test_Only_Valid_Email_And_Identifier",
			args{user, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			true,
			false,
		},
		{
			"Test_Only_Valid_Email",
			args{users[0], users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Only_Valid_Identifier",
			args{users[1], users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Invalid_Email",
			args{userEmailInvalid, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_Invalid_Identifier",
			args{userIdentifierInvalid, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_And_Email_Required",
			args{userWithoutEmail, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
		{
			"Test_And_Identifier_Required",
			args{userWithoutIdentifier, users, [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUserValid(tt.args.user, tt.args.users, tt.args.oldValues)
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
