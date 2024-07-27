package cmd

import (
	"fmt"

	"github.com/phantompunk/gupi/gupi"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a template",
	Long:  "Edits an existing template",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			errAndExit("Needs a file name")
		}

		templateName := args[0]
		err := gupi.EditTemplate(templateName)
		if err != nil {
			errAndExit("Unable to edit template")
		}
		fmt.Printf("gupi: Template '%s' was edited\n", templateName)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
