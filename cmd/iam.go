package cmd

import (
	"github.com/Method-Security/methodazure/internal/iam"
	"github.com/spf13/cobra"
)

// InitIAMCommand initializes the `methodazure iam` subcommand that deals with enumerating IAM related resources in the Azure environment.
func (a *MethodAzure) InitIAMCommand() {
	iamCmd := &cobra.Command{
		Use:   "iam",
		Short: "Audit and command IAM related resources",
		Long:  `Audit and command IAM related resourcess`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate IAM related resources in a given subscription",
		Long:  `Enumerate IAM related resources; retreives roles and role assignments in a given subscription`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := iam.EnumerateIAMResources(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	iamCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(iamCmd)
}
