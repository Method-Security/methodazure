package cmd

import (
	"github.com/Method-Security/methodazure/internal/vm"
	"github.com/spf13/cobra"
)

// InitVMCommand initializes the `methodazure vm` subcommand that deals with enumerating Virtual Machines in the
// Azure environment.
func (a *MethodAzure) InitVMCommand() {
	vmCmd := &cobra.Command{
		Use:   "vm",
		Short: "Audit and command Virtual Machines",
		Long:  `Audit and command Virtual Machines`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Virtual Machines",
		Long:  `Enumerate Virtual Machines`,
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

			report, err := vm.EnumerateVMs(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.PersistentFlags().StringP("subscription-id", "s", "", "Azure subscription ID")

	vmCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(vmCmd)
}
