package cmd

import (
	"github.com/spf13/cobra"
)

var templateName string
var fileName string
var outputPath string

var newCmd = &cobra.Command{
	Use: "new",
	Short: "Add a new template",
	Long: "Add a new template from a file path or URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			errAndExit("Needs a file name")
		}

		fileName = args[0]
		err := editor.New(fileName, outputPath, templateName)
		if err != nil {
			errAndExit("Not able to add template" + err.Error())
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&templateName, "template", "t", "", "")
	newCmd.Flags().StringVarP(&outputPath, "output", "o", ".", "")
	rootCmd.AddCommand(newCmd)
}
