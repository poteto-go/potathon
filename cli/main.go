package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "potathon",
		Short: "A brief description of your application",
	}

	var kubeCmd = &cobra.Command{
		Use:   "kube",
		Short: "Kubernetes related commands",
	}

	var teardownCmd = &cobra.Command{
		Use:   "teardown",
		Short: "Teardown Kubernetes resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmdStr := "kubectl delete rs calculator-pods"
			fmt.Println(cmdStr)
			exec.Command("bash", "-c", cmdStr).Run()
		},
	}

	kubeCmd.AddCommand(teardownCmd)
	rootCmd.AddCommand(kubeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
