package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a servicePrincipal object.
type ServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalItemRequestBuilderGetQueryParameters
}
// ServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddKey provides operations to call the addKey method.
// returns a *ItemAddkeyAddKeyRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AddKey()(*ItemAddkeyAddKeyRequestBuilder) {
    return NewItemAddkeyAddKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddPassword provides operations to call the addPassword method.
// returns a *ItemAddpasswordAddPasswordRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AddPassword()(*ItemAddpasswordAddPasswordRequestBuilder) {
    return NewItemAddpasswordAddPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddTokenSigningCertificate provides operations to call the addTokenSigningCertificate method.
// returns a *ItemAddtokensigningcertificateAddTokenSigningCertificateRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AddTokenSigningCertificate()(*ItemAddtokensigningcertificateAddTokenSigningCertificateRequestBuilder) {
    return NewItemAddtokensigningcertificateAddTokenSigningCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AppManagementPolicies()(*ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) {
    return NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignedTo provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemApproleassignedtoAppRoleAssignedToRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*ItemApproleassignedtoAppRoleAssignedToRequestBuilder) {
    return NewItemApproleassignedtoAppRoleAssignedToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemApproleassignmentsAppRoleAssignmentsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignments()(*ItemApproleassignmentsAppRoleAssignmentsRequestBuilder) {
    return NewItemApproleassignmentsAppRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
// returns a *ItemCheckmembergroupsCheckMemberGroupsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) CheckMemberGroups()(*ItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    return NewItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
// returns a *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) CheckMemberObjects()(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) {
    return NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemCreatedobjectsCreatedObjectsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) CreatedObjects()(*ItemCreatedobjectsCreatedObjectsRequestBuilder) {
    return NewItemCreatedobjectsCreatedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DelegatedPermissionClassifications provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) {
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a servicePrincipal object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-delete?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemEndpointsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) Endpoints()(*ItemEndpointsRequestBuilder) {
    return NewItemEndpointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentials()(*ItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilder) {
    return NewItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederatedIdentityCredentialsWithName provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentialsWithName(name *string)(*ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) {
    return NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, name)
}
// Get retrieve the properties and relationships of a servicePrincipal object.
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-get?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
// returns a *ItemGetmembergroupsGetMemberGroupsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) GetMemberGroups()(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    return NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
// returns a *ItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) GetMemberObjects()(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemMemberofMemberOfRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) MemberOf()(*ItemMemberofMemberOfRequestBuilder) {
    return NewItemMemberofMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemOwnedobjectsOwnedObjectsRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) OwnedObjects()(*ItemOwnedobjectsOwnedObjectsRequestBuilder) {
    return NewItemOwnedobjectsOwnedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemOwnersRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch create a new servicePrincipal object if it doesn't exist, or update the properties of an existing servicePrincipal object.
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-upsert?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// RemoteDesktopSecurityConfiguration provides operations to manage the remoteDesktopSecurityConfiguration property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemRemotedesktopsecurityconfigurationRemoteDesktopSecurityConfigurationRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) RemoteDesktopSecurityConfiguration()(*ItemRemotedesktopsecurityconfigurationRemoteDesktopSecurityConfigurationRequestBuilder) {
    return NewItemRemotedesktopsecurityconfigurationRemoteDesktopSecurityConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveKey provides operations to call the removeKey method.
// returns a *ItemRemovekeyRemoveKeyRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) RemoveKey()(*ItemRemovekeyRemoveKeyRequestBuilder) {
    return NewItemRemovekeyRemoveKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemovePassword provides operations to call the removePassword method.
// returns a *ItemRemovepasswordRemovePasswordRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) RemovePassword()(*ItemRemovepasswordRemovePasswordRequestBuilder) {
    return NewItemRemovepasswordRemovePasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemRestoreRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Synchronization provides operations to manage the synchronization property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemSynchronizationRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) Synchronization()(*ItemSynchronizationRequestBuilder) {
    return NewItemSynchronizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a servicePrincipal object.
// returns a *RequestInformation when successful
func (m *ServicePrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a servicePrincipal object.
// returns a *RequestInformation when successful
func (m *ServicePrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*ItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPatchRequestInformation create a new servicePrincipal object if it doesn't exist, or update the properties of an existing servicePrincipal object.
// returns a *RequestInformation when successful
func (m *ServicePrincipalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemTransitivememberofTransitiveMemberOfRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*ItemTransitivememberofTransitiveMemberOfRequestBuilder) {
    return NewItemTransitivememberofTransitiveMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServicePrincipalItemRequestBuilder when successful
func (m *ServicePrincipalItemRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalItemRequestBuilder) {
    return NewServicePrincipalItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
