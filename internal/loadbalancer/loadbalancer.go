// Package loadbalancer provides functions and data structures to interact with Azure Load Balancer resources.
package loadbalancer

import (
	"context"
	"fmt"
	"strings"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	
	methodazure "github.com/Method-Security/methodazure/generated/go"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// EnumerateLoadBalancers enumerates all Load Balancers in the subscription, returning a report of the Load Balancers and any
// non-fatal errors encountered.
func EnumerateLoadBalancers(ctx context.Context, cfg config.AzureConfig) methodazure.LoadBalancerReport {
	loadBalancers := []*methodazure.LoadBalancer{}
	errors := []string{}

	// Create a new client to interact with the storage resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armnetwork.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		errors = append(errors, err.Error())
		return methodazure.LoadBalancerReport{
			SubscriptionId: cfg.SubID,
			TenantId:       cfg.TenantID,
			LoadBalancers:  loadBalancers,
			Errors:         errors,
		}
	}

	// Create a pager to list all Load Balancers in the subscription
	lbPager := clientFactory.NewLoadBalancersClient().NewListAllPager(nil)
	for lbPager.More() {
		page, err := lbPager.NextPage(ctx)
		if err != nil {
			errors = append(errors, err.Error())
			return methodazure.LoadBalancerReport{
				SubscriptionId: cfg.SubID,
				TenantId:       cfg.TenantID,
				LoadBalancers:  loadBalancers,
				Errors:         errors,
			}
		}
		for _, lb := range page.Value {
			loadBalancer := methodazure.LoadBalancer{
				Id:            *lb.ID,
				Name:          *lb.Name,
				Location:      *lb.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*lb.ID),
				Sku:           convertLoadBalancerSkuType(lb.SKU),
			}
			loadBalancer.ResourceGroupId = azure.GetResourceGroupIDFromName(cfg.SubID, loadBalancer.ResourceGroup)

			// Backend address pools
			enrichedBackendAddressPools, poolErrors := enrichBackendAddressPools(ctx, *clientFactory, lb.Properties.BackendAddressPools)
			if len(poolErrors) > 0 {
				errors = append(errors, poolErrors...)
			}
			loadBalancer.BackendAddressPools = convertBackendAddressPools(enrichedBackendAddressPools)
			
			// Frontend IP configurations
			enrichedFrontendIPConfigs, frontendErrors := enrichFrontendIPConfigurations(ctx, *clientFactory, lb.Properties.FrontendIPConfigurations)
			if len(frontendErrors) > 0 {
				errors = append(errors, frontendErrors...)
			}
			loadBalancer.FrontendIpConfigurations = convertFrontendIPConfigurations(enrichedFrontendIPConfigs)

			// Load Balancing Rules
			loadBalancer.LoadBalancingRules = convertLoadBalancingRules(lb.Properties.LoadBalancingRules)

			loadBalancers = append(loadBalancers, &loadBalancer)
		}
	}

	return methodazure.LoadBalancerReport{
		SubscriptionId: cfg.SubID,
		TenantId:       cfg.TenantID,
		LoadBalancers:  loadBalancers,
		Errors:         errors,
	}
}

// enrichBackendAddressPools enriches BackendAddressPools with the details of the BackendIPConfigurations.
// This enrichment is done inline and keeps the original ID only information if enrichment is not possible.
func enrichBackendAddressPools(ctx context.Context, clientFactory armnetwork.ClientFactory, azurePools []*armnetwork.BackendAddressPool) ([]*armnetwork.BackendAddressPool, []string) {
	errors := []string{}

	for azurePoolIndex, azurePool := range azurePools {
		if azurePool == nil {
			continue
		}

		// Loop over every BackendIPConfigurations and query for the details, enrich when possible, otherwise keep just the ID
		if azurePool.Properties.BackendIPConfigurations != nil {
			for backendIPConfigIndex, backendIPConfig := range azurePool.Properties.BackendIPConfigurations {
				if backendIPConfig == nil {
					continue
				}

				// Attempt to enrich the BackendIPConfiguration with the details
				if backendIPConfig.ID != nil {
					ipConfigDetails, err := azure.GetIPConfurationDetailsFromID(ctx, clientFactory, *backendIPConfig.ID)
					if err != nil {
						errors = append(errors, err.Error())
					} else {
						if ipConfigDetails != nil {
							azurePools[azurePoolIndex].Properties.BackendIPConfigurations[backendIPConfigIndex] = ipConfigDetails

							// Attempt to enrich PublicIPAddress with the details
							if ipConfigDetails.Properties.PublicIPAddress != nil {
								publicIPDetails, err := azure.GetPublicIPAddressDetailsFromID(ctx, clientFactory, *ipConfigDetails.Properties.PublicIPAddress.ID)
								if err != nil {
									errors = append(errors, err.Error())
								} else {
									if publicIPDetails != nil {
										azurePools[azurePoolIndex].Properties.BackendIPConfigurations[backendIPConfigIndex].Properties.PublicIPAddress = publicIPDetails
									}
								}
							}

							// Attempt to enrich Subnet with the details
							if ipConfigDetails.Properties.Subnet != nil {
								subnetDetails, err := azure.GetSubnetDetailsFromID(ctx, clientFactory, *ipConfigDetails.Properties.Subnet.ID)
								if err != nil {
									errors = append(errors, err.Error())
								} else {
									if subnetDetails != nil {
										azurePools[azurePoolIndex].Properties.BackendIPConfigurations[backendIPConfigIndex].Properties.Subnet = subnetDetails
									}
								}
							}
						}
					}

				}
			}
		}
	}

	return azurePools, errors
}

