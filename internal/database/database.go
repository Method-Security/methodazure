// Package database provides functions and data structures to interact with Azure Database resources.
package database

import (
	"context"
	"fmt"
	"log"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// SQLInstanceDetails contains details about a single SQL instance.
type SQLInstanceDetails struct {
	InstanceName    string                 `json:"instance_name" yaml:"instance_name"`
	Location        string                 `json:"location" yaml:"location"`
	ResourceGroup   string                 `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                 `json:"resource_group_id" yaml:"resource_group_id"`
	VNetID          string                 `json:"vnet_id" yaml:"vnet_id"`
	Details         armsql.ManagedInstance `json:"details" yaml:"details"`
}

// PostgresInstanceDetails contains details about a single PostgreSQL instance.
type PostgresInstanceDetails struct {
	InstanceName    string               `json:"instance_name" yaml:"instance_name"`
	Location        string               `json:"location" yaml:"location"`
	ResourceGroup   string               `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string               `json:"resource_group_id" yaml:"resource_group_id"`
	VNetID          string               `json:"vnet_id" yaml:"vnet_id"`
	Details         armpostgresql.Server `json:"details" yaml:"details"`
}

// PostgresFlexibleInstanceDetails contains details about a single PostgreSQL Flexible instance.
type PostgresFlexibleInstanceDetails struct {
	InstanceName    string                              `json:"instance_name" yaml:"instance_name"`
	Location        string                              `json:"location" yaml:"location"`
	ResourceGroup   string                              `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                              `json:"resource_group_id" yaml:"resource_group_id"`
	VNetID          string                              `json:"vnet_id" yaml:"vnet_id"`
	Details         armpostgresqlflexibleservers.Server `json:"details" yaml:"details"`
}

// AzureResources contains details about all database instances in the subscription.
type AzureResources struct {
	SubscriptionID            string                            `json:"subscription_id" yaml:"subscription_id"`
	TenantID                  string                            `json:"tenant_id" yaml:"tenant_id"`
	SQLInstances              []SQLInstanceDetails              `json:"sql_instances" yaml:"sql_instances"`
	PostgresInstances         []PostgresInstanceDetails         `json:"postgres_instances" yaml:"postgres_instances"`
	PostgresFlexibleInstances []PostgresFlexibleInstanceDetails `json:"postgres_flexible_instances" yaml:"postgres_flexible_instances"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateDatabaseInstances enumerates all managed Database instances in the subscription, returning a report of the
// instances and any non-fatal errors encountered.
func EnumerateDatabaseInstances(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	errors := []string{}

	sqlInstances, err := listSQLInstances(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.SQLInstances = sqlInstances
	}

	postgresInstances, err := listPostgresInstances(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.PostgresInstances = postgresInstances
	}

	postgresFlexibleInstances, err := listPostgresFlexibleInstances(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.PostgresFlexibleInstances = postgresFlexibleInstances
	}

	resources.SubscriptionID = cfg.SubID
	resources.TenantID = cfg.TenantID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}

func listSQLInstances(ctx context.Context, cfg config.AzureConfig) ([]SQLInstanceDetails, error) {
	instances := []SQLInstanceDetails{}

	// Create a new client to interact with the SQL resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armsql.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return instances, fmt.Errorf("failed to create client factory: %v", err)
	}

	// Create a pager to list all SQL instances in the subscription
	pager := clientFactory.NewManagedInstancesClient().NewListPager(nil)

	// Loop through the pages and save each SQL instance
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return instances, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, sqlInstance := range page.Value {
			sqlInstanceDetails := SQLInstanceDetails{
				InstanceName:  *sqlInstance.Name,
				Location:      *sqlInstance.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*sqlInstance.ID),
				Details:       *sqlInstance,
			}
			sqlInstanceDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, sqlInstanceDetails.ResourceGroup)

			// Get VNet ID
			vNetName := azure.GetVNetNameFromID(*sqlInstance.Properties.SubnetID)
			vNetID, err := azure.GetVNetIDFromVNetName(ctx, cfg, sqlInstanceDetails.ResourceGroup, vNetName)
			if err != nil {
				log.Printf("failed to get VNet ID: %v", err)
			} else {
				sqlInstanceDetails.VNetID = vNetID
			}

			instances = append(instances, sqlInstanceDetails)
		}
	}

	return instances, nil
}

func listPostgresInstances(ctx context.Context, cfg config.AzureConfig) ([]PostgresInstanceDetails, error) {
	instances := []PostgresInstanceDetails{}

	// Create a new client to interact with the PostgreSQL resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armpostgresql.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return instances, fmt.Errorf("failed to create client factory: %v", err)
	}

	// Create a pager to list all PostgreSQL instances in the subscription
	pager := clientFactory.NewServersClient().NewListPager(nil)

	// Loop through the pages and save each PostgreSQL instance
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return instances, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, postgresInstance := range page.Value {
			postgresInstanceDetails := PostgresInstanceDetails{
				InstanceName:  *postgresInstance.Name,
				Location:      *postgresInstance.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*postgresInstance.ID),
				Details:       *postgresInstance,
			}
			postgresInstanceDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, postgresInstanceDetails.ResourceGroup)

			// VNet ID is not readily accessible

			instances = append(instances, postgresInstanceDetails)
		}
	}

	return instances, nil
}

func listPostgresFlexibleInstances(ctx context.Context, cfg config.AzureConfig) ([]PostgresFlexibleInstanceDetails, error) {
	instances := []PostgresFlexibleInstanceDetails{}

	// Create a new client to interact with the PostgreSQL resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return instances, fmt.Errorf("failed to create client factory: %v", err)
	}

	// Create a pager to list all PostgreSQL instances in the subscription
	pager := clientFactory.NewServersClient().NewListPager(nil)

	// Loop through the pages and save each PostgreSQL instance
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return instances, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, postgresInstance := range page.Value {
			postgresInstanceDetails := PostgresFlexibleInstanceDetails{
				InstanceName:  *postgresInstance.Name,
				Location:      *postgresInstance.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*postgresInstance.ID),
				Details:       *postgresInstance,
			}
			postgresInstanceDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, postgresInstanceDetails.ResourceGroup)

			// Get VNet ID
			if postgresInstance.Properties.Network == nil || postgresInstance.Properties.Network.DelegatedSubnetResourceID == nil {
				log.Printf("no VNet ID found for PostgreSQL Flexible instance %s", *postgresInstance.Name)
				continue
			}
			vNetName := azure.GetVNetNameFromID(*postgresInstance.Properties.Network.DelegatedSubnetResourceID)
			vNetID, err := azure.GetVNetIDFromVNetName(ctx, cfg, postgresInstanceDetails.ResourceGroup, vNetName)
			if err != nil {
				log.Printf("failed to get VNet ID: %v", err)
			} else {
				postgresInstanceDetails.VNetID = vNetID
			}

			instances = append(instances, postgresInstanceDetails)
		}
	}

	return instances, nil
}
