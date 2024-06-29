// Package dns provides functions and data structures to interact with Azure DNS resources.
package dns

import (
	"context"
	"fmt"
	"log"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// ZoneDetails contains details about a single DNS zone and the record sets in the zone.
type ZoneDetails struct {
	ZoneName        string             `json:"zone_name" yaml:"zone_name"`
	ResourceGroup   string             `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string             `json:"resource_group_id" yaml:"resource_group_id"`
	Details         armdns.Zone        `json:"details" yaml:"details"`
	RecordSets      []armdns.RecordSet `json:"record_sets" yaml:"record_sets"`
}

// TrafficManagerProfileDetails contains details about a single Traffic Manager profile.
type TrafficManagerProfileDetails struct {
	ProfileName     string                    `json:"profile_name" yaml:"profile_name"`
	ResourceGroup   string                    `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                    `json:"resource_group_id" yaml:"resource_group_id"`
	Details         armtrafficmanager.Profile `json:"details" yaml:"details"`
}

// AzureResources contains details about all DNS related resources in the subscription.
type AzureResources struct {
	SubscriptionID         string                         `json:"subscription_id" yaml:"subscription_id"`
	TenantID               string                         `json:"tenant_id" yaml:"tenant_id"`
	DNSZones               []ZoneDetails                  `json:"dns_zones" yaml:"dns_zones"`
	TrafficManagerProfiles []TrafficManagerProfileDetails `json:"traffic_manager_profiles" yaml:"traffic_manager_profiles"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateDNSResources enumerates all DNS related resources in the subscription, returning a report of the resources and
// any non-fatal errors encountered.
func EnumerateDNSResources(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	errors := []string{}

	// List all DNS Zones
	dnsZones, err := listDNSZones(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.DNSZones = dnsZones
	}

	// List all Traffic Manager Profiles
	trafficManagerProfiles, err := listTrafficManagerProfiles(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.TrafficManagerProfiles = trafficManagerProfiles
	}

	// Prepare report
	resources.SubscriptionID = cfg.SubID
	resources.TenantID = cfg.TenantID

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}

func listDNSZones(ctx context.Context, cfg config.AzureConfig) ([]ZoneDetails, error) {
	var dnsZones []ZoneDetails

	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armdns.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create DNS client factory: %v", err)
	}

	dnsZonePager := clientFactory.NewZonesClient().NewListPager(nil)
	for dnsZonePager.More() {
		page, err := dnsZonePager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, zone := range page.Value {
			resourceGroup := azure.GetResourceGroupFromID(*zone.ID)
			zoneDetails := ZoneDetails{
				ZoneName:      *zone.Name,
				ResourceGroup: resourceGroup,
				Details:       *zone,
			}
			zoneDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, zoneDetails.ResourceGroup)

			// List record sets in the DNS zone
			recordSets, err := listRecordSets(ctx, cfg, resourceGroup, *zone.Name)
			if err != nil {
				log.Printf("failed to list record sets for DNS zone %s: %v", *zone.Name, err)
			} else {
				zoneDetails.RecordSets = recordSets
			}

			dnsZones = append(dnsZones, zoneDetails)
		}
	}

	return dnsZones, nil
}

func listRecordSets(ctx context.Context, cfg config.AzureConfig, resourceGroup string, dnsZoneName string) ([]armdns.RecordSet, error) {
	var recordSets []armdns.RecordSet

	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armdns.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create DNS client factory: %v", err)
	}

	recordSetPager := clientFactory.NewRecordSetsClient().NewListByDNSZonePager(resourceGroup, dnsZoneName, nil)
	for recordSetPager.More() {
		page, err := recordSetPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, recordSet := range page.Value {
			recordSets = append(recordSets, *recordSet)
		}
	}

	return recordSets, nil
}

func listTrafficManagerProfiles(ctx context.Context, cfg config.AzureConfig) ([]TrafficManagerProfileDetails, error) {
	var trafficManagerProfiles []TrafficManagerProfileDetails

	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armtrafficmanager.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create traffic manager client factory: %v", err)
	}

	tmProfilePager := clientFactory.NewProfilesClient().NewListBySubscriptionPager(nil)
	for tmProfilePager.More() {
		page, err := tmProfilePager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, profile := range page.Value {
			resourceGroup := azure.GetResourceGroupFromID(*profile.ID)
			profileDetails := TrafficManagerProfileDetails{
				ProfileName:   *profile.Name,
				ResourceGroup: resourceGroup,
				Details:       *profile,
			}
			profileDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, profileDetails.ResourceGroup)

			trafficManagerProfiles = append(trafficManagerProfiles, profileDetails)
		}
	}

	return trafficManagerProfiles, nil
}
