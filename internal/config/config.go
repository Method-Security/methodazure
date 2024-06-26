// Package config contains common configuration values that are used by the various commands and subcommands in the CLI.
package config

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type RootFlags struct {
	Quiet   bool
	Verbose bool
}

type AzureConfig struct {
	Cred                 *azidentity.DefaultAzureCredential
	TenantID             string
	SubID                string
	GraphServiceEndpoint string
}
