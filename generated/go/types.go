// This file was auto-generated by Fern from our API Definition.

package methodazure

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/Method-Security/methodazure/generated/go/core"
)

type ResourceGroup struct {
	Id   string `json:"id" url:"id"`
	Name string `json:"name" url:"name"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResourceGroup) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResourceGroup) UnmarshalJSON(data []byte) error {
	type unmarshaler ResourceGroup
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResourceGroup(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResourceGroup) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type Subresource struct {
	Id string `json:"id" url:"id"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Subresource) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Subresource) UnmarshalJSON(data []byte) error {
	type unmarshaler Subresource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Subresource(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Subresource) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type InterfaceIpConfiguration struct {
	Id               string           `json:"id" url:"id"`
	Name             string           `json:"name" url:"name"`
	Type             *string          `json:"type,omitempty" url:"type,omitempty"`
	PrivateIpAddress *string          `json:"privateIpAddress,omitempty" url:"privateIpAddress,omitempty"`
	PublicIpAddress  *PublicIpAddress `json:"publicIpAddress,omitempty" url:"publicIpAddress,omitempty"`
	Subnet           *Subnet          `json:"subnet,omitempty" url:"subnet,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *InterfaceIpConfiguration) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InterfaceIpConfiguration) UnmarshalJSON(data []byte) error {
	type unmarshaler InterfaceIpConfiguration
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InterfaceIpConfiguration(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *InterfaceIpConfiguration) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type NetworkInterface struct {
	Id                     string  `json:"id" url:"id"`
	Name                   string  `json:"name" url:"name"`
	NetworkSecurityGroupId *string `json:"networkSecurityGroupID,omitempty" url:"networkSecurityGroupID,omitempty"`
	MacAddress             *string `json:"macAddress,omitempty" url:"macAddress,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (n *NetworkInterface) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NetworkInterface) UnmarshalJSON(data []byte) error {
	type unmarshaler NetworkInterface
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NetworkInterface(value)

	extraProperties, err := core.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties

	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *NetworkInterface) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type PublicIpAddressDnsSettings struct {
	DomainNameLabel string `json:"domainNameLabel" url:"domainNameLabel"`
	Fqdn            string `json:"fqdn" url:"fqdn"`
	ReverseFqdn     string `json:"reverseFqdn" url:"reverseFqdn"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PublicIpAddressDnsSettings) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PublicIpAddressDnsSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler PublicIpAddressDnsSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PublicIpAddressDnsSettings(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PublicIpAddressDnsSettings) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PublicIpAddress struct {
	Id          string                      `json:"id" url:"id"`
	Location    string                      `json:"location" url:"location"`
	IpAddress   string                      `json:"ipAddress" url:"ipAddress"`
	DnsSettings *PublicIpAddressDnsSettings `json:"dnsSettings,omitempty" url:"dnsSettings,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PublicIpAddress) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PublicIpAddress) UnmarshalJSON(data []byte) error {
	type unmarshaler PublicIpAddress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PublicIpAddress(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PublicIpAddress) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type Subnet struct {
	Id                     string   `json:"id" url:"id"`
	Name                   string   `json:"name" url:"name"`
	Type                   *string  `json:"type,omitempty" url:"type,omitempty"`
	AddressPrefix          *string  `json:"addressPrefix,omitempty" url:"addressPrefix,omitempty"`
	AddressPrefixes        []string `json:"addressPrefixes,omitempty" url:"addressPrefixes,omitempty"`
	NetworkSecurityGroupId *string  `json:"networkSecurityGroupID,omitempty" url:"networkSecurityGroupID,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Subnet) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Subnet) UnmarshalJSON(data []byte) error {
	type unmarshaler Subnet
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Subnet(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Subnet) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type TransportProtocol string

const (
	TransportProtocolTcp TransportProtocol = "Tcp"
	TransportProtocolUdp TransportProtocol = "Udp"
	TransportProtocolAll TransportProtocol = "All"
)

func NewTransportProtocolFromString(s string) (TransportProtocol, error) {
	switch s {
	case "Tcp":
		return TransportProtocolTcp, nil
	case "Udp":
		return TransportProtocolUdp, nil
	case "All":
		return TransportProtocolAll, nil
	}
	var t TransportProtocol
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TransportProtocol) Ptr() *TransportProtocol {
	return &t
}

// Collection of backend address pools used by the load balancer:
// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5#BackendAddressPool
type BackendAddressPool struct {
	Id                           string                        `json:"id" url:"id"`
	Name                         string                        `json:"name" url:"name"`
	Type                         string                        `json:"type" url:"type"`
	LoadBalancerBackendAddresses []*LoadBalancerBackendAddress `json:"loadBalancerBackendAddresses,omitempty" url:"loadBalancerBackendAddresses,omitempty"`
	Location                     *string                       `json:"location,omitempty" url:"location,omitempty"`
	SyncMode                     *SyncMode                     `json:"syncMode,omitempty" url:"syncMode,omitempty"`
	VirtualNetwork               *Subresource                  `json:"virtualNetwork,omitempty" url:"virtualNetwork,omitempty"`
	BackendIpConfigurations      []*InterfaceIpConfiguration   `json:"backendIpConfigurations,omitempty" url:"backendIpConfigurations,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BackendAddressPool) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BackendAddressPool) UnmarshalJSON(data []byte) error {
	type unmarshaler BackendAddressPool
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BackendAddressPool(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BackendAddressPool) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// LoadBalancer represents an Azure Load Balancer as defined in the Azure Go SDK:
// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5#LoadBalancer
type LoadBalancer struct {
	Id                       string                      `json:"id" url:"id"`
	Name                     string                      `json:"name" url:"name"`
	Location                 string                      `json:"location" url:"location"`
	ResourceGroup            string                      `json:"resourceGroup" url:"resourceGroup"`
	ResourceGroupId          string                      `json:"resourceGroupId" url:"resourceGroupId"`
	Sku                      *LoadBalancerSku            `json:"sku,omitempty" url:"sku,omitempty"`
	BackendAddressPools      []*BackendAddressPool       `json:"backendAddressPools,omitempty" url:"backendAddressPools,omitempty"`
	FrontendIpConfigurations []*InterfaceIpConfiguration `json:"frontendIPConfigurations,omitempty" url:"frontendIPConfigurations,omitempty"`
	LoadBalancingRules       []*LoadBalancingRule        `json:"loadBalancingRules,omitempty" url:"loadBalancingRules,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *LoadBalancer) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoadBalancer) UnmarshalJSON(data []byte) error {
	type unmarshaler LoadBalancer
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoadBalancer(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoadBalancer) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LoadBalancerBackendAddress struct {
	Name                                string                                `json:"name" url:"name"`
	AdminState                          *LoadBalancerBackendAddressAdminState `json:"adminState,omitempty" url:"adminState,omitempty"`
	IpAddress                           *string                               `json:"ipAddress,omitempty" url:"ipAddress,omitempty"`
	LoadBalancerFrontendIpConfiguration *Subresource                          `json:"loadBalancerFrontendIPConfiguration,omitempty" url:"loadBalancerFrontendIPConfiguration,omitempty"`
	Subnet                              *Subresource                          `json:"subnet,omitempty" url:"subnet,omitempty"`
	VirtualNetwork                      *Subresource                          `json:"virtualNetwork,omitempty" url:"virtualNetwork,omitempty"`
	InboundNatRulesPortMapping          []*NatRulePortMapping                 `json:"inboundNatRulesPortMapping,omitempty" url:"inboundNatRulesPortMapping,omitempty"`
	NetworkInterfaceIpConfigurations    *Subresource                          `json:"networkInterfaceIpConfigurations,omitempty" url:"networkInterfaceIpConfigurations,omitempty"`
	NetworkInterface                    *NetworkInterface                     `json:"networkInterface,omitempty" url:"networkInterface,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *LoadBalancerBackendAddress) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoadBalancerBackendAddress) UnmarshalJSON(data []byte) error {
	type unmarshaler LoadBalancerBackendAddress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoadBalancerBackendAddress(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoadBalancerBackendAddress) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LoadBalancerBackendAddressAdminState string

const (
	LoadBalancerBackendAddressAdminStateDown LoadBalancerBackendAddressAdminState = "Down"
	LoadBalancerBackendAddressAdminStateNone LoadBalancerBackendAddressAdminState = "None"
	LoadBalancerBackendAddressAdminStateUp   LoadBalancerBackendAddressAdminState = "Up"
)

func NewLoadBalancerBackendAddressAdminStateFromString(s string) (LoadBalancerBackendAddressAdminState, error) {
	switch s {
	case "Down":
		return LoadBalancerBackendAddressAdminStateDown, nil
	case "None":
		return LoadBalancerBackendAddressAdminStateNone, nil
	case "Up":
		return LoadBalancerBackendAddressAdminStateUp, nil
	}
	var t LoadBalancerBackendAddressAdminState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LoadBalancerBackendAddressAdminState) Ptr() *LoadBalancerBackendAddressAdminState {
	return &l
}

type LoadBalancerReport struct {
	SubscriptionId string          `json:"subscriptionId" url:"subscriptionId"`
	TenantId       string          `json:"tenantId" url:"tenantId"`
	LoadBalancers  []*LoadBalancer `json:"loadBalancers,omitempty" url:"loadBalancers,omitempty"`
	Errors         []string        `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *LoadBalancerReport) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoadBalancerReport) UnmarshalJSON(data []byte) error {
	type unmarshaler LoadBalancerReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoadBalancerReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoadBalancerReport) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LoadBalancerSkuName string

const (
	LoadBalancerSkuNameBasic    LoadBalancerSkuName = "Basic"
	LoadBalancerSkuNameGateway  LoadBalancerSkuName = "Gateway"
	LoadBalancerSkuNameStandard LoadBalancerSkuName = "Standard"
	LoadBalancerSkuNameUnknown  LoadBalancerSkuName = "Unknown"
)

func NewLoadBalancerSkuNameFromString(s string) (LoadBalancerSkuName, error) {
	switch s {
	case "Basic":
		return LoadBalancerSkuNameBasic, nil
	case "Gateway":
		return LoadBalancerSkuNameGateway, nil
	case "Standard":
		return LoadBalancerSkuNameStandard, nil
	case "Unknown":
		return LoadBalancerSkuNameUnknown, nil
	}
	var t LoadBalancerSkuName
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LoadBalancerSkuName) Ptr() *LoadBalancerSkuName {
	return &l
}

type LoadBalancerSkuTier string

const (
	LoadBalancerSkuTierGlobal   LoadBalancerSkuTier = "Global"
	LoadBalancerSkuTierRegional LoadBalancerSkuTier = "Regional"
	LoadBalancerSkuTierUnknown  LoadBalancerSkuTier = "Unknown"
)

func NewLoadBalancerSkuTierFromString(s string) (LoadBalancerSkuTier, error) {
	switch s {
	case "Global":
		return LoadBalancerSkuTierGlobal, nil
	case "Regional":
		return LoadBalancerSkuTierRegional, nil
	case "Unknown":
		return LoadBalancerSkuTierUnknown, nil
	}
	var t LoadBalancerSkuTier
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LoadBalancerSkuTier) Ptr() *LoadBalancerSkuTier {
	return &l
}

type LoadBalancerSku struct {
	Name LoadBalancerSkuName `json:"name" url:"name"`
	Tier LoadBalancerSkuTier `json:"tier" url:"tier"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *LoadBalancerSku) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoadBalancerSku) UnmarshalJSON(data []byte) error {
	type unmarshaler LoadBalancerSku
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoadBalancerSku(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoadBalancerSku) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// LoadBalancingRule represents an Azure Load Balancing Rule as defined in the Azure Go SDK:
// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5#LoadBalancingRule
type LoadBalancingRule struct {
	Id                      string            `json:"id" url:"id"`
	Name                    string            `json:"name" url:"name"`
	FrontendPort            int               `json:"frontendPort" url:"frontendPort"`
	Protocol                TransportProtocol `json:"protocol" url:"protocol"`
	BackendAddressPool      *Subresource      `json:"backendAddressPool,omitempty" url:"backendAddressPool,omitempty"`
	BackendAddressPools     []*Subresource    `json:"backendAddressPools,omitempty" url:"backendAddressPools,omitempty"`
	BackendPort             int               `json:"backendPort" url:"backendPort"`
	FrontendIpConfiguration *Subresource      `json:"frontendIPConfiguration,omitempty" url:"frontendIPConfiguration,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *LoadBalancingRule) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoadBalancingRule) UnmarshalJSON(data []byte) error {
	type unmarshaler LoadBalancingRule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoadBalancingRule(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoadBalancingRule) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type NatRulePortMapping struct {
	BackendPort        int    `json:"backendPort" url:"backendPort"`
	FrontendPort       int    `json:"frontendPort" url:"frontendPort"`
	InboundNatRuleName string `json:"inboundNatRuleName" url:"inboundNatRuleName"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (n *NatRulePortMapping) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NatRulePortMapping) UnmarshalJSON(data []byte) error {
	type unmarshaler NatRulePortMapping
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NatRulePortMapping(value)

	extraProperties, err := core.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties

	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *NatRulePortMapping) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type SyncMode string

const (
	SyncModeAutomatic SyncMode = "Automatic"
	SyncModeManual    SyncMode = "Manual"
)

func NewSyncModeFromString(s string) (SyncMode, error) {
	switch s {
	case "Automatic":
		return SyncModeAutomatic, nil
	case "Manual":
		return SyncModeManual, nil
	}
	var t SyncMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SyncMode) Ptr() *SyncMode {
	return &s
}
