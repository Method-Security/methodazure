package cmd

import (
	"github.com/Method-Security/methodazure/internal/vm"
	"github.com/spf13/cobra"
)

// InitVMCommand initializes the `methodazure vm` subcommand that deals with enumerating Virtual Machines in the
// Azure environment.
func (a *MethodAzure) InitVMCommand() {
	a.VMCmd = &cobra.Command{
		Use:   "vm",
		Short: "Audit and command Virtual Machines",
		Long:  `Audit and command Virtual Machines`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Virtual Machines",
		Long:  `Enumerate Virtual Machines`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := vm.EnumerateVMs(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	a.VMCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(a.VMCmd)
}
