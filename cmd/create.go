package cmd

import (
	"github.com/spf13/cobra"
)

var sampleTemplate bool
var pathToTemplate string

var createFunc = func(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errAndExit("Needs a file name")
	}

	templateName = args[0]
	err := editor.Create(templateName, pathToTemplate, sampleTemplate)
	if err != nil {
		errAndExit("Not able to create a template " + err.Error())
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new template",
	Run:   createFunc,
	Args: cobra.ExactArgs(1),
}

func init() {
	createCmd.Flags().BoolVarP(&sampleTemplate, "sample", "s", false, "Create from a sample template")
	createCmd.Flags().StringVarP(&pathToTemplate, "file", "f", "", "Filename, existing template name, or URL to file to create the template")
	rootCmd.AddCommand(createCmd)
}
