package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder provides operations to call the getEmailAppUsageUserDetail method.
type GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderInternal instantiates a new GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    m := &GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getEmailAppUsageUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder instantiates a new GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get details about which activities users performed on the various email apps.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get details about which activities users performed on the various email apps.
// returns a *RequestInformation when successful
func (m *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder when successful
func (m *GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewGetemailappusageuserdetailwithdateGetEmailAppUsageUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}