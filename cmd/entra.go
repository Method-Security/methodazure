package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Method-Security/methodazure/internal/entra"
)

// InitDNSCommand initializes the `methodazure dns` subcommand that deals with enumerating DNS related resources in the Azure environment.
func (a *MethodAzure) InitEntraCommand() {
	entraCmd := &cobra.Command{
		Use:   "entra",
		Short: "Audit and command Azure Entra ID",
		Long:  `Audit and command Azure Entra ID`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Entra ID users, groups, rols, and role assignments",
		Long:  `Enumerate Entra ID users, groups, rols, and role assignments`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := entra.EnumerateEntra(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	entraCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(entraCmd)
}
