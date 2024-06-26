package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the getGroupArchivedPrintJobs method.
type GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters get a list of archived print jobs for a particular group.
type GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getGroupArchivedPrintJobs(groupId='{groupId}',startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if groupId != nil {
        m.BaseRequestBuilder.PathParameters["groupId"] = *groupId
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get a list of archived print jobs for a particular group.
// Deprecated: This method is obsolete. Use GetAsGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getgrouparchivedprintjobs?view=graph-rest-1.0
func (m *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponse get a list of archived print jobs for a particular group.
// returns a GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getgrouparchivedprintjobs?view=graph-rest-1.0
func (m *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation get a list of archived print jobs for a particular group.
// returns a *RequestInformation when successful
func (m *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*GetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
