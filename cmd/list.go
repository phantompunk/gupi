package cmd

import (
	"fmt"
	"os"

	"github.com/phantompunk/gupi/gupi"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Long: "List all currently avaible templates",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gupi.Display(); err != nil {
			fmt.Fprint(os.Stderr, err.Error(), "\n")
			os.Exit(1)
		}
	},
}
