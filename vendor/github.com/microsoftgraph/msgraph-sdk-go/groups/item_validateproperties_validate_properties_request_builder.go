package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemValidatepropertiesValidatePropertiesRequestBuilder provides operations to call the validateProperties method.
type ItemValidatepropertiesValidatePropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemValidatepropertiesValidatePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemValidatepropertiesValidatePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemValidatepropertiesValidatePropertiesRequestBuilderInternal instantiates a new ItemValidatepropertiesValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatepropertiesValidatePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatepropertiesValidatePropertiesRequestBuilder) {
    m := &ItemValidatepropertiesValidatePropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/validateProperties", pathParameters),
    }
    return m
}
// NewItemValidatepropertiesValidatePropertiesRequestBuilder instantiates a new ItemValidatepropertiesValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatepropertiesValidatePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatepropertiesValidatePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemValidatepropertiesValidatePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use this API to determine whether a display name or mail nickname is valid before trying to update a Microsoft 365 group. To validate the properties before creating a group, use the directoryobject:validateProperties function. The following policy validations are performed for the display name and mail nickname properties: This API only returns the first validation failure that is encountered. If the properties fail multiple validations, only the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy. To learn more about configuring naming policies, see Configure naming policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-validateproperties?view=graph-rest-1.0
func (m *ItemValidatepropertiesValidatePropertiesRequestBuilder) Post(ctx context.Context, body ItemValidatepropertiesValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatepropertiesValidatePropertiesRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation validate that a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use this API to determine whether a display name or mail nickname is valid before trying to update a Microsoft 365 group. To validate the properties before creating a group, use the directoryobject:validateProperties function. The following policy validations are performed for the display name and mail nickname properties: This API only returns the first validation failure that is encountered. If the properties fail multiple validations, only the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy. To learn more about configuring naming policies, see Configure naming policy.
// returns a *RequestInformation when successful
func (m *ItemValidatepropertiesValidatePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemValidatepropertiesValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatepropertiesValidatePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemValidatepropertiesValidatePropertiesRequestBuilder when successful
func (m *ItemValidatepropertiesValidatePropertiesRequestBuilder) WithUrl(rawUrl string)(*ItemValidatepropertiesValidatePropertiesRequestBuilder) {
    return NewItemValidatepropertiesValidatePropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
