package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GraphroomlistGraphRoomListRequestBuilder casts the previous resource to roomList.
type GraphroomlistGraphRoomListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GraphroomlistGraphRoomListRequestBuilderGetQueryParameters get the items of type microsoft.graph.roomList in the microsoft.graph.place collection
type GraphroomlistGraphRoomListRequestBuilderGetQueryParameters struct {
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
// GraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GraphroomlistGraphRoomListRequestBuilderGetQueryParameters
}
// NewGraphroomlistGraphRoomListRequestBuilderInternal instantiates a new GraphroomlistGraphRoomListRequestBuilder and sets the default values.
func NewGraphroomlistGraphRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphroomlistGraphRoomListRequestBuilder) {
    m := &GraphroomlistGraphRoomListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/graph.roomList{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGraphroomlistGraphRoomListRequestBuilder instantiates a new GraphroomlistGraphRoomListRequestBuilder and sets the default values.
func NewGraphroomlistGraphRoomListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphroomlistGraphRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGraphroomlistGraphRoomListRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GraphroomlistCountRequestBuilder when successful
func (m *GraphroomlistGraphRoomListRequestBuilder) Count()(*GraphroomlistCountRequestBuilder) {
    return NewGraphroomlistCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.roomList in the microsoft.graph.place collection
// returns a RoomListCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GraphroomlistGraphRoomListRequestBuilder) Get(ctx context.Context, requestConfiguration *GraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoomListCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRoomListCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoomListCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.roomList in the microsoft.graph.place collection
// returns a *RequestInformation when successful
func (m *GraphroomlistGraphRoomListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GraphroomlistGraphRoomListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GraphroomlistGraphRoomListRequestBuilder when successful
func (m *GraphroomlistGraphRoomListRequestBuilder) WithUrl(rawUrl string)(*GraphroomlistGraphRoomListRequestBuilder) {
    return NewGraphroomlistGraphRoomListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}