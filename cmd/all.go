package cmd

import (
	"github.com/spf13/cobra"
	"uploader/cmd/help"
	"uploader/constants"
	"uploader/internal/adapters/primary/cmd/all"
)

// allCmd represents the allCmd command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Processes all",
	Long: `Processes all files in the (./transfer/processed/) folder
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hasH, err := help.GetBool(cmd, constants.HASHEADER)
		if err != nil {
			return err
		}
		return processAllFiles(hasH)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().BoolVarP(
		&hasHeader,
		constants.HASHEADER,
		string(constants.SHORTHASHEADER),
		true,
		"If the csv file has a col title header or not",
	)
}

func processAllFiles(hasH bool) error {
	err := primaryAll.ExecAll(hasH)
	if err != nil {
		return err
	}
	return nil
}
