package cmd

import (
	"github.com/spf13/cobra"
)

var templateName string
var fileName string
var outputPath string
var nextWeek bool

var newCmd = &cobra.Command{
	Use: "new",
	Short: "Create a file based on a template",
	Long: "Create a file based on a template",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			errAndExit("Needs a file name")
		}

		fileName = args[0]
		err := editor.New(fileName, outputPath, templateName, nextWeek)
		if err != nil {
			errAndExit("Not able to add template" + err.Error())
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&templateName, "template", "t", "", "")
	newCmd.Flags().StringVarP(&outputPath, "output", "o", ".", "")
	newCmd.Flags().BoolVarP(&nextWeek, "next-week", "n", false, "Modify internal date 1 week ahead")
	rootCmd.AddCommand(newCmd)
}
