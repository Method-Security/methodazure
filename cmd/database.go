package cmd

import (
	"github.com/Method-Security/methodazure/internal/database"
	"github.com/spf13/cobra"
)

// InitDatabaseCommand initializes the `methodazure database` subcommand that deals with enumerating managed Database
// instances in the Azure environment.
func (a *MethodAzure) InitDatabaseCommand() {
	databaseCmd := &cobra.Command{
		Use:   "database",
		Short: "Audit and command managed Database instances",
		Long:  `Audit and command managed Database instances`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate managed Database instances",
		Long:  `Enumerate managed Database instances; retreives managed SQL, Postgres, and Postgres Flexible instance details`,
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

			report, err := database.EnumerateDatabaseInstances(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	databaseCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(databaseCmd)
}
