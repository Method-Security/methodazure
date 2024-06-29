package cmd

import (
	"strings"

	"github.com/Method-Security/methodazure/internal/entra"
	"github.com/spf13/cobra"
)

// InitEntraCommand initializes the `methodazure entra` subcommand that deals with enumerating Entra related resources in the Azure environment.
func (a *MethodAzure) InitEntraCommand() {
	entraCmd := &cobra.Command{
		Use:   "entra",
		Short: "Audit and command Azure Entra ID",
		Long:  `Audit and command Azure Entra ID`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Entra ID users, groups, and service principals in a given Tenant",
		Long:  `Enumerate Entra ID users, groups, and service principals in a given Tenant`,
		Run: func(cmd *cobra.Command, args []string) {
			// If the graph-service-endpoint flag is not set, default to the --cloud-config mapped value
			graphServiceEndpoint, err := cmd.Flags().GetString("graph-service-endpoint")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			if graphServiceEndpoint == "" {
				if strings.Contains(a.AzureConfig.CloudConfig.ActiveDirectoryAuthorityHost, ".com") {
					graphServiceEndpoint = "https://graph.microsoft.com/.default"
				} else if strings.Contains(a.AzureConfig.CloudConfig.ActiveDirectoryAuthorityHost, ".us") {
					graphServiceEndpoint = "https://graph.microsoft.us/.default"
				} else if strings.Contains(a.AzureConfig.CloudConfig.ActiveDirectoryAuthorityHost, ".cn") {
					graphServiceEndpoint = "https://microsoftgraph.chinacloudapi.cn/.default"
				}
			}
			a.AzureConfig.GraphServiceEndpoint = graphServiceEndpoint

			report, err := entra.EnumerateEntra(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("graph-service-endpoint", "g", "", "Scope of Microsoft Graph Service Endpoint (e.g. https://graph.microsoft.com/.default), this is automatically defaulted based on --cloud-config value")

	entraCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(entraCmd)
}
