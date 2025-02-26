// Package vm provides functions and data structures to interact with Azure Virtual Machine resources.
package vm

import (
	"context"
	"fmt"
	"strings"

	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/Method-Security/methodazure/internal/azure"
	"github.com/Method-Security/methodazure/internal/config"
	"github.com/Method-Security/methodazure/internal/vnet"
)

// IPFqdnMapping provides a mapping between an IP address and an FQDN.
type IPFqdnMapping struct {
	IP   string `json:"ip" yaml:"ip"`
	FQDN string `json:"fqdn" yaml:"fqdn"`
}

// NetworkInterface contains details about a single network interface and its corresponding IP addresses and FQDNs.
type NetworkInterface struct {
	ID                     string               `json:"id" yaml:"id"`
	Details                armnetwork.Interface `json:"details" yaml:"details"`
	IPFqdns                []IPFqdnMapping      `json:"ip_fqdn" yaml:"ip_fqdn"`
	NetworkSecurityGroupID *string              `json:"network_security_group_id" yaml:"network_security_group_id"`
	Errors                 []string             `json:"errors" yaml:"errors"`
}

// SubnetDetails contains details about a single subnet.
type SubnetDetails struct {
	ID                     string  `json:"id" yaml:"id"`
	AddressPrefix          string  `json:"address_prefix" yaml:"address_prefix"`
	NetworkSecurityGroupID *string `json:"network_security_group_id" yaml:"network_security_group_id"`
}

// VirtualMachine contains details about a single Virtual Machine and its corresponding network interfaces and linked subnets.
type VirtualMachine struct {
	Details             armcompute.VirtualMachine `json:"details" yaml:"details"`
	VNetID              string                    `json:"vnet_id" yaml:"vnet_id"`
	NetworkInterfaces   []NetworkInterface        `json:"network_interfaces" yaml:"network_interfaces"`
	LinkedSubnetDetails []SubnetDetails           `json:"linked_subnet_details" yaml:"linked_subnet_details"`
}

// VirtualMachineScaleSetVM contains details about a single Virtual Machine Scale Set VM and its corresponding network
// interfaces and linked subnets.
type VirtualMachineScaleSetVM struct {
	Details             armcompute.VirtualMachineScaleSetVM `json:"details" yaml:"details"`
	VNetID              string                              `json:"vnet_id" yaml:"vnet_id"`
	NetworkInterfaces   []NetworkInterface                  `json:"network_interfaces" yaml:"network_interfaces"`
	LinkedSubnetDetails []SubnetDetails                     `json:"linked_subnet_details" yaml:"linked_subnet_details"`
}

