package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
type GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getTeamsDeviceUsageDistributionUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the number of Microsoft Teams unique users by device type over the selected time period.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getteamsdeviceusagedistributionusercounts?view=graph-rest-1.0
func (m *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the number of Microsoft Teams unique users by device type over the selected time period.
// returns a *RequestInformation when successful
func (m *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder when successful
func (m *GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetteamsdeviceusagedistributionusercountswithperiodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}