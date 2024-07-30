package cmd

import (
	"github.com/spf13/cobra"
)

var templateName string
var fileName string

var newCmd = &cobra.Command{
	Use: "new",
	Short: "Add a new template",
	Long: "Add a new template from a file path or URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			errAndExit("Needs a file name")
		}

		// if len(pathToFile) > 0 && len(urlToFile) > 0 {
		// 	errAndExit("Only use filepath or url, not both")
		// }
		//
		// templateName = args[0]
		// fmt.Println("Name:" + templateName)
		// fmt.Println("Path:" + pathToFile)
		// fmt.Println("URL:" + urlToFile)

		fileName = args[0]
		err := editor.New(fileName, templateName)
		if err != nil {
			errAndExit("Not able to add template")
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&templateName, "template", "t", "", "")
	// newCmd.Flags().StringVarP(&pathToFile, "file", "f", "", "Path to file")
	// newCmd.Flags().StringVarP(&urlToFile, "url", "u", "", "Url to file")
	rootCmd.AddCommand(newCmd)
}
