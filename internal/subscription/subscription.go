// Package subscription provides functions and data structures to interact with Azure Subscriptions.
package subscription

import (
	"context"
	"fmt"
	"log"

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

func EnumerateSubscriptions(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	subscriptions := []armsubscriptions.Subscription{}
	errors := []string{}

	clientFactory, err := armsubscriptions.NewClientFactory(cfg.Cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription client factory: %v", err)
	}

	// Create a pager to list all Tenants this credential has haccess to
	pager := clientFactory.NewClient().NewListPager(nil)

	// Loop through the pages and get Tenants
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			errors = append(errors, err.Error())
		}
		for _, subscription := range page.Value {
			subscriptions = append(subscriptions, *subscription)
		}
	}

	resources.Subscriptions = subscriptions
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
