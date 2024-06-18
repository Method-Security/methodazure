package cmd

import (
	"github.com/Method-Security/methodazure/internal/resourcegroup"
	"github.com/spf13/cobra"
)

// InitResourceGroupCommand initializes the `methodazure resourcegroup` subcommand that deals with enumerating
// Resource Groups in the Azure environment.
func (a *MethodAzure) InitResourceGroupCommand() {
	a.ResourceGroupCmd = &cobra.Command{
		Use:   "resourcegroup",
		Short: "Audit and command Resource Groups",
		Long:  `Audit and command Resource Groups`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Resource Groups",
		Long:  `Enumerate Resource Groups`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := resourcegroup.EnumerateResourceGroups(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	a.ResourceGroupCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(a.ResourceGroupCmd)
}
