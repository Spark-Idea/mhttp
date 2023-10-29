package cmd

import (
	"fmt"
	"mhttp/defaults"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows application version",
	Long:  `Shows application version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(defaults.Defaults.AppVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
