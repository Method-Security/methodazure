package loadbalancer

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	methodazure "github.com/Method-Security/methodazure/generated/go"
	"github.com/Method-Security/methodazure/internal/azure"
	azuretransformers "github.com/Method-Security/methodazure/internal/azure/transformers"
)

// General helpers
func convertSubResource(azureSubResource *armnetwork.SubResource) *methodazure.SubResource {
	if azureSubResource == nil {
		return nil
	}

	return &methodazure.SubResource{
		Id: azure.GetStringPtrValue(azureSubResource.ID),
	}
}

// SKU transformers
func convertLoadBalancerSkuType(sku *armnetwork.LoadBalancerSKU) *methodazure.LoadBalancerSku {
	if sku == nil {
		return &methodazure.LoadBalancerSku{}
	}

	var name methodazure.LoadBalancerSkuName
	var tier methodazure.LoadBalancerSkuTier
	var err error

	name, err = methodazure.NewLoadBalancerSkuNameFromString(azure.GetStringEnumPtrValue(sku.Name))
	if err != nil {
		name = methodazure.LoadBalancerSkuNameUnknown
	}

	tier, err = methodazure.NewLoadBalancerSkuTierFromString(azure.GetStringEnumPtrValue(sku.Tier))
	if err != nil {
		tier = methodazure.LoadBalancerSkuTierUnknown
	}

	return &methodazure.LoadBalancerSku{
		Name: name,
		Tier: tier,
	}
}

// BackendAddressPool transformers
func convertBackendAddressPools(azurePools []*armnetwork.BackendAddressPool) []*methodazure.BackendAddressPool {
	var pools []*methodazure.BackendAddressPool

	for _, azurePool := range azurePools {
		if azurePool == nil {
			continue
		}

		syncMode, _ := methodazure.NewSyncModeFromString(azure.GetStringEnumPtrValue(azurePool.Properties.SyncMode))
		customPool := &methodazure.BackendAddressPool{
			Id:                           azure.GetStringPtrValue(azurePool.ID),
			Name:                         azure.GetStringPtrValue(azurePool.Name),
			Type:                         azure.GetStringPtrValue(azurePool.Type),
			LoadBalancerBackendAddresses: convertBackendAddresses(azurePool.Properties.LoadBalancerBackendAddresses),
			Location:                     azure.GetStringPtrValue(azurePool.Properties.Location),
			SyncMode:                     syncMode,
			BackendIpConfigurations:      azuretransformers.ConvertInterfaceIPConfigurations(azurePool.Properties.BackendIPConfigurations),
		}

		pools = append(pools, customPool)
	}

	return pools
}

func convertBackendAddresses(azureAddresses []*armnetwork.LoadBalancerBackendAddress) []*methodazure.LoadBalancerBackendAddress {
	var addresses []*methodazure.LoadBalancerBackendAddress

	for _, azureAddress := range azureAddresses {
		if azureAddress == nil {
			continue
		}

		adminState, _ := methodazure.NewLoadBalancerBackendAddressAdminStateFromString(azure.GetStringEnumPtrValue(azureAddress.Properties.AdminState))
		address := &methodazure.LoadBalancerBackendAddress{
			Name:                                azure.GetStringPtrValue(azureAddress.Name),
			AdminState:                          adminState,
			IpAddress:                           azure.GetStringPtrValue(azureAddress.Properties.IPAddress),
			LoadBalancerFrontendIpConfiguration: convertSubResource(azureAddress.Properties.LoadBalancerFrontendIPConfiguration),
			Subnet:                              convertSubResource(azureAddress.Properties.Subnet),
			VirtualNetwork:                      convertSubResource(azureAddress.Properties.VirtualNetwork),
			InboundNatRulesPortMapping:          convertNatRulePortMappings(azureAddress.Properties.InboundNatRulesPortMapping),
		}

		addresses = append(addresses, address)
	}

	return addresses
}

func convertNatRulePortMappings(azureMappings []*armnetwork.NatRulePortMapping) []*methodazure.NatRulePortMapping {
	var mappings []*methodazure.NatRulePortMapping

	for _, azureMapping := range azureMappings {
		if azureMapping == nil {
			continue
		}

		mapping := &methodazure.NatRulePortMapping{
			BackendPort:        int(*azureMapping.BackendPort),
			FrontendPort:       int(*azureMapping.FrontendPort),
			InboundNatRuleName: azure.GetStringPtrValue(azureMapping.InboundNatRuleName),
		}

		mappings = append(mappings, mapping)
	}

	return mappings
}

// FrontendIPConfiguration transformers
func convertFrontendIPConfigurations(azureConfigs []*armnetwork.FrontendIPConfiguration) []*methodazure.InterfaceIpConfiguration {
	var configs []*methodazure.InterfaceIpConfiguration

	for _, azureConfig := range azureConfigs {
		if azureConfig == nil {
			continue
		}

		config := &methodazure.InterfaceIpConfiguration{
			Id:               azure.GetStringPtrValue(azureConfig.ID),
			Name:             azure.GetStringPtrValue(azureConfig.Name),
			PrivateIpAddress: azure.GetStringPtrValue(azureConfig.Properties.PrivateIPAddress),
			PublicIpAddress:  azuretransformers.ConvertPublicIPAddress(azureConfig.Properties.PublicIPAddress),
			Subnet:           azuretransformers.ConvertSubnet(azureConfig.Properties.Subnet),
		}

		configs = append(configs, config)
	}

	return configs
}
