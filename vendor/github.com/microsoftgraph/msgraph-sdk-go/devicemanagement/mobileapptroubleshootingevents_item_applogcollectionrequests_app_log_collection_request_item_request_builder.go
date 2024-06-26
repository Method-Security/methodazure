package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters read properties and relationships of the appLogCollectionRequest object.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    m := &MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
// returns a *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) CreateDownloadUrl()(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes a appLogCollectionRequest.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applogcollectionrequest-delete?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the appLogCollectionRequest object.
// returns a AppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applogcollectionrequest-get?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable), nil
}
// Patch update the properties of a appLogCollectionRequest object.
// returns a AppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-applogcollectionrequest-update?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable), nil
}
// ToDeleteRequestInformation deletes a appLogCollectionRequest.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the appLogCollectionRequest object.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a appLogCollectionRequest object.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) WithUrl(rawUrl string)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
