package transformers

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	methodazure "github.com/Method-Security/methodazure/generated/go"
	"github.com/Method-Security/methodazure/internal/azure"
)

func ConvertSubnet(azureSubnet *armnetwork.Subnet) *methodazure.Subnet {
	if azureSubnet == nil {
		return nil
	}

	var addressPrefixes []string
	for _, prefix := range azureSubnet.Properties.AddressPrefixes {
		addressPrefixes = append(addressPrefixes, *prefix)
	}

	subnet := &methodazure.Subnet{
		Id:              azure.GetStringPtrValue(azureSubnet.ID),
		Name:            azure.GetStringPtrValue(azureSubnet.Name),
		AddressPrefix:   azureSubnet.Properties.AddressPrefix,
		AddressPrefixes: addressPrefixes,
	}

	return subnet
}

func ConvertPublicIPAddressDNSSettings(azureDNSSettings *armnetwork.PublicIPAddressDNSSettings) *methodazure.PublicIpAddressDnsSettings {
	if azureDNSSettings == nil {
		return nil
	}
	dnsSettings := &methodazure.PublicIpAddressDnsSettings{
		DomainNameLabel: azure.GetStringPtrValue(azureDNSSettings.DomainNameLabel),
		Fqdn:            azure.GetStringPtrValue(azureDNSSettings.Fqdn),
	}
	return dnsSettings
}

func ConvertPublicIPAddress(azurePublicIP *armnetwork.PublicIPAddress) *methodazure.PublicIpAddress {
	if azurePublicIP == nil {
		return nil
	}

	publicIP := &methodazure.PublicIpAddress{
		Id:          azure.GetStringPtrValue(azurePublicIP.ID),
		Location:    azure.GetStringPtrValue(azurePublicIP.Location),
		IpAddress:   azure.GetStringPtrValue(azurePublicIP.Properties.IPAddress),
		DnsSettings: ConvertPublicIPAddressDNSSettings(azurePublicIP.Properties.DNSSettings),
	}

	return publicIP
}

func ConvertInterfaceIPConfigurations(azureIPConfigs []*armnetwork.InterfaceIPConfiguration) []*methodazure.InterfaceIpConfiguration {
	var ipConfigs []*methodazure.InterfaceIpConfiguration

	for _, azureIPConfig := range azureIPConfigs {
		if azureIPConfig == nil {
			continue
		}

		ipConfig := &methodazure.InterfaceIpConfiguration{
			Id:               azure.GetStringPtrValue(azureIPConfig.ID),
			Name:             azure.GetStringPtrValue(azureIPConfig.Name),
			PrivateIpAddress: azure.GetStringPtrValue(azureIPConfig.Properties.PrivateIPAddress),
			PublicIpAddress:  ConvertPublicIPAddress(azureIPConfig.Properties.PublicIPAddress),
			Subnet:           ConvertSubnet(azureIPConfig.Properties.Subnet),
		}

		ipConfigs = append(ipConfigs, ipConfig)
	}

	return ipConfigs
}
