package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder provides operations to call the cleanWindowsDevice method.
type ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal instantiates a new ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    m := &ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/cleanWindowsDevice", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder instantiates a new ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clean Windows device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-cleanwindowsdevice?view=graph-rest-1.0
func (m *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) Post(ctx context.Context, body ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDevicePostRequestBodyable, requestConfiguration *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation clean Windows device
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDevicePostRequestBodyable, requestConfiguration *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
