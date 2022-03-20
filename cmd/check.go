package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"uploader/cmd/help"
	"uploader/constants"
	primaryCheck "uploader/internal/adapters/primary/cmd/check"
)

// checkCmd represents the checkCmd command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "The check command process a specific file",
	Long:  "The check command process a specific file with validations and generates new valid files in invalid ones",
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, err := help.GetString(cmd, constants.FILE)
		if err != nil {
			return err
		}
		hasH, err := help.GetBool(cmd, constants.HASHEADER)
		if err != nil {
			return err
		}
		fileType, err := help.GetString(cmd, constants.FILETYPE)
		if err != nil {
			return err
		}
		return checkFile(fileName, hasH, fileType)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(
		&file,
		constants.FILE,
		string(constants.SHORTFILE),
		"",
		"The csv file to be validated",
	)
	checkCmd.Flags().BoolVarP(
		&hasHeader,
		constants.HASHEADER,
		string(constants.SHORTHASHEADER),
		true,
		"If the csv file has a col title header or not",
	)
	checkCmd.Flags().StringVarP(
		&fileType,
		constants.FILETYPE,
		string(constants.SHORTFILETYPE),
		"csv",
		"sets the file type",
	)
	checkCmd.MarkFlagRequired(constants.FILE)
}

func checkFile(filename string, hasH bool, fileType string) error {
	err := primaryCheck.Exec(filename, hasH, strings.ToLower(fileType))
	if err != nil {
		return err
	}
	return nil
}
