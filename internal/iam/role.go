// Package iam provides functions and data structures to interact with Azure IAM resources.
package iam

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/Method-Security/methodazure/internal/config"
)

func listRoles(ctx context.Context, cfg config.AzureConfig) ([]RoleDetails, error) {
	roles := []RoleDetails{}

	// Create a new client to interact with the Authorization resource provider
	clientFactory, err := armauthorization.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		return roles, fmt.Errorf("failed to create role definitions client: %v", err)
	}

	// The scope for listing role definitions (in this case, the subscription)
	scope := fmt.Sprintf("/subscriptions/%s", cfg.SubID)

	// Create a pager to list all roles in the subscription
	pager := clientFactory.NewRoleDefinitionsClient().NewListPager(scope, nil)

	// Loop through the pages and save each role definition
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return roles, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, role := range page.Value {
			roleDetails := RoleDetails{
				ID:             *role.ID,
				Name:           *role.Name,
				RoleDefinition: *role,
			}

			roles = append(roles, roleDetails)
		}
	}

	return roles, nil
}
