package cmd

import (
	"github.com/Method-Security/methodazure/internal/tenant"
	"github.com/spf13/cobra"
)

// InitTenantCommand initializes the `methodazure tenant` subcommand that deals with enumerating Azure Tenant related details.
func (a *MethodAzure) InitTenantCommand() {
	tenantCmd := &cobra.Command{
		Use:   "tenant",
		Short: "Audit and command Azure Tenants",
		Long:  `Audit and command Azure Tenants`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Tenants for the provided credentials",
		Long:  `Enumerate Tenants for the provided credentials`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := tenant.EnumerateTenants(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	tenantCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(tenantCmd)
}
