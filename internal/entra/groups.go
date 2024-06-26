// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func listGroups(ctx context.Context, client *msgraphsdk.GraphServiceClient) ([]GroupDetails, error) {
	var groupDetails []GroupDetails

	// Set up query parameters for groups
	queryParams := &groups.GroupsRequestBuilderGetQueryParameters{
		Select: []string{"id", "displayName"},
	}
	*queryParams.Top = 50 // Controls pagination size

	// Set up request configuration for groups
	requestConfig := &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: queryParams,
	}

	// Get the first page of results
	result, err := client.Groups().Get(ctx, requestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	for {
		if result.GetValue() != nil {
			for _, group := range result.GetValue() {
				fmt.Printf("Group ID: %s, Display Name: %s\n", *group.GetId(), *group.GetDisplayName())
				
				// Get members for this group
				getGroupMembers(graphClient, *group.GetId())
				
				fmt.Println() // Add a blank line between groups
			}
		}

		nextLink := result.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		// Create a new request for the next page
		nextReq := groups.NewGroupsRequestBuilder(*nextLink, graphClient.GetAdapter())
		result, err = nextReq.Get(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to get next page of groups: %v", err)
		}
	}
}

func getGroupMembers(graphClient *msgraphsdk.GraphServiceClient, groupID string) {
	ctx := context.Background()

	// Set up query parameters for group members
	queryParams := &groups.ItemMembersRequestBuilderGetQueryParameters{
		Select: []string{"id"},
	}

	// Set up request configuration for group members
	requestConfig := &groups.ItemMembersRequestBuilderGetRequestConfiguration{
		QueryParameters: queryParams,
	}

	// Get the first page of results
	result, err := graphClient.GroupsById(groupID).Members().Get(ctx, requestConfig)
	if err != nil {
		log.Printf("Failed to get members for group %s: %v", groupID, err)
		return
	}

	fmt.Println("  Members:")
	for {
		if result.GetValue() != nil {
			for _, member := range result.GetValue() {
				fmt.Printf("    Member ID: %s\n", *member.GetId())
			}
		}

		nextLink := result.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		// Create a new request for the next page
		nextReq := groups.NewItemMembersRequestBuilder(*nextLink, graphClient.GetAdapter())
		result, err = nextReq.Get(ctx, nil)
		if err != nil {
			log.Printf("Failed to get next page of members for group %s: %v", groupID, err)
			return
		}
	}
}