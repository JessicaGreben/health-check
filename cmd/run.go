package cmd

import (
	"fmt"

	"github.com/jessicagreben/health-check/pkg/checks"
	"github.com/jessicagreben/health-check/pkg/report"
	"github.com/jessicagreben/health-check/pkg/types"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Check the health of the current Kubernetes cluster",
	Long:  `Run a health check against a Kubernetes cluster. A report is generated describing the results of the checks.`,
	Run: func(cmd *cobra.Command, args []string) {

		podResults, err := checks.Pods()
		if err != nil {
			fmt.Print("checks.Pods err: ", err)
		}

		deployResults, err := checks.Deploys()
		if err != nil {
			fmt.Print("checks.Pods err: ", err)
		}

		results := types.Results{
			Pods:    podResults,
			Deploys: deployResults,
		}

		if err := report.Render(results); err != nil {
			fmt.Print("report.Render err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
