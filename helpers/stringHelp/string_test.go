package stringHelp

import "testing"

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
