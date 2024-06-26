package cmd

import (
	"github.com/Method-Security/methodazure/internal/resourcegroup"
	"github.com/spf13/cobra"
)

// InitResourceGroupCommand initializes the `methodazure resourcegroup` subcommand that deals with enumerating
// Resource Groups in the Azure environment.
func (a *MethodAzure) InitResourceGroupCommand() {
	resourceGroupCmd := &cobra.Command{
		Use:   "resourcegroup",
		Short: "Audit and command Resource Groups",
		Long:  `Audit and command Resource Groups`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Resource Groups",
		Long:  `Enumerate Resource Groups`,
		Run: func(cmd *cobra.Command, args []string) {
			subscriptionID, err := cmd.Flags().GetString("subscription-id")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			if subscriptionID == "" {
				errorMessage := "subscription-id is not set"
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.AzureConfig.SubID = subscriptionID

			report, err := resourcegroup.EnumerateResourceGroups(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	resourceGroupCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(resourceGroupCmd)
}
