package cmd

import (
	"github.com/Method-Security/methodazure/internal/dns"
	"github.com/spf13/cobra"
)

// InitDNSCommand initializes the `methodazure dns` subcommand that deals with enumerating DNS related resources in the Azure environment.
func (a *MethodAzure) InitDNSCommand() {
	dnsCmd := &cobra.Command{
		Use:   "dns",
		Short: "Audit and command DNS related resources",
		Long:  `Audit and command DNS related resourcess`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate DNS related resources",
		Long:  `Enumerate DNS related resources, retreives DNS Zones and Traiffic Manager details`,
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

			report, err := dns.EnumerateDNSResources(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	dnsCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(dnsCmd)
}
