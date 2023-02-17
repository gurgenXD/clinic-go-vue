package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "clinic",
	Version: "0.1.0",
}

func Execute() error {
	rootCmd.AddCommand(migrationCmd)

	return rootCmd.Execute()
}
