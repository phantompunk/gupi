package cmd
//
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
//
// 	"github.com/spf13/cobra"
// )
//
// var pathToFile string
//
// func init() {
// 	newCmd.Flags().StringVarP(&pathToFile, "file", "f", "", "Path to file")
// 	rootCmd.AddCommand(newCmd)
// }
//
// var addFunc = func(cmd *cobra.Command, args []string) {
// 	if len(pathToFile) == 0 {
// 		errAndExit("File path required")
// 	}
//
// 	if _, err := os.Stat(pathToFile); err != nil {
// 		errAndExit("File does not exist")
// 	}
//
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		errAndExit("Failed to return user's home directory")
// 	}
//
// 	fileDir := filepath.Join(homeDir, ".gupi", "template")
// 	if _, err := os.Stat(fileDir); err != nil {
// 		fmt.Println("Creating template folder at", fileDir)
// 		os.Mkdir(fileDir, 0755)
// 	}
//
// 	f, err := os.ReadFile(pathToFile)
// 	if err != nil {
// 		errAndExit(fmt.Sprintf("failed to read from %s\n", pathToFile))
// 	}
//
// 	file_name := filepath.Base(pathToFile)
// 	fileOutPath := filepath.Join(fileDir, file_name)
// 	out, err := os.Create(fileOutPath)
// 	if err != nil {
// 		errAndExit("Failed to write file " + fileOutPath)
// 	}
// 	defer out.Close()
//
// 	out.WriteString(string(f))
// 	fmt.Printf("gupi: Template '%s' was added\n", file_name)
// }
//
// var newCmd = &cobra.Command{
// 	Use:   "new",
// 	Short: "Add a new template",
// 	Long:  "Add a template from a file path or URL",
// 	Run:   func(cmd *cobra.Command, args []string) {
// 		if len(args) < 1 {
// 			
// 		}
// 		fmt.Println(args)
// 		// gupi.AddTemplate(pathToFile)
// 	},
// }
