// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"context"
	"fmt"
	"log"

	"github.com/Method-Security/methodazure/internal/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func listGroups(ctx context.Context, client *msgraphsdk.GraphServiceClient) ([]GroupDetails, error) {
	groupsDetailsList := []GroupDetails{}

	// Set up query parameters for groups
	queryParams := &groups.GroupsRequestBuilderGetQueryParameters{
		Select: []string{"id", "displayName", "description", "createdDateTime", "members"},
		Expand: []string{"members($select=id)"},
		Top:    new(int32),
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
				groupDetails := GroupDetails{
					ID:              azure.GetStringPtrValue(group.GetId()),
					DisplayName:     azure.GetStringPtrValue(group.GetDisplayName()),
					Description:     azure.GetStringPtrValue(group.GetDescription()),
					CreatedDateTime: *group.GetCreatedDateTime(),
				}

				// Get members for this group
				members, err := getGroupMembers(group)
				if err != nil {
					fmt.Println("Failed to get members for group", *group.GetId(), ":", err)
				} else {
					groupDetails.Members = members
				}

				groupsDetailsList = append(groupsDetailsList, groupDetails)

			}
		}

		nextLink := result.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		// Create a new request for the next page
		nextReq := groups.NewGroupsRequestBuilder(*nextLink, client.GetAdapter())
		result, err = nextReq.Get(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to get next page of groups: %v", err)
		}
	}

	return groupsDetailsList, nil
}

func getGroupMembers(group models.Groupable) ([]GroupMember, error) {
	groupMembers := []GroupMember{}

	members := group.GetMembers()
	if members == nil {
		return nil, fmt.Errorf("group %s has no members", *group.GetId())
	}

	for _, member := range members {
		// Check if the member object is not nil
		if member == nil {
			continue
		}

		id := member.GetId()
		if id == nil {
			continue
		} else {
			groupMembers = append(groupMembers, GroupMember{
				ID: *id,
			})
		}
	}

	return groupMembers, nil
}