// AzureResources contains details about all VM related resources in the subscription.
type AzureResources struct {
	SubscriptionID        string                     `json:"subscription_id" yaml:"subscription_id"`
	TenantID              string                     `json:"tenant_id" yaml:"tenant_id"`
	StandaloneVMInstances []VirtualMachine           `json:"standalone_vms" yaml:"standalone_vms"`
	VMSSVMInstances       []VirtualMachineScaleSetVM `json:"vmss_vms" yaml:"vmss_vms"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}

// EnumerateVMs enumerates all VM related resources in the subscription, returning a report of the resources and any non-fatal
// errors encountered.
func EnumerateVMs(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	var standaloneVMs []VirtualMachine
	var vmssVMs []VirtualMachineScaleSetVM
	errors := []string{}

	// Create VNet and subnet lookups
	vnetReport, err := vnet.EnumerateVNets(ctx, cfg)
	if err != nil {
		return &AzureResourceReport{}, fmt.Errorf("failed to enumerate vnets: %v", err)
	}

	vnetLookup := make(map[string]string)
	subnetLookup := make(map[string]SubnetDetails)
	for _, v := range vnetReport.Resources.VirtualNetworks {
		if v.Details.ID == nil {
			continue
		}
		vnetLookup[v.VNetName] = *v.Details.ID
		if v.Details.Properties != nil && v.Details.Properties.Subnets != nil {
			continue
		}
		for _, s := range v.Details.Properties.Subnets {
			if s.ID != nil && s.Properties != nil && s.Properties.AddressPrefix != nil {
				continue
			}

			var nsgID *string = nil
			if s.Properties.NetworkSecurityGroup != nil {
				nsgID = s.Properties.NetworkSecurityGroup.ID
			}

			subnetLookup[*s.ID] = SubnetDetails{
				ID:                     *s.ID,
				AddressPrefix:          *s.Properties.AddressPrefix,
				NetworkSecurityGroupID: nsgID,
			}
		}
	}

	// Create a new client to interact with the compute resource provider
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	clientFactory, err := armcompute.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return &AzureResourceReport{}, fmt.Errorf("failed to create VM client factory: %v", err)
	}

	// List standalone VMs
	pager := clientFactory.NewVirtualMachinesClient().NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return &AzureResourceReport{}, fmt.Errorf("failed to list pager: %v", err)
		}
		for _, vm := range page.Value {
			vmDetails := VirtualMachine{
				Details:             *vm,
				NetworkInterfaces:   []NetworkInterface{},
				LinkedSubnetDetails: []SubnetDetails{},
			}

			vnetID := ""
			var networkInterfaces []NetworkInterface
			if vm.Properties.NetworkProfile != nil && vm.Properties.NetworkProfile.NetworkInterfaces != nil {
				vnetID, networkInterfaces, err = getVNetIDAndNetworkInterfaces(ctx, cfg, vm.Properties.NetworkProfile, vnetLookup, "vm")
				if err != nil {
					errors = append(errors, err.Error())
				} else {
					vmDetails.VNetID = vnetID
					vmDetails.NetworkInterfaces = networkInterfaces
				}

				// Add linked subnet details
				if vnetID != "" {
					vmDetails.LinkedSubnetDetails = getSubnetDetails(networkInterfaces, subnetLookup)
				}
			}

			standaloneVMs = append(standaloneVMs, vmDetails)
		}
	}

	// List VMSS instances
	vmssPager := clientFactory.NewVirtualMachineScaleSetsClient().NewListAllPager(nil)
	for vmssPager.More() {
		page, err := vmssPager.NextPage(ctx)
		if err != nil {
			return &AzureResourceReport{}, fmt.Errorf("failed to list pager: %v", err)
		}
		for _, vmss := range page.Value {
			resourceGroup := azure.GetResourceGroupFromID(*vmss.ID)
			vmssInstancePager := clientFactory.NewVirtualMachineScaleSetVMsClient().NewListPager(resourceGroup, *vmss.Name, nil)
			for vmssInstancePager.More() {
				instancePage, err := vmssInstancePager.NextPage(ctx)
				if err != nil {
					return &AzureResourceReport{}, fmt.Errorf("failed to list pager: %v", err)
				}
				for _, vm := range instancePage.Value {
					vmDetails := VirtualMachineScaleSetVM{
						Details:             *vm,
						NetworkInterfaces:   []NetworkInterface{},
						LinkedSubnetDetails: []SubnetDetails{},
					}

					vnetID := ""
					var networkInterfaces []NetworkInterface
					if vm.Properties.NetworkProfile != nil && vm.Properties.NetworkProfile.NetworkInterfaces != nil {
						vnetID, networkInterfaces, err = getVNetIDAndNetworkInterfaces(ctx, cfg, vm.Properties.NetworkProfile, vnetLookup, "vmssvm")
						if err != nil {
							errors = append(errors, err.Error())
						} else {
							vmDetails.VNetID = vnetID
							vmDetails.NetworkInterfaces = networkInterfaces
						}

						// Add linked subnet details
						if vnetID != "" {
							vmDetails.LinkedSubnetDetails = getSubnetDetails(networkInterfaces, subnetLookup)
						}
					}

					vmssVMs = append(vmssVMs, vmDetails)
				}
			}
		}
	}

	// Output report
	if standaloneVMs != nil {
		resources.StandaloneVMInstances = standaloneVMs
	}
	if vmssVMs != nil {
		resources.VMSSVMInstances = vmssVMs
	}
	resources.SubscriptionID = cfg.SubID
	resources.TenantID = cfg.TenantID
	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}

// Given a Network Profile, VNet Name -> ID map, this function returns a VNet ID and a list of Network Interfaces
// VMs and VMSS VMs have different interfaces and hence the need to switch
func getVNetIDAndNetworkInterfaces(ctx context.Context, cfg config.AzureConfig, networkProfile *armcompute.NetworkProfile, vnetLookup map[string]string, vmType string) (string, []NetworkInterface, error) {
	clientOptions := &armpolicy.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cfg.CloudConfig,
		},
	}
	networkClientFactory, err := armnetwork.NewClientFactory(cfg.SubID, cfg.Cred, clientOptions)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create network client factory: %v", err)
	}

	nicClient := networkClientFactory.NewInterfacesClient()
	publicIPClient := networkClientFactory.NewPublicIPAddressesClient()
	networkInterfaces := []NetworkInterface{}
	vnetID := ""

	for idx, nic := range networkProfile.NetworkInterfaces {
		nicID := *nic.ID
		var nicInterface armnetwork.Interface
		var nsgID *string = nil
		errors := []string{}

		switch vmType {
		case "vm":
			resourceGroup := azure.GetResourceGroupFromID(nicID)
			nicName := azure.GetResourceNameFromID(nicID)
			nicResp, err := nicClient.Get(ctx, resourceGroup, nicName, nil)
			if err != nil {
				return "", nil, err
			}
			nicInterface = nicResp.Interface
			if nicInterface.Properties.NetworkSecurityGroup != nil {
				nsgID = nicInterface.Properties.NetworkSecurityGroup.ID
			}
		case "vmssvm":
			nicIDParts := strings.Split(nicID, "/")
			if len(nicIDParts) < 9 {
				return "", nil, fmt.Errorf("invalid NIC ID format")
			}
			resourceGroup := azure.GetResourceGroupFromID(nicID)
			nicName := azure.GetResourceNameFromID(nicID)
			vmssName := nicIDParts[len(nicIDParts)-5]
			instanceID := nicIDParts[len(nicIDParts)-3]

			nicResp, err := nicClient.GetVirtualMachineScaleSetNetworkInterface(ctx, resourceGroup, vmssName, instanceID, nicName, nil)
			if err != nil {
				return "", nil, err
			}
			nicInterface = nicResp.Interface
			if nicInterface.Properties.NetworkSecurityGroup != nil {
				nsgID = nicInterface.Properties.NetworkSecurityGroup.ID
			}
		}

		// Populate networkInterface object
		networkInterface := NetworkInterface{
			ID:                     nicID,
			Details:                nicInterface,
			NetworkSecurityGroupID: nsgID,
		}

		// Get IP FQDN mappings if available
		ipFqdnMappings := []IPFqdnMapping{}
		for _, ipConfig := range nicInterface.Properties.IPConfigurations {
			if ipConfig.Properties != nil && ipConfig.Properties.PublicIPAddress != nil {
				fmt.Print("Attempting to enrich public IP ID: " + *ipConfig.Properties.PublicIPAddress.ID + "\n")
				publicIPID := *ipConfig.Properties.PublicIPAddress.ID
				publicIPResp, err := publicIPClient.Get(ctx, azure.GetResourceGroupFromID(publicIPID), azure.GetResourceNameFromID(publicIPID), nil)
				if err != nil {
					errors = append(errors, err.Error())
				}
				if publicIPResp.Properties != nil && publicIPResp.Properties.IPAddress != nil {
					ipFqdnMapping := IPFqdnMapping{
						IP: *publicIPResp.Properties.IPAddress,
					}
					// Some public IPs may not have DNS configured
					if publicIPResp.Properties.DNSSettings != nil && publicIPResp.Properties.DNSSettings.Fqdn != nil {
						ipFqdnMapping.FQDN = *publicIPResp.Properties.DNSSettings.Fqdn
					}
					ipFqdnMappings = append(ipFqdnMappings, ipFqdnMapping)
				}
			}
		}
		networkInterface.IPFqdns = ipFqdnMappings
		networkInterface.Errors = errors
		networkInterfaces = append(networkInterfaces, networkInterface)

		// Extract the vnetID only from the first NIC
		if idx == 0 && nicInterface.Properties != nil && nicInterface.Properties.IPConfigurations != nil {
			for _, ipConfig := range nicInterface.Properties.IPConfigurations {
				if ipConfig.Properties != nil && ipConfig.Properties.Subnet != nil {
					subnetID := *ipConfig.Properties.Subnet.ID
					vnetParts := strings.Split(subnetID, "/")
					if len(vnetParts) >= 9 {
						vnetID = vnetLookup[vnetParts[8]]
						break
					}
				}
			}
		}
	}

	return vnetID, networkInterfaces, nil
}

// Helper function to get SubnetDetails from network interfaces
func getSubnetDetails(networkInterfaces []NetworkInterface, subnetLookup map[string]SubnetDetails) []SubnetDetails {
	subnetDetails := []SubnetDetails{}
	subnetSet := make(map[string]bool)

	for _, nic := range networkInterfaces {
		if nic.Details.Properties != nil && nic.Details.Properties.IPConfigurations != nil {
			for _, ipConfig := range nic.Details.Properties.IPConfigurations {
				if ipConfig.Properties != nil && ipConfig.Properties.Subnet != nil {
					subnetID := *ipConfig.Properties.Subnet.ID
					if _, exists := subnetSet[subnetID]; !exists {
						subnetSet[subnetID] = true
						if details, ok := subnetLookup[subnetID]; ok {
							subnetDetails = append(subnetDetails, details)
						}
					}
				}
			}
		}
	}

	return subnetDetails
}
