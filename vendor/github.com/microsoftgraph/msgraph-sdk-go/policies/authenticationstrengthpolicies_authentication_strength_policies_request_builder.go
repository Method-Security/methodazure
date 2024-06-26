package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
type AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetQueryParameters get a list of the authenticationStrengthPolicy objects and their properties. This API returns both built-in and custom policies.
type AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetQueryParameters struct {
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
// AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetQueryParameters
}
// AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationStrengthPolicyId provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
// returns a *AuthenticationstrengthpoliciesAuthenticationStrengthPolicyItemRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) ByAuthenticationStrengthPolicyId(authenticationStrengthPolicyId string)(*AuthenticationstrengthpoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationStrengthPolicyId != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = authenticationStrengthPolicyId
    }
    return NewAuthenticationstrengthpoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderInternal instantiates a new AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) {
    m := &AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder instantiates a new AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationstrengthpoliciesCountRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) Count()(*AuthenticationstrengthpoliciesCountRequestBuilder) {
    return NewAuthenticationstrengthpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the authenticationStrengthPolicy objects and their properties. This API returns both built-in and custom policies.
// returns a AuthenticationStrengthPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthroot-list-policies?view=graph-rest-1.0
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationStrengthPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyCollectionResponseable), nil
}
// Post create a new custom authenticationStrengthPolicy object.
// returns a AuthenticationStrengthPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthroot-post-policies?view=graph-rest-1.0
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable), nil
}
// ToGetRequestInformation get a list of the authenticationStrengthPolicy objects and their properties. This API returns both built-in and custom policies.
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new custom authenticationStrengthPolicy object.
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder) {
    return NewAuthenticationstrengthpoliciesAuthenticationStrengthPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
