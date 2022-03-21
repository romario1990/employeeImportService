package userUsecase

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"uploader/constants"
	"uploader/helpers/fileHelp"
	"uploader/internal/core/domains/model/userModel"
	"uploader/internal/core/domains/repositories/userRepo"
	"uploader/internal/repositories/adapters/file/csv"
	"uploader/internal/repositories/ports"
)

func tearDown() {
	os.Remove("./../../../../../mock/transfer/error/employeeinvalid.csv")
	os.Remove("./../../../../../mock/transfer/success/employeesucess.csv")
}

func TestFormatExport(t *testing.T) {
	inputUser := []string{"Matthew", "Vargas", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeader := []string{"FirstName", "MiddleName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	inputUserFullName := []string{"Matthew Doe", "Matthew", "Doe", "matthew.doe@test.com", "2,451.45", "RT6", ""}
	inputHeaderFullName := []string{"FullName", "FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	validUserExit := userModel.ConfigurationHeaderExport{
		Name:       "Matthew Vargas Doe",
		Email:      "matthew.doe@test.com",
		Salary:     "2,451.45",
		Identifier: "RT6",
	}
	validUserExitFullName := userModel.ConfigurationHeaderExport{
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
		want userModel.ConfigurationHeaderExport
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
			tearDown()
		})
	}
}

func Test_userService_CheckUserValid(t *testing.T) {
	filename := "./../../../../../mock/transfer/pending/roster1.csv"
	var userService2 UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	filePending = csv.NewFileCSVRepo(filename, "./../../../../../mock/transfer/processed/")
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/"+constants.ERRORPATHNAME, "./../../../../../mock/transfer/processedError/")
	userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo("./../../../../../mock/"+constants.SUCCESSPATHNAME, "./../../../../../mock/transfer/processed/")
	userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	configHeader := userModel.ConfigurationHeader{
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
	userService2, _ = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, true)
	users := []userModel.ConfigurationHeaderExport{
		{Name: "Mary Jane", Email: "Mary@tes.com", Salary: "$15", Identifier: "2"},
		{Name: "Max Topperson", Email: "mary@tes.com", Salary: "$11", Identifier: "3"},
		{Name: "Alfred Donald", Email: "", Salary: "$11.5", Identifier: "3"},
		{Name: "Jane Doe", Email: "jane_doe@test.com", Salary: "$8.45", Identifier: "5"},
	}
	user := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userEmailInvalid := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: "1"}
	userWithoutEmail := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "", Salary: "$10.00", Identifier: "1"}
	userWithoutIdentifier := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doetest.com", Salary: "$10.00", Identifier: ""}
	userIdentifierInvalid := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doet@est.com", Salary: "$10.00", Identifier: "2"}
	type fields struct {
		useCases        *UserService
		userSuccessRepo userRepo.UserRepository
		userErrorRepo   userRepo.UserRepository
		filePending     ports.FileRepository
		configHeader    userModel.ConfigurationHeader
	}
	type args struct {
		user  userModel.ConfigurationHeaderExport
		users []userModel.ConfigurationHeaderExport
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Test_Only_Valid_Email_And_Identifier",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{user, users},
			true,
		},
		{
			"Test_Only_Valid_Email",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{users[0], users},
			false,
		},
		{
			"Test_Only_Valid_Identifier",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{users[1], users},
			false,
		},
		{
			"Test_Invalid_Email",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{userEmailInvalid, users},
			false,
		},
		{
			"Test_Invalid_Identifier",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{userIdentifierInvalid, users},
			false,
		},
		{
			"Test_And_Email_Required",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{userWithoutEmail, users},
			false,
		},
		{
			"Test_And_Identifier_Required",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{userWithoutIdentifier, users},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService2 := &userService{
				useCases:        tt.fields.useCases,
				userSuccessRepo: tt.fields.userSuccessRepo,
				userErrorRepo:   tt.fields.userErrorRepo,
				filePending:     tt.fields.filePending,
				configHeader:    tt.fields.configHeader,
			}
			if got := userService2.CheckUserValid(tt.args.user, tt.args.users); got != tt.want {
				t.Errorf("CheckUserValid() = %v, want %v", got, tt.want)
			}
		})
		fileHelp.MoveFile("./../../../../../mock/transfer/processed/roster1.csv", filename)
		tearDown()
	}
}

