package cmd

import (
	"fmt"
	"os"

	store "github.com/phantompunk/gupi/internal"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var editor *store.Editor
var rootCmd = &cobra.Command{
	Use:   "gupi",
	Short: "gupi is a minimal template renderer",
	Long: `Usage: gupi command [options]
A simple tool to generate and manage custom templates
`,
}

func Execute() {
	fileSystem := store.NewFileStore("/Users/rigo/.gupi/templates", afero.NewOsFs())
	editor = store.NewEditor(fileSystem)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func errAndExit(msg string) {
	fmt.Fprint(os.Stderr, msg, "\n")
	os.Exit(1)
}
