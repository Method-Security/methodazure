package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivity entity.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters exposes the driveItem that was the target of this activity.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}/driveItem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) Content()(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the driveItem that was the target of this activity.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ToGetRequestInformation exposes the driveItem that was the target of this activity.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}