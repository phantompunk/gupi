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
	Short: "Render a new file based on a template",
	Long: "Render a new file based on a template",
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
	Args: cobra.ExactArgs(1),
}

func init() {
	newCmd.Flags().StringVarP(&templateName, "template", "t", "", "Template name to render")
	newCmd.Flags().StringVarP(&outputPath, "output", "o", ".", "Write to output path")
	newCmd.Flags().BoolVarP(&nextWeek, "next-week", "n", false, "Modify internal date 1 week ahead")
	rootCmd.AddCommand(newCmd)
}
