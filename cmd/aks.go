package cmd

import (
	"github.com/Method-Security/methodazure/internal/aks"
	"github.com/spf13/cobra"
)

// InitAKSCommand initializes the `methodazure aks` subcommand that deals with enumerating AKS clusters in the Azure environment.
func (a *MethodAzure) InitAKSCommand() {
	aksCmd := &cobra.Command{
		Use:   "aks",
		Short: "Audit and command AKS clusters",
		Long:  `Audit and command AKS clusters`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate AKS clusters",
		Long:  `Enumerate AKS clusters`,
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

			report, err := aks.EnumerateAKSClusters(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	aksCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(aksCmd)
}
