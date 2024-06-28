package cmd

import (
	"github.com/Method-Security/methodazure/internal/vnet"
	"github.com/spf13/cobra"
)

// InitVNetCommand initializes the `methodazure vnet` subcommand that deals with enumerating VNets in the Azure environment.
func (a *MethodAzure) InitVNetCommand() {
	vnetCmd := &cobra.Command{
		Use:   "vnet",
		Short: "Audit and command VNets",
		Long:  `Audit and command VNets`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate VNets",
		Long:  `Enumerate VNets`,
		Run: func(cmd *cobra.Command, args []string) {
			subscriptionID, err := cmd.Flags().GetString("subscription-id")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			if subscriptionID == "" {
				errorMessage := "subscription-id is not set"
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			a.AzureConfig.SubID = subscriptionID

			report, err := vnet.EnumerateVNets(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	vnetCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(vnetCmd)
}