// enrichFrontendIPConfigurations enriches FrontendIPConfigurations with the details of the PublicIPAddress and Subnet.
// This enrichment is done inline and keeps the original ID only information if enrichment is not possible.
func enrichFrontendIPConfigurations(ctx context.Context, clientFactory armnetwork.ClientFactory, frontendIPConfigs []*armnetwork.FrontendIPConfiguration) ([]*armnetwork.FrontendIPConfiguration, []string) {
	errors := []string{}

	for frontendIPConfigIndex, frontendIPConfig := range frontendIPConfigs {
		if frontendIPConfig == nil {
			continue
		}

		// Attempt to enrich the FrontendIPConfiguration with the details
		if frontendIPConfig.ID != nil {
			frontendIPConfigDetails, err := getFrontendIPConfigurationDetailsFromID(ctx, clientFactory, *frontendIPConfig.ID)
			if err != nil {
				errors = append(errors, err.Error())
			} else {
				if frontendIPConfigDetails != nil {
					frontendIPConfigs[frontendIPConfigIndex] = frontendIPConfigDetails

					// Attempt to enrich PublicIPAddress with the details
					if frontendIPConfigDetails.Properties.PublicIPAddress != nil {
						publicIPDetails, err := azure.GetPublicIPAddressDetailsFromID(ctx, clientFactory, *frontendIPConfigDetails.Properties.PublicIPAddress.ID)
						if err != nil {
							errors = append(errors, err.Error())
						} else {
							if publicIPDetails != nil {
								frontendIPConfigs[frontendIPConfigIndex].Properties.PublicIPAddress = publicIPDetails
							}
						}
					}

					// Attempt to enrich Subnet with the details
					if frontendIPConfigDetails.Properties.Subnet != nil {
						subnetDetails, err := azure.GetSubnetDetailsFromID(ctx, clientFactory, *frontendIPConfigDetails.Properties.Subnet.ID)
						if err != nil {
							errors = append(errors, err.Error())
						} else {
							if subnetDetails != nil {
								frontendIPConfigs[frontendIPConfigIndex].Properties.Subnet = subnetDetails
							}
						}
					}
				}
			}

		}
	}

	return frontendIPConfigs, errors
}

// getFrontendIpConfigurationDetailsFromID returns the frontend IP configuration details from a given frontend IP configuration resource ID.
// FrontendIPConfiguration is specific to load balancer so keep this function in loadbalancer package
func getFrontendIPConfigurationDetailsFromID(ctx context.Context, clientFactory armnetwork.ClientFactory, id string) (*armnetwork.FrontendIPConfiguration, error) {
	// Create a Frontend IP Configurations client
	client := clientFactory.NewLoadBalancerFrontendIPConfigurationsClient()

	resourceGroup := azure.GetResourceGroupFromID(id)
	loadBalancerName := getLoadBalancerNameFromID(id)
	frontendIPConfigName := azure.GetResourceNameFromID(id)
	// Get the frontend IP configuration details
	resp, err := client.Get(ctx, resourceGroup, loadBalancerName, frontendIPConfigName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get frontend IP configuration details: %v", err)
	}

	return &resp.FrontendIPConfiguration, nil
}

func getLoadBalancerNameFromID(id string) string {
	parts := strings.Split(id, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "loadBalancers") && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
