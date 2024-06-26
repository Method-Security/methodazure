// Package resourcegroup provides functions and data structures to interact with Azure Resource Group resources.
package resourcegroup

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Method-Security/methodazure/internal/config"
)

// Details contains details about a single Resource Group.
type Details struct {
	ResourceGroupName string                     `json:"resource_group_name" yaml:"resource_group_name"`
	Location          string                     `json:"location" yaml:"location"`
	Details           armresources.ResourceGroup `json:"details" yaml:"details"`
}

// AzureResources contains details about all Resource Groups in the subscription.
type AzureResources struct {
	SubscriptionID string    `json:"subscription_id" yaml:"subscription_id"`
	ResourceGroups []Details `json:"resource_groups" yaml:"resource_groups"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateResourceGroups enumerates all Resource Groups in the subscription, returning a report of the Resource Groups and any
// non-fatal errors encountered.
func EnumerateResourceGroups(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var resourceGroups []Details
	errors := []string{}

	clientFactory, err := armresources.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		log.Fatalf("failed to create resources client factory: %v", err)
		errors = append(errors, err.Error())
	}

	// Create a pager to list all Resource Groups in the subscription
	rgPager := clientFactory.NewResourceGroupsClient().NewListPager(nil)
	for rgPager.More() {
		page, err := rgPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			errors = append(errors, err.Error())
		}
		for _, rg := range page.Value {
			rgDetails := Details{
				ResourceGroupName: *rg.Name,
				Location:          *rg.Location,
				Details:           *rg,
			}

			resourceGroups = append(resourceGroups, rgDetails)
		}
	}

	if resourceGroups != nil {
		resources.ResourceGroups = resourceGroups
	}
	resources.SubscriptionID = cfg.SubID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
