// Package vnet provides functions and data structures to interact with Azure VNet resources.
package vnet

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// Details contains details about a single VNet.
type Details struct {
	VNetName        string                    `json:"vnet_name"`
	Location        string                    `json:"location"`
	ResourceGroup   string                    `json:"resource_group"`
	ResourceGroupID string                    `json:"resource_group_id"`
	Details         armnetwork.VirtualNetwork `json:"details"`
}

// AzureResources contains details about all VNets in the subscription.
type AzureResources struct {
	SubscriptionID  string    `json:"subscription_id"`
	VirtualNetworks []Details `json:"virtual_networks"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources"`
	Errors    []string       `json:"errors"`
}

// EnumerateVNets enumerates all VNets in the subscription, returning a report of the VNets and any non-fatal errors encountered.
func EnumerateVNets(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var virtualNetworks []Details
	errors := []string{}

	clientFactory, err := armnetwork.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		log.Fatalf("failed to create network client factory: %v", err)
		errors = append(errors, err.Error())
	}

	// Create a pager to list all VNets in the subscription
	vnetPager := clientFactory.NewVirtualNetworksClient().NewListAllPager(nil)
	for vnetPager.More() {
		page, err := vnetPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			errors = append(errors, err.Error())
		}
		for _, vnet := range page.Value {
			resourceGroup := azure.GetResourceGroupFromID(*vnet.ID)
			vnetDetails := Details{
				VNetName:      *vnet.Name,
				Location:      *vnet.Location,
				ResourceGroup: resourceGroup,
				Details:       *vnet,
			}
			vnetDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, vnetDetails.ResourceGroup)

			virtualNetworks = append(virtualNetworks, vnetDetails)
		}
	}

	if virtualNetworks != nil {
		resources.VirtualNetworks = virtualNetworks
	}
	resources.SubscriptionID = cfg.SubID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
