// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/Method-Security/methodazure/internal/azure"
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

	// Create a pager to list all users
	result, err := client.Users().Get(ctx, requestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	// Process and print user details, goes through first pagination page
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

	// Handle pagination if there are more results
	nextLink := result.GetOdataNextLink()
	for nextLink != nil {
		// Create a new request for the next page
		nextReq := users.NewUsersRequestBuilder(*nextLink, client.GetAdapter())
		result, err = nextReq.Get(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of users: %v", err)
		}

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

		nextLink = result.GetOdataNextLink()
	}

	return userDetails, nil
}