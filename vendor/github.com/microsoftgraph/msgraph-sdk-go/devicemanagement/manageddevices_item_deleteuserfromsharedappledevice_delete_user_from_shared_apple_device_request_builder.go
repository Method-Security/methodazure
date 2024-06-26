package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder provides operations to call the deleteUserFromSharedAppleDevice method.
type ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal instantiates a new ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    m := &ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deleteUserFromSharedAppleDevice", pathParameters),
    }
    return m
}
// NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder instantiates a new ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete user from shared Apple device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-deleteuserfromsharedappledevice?view=graph-rest-1.0
func (m *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) Post(ctx context.Context, body ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDevicePostRequestBodyable, requestConfiguration *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation delete user from shared Apple device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDevicePostRequestBodyable, requestConfiguration *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
