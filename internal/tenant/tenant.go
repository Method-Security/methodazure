// Package tenant provides functions and data structures to interact with Azure Tenants.
package tenant

import (
	"context"
	"fmt"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Method-Security/methodazure/internal/config"
)

// AzureResources contains details about all Tenants.
type AzureResources struct {
	Tenants []armsubscriptions.TenantIDDescription `json:"tenants" yaml:"tenants"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

func EnumerateTenants(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	tenants := []armsubscriptions.TenantIDDescription{}
	errors := []string{}

	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armsubscriptions.NewClientFactory(cfg.Cred, clientOptions)
	if err != nil {
		return &AzureResourceReport{}, fmt.Errorf("failed to create subscription client factory: %v", err)
	}

	// Create a pager to list all Tenants this credential has haccess to
	pager := clientFactory.NewTenantsClient().NewListPager(nil)

	// Loop through the pages and get Tenants
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return &AzureResourceReport{}, fmt.Errorf("failed to list pager: %v", err)
		}
		for _, tenant := range page.Value {
			tenants = append(tenants, *tenant)
		}
	}

	resources.Tenants = tenants
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
