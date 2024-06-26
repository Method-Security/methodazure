package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetQueryParameters retrieve the list of delegatedPermissionClassification currently configured for the delegated permissions exposed by an API.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetQueryParameters struct {
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
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetQueryParameters
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDelegatedPermissionClassificationId provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) ByDelegatedPermissionClassificationId(delegatedPermissionClassificationId string)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if delegatedPermissionClassificationId != "" {
        urlTplParams["delegatedPermissionClassification%2Did"] = delegatedPermissionClassificationId
    }
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderInternal instantiates a new ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder and sets the default values.
func NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) {
    m := &ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/delegatedPermissionClassifications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder instantiates a new ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder and sets the default values.
func NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDelegatedpermissionclassificationsCountRequestBuilder when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) Count()(*ItemDelegatedpermissionclassificationsCountRequestBuilder) {
    return NewItemDelegatedpermissionclassificationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the list of delegatedPermissionClassification currently configured for the delegated permissions exposed by an API.
// returns a DelegatedPermissionClassificationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-delegatedpermissionclassifications?view=graph-rest-1.0
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedPermissionClassificationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationCollectionResponseable), nil
}
// Post classify a delegated permission by adding a delegatedPermissionClassification to the servicePrincipal representing the API.
// returns a DelegatedPermissionClassificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-post-delegatedpermissionclassifications?view=graph-rest-1.0
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationable, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedPermissionClassificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationable), nil
}
// ToGetRequestInformation retrieve the list of delegatedPermissionClassification currently configured for the delegated permissions exposed by an API.
// returns a *RequestInformation when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation classify a delegated permission by adding a delegatedPermissionClassification to the servicePrincipal representing the API.
// returns a *RequestInformation when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedPermissionClassificationable, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) WithUrl(rawUrl string)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder) {
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
