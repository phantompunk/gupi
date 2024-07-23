package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteUsage = ``

var deleteFunc = func(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		errAndExit("Template name is required")
	}
	file_name := args[0]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("Failed to read home directory")
	}

	filePath := filepath.Join(homeDir, ".gupi", "template", file_name)
	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
		fmt.Printf("gupi: Template '%s' was deleted\n", file_name)
	}
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a template",
	Long:  deleteUsage,
	Run:   deleteFunc,
}
