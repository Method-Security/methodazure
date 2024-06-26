// Package nsg provides functions and data structures to interact with Azure Network Security Group resources.
package nsg

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// Details contains details about a single NSG and the rules in the NSG.
type Details struct {
	NSGName         string                   `json:"nsg_name" yaml:"nsg_name"`
	Location        string                   `json:"location" yaml:"location"`
	ResourceGroup   string                   `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                   `json:"resource_group_id" yaml:"resource_group_id"`
	Details         armnetwork.SecurityGroup `json:"details" yaml:"details"`
}

// AzureResources contains details about all NSGs in the subscription.
type AzureResources struct {
	SubscriptionID        string    `json:"subscription_id" yaml:"subscription_id"`
	NetworkSecurityGroups []Details `json:"network_security_groups" yaml:"network_security_groups"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateNSGs enumerates all NSGs in the subscription, returning a report of the NSGs and any non-fatal errors encountered.
func EnumerateNSGs(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var networkSecurityGroups []Details
	errors := []string{}

	clientFactory, err := armnetwork.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		log.Fatalf("failed to create network client factory: %v", err)
		errors = append(errors, err.Error())
	}

	// Create a pager to list all NSGs in the subscription
	nsgPager := clientFactory.NewSecurityGroupsClient().NewListAllPager(nil)
	for nsgPager.More() {
		page, err := nsgPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			errors = append(errors, err.Error())
		}
		for _, nsg := range page.Value {
			nsgDetails := Details{
				NSGName:       *nsg.Name,
				Location:      *nsg.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*nsg.ID),
				Details:       *nsg,
			}
			nsgDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, nsgDetails.ResourceGroup)

			networkSecurityGroups = append(networkSecurityGroups, nsgDetails)
		}
	}

	if networkSecurityGroups != nil {
		resources.NetworkSecurityGroups = networkSecurityGroups
	}
	resources.SubscriptionID = cfg.SubID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
