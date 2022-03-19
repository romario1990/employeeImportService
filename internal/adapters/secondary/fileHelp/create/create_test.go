package secondaryCreateFile

import (
	"os"
	"testing"
	"uploader/constants"
)

func TestCreateDefaultFiles(t *testing.T) {
	defaultPath := "./../../../../../transfer/mock/"
	type args struct {
		header      [][]string
		defaultPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test_Creation_Of_Standard_Files_Of_Success_And_Error",
			args{constants.HEADER, defaultPath},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDefaultFiles(tt.args.header, tt.args.defaultPath); (err != nil) != tt.wantErr {
				t.Errorf("CreateDefaultFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(defaultPath + constants.SUCCESSPATHNAME); err != nil {
				t.Errorf("CreateDefaultFiles() error, success file not created = %v", constants.SUCCESSNAMEFILE)
			}
			if _, err := os.Stat(defaultPath + constants.ERRORPATHNAME); err != nil {
				t.Errorf("CreateDefaultFiles() error, error file not created = %v", constants.ERRORNAMEFILE)
			}
			e := os.Remove(defaultPath + constants.SUCCESSPATHNAME)
			if e != nil {
				t.Errorf("CreateDefaultFiles() error, success file not deleted = %v", constants.ERRORNAMEFILE)
			}
			e = os.Remove(defaultPath + constants.ERRORPATHNAME)
			if e != nil {
				t.Errorf("CreateDefaultFiles() error, error file not deleted = %v", constants.ERRORNAMEFILE)
			}
		})
	}
}
