// Package azure provides utility functions for interacting with Azure resources and the Azure SDK.
package azure

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
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

// GetResourceNameFromID returns the resource name from a given resource ID, assuming the last element
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

// GetNetworkInterfaceNameFromID returns the network interface name from a given resource ID.
func GetNetworkInterfaceNameFromID(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "networkInterfaces") && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// GetVNetIDFromVNetName returns the VNet ID from a given resource group name and VNet name.
func GetVNetIDFromVNetName(ctx context.Context, cfg config.AzureConfig, resourceGroupName string, vnetName string) (string, error) {
	// Create a new client to interact with the storage resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}

	// Create a Virtual Networks client
	vnetClient, err := armnetwork.NewVirtualNetworksClient(cfg.SubID, cfg.Cred, clientOptions)
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

// GetIPConfurationDetailsFromID returns the IP configuration details from a given IpConfiguration resource ID.
func GetIPConfurationDetailsFromID(ctx context.Context, clientFactory armnetwork.ClientFactory, id string) (*armnetwork.InterfaceIPConfiguration, error) {
	// Create a Virtual Networks client
	client := clientFactory.NewInterfaceIPConfigurationsClient()

	resourceGroup := GetResourceGroupFromID(id)
	networkInterfaceName := GetNetworkInterfaceNameFromID(id)
	ipConfigurationName := GetResourceNameFromID(id)
	// Get the IP configuration details
	resp, err := client.Get(ctx, resourceGroup, networkInterfaceName, ipConfigurationName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get IP configuration details: %v", err)
	}

	return &resp.InterfaceIPConfiguration, nil
}

// GetPublicIPAddressDetailsFromID returns the public IP address details from a given public IP address resource ID.
func GetPublicIPAddressDetailsFromID(ctx context.Context, clientFactory armnetwork.ClientFactory, id string) (*armnetwork.PublicIPAddress, error) {
	// Create a Public IP Addresses client
	client := clientFactory.NewPublicIPAddressesClient()

	resourceGroup := GetResourceGroupFromID(id)
	publicIPAddressName := GetResourceNameFromID(id)
	// Get the public IP address details
	resp, err := client.Get(ctx, resourceGroup, publicIPAddressName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get public IP address details: %v", err)
	}

	return &resp.PublicIPAddress, nil
}

// GetSubnetDetailsFromID returns the subnet details from a given subnet resource ID.
func GetSubnetDetailsFromID(ctx context.Context, clientFactory armnetwork.ClientFactory, id string) (*armnetwork.Subnet, error) {
	// Create a Subnets client
	client := clientFactory.NewSubnetsClient()

	resourceGroup := GetResourceGroupFromID(id)
	vnetName := GetVNetNameFromID(id)
	subnetName := GetResourceNameFromID(id)
	// Get the subnet details
	resp, err := client.Get(ctx, resourceGroup, vnetName, subnetName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get subnet details: %v", err)
	}

	return &resp.Subnet, nil
}

// GetStringPtrValue safely returns the value of a string pointer or an empty string if it's nil
func GetStringPtrValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// GetStringEnumPtrValue safely returns the value of a ptr string enum type or an empty string if it's nil
func GetStringEnumPtrValue(value interface{}) string {
	if value == nil {
		return ""
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		fmt.Print("warning, value provided to GetStringEnumValue is not a ptr")
		return ""
	}

	// Dereference the pointer
	v = v.Elem()
	if v.Kind() != reflect.String {
		fmt.Println("warning, value provided to GetStringEnumPtrValue is not pointing to a string")
		return ""
	}

	return v.String()
}
