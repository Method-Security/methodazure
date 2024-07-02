package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemValidatepermissionValidatePermissionRequestBuilder provides operations to call the validatePermission method.
type ItemItemsItemValidatepermissionValidatePermissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemValidatepermissionValidatePermissionRequestBuilderInternal instantiates a new ItemItemsItemValidatepermissionValidatePermissionRequestBuilder and sets the default values.
func NewItemItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    m := &ItemItemsItemValidatepermissionValidatePermissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/validatePermission", pathParameters),
    }
    return m
}
// NewItemItemsItemValidatepermissionValidatePermissionRequestBuilder instantiates a new ItemItemsItemValidatepermissionValidatePermissionRequestBuilder and sets the default values.
func NewItemItemsItemValidatepermissionValidatePermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validatePermission
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) Post(ctx context.Context, body ItemItemsItemValidatepermissionValidatePermissionPostRequestBodyable, requestConfiguration *ItemItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action validatePermission
// returns a *RequestInformation when successful
func (m *ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemValidatepermissionValidatePermissionPostRequestBodyable, requestConfiguration *ItemItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemValidatepermissionValidatePermissionRequestBuilder when successful
func (m *ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    return NewItemItemsItemValidatepermissionValidatePermissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}