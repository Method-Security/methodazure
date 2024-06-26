package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
type ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters get homeRealmDiscoveryPolicies from applications
type ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters
}
// ByHomeRealmDiscoveryPolicyId provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
// returns a *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder when successful
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) ByHomeRealmDiscoveryPolicyId(homeRealmDiscoveryPolicyId string)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if homeRealmDiscoveryPolicyId != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = homeRealmDiscoveryPolicyId
    }
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal instantiates a new ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    m := &ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/homeRealmDiscoveryPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder instantiates a new ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemHomerealmdiscoverypoliciesCountRequestBuilder when successful
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) Count()(*ItemHomerealmdiscoverypoliciesCountRequestBuilder) {
    return NewItemHomerealmdiscoverypoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get homeRealmDiscoveryPolicies from applications
// returns a HomeRealmDiscoveryPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHomeRealmDiscoveryPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable), nil
}
// ToGetRequestInformation get homeRealmDiscoveryPolicies from applications
// returns a *RequestInformation when successful
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder when successful
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
