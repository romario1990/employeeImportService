package secondaryArrayHelp

import (
	"testing"
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
