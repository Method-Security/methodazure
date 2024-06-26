package cmd

import (
	"github.com/Method-Security/methodazure/internal/subscription"
	"github.com/spf13/cobra"
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
			report, err := subscription.EnumerateSubscriptions(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	subscriptionCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(subscriptionCmd)
}
