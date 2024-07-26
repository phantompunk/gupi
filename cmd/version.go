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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the app version and build info of gupi",
	Long:  "Print the app version and build info for the current context",
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Printf("gupi version: v%s", version)
		} else {
			fmt.Printf("gupi version: v%s, build: %s", version, build)
		}
		os.Exit(0)
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&short, "short", "s", short, "")
	rootCmd.AddCommand(versionCmd)
}
