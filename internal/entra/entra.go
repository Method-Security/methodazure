// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/Method-Security/methodazure/internal/config"
)

func EnumerateEntra(ctx context.Context, cfg config.AzureConfig) (AzureResourceReport, error) {
	resources := AzureResourceReport{}
	errors := []string{}

	// Create client to interact with Microsoft Graph
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cfg.Cred, []string{cfg.GraphServiceEndpoint})
	if err != nil {
		return resources, fmt.Errorf("failed to create Graph client: %v", err)
	}

	// Get Users
	users, err := listUsers(ctx, graphClient)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.Users = users
	}

	// Get Groups
	resources.Errors = errors
	return resources, nil
}