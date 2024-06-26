package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DetectedappsItemManageddevicesManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.detectedApp entity.
type DetectedappsItemManageddevicesManagedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetQueryParameters the devices that have the discovered application installed
type DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetQueryParameters struct {
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
// DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetQueryParameters
}
// ByManagedDeviceId provides operations to manage the managedDevices property of the microsoft.graph.detectedApp entity.
// returns a *DetectedappsItemManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *DetectedappsItemManageddevicesManagedDevicesRequestBuilder) ByManagedDeviceId(managedDeviceId string)(*DetectedappsItemManageddevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceId != "" {
        urlTplParams["managedDevice%2Did"] = managedDeviceId
    }
    return NewDetectedappsItemManageddevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDetectedappsItemManageddevicesManagedDevicesRequestBuilderInternal instantiates a new DetectedappsItemManageddevicesManagedDevicesRequestBuilder and sets the default values.
func NewDetectedappsItemManageddevicesManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DetectedappsItemManageddevicesManagedDevicesRequestBuilder) {
    m := &DetectedappsItemManageddevicesManagedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/detectedApps/{detectedApp%2Did}/managedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDetectedappsItemManageddevicesManagedDevicesRequestBuilder instantiates a new DetectedappsItemManageddevicesManagedDevicesRequestBuilder and sets the default values.
func NewDetectedappsItemManageddevicesManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DetectedappsItemManageddevicesManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDetectedappsItemManageddevicesManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DetectedappsItemManageddevicesCountRequestBuilder when successful
func (m *DetectedappsItemManageddevicesManagedDevicesRequestBuilder) Count()(*DetectedappsItemManageddevicesCountRequestBuilder) {
    return NewDetectedappsItemManageddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the devices that have the discovered application installed
// returns a ManagedDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DetectedappsItemManageddevicesManagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceCollectionResponseable), nil
}
// ToGetRequestInformation the devices that have the discovered application installed
// returns a *RequestInformation when successful
func (m *DetectedappsItemManageddevicesManagedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DetectedappsItemManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DetectedappsItemManageddevicesManagedDevicesRequestBuilder when successful
func (m *DetectedappsItemManageddevicesManagedDevicesRequestBuilder) WithUrl(rawUrl string)(*DetectedappsItemManageddevicesManagedDevicesRequestBuilder) {
    return NewDetectedappsItemManageddevicesManagedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}