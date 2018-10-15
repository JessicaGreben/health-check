package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print health-check version",
	Run:   getVersion,
}

func getVersion(cmd *cobra.Command, args []string) {
	fmt.Println("Current API version:", Version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
