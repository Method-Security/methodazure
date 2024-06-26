package cmd

import (
	"github.com/Method-Security/methodazure/internal/nsg"
	"github.com/spf13/cobra"
)

// InitNSGCommand initializes the `methodazure nsg` subcommand that deals with enumerating Network Security Groups in the Azure environment.
func (a *MethodAzure) InitNSGCommand() {
	nsgCmd := &cobra.Command{
		Use:   "nsg",
		Short: "Audit and command Network Security Groups",
		Long:  `Audit and command Network Security Groups`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Network Security Groups",
		Long:  `Enumerate Network Security Groups`,
		Run: func(cmd *cobra.Command, args []string) {
			report, err := nsg.EnumerateNSGs(cmd.Context(), a.AzureConfig)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}

	nsgCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(nsgCmd)
}
