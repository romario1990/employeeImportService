package help

import (
	"github.com/spf13/cobra"
)

func GetString(c *cobra.Command, flagName string) (string, error) {
	value, err := c.Flags().GetString(flagName)
	if err != nil {
		return "", err
	}
	return value, nil
}

func GetBool(c *cobra.Command, flagName string) (bool, error) {
	value, err := c.Flags().GetBool(flagName)
	if err != nil {
		return false, err
	}
	return value, nil
}
