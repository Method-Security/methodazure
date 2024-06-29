// Package aks provides functions and data structures to interact with Azure Kubernetes Service (AKS) resources.
package aks

import (
	"context"
	"fmt"
	"log"
	"strings"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
)

// NodePoolDetails contains details about a single AKS node pool and the VMs in the pool.
type NodePoolDetails struct {
	Details armcontainerservice.AgentPool `json:"details" yaml:"details"`
	VMIDs   []string                      `json:"vm_ids" yaml:"vm_ids"`
}

// ClusterDetails contains details about a single AKS cluster, including the cluster details and node pools.
type ClusterDetails struct {
	Name            string                             `json:"cluster_name" yaml:"cluster_name"`
	ResourceGroup   string                             `json:"resource_group" yaml:"resource_group"`
	ResourceGroupID string                             `json:"resource_group_id" yaml:"resource_group_id"`
	Details         armcontainerservice.ManagedCluster `json:"details" yaml:"details"`
	NodePools       []NodePoolDetails                  `json:"node_pools" yaml:"node_pools"`
}

// AzureResources contains details about all AKS clusters in the subscription.
type AzureResources struct {
	SubscriptionID string           `json:"subscription_id" yaml:"subscription_id"`
	TenantID       string           `json:"tenant_id" yaml:"tenant_id"`
	AKSClusters    []ClusterDetails `json:"aks_clusters" yaml:"aks_clusters"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateAKSClusters enumerates all AKS clusters in the subscription, returning a report of the clusters and any
// non-fatal errors encountered.
func EnumerateAKSClusters(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var aksClusters []ClusterDetails
	errors := []string{}

	// Create a new client to interact with the container service resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armcontainerservice.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return &AzureResourceReport{}, fmt.Errorf("failed to create AKS client factory: %v", err)
	}

	// Create a pager to list all AKS clusters in the subscription
	pager := clientFactory.NewManagedClustersClient().NewListPager(nil)

	// Loop through the pages and get AKS clusters
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return &AzureResourceReport{}, fmt.Errorf("failed to list pager: %v", err)
		}
		for _, cluster := range page.Value {
			resourceGroup := azure.GetResourceGroupFromID(*cluster.ID)
			clusterDetails := ClusterDetails{
				Name:          *cluster.Name,
				ResourceGroup: resourceGroup,
				Details:       *cluster,
			}
			clusterDetails.ResourceGroupID = azure.GetResourceGroupIDFromName(cfg.SubID, clusterDetails.ResourceGroup)

			// List node pools in the cluster
			nodePools, err := listNodePools(ctx, cfg, *clientFactory, resourceGroup, *cluster.Name, *cluster.Properties.NodeResourceGroup)
			if err != nil {
				log.Printf("failed to list node pools for AKS cluster %s: %v", *cluster.Name, err)
				errors = append(errors, err.Error())
			} else {
				clusterDetails.NodePools = nodePools
			}

			aksClusters = append(aksClusters, clusterDetails)
		}
	}

	if aksClusters != nil {
		resources.AKSClusters = aksClusters
	}
	resources.SubscriptionID = cfg.SubID
	resources.TenantID = cfg.TenantID
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}

func listNodePools(ctx context.Context, cfg config.AzureConfig, clientFactory armcontainerservice.ClientFactory, resourceGroup string, clusterName string, nodeResourceGroup string) ([]NodePoolDetails, error) {
	var nodePools []NodePoolDetails

	// List node pools in the AKS cluster
	pager := clientFactory.NewAgentPoolsClient().NewListPager(resourceGroup, clusterName, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, pool := range page.Value {
			nodePool := NodePoolDetails{
				Details: *pool,
			}

			// List VM IDs in the node pool
			vmIDs, err := listVMIDsInNodePool(ctx, cfg, nodeResourceGroup, *pool.Name)
			if err != nil {
				return nil, fmt.Errorf("failed to list VMs for node pool %s: %v", *pool.Name, err)
			}
			nodePool.VMIDs = vmIDs
			nodePools = append(nodePools, nodePool)
		}
	}

	return nodePools, nil
}

func listVMIDsInNodePool(ctx context.Context, cfg config.AzureConfig, nodeResourceGroup string, nodePoolName string) ([]string, error) {
	var vmIDs []string

	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armcompute.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client factory: %v", err)
	}

	vmssPager := clientFactory.NewVirtualMachineScaleSetsClient().NewListPager(nodeResourceGroup, nil)
	for vmssPager.More() {
		page, err := vmssPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		for _, vmss := range page.Value {
			// VMSS names associated with node pools are in the format "aks-<nodePoolName>-<random>-vmss"
			// This name check ensures only VMSS, and therefore VMs, associated with the node pool are included
			if strings.Contains("-"+*vmss.Name+"-", nodePoolName) {
				vmssInstancePager := clientFactory.NewVirtualMachineScaleSetVMsClient().NewListPager(nodeResourceGroup, *vmss.Name, nil)
				for vmssInstancePager.More() {
					instancePage, err := vmssInstancePager.NextPage(ctx)
					if err != nil {
						return nil, fmt.Errorf("failed to advance page: %v", err)
					}
					for _, vm := range instancePage.Value {
						vmIDs = append(vmIDs, *vm.ID)
					}
				}
			}
		}
	}

	return vmIDs, nil
}
