// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"

	"github.com/Method-Security/methodazure/internal/config"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func EnumerateEntra(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	errors := []string{}

	// Create client to interact with Microsoft Graph
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cfg.Cred, []string{cfg.GraphServiceEndpoint})
	if err != nil {
		return &AzureResourceReport{}, fmt.Errorf("failed to create Graph client: %v", err)
	}

	// Get Users
	users, err := listUsers(ctx, graphClient)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.Users = users
	}

	// Get Groups
	groups, err := listGroups(ctx, graphClient)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.Groups = groups
	}

	// Get Service Principals
	servicePrincipals, err := listServicePrincipals(ctx, graphClient)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.ServicePrincipals = servicePrincipals
	}

	resources.TenantID = cfg.TenantID
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
