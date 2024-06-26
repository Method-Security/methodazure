package cmd

import (
	"github.com/Method-Security/methodazure/internal/storage"
	"github.com/spf13/cobra"
)

// InitStorageAccountCommand initializes the `methodazure storage` subcommand that deals with enumerating Storage
// Accounts in the Azure environment.
func (a *MethodAzure) InitStorageAccountCommand() {
	storageAccountCmd := &cobra.Command{
		Use:   "storage",
		Short: "Audit and command Storage Accounts",
		Long:  `Audit and command Storage Accounts`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Storage Accounts",
		Long:  `Enumerate Storage Accounts`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := storage.EnumerateStorageAccounts(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	storageAccountCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(storageAccountCmd)
}
