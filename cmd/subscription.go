package cmd

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/spf13/cobra"

	"github.com/Method-Security/methodazure/internal/subscription"
)

// InitSubscriptionCommand initializes the `methodazure subscription` subcommand that deals with enumerating Azure Subscription related details.
func (a *MethodAzure) InitSubscriptionCommand() {
	subscriptionCmd := &cobra.Command{
		Use:   "subscription",
		Short: "Audit and command Azure Subscriptions",
		Long:  `Audit and command Azure Subscriptions`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Subscriptions for the provided credentials",
		Long:  `Enumerate Subscriptions for the provided credentials`,
		Run: func(cmd *cobra.Command, args []string) {
			clouds := []cloud.Configuration{cloud.AzurePublic}
			tryAllClouds, err := cmd.Flags().GetBool("try-all-clouds")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			if tryAllClouds == true {
				clouds = []cloud.Configuration{cloud.AzurePublic, cloud.AzureGovernment, cloud.AzureChina}
			} else {
				cloudName, err := cmd.Flags().GetString("cloud")
				if err != nil {
					errorMessage := err.Error()
					a.OutputSignal.ErrorMessage = &errorMessage
					a.OutputSignal.Status = 1
					return
				}
				switch cloudName {
				case "AzurePublic":
					clouds = []cloud.Configuration{cloud.AzurePublic}
				case "AzureGovernment":
					clouds = []cloud.Configuration{cloud.AzureGovernment}
				case "AzureChina":
					clouds = []cloud.Configuration{cloud.AzureChina}
				default:
					errorMessage := "Invalid cloud name provided"
					a.OutputSignal.ErrorMessage = &errorMessage
					a.OutputSignal.Status = 1
					return
				}
			}

			report, err := subscription.EnumerateSubscriptions(cmd.Context(), a.AzureConfig, clouds)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().BoolP("try-all-clouds", "t", false, "Attempt to list subscriptions in AzurePublic, Azure Government, and AzureChina; if true overrides cloud flag")
	enumerateCmd.PersistentFlags().StringP("cloud", "c", "AzurePublic", "Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina)")

	subscriptionCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(subscriptionCmd)
}
