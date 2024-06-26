// Package azure provides utility functions for interacting with Azure resources and the Azure SDK.
package azure

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/Method-Security/methodazure/internal/config"
)

// GetResourceGroupFromID returns the resource group name from a given resource ID.
func GetResourceGroupFromID(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "resourceGroups") && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// GetResourceNameFromID returns the resource name from a given resource ID.
func GetResourceNameFromID(resourceID string) string {
	// Split the resource ID by '/'
	parts := strings.Split(resourceID, "/")
	// The resource name is typically the last part of the ID
	return parts[len(parts)-1]
}

// GetResourceGroupIDFromName creates the resource group ID from a given subscription ID and resource group name.
func GetResourceGroupIDFromName(subID string, resourceGroupName string) string {
	return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subID, resourceGroupName)
}

// GetVNetNameFromID returns the VNet name from a given resource ID.
func GetVNetNameFromID(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "virtualNetworks") && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// GetVNetIDFromVNetName returns the VNet ID from a given resource group name and VNet name.
func GetVNetIDFromVNetName(ctx context.Context, cfg config.AzureConfig, resourceGroupName string, vnetName string) (string, error) {
	// Create a Virtual Networks client
	vnetClient, err := armnetwork.NewVirtualNetworksClient(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create virtual networks client: %v", err)
	}

	// Get the VNet
	vnet, err := vnetClient.Get(ctx, resourceGroupName, vnetName, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get VNet: %v", err)
	}

	// Return the VNet ID
	return *vnet.ID, nil
}

// GetStringPtrValue safely returns the value of a string pointer or an empty string if it's nil
func GetStringPtrValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
