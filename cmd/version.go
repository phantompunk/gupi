package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	build   = "???"
	version = "???"
	short   = false
)

var versionUsage = `Print the app version and build info for the current context.

Usage: gupi version [options]

Options:
  --short  If true, print just the version number. Default false.
`

func init() {
	versionCmd.Flags().BoolVarP(&short, "short", "s", short, "")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the app version and build info of gupi",
	Long:  versionUsage,
	Run:   versionFunc,
}

var versionFunc = func(cmd *cobra.Command, args []string) {
	if short {
		fmt.Printf("gupi version: v%s", version)
	} else {
		fmt.Printf("gupi version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