func Test_userService_ConvertDataToHeaderExport(t *testing.T) {
	filename := "./../../../../../mock/transfer/pending/roster1.csv"
	var userService2 UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	filePending = csv.NewFileCSVRepo(filename, "./../../../../../mock/transfer/processed/")
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/"+constants.ERRORPATHNAME, "./../../../../../mock/transfer/processedError/")
	userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo("./../../../../../mock/"+constants.SUCCESSPATHNAME, "./../../../../../mock/transfer/processed/")
	userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	configHeader := userModel.ConfigurationHeader{
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
	userService2, _ = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, true)
	type fields struct {
		useCases        *UserService
		userSuccessRepo userRepo.UserRepository
		userErrorRepo   userRepo.UserRepository
		filePending     ports.FileRepository
		configHeader    userModel.ConfigurationHeader
	}
	type args struct {
		data   [][]string
		header []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []userModel.ConfigurationHeaderExport
		wantErr bool
	}{
		{
			"Test_Convert_Header_export",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{[][]string{[]string{"name", "email", "salary", "identifier"}}, []string{"FullName", "Email", "Salary", "Identifier"}},
			[]userModel.ConfigurationHeaderExport{{"name", "email", "salary", "identifier", "", ""}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := &userService{
				useCases:        tt.fields.useCases,
				userSuccessRepo: tt.fields.userSuccessRepo,
				userErrorRepo:   tt.fields.userErrorRepo,
				filePending:     tt.fields.filePending,
				configHeader:    tt.fields.configHeader,
			}
			got, err := userService.ConvertDataToHeaderExport(tt.args.data, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertDataToHeaderExport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertDataToHeaderExport() got = %v, want %v", got, tt.want)
			}
			tearDown()
		})
	}
}

func Test_userService_FormatHeader(t *testing.T) {
	filename := "./../../../../../mock/transfer/pending/roster1.csv"
	var userService2 UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	filePending = csv.NewFileCSVRepo(filename, "./../../../../../mock/transfer/processed/")
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/"+constants.ERRORPATHNAME, "./../../../../../mock/transfer/processedError/")
	userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo("./../../../../../mock/"+constants.SUCCESSPATHNAME, "./../../../../../mock/transfer/processed/")
	userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	configHeader := userModel.ConfigurationHeader{
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
	userService2, _ = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, true)
	headerFirstEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone"}
	headerSecondEntry := []string{"f. name", "l. name", "email,", "wage", "emp id", "phone", "test"}
	headerFirstOutput := []string{"FirstName", "LastName", "Email", "Salary", "Identifier", "Phone"}
	type fields struct {
		useCases        *UserService
		userSuccessRepo userRepo.UserRepository
		userErrorRepo   userRepo.UserRepository
		filePending     ports.FileRepository
		configHeader    userModel.ConfigurationHeader
	}
	type args struct {
		header           []string
		configHeader     userModel.ConfigurationHeader
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
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{headerFirstEntry, configHeader, 9},
			headerFirstOutput,
		},
		{
			"Test_Unmapped_Column",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{headerSecondEntry, configHeader, 9},
			headerFirstOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := &userService{
				useCases:        tt.fields.useCases,
				userSuccessRepo: tt.fields.userSuccessRepo,
				userErrorRepo:   tt.fields.userErrorRepo,
				filePending:     tt.fields.filePending,
				configHeader:    tt.fields.configHeader,
			}
			if got := userService.FormatHeader(tt.args.header, tt.args.configHeader, tt.args.sizeStructHeader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatHeader() = %v, want %v", got, tt.want)
			}
			tearDown()
		})
	}
}

func Test_userService_GetField(t *testing.T) {
	filename := "./../../../../../mock/transfer/pending/roster1.csv"
	var userService2 UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	filePending = csv.NewFileCSVRepo(filename, "./../../../../../mock/transfer/processed/")
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/"+constants.ERRORPATHNAME, "./../../../../../mock/transfer/processedError/")
	userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo("./../../../../../mock/"+constants.SUCCESSPATHNAME, "./../../../../../mock/transfer/processed/")
	userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	configHeader := userModel.ConfigurationHeader{
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
	userService2, _ = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, true)
	type fields struct {
		useCases *UserService
	}
	inputField := "FullName"
	inputFieldNonExistent := "Instagram"
	output := []string{"name", "fullname"}

	type args struct {
		v     userModel.ConfigurationHeader
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
			fields{&userService2},
			args{configHeader, inputField},
			output,
			false,
		},
		{
			"Test_Non-existent_Field_Setting_In_Search",
			fields{&userService2},
			args{configHeader, inputFieldNonExistent},
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := &userService{
				useCases: tt.fields.useCases,
			}
			got, err := userService.GetField(tt.args.v, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetField() got = %v, want %v", got, tt.want)
			}
			tearDown()
		})
	}
}

