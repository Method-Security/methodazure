package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityDetail method.
type Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal instantiates a new Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    m := &Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder instantiates a new Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getOffice365GroupsActivityDetail
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder when successful
func (m *Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewGetoffice365groupsactivitydetailwithperiodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
