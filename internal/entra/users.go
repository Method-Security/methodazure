// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"

	"github.com/Method-Security/methodazure/internal/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func listUsers(ctx context.Context, client *msgraphsdk.GraphServiceClient) ([]UserDetails, error) {
	var userDetails []UserDetails

	// Set up query parameters
	queryParams := &users.UsersRequestBuilderGetQueryParameters{
		Select: []string{"id", "displayName", "userPrincipalName", "mail"},
		Top:    new(int32),
	}
	*queryParams.Top = 50 // Controls pagination size

	// Set up request configuration
	requestConfig := &users.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: queryParams,
	}

	// Get the first page of results
	result, err := client.Users().Get(ctx, requestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	for {
		if result.GetValue() != nil {
			for _, user := range result.GetValue() {
				userDetails = append(userDetails, UserDetails{
					ID:                azure.GetStringPtrValue(user.GetId()),
					DisplayName:       azure.GetStringPtrValue(user.GetDisplayName()),
					UserPrincipalName: azure.GetStringPtrValue(user.GetUserPrincipalName()),
					Mail:              azure.GetStringPtrValue(user.GetMail()),
				})
			}
		}

		nextLink := result.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		// Create a new request for the next page
		nextReq := users.NewUsersRequestBuilder(*nextLink, client.GetAdapter())
		result, err = nextReq.Get(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of users: %v", err)
		}
	}

	return userDetails, nil
}