func Test_userService_SaveUsers(t *testing.T) {
	filename := "./../../../../../mock/transfer/pending/roster1.csv"
	var userService2 UserService
	var filePending ports.FileRepository
	var userSuccessRepo userRepo.UserRepository
	var userErrorRepo userRepo.UserRepository
	filePending = csv.NewFileCSVRepo(filename, "./../../../../../mock/transfer/processed/")
	var fileRepoE ports.FileRepository
	fileRepoE = csv.NewFileCSVRepo("./../../../../../mock/"+constants.ERRORPATHNAME, "./../../../../../mock/transfer/processedError/")
	userErrorRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoE)
	var fileRepoS ports.FileRepository
	fileRepoS = csv.NewFileCSVRepo("./../../../../../mock/"+constants.SUCCESSPATHNAME, "./../../../../../mock/transfer/processed/")
	userSuccessRepo = userRepo.NewUserRepo(constants.HEADER, fileRepoS)
	configHeader := userModel.ConfigurationHeader{
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
	userService2, _ = NewUserService(filePending, userSuccessRepo, userErrorRepo, configHeader, true)
	validUsers := []userModel.ConfigurationHeaderExport{
		{Name: "Jane Doe", Email: "doe@test.com", Salary: "$8.45", Identifier: "5", Phone: "", Mobile: ""},
	}
	valid := [][]string{
		{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		{"Jane Doe", "doe@test.com", "$8.45", "5", "", ""},
	}
	invalidUsers := []userModel.ConfigurationHeaderExport{
		{Name: "Jane Doe", Email: "", Salary: "$8.45", Identifier: "5", Phone: "", Mobile: ""},
	}
	invalid := [][]string{
		{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"},
		{"Jane Doe", "", "$8.45", "5", "", ""},
	}
	type fields struct {
		useCases        *UserService
		userSuccessRepo userRepo.UserRepository
		userErrorRepo   userRepo.UserRepository
		filePending     ports.FileRepository
		configHeader    userModel.ConfigurationHeader
	}
	type args struct {
		users     []userModel.ConfigurationHeaderExport
		userValid bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test_UserValid",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{validUsers, true},
			false,
		},
		{
			"Test_UserInvalid",
			fields{&userService2, userSuccessRepo, userErrorRepo, filePending, configHeader},
			args{invalidUsers, false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := &userService{
				useCases:        tt.fields.useCases,
				userSuccessRepo: tt.fields.userSuccessRepo,
				userErrorRepo:   tt.fields.userErrorRepo,
				filePending:     tt.fields.filePending,
				configHeader:    tt.fields.configHeader,
			}
			if err := userService.SaveUsers(tt.args.users, tt.args.userValid); (err != nil) != tt.wantErr {
				t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("AQUII", tt.args.userValid)
			if tt.args.userValid {
				got, err := userService.userSuccessRepo.GetList()
				if (err != nil) != tt.wantErr {
					t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, valid) {
					t.Errorf("SaveUsers() got = %v, want %v", got, valid)
				}
			} else {
				got, err := userService.userErrorRepo.GetList()
				if (err != nil) != tt.wantErr {
					t.Errorf("SaveUsers() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, invalid) {
					t.Errorf("SaveUsers() got = %v, want %v", got, invalid)
				}
			}
		})
	}
	tearDown()
}

func Test_validateFieldsAlreadyRegistered(t *testing.T) {
	userWithAlreadyRegisteredEmail := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "1"}
	userWithAlreadyRegisteredIdentifier := userModel.ConfigurationHeaderExport{Name: "John Doe", Email: "doe@test.com", Salary: "$10.00", Identifier: "2"}
	values := []string{"doe@test.com", "2"}

	type args struct {
		newUser userModel.ConfigurationHeaderExport
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
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFieldsAlreadyRegistered(tt.args.newUser, tt.args.user); got != tt.want {
				t.Errorf("validateFieldsAlreadyRegistered() = %v, want %v", got, tt.want)
			}
			tearDown()
		})
	}
}
