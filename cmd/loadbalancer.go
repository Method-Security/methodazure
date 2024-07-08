package cmd

import (
	"github.com/Method-Security/methodazure/internal/loadbalancer"
	"github.com/spf13/cobra"
)

// InitLoadBalancerCommand initializes the `methodazure loadbalancer` subcommand that deals with enumerating Load Balanccers in the Azure environment.
func (a *MethodAzure) InitLoadBalancerCommand() {
	lbCmd := &cobra.Command{
		Use:   "loadbalancer",
		Short: "Audit and command Load Balanccers",
		Long:  `Audit and command Load Balanccers`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Network Security Groups",
		Long:  `Enumerate Network Security Groups`,
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

			report := loadbalancer.EnumerateLoadBalancers(cmd.Context(), a.AzureConfig)

			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	lbCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(lbCmd)
}
