package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Fail bool

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.PersistentFlags().BoolVarP(&Fail, "fail", "f", false, "Fail fast with no output when check fails with a non-0 exit code")
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check kube status",
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()

		list, err := getDeploymentsList()

		if err != nil {
			return err
		}

		var healthStatus = true

		fmt.Fprintln(out, "| Deployment | Ready | Required |")
		fmt.FPrintln(out, "| --- | --- | --- |")

		for _, d := range list.Items {
			fmt.FPrintf(out, "| %s | %d | %d |\n", d.Name, d.Status.ReadyReplicas, d.Status.Replicas)

			if d.Status.ReadyReplicas < d.Status.Replicas {
				healthStatus = false
			}
		}

		fmt.FPrint(out, "\n")

		if healthStatus {
			fmt.FPrintln("The Cluster is **Healthy**")
		} else {
			fmt.FPrintln("The Cluster is **Unhealthy**")

			if Fail {
				os.Exit(1)
			}
		}

		return nil
	},
}
