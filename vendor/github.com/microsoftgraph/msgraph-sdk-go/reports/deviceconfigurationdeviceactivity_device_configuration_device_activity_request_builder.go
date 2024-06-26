package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder provides operations to call the deviceConfigurationDeviceActivity method.
type DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderInternal instantiates a new DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder and sets the default values.
func NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) {
    m := &DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/deviceConfigurationDeviceActivity()", pathParameters),
    }
    return m
}
// NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder instantiates a new DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder and sets the default values.
func NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get metadata for the device configuration device activity report
// returns a Reportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-reportroot-deviceconfigurationdeviceactivity?view=graph-rest-1.0
func (m *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Reportable), nil
}
// ToGetRequestInformation metadata for the device configuration device activity report
// returns a *RequestInformation when successful
func (m *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder when successful
func (m *DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewDeviceconfigurationdeviceactivityDeviceConfigurationDeviceActivityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
