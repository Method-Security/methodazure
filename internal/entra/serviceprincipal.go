// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"

	"github.com/Method-Security/methodazure/internal/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
)

func listServicePrincipals(ctx context.Context, client *msgraphsdk.GraphServiceClient) ([]ServicePrincipalDetails, error) {
	var spDetails []ServicePrincipalDetails

	// Set up query parameters
	queryParams := &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
		Select: []string{"id", "displayName", "appId", "servicePrincipalType",
			"appRoleAssignmentRequired", "accountEnabled"},
		Top: new(int32),
	}
	*queryParams.Top = 50 // Controls pagination size

	// Set up request configuration
	requestConfig := &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: queryParams,
	}

	// Create a pager to list all users
	result, err := client.ServicePrincipals().Get(ctx, requestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	for {
		if result.GetValue() != nil {
			for _, serviceprincipal := range result.GetValue() {
				spDetails = append(spDetails, ServicePrincipalDetails{
					ID:                        azure.GetStringPtrValue(serviceprincipal.GetId()),
					DisplayName:               azure.GetStringPtrValue(serviceprincipal.GetDisplayName()),
					AppID:                     azure.GetStringPtrValue(serviceprincipal.GetAppId()),
					ServicePrincipalType:      azure.GetStringPtrValue(serviceprincipal.GetServicePrincipalType()),
					AppRoleAssignmentRequired: *serviceprincipal.GetAppRoleAssignmentRequired(),
					AccountEnabled:            *serviceprincipal.GetAccountEnabled(),
				})
			}
		}

		nextLink := result.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		// Create a new request for the next page
		nextReq := serviceprincipals.NewServicePrincipalsRequestBuilder(*nextLink, client.GetAdapter())
		result, err = nextReq.Get(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of serviceprincipals: %v", err)
		}
	}

	return spDetails, nil
}
