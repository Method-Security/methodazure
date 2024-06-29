// Package storage provides functions and data structures to interact with Azure Storage resources.
package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// Details contains details about a single Storage Account.
type Details struct {
	AccountName     string                         `json:"account_name" yaml:"account_name"`
	Location        string                         `json:"location" yaml:"location"`
	ResourceGroup   string                         `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                         `json:"resource_group_id" yaml:"resource_group_id"`
	Details         armstorage.Account             `json:"details" yaml:"details"`
	VNetID          string                         `json:"vnet_id" yaml:"vnet_id"`
	BlobContainers  []armstorage.ListContainerItem `json:"blob_containers" yaml:"blob_containers"`
}

// AzureResources contains details about all Storage Accounts in the subscription.
type AzureResources struct {
	SubscriptionID  string    `json:"subscription_id" yaml:"subscription_id"`
	TenantID       	string    `json:"tenant_id" yaml:"tenant_id"`
	StorageAccounts []Details `json:"storage_accounts" yaml:"storage_accounts"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateStorageAccounts enumerates all Storage Accounts in the subscription, returning a report of the accounts and any
// non-fatal errors encountered.
func EnumerateStorageAccounts(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var storageAccounts []Details
	errors := []string{}

	// Create a new client to interact with the storage resource provider
	clientFactory, err := armstorage.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		log.Fatalf("failed to create client factory: %v", err)
		errors = append(errors, err.Error())
	}

	// Create a pager to list all storage accounts in the subscription
	pager := clientFactory.NewAccountsClient().NewListPager(nil)

	// Loop through the pages and save each storage account
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			errors = append(errors, err.Error())
		}
		for _, storageAccount := range page.Value {
			storageAccountDetails := Details{
				AccountName:   *storageAccount.Name,
				Location:      *storageAccount.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*storageAccount.ID),
				Details:       *storageAccount,
			}
			storageAccountDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, storageAccountDetails.ResourceGroup)

			// Get the VNet ID for the storage account
			if storageAccount.Properties.NetworkRuleSet != nil {
				var vNetName string
				if len(storageAccount.Properties.NetworkRuleSet.VirtualNetworkRules) > 0 {
					vNetName = azure.GetVNetNameFromID(*storageAccount.Properties.NetworkRuleSet.VirtualNetworkRules[0].VirtualNetworkResourceID)
				} else {
					log.Printf("no virtual network rules found for storage account %s", *storageAccount.Name)
					errors = append(errors, fmt.Sprintf("no virtual network rules found for storage account %s", *storageAccount.Name))
				}
				vNetID, err := azure.GetVNetIDFromVNetName(ctx, cfg, storageAccountDetails.ResourceGroup, vNetName)
				if err != nil {
					log.Printf("failed to get VNet ID: %v", err)
					errors = append(errors, err.Error())
				} else {
					storageAccountDetails.VNetID = vNetID
				}
			}

			// List blob containers in the storage account
			blobContainers, err := listBlobContainers(ctx, cfg, storageAccountDetails.AccountName, storageAccountDetails.ResourceGroup)
			if err != nil {
				log.Printf("failed to list blob containers for account %s: %v", *storageAccount.Name, err)
				errors = append(errors, err.Error())
			} else {
				storageAccountDetails.BlobContainers = blobContainers
			}

			storageAccounts = append(storageAccounts, storageAccountDetails)
		}
	}

	if storageAccounts != nil {
		resources.StorageAccounts = storageAccounts
	}
	resources.SubscriptionID = cfg.SubID
	resources.TenantID = cfg.TenantID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}

func listBlobContainers(ctx context.Context, cfg config.AzureConfig, accountName string, resourceGroup string) ([]armstorage.ListContainerItem, error) {
	var blobContainers []armstorage.ListContainerItem

	clientFactory, err := armstorage.NewClientFactory(cfg.SubID, cfg.Cred, nil)
	if err != nil {
		return []armstorage.ListContainerItem{}, fmt.Errorf("failed to create client factory: %v", err)
	}

	// Get the blob containers pager for the specific storage account
	pager := clientFactory.NewBlobContainersClient().NewListPager(resourceGroup, accountName, nil)

	// Loop through the pages and save each storage account
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return []armstorage.ListContainerItem{}, fmt.Errorf("failed to advance page: %v", err)
		}

		for _, blobContainer := range page.Value {
			blobContainers = append(blobContainers, *blobContainer)
		}
	}

	return blobContainers, nil
}
