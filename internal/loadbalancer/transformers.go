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

		// Id, Name, Type are always present, regardless of ability to get more details
		// The other fields need enrichment of some kind which may not always work; the cnversion functions below return nil if no enrichment
		pool := &methodazure.BackendAddressPool{
			Id:                           azure.GetStringPtrValue(azurePool.ID),
			Name:                         azure.GetStringPtrValue(azurePool.Name),
			Type:                         azure.GetStringPtrValue(azurePool.Type),
			LoadBalancerBackendAddresses: convertBackendAddresses(azurePool.Properties.LoadBalancerBackendAddresses),
			BackendIpConfigurations:      azuretransformers.ConvertInterfaceIPConfigurations(azurePool.Properties.BackendIPConfigurations),
		}
		syncMode, err := methodazure.NewSyncModeFromString(azure.GetStringEnumPtrValue(azurePool.Properties.SyncMode))
		if err != nil && syncMode != ""{
			pool.SyncMode = &syncMode
		}
		location := azure.GetStringPtrValue(azurePool.Properties.Location)
		if location != "" {
			pool.Location = &location
		}

		pools = append(pools, pool)
	}

	return pools
}

func convertBackendAddresses(azureAddresses []*armnetwork.LoadBalancerBackendAddress) []*methodazure.LoadBalancerBackendAddress {
	var addresses []*methodazure.LoadBalancerBackendAddress

	for _, azureAddress := range azureAddresses {
		if azureAddress == nil {
			continue
		}

		// Name is the only guaranteed field
		// The properties from conversions return nil if no enrichment is possible
		address := &methodazure.LoadBalancerBackendAddress{
			Name:                                azure.GetStringPtrValue(azureAddress.Name),
			LoadBalancerFrontendIpConfiguration: convertSubResource(azureAddress.Properties.LoadBalancerFrontendIPConfiguration),
			Subnet:                              convertSubResource(azureAddress.Properties.Subnet),
			VirtualNetwork:                      convertSubResource(azureAddress.Properties.VirtualNetwork),
			InboundNatRulesPortMapping:          convertNatRulePortMappings(azureAddress.Properties.InboundNatRulesPortMapping),
		}
		adminState, err := methodazure.NewLoadBalancerBackendAddressAdminStateFromString(azure.GetStringEnumPtrValue(azureAddress.Properties.AdminState))
		if err != nil && adminState != "" {
			address.AdminState = &adminState
		}
		ipAddress := azure.GetStringPtrValue(azureAddress.Properties.IPAddress)
		if ipAddress != "" {
			address.IpAddress = &ipAddress
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
			PrivateIpAddress: azureConfig.Properties.PrivateIPAddress,
			PublicIpAddress:  azuretransformers.ConvertPublicIPAddress(azureConfig.Properties.PublicIPAddress),
			Subnet:           azuretransformers.ConvertSubnet(azureConfig.Properties.Subnet),
		}

		configs = append(configs, config)
	}

	return configs
}

// Load Balancer rule transformers
func convertLoadBalancingRules(azureRules []*armnetwork.LoadBalancingRule) []*methodazure.LoadBalancingRule {
	var rules []*methodazure.LoadBalancingRule

	for _, azureRule := range azureRules {
		if azureRule == nil {
			continue
		}

		var backendAddressPools []*methodazure.SubResource
		for _, pool := range azureRule.Properties.BackendAddressPools {
			backendAddressPools = append(backendAddressPools, convertSubResource(pool))
		}

		rule := &methodazure.LoadBalancingRule{
			Id:                 azure.GetStringPtrValue(azureRule.ID),
			Name:               azure.GetStringPtrValue(azureRule.Name),
			FrontendPort:       int(*azureRule.Properties.FrontendPort),
			BackendAddressPool: convertSubResource(azureRule.Properties.BackendAddressPool),
			BackendAddressPools: backendAddressPools,
			BackendPort:        int(*azureRule.Properties.BackendPort),
			FrontendIpConfiguration: convertSubResource(azureRule.Properties.FrontendIPConfiguration),
		}
		protocol, err := methodazure.NewTransportProtocolFromString(azure.GetStringEnumPtrValue(azureRule.Properties.Protocol))
		if err != nil && protocol != "" {
			rule.Protocol = protocol
		}
		

		rules = append(rules, rule)
	}

	return rules
}