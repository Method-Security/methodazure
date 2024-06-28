// Package subscription provides functions and data structures to interact with Azure Subscriptions.
package subscription

import (
	"context"
	"fmt"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Method-Security/methodazure/internal/config"
)

// AzureResources contains details about all Subscriptions.
type AzureResources struct {
	Subscriptions []armsubscriptions.Subscription `json:"subscriptions" yaml:"subscriptions"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateSubscriptions enumerates all managed Subscriptions acrtoss a specified list of Azure Clouds (e.g. Public, Government),
// returning a report of the subscriptions and any non-fatal errors encountered.
func EnumerateSubscriptions(ctx context.Context, cfg config.AzureConfig, specifiedClouds []cloud.Configuration) (*AzureResourceReport, error) {
	resources := AzureResources{}
	subscriptions := []armsubscriptions.Subscription{}
	errors := []string{}

	for _, specifiedCloud := range specifiedClouds {
		clientOptions := &armpolicy.ClientOptions{
			ClientOptions: policy.ClientOptions{
				Cloud: specifiedCloud,
			},
		}
		clientFactory, err := armsubscriptions.NewClientFactory(cfg.Cred, clientOptions)
		if err != nil {
			return nil, fmt.Errorf("failed to create subscription client factory: %v", err)
		}

		// Create a pager to list all Tenants this credential has haccess to
		pager := clientFactory.NewClient().NewListPager(nil)

		// Loop through the pages and get Tenants
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				errors = append(errors, err.Error())
				break // This likely indicates that the tenant has no haccess to any subscriptions
			}
			for _, subscription := range page.Value {
				subscriptions = append(subscriptions, *subscription)
			}
		}
	}

	resources.Subscriptions = subscriptions
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
