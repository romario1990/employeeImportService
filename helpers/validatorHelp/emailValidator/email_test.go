package secondaryEmailValidator

import "testing"

func TestValidateEmail(t *testing.T) {
	validEmail := "romario.getulio@gmail.com"
	invalidEmail := "romario.getuliogmail.com"
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_Valid_Email",
			args{validEmail},
			true,
		},
		{
			"Test_Valid_Email",
			args{invalidEmail},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.args.email); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
