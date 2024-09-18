package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteFunc = func(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		errAndExit("Template name is required")
	}

	fileName := args[0]
	err := editor.Delete(fileName)
	if err != nil {
		errAndExit("Failed to remove template: " + fileName)
	}
	fmt.Printf("gupi: Deleted Template '%s'\n", fileName)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"remove", "rm"},
	Short: "Remove a template",
	Long:  "Removes a specific templates from the saved directory",
	Run:   deleteFunc,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
