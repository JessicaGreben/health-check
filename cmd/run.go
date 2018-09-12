package cmd

import (
	"github.com/jessicagreben/pratique/pkg/checks"
	"github.com/spf13/cobra"
	// "k8s.io/api/core/v1"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Check the health of the current Kubernetes cluster",
	Long:  `Run a health check against a Kubernetes cluster. A report is generated describing the results of the checks.`,
	Run:   runHealthCheck,
}

func runHealthCheck(cmd *cobra.Command, args []string) {
	checks.RunPodHealthChecks()
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
