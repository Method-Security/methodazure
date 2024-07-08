// Package loadbalancer provides functions and data structures to interact with Azure Load Balancer resources.
package loadbalancer

import (
	"context"

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
				Name: 		*lb.Name,
				Location:   *lb.Location,
				ResourceGroup: azure.GetResourceGroupFromID(*lb.ID),
			}
			loadBalancer.ResourceGroupId = azure.GetResourceGroupIDFromName(cfg.SubID, loadBalancer.ResourceGroup)

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
