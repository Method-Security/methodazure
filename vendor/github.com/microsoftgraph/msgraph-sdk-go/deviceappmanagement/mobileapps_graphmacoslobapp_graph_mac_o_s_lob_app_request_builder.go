package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder casts the previous resource to macOSLobApp.
type MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters get the items of type microsoft.graph.macOSLobApp in the microsoft.graph.mobileApp collection
type MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters struct {
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
// MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters
}
// NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal instantiates a new MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    m := &MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/graph.macOSLobApp{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder instantiates a new MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsGraphmacoslobappCountRequestBuilder when successful
func (m *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) Count()(*MobileappsGraphmacoslobappCountRequestBuilder) {
    return NewMobileappsGraphmacoslobappCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.macOSLobApp in the microsoft.graph.mobileApp collection
// returns a MacOSLobAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSLobAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMacOSLobAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSLobAppCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.macOSLobApp in the microsoft.graph.mobileApp collection
// returns a *RequestInformation when successful
func (m *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder when successful
func (m *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    return NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
