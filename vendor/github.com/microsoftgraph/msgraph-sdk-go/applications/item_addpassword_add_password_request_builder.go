package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAddpasswordAddPasswordRequestBuilder provides operations to call the addPassword method.
type ItemAddpasswordAddPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAddpasswordAddPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAddpasswordAddPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAddpasswordAddPasswordRequestBuilderInternal instantiates a new ItemAddpasswordAddPasswordRequestBuilder and sets the default values.
func NewItemAddpasswordAddPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddpasswordAddPasswordRequestBuilder) {
    m := &ItemAddpasswordAddPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/addPassword", pathParameters),
    }
    return m
}
// NewItemAddpasswordAddPasswordRequestBuilder instantiates a new ItemAddpasswordAddPasswordRequestBuilder and sets the default values.
func NewItemAddpasswordAddPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddpasswordAddPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddpasswordAddPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post adds a strong password or secret to an application. You can also add passwords while creating the application.
// returns a PasswordCredentialable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-addpassword?view=graph-rest-1.0
func (m *ItemAddpasswordAddPasswordRequestBuilder) Post(ctx context.Context, body ItemAddpasswordAddPasswordPostRequestBodyable, requestConfiguration *ItemAddpasswordAddPasswordRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePasswordCredentialFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable), nil
}
// ToPostRequestInformation adds a strong password or secret to an application. You can also add passwords while creating the application.
// returns a *RequestInformation when successful
func (m *ItemAddpasswordAddPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddpasswordAddPasswordPostRequestBodyable, requestConfiguration *ItemAddpasswordAddPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAddpasswordAddPasswordRequestBuilder when successful
func (m *ItemAddpasswordAddPasswordRequestBuilder) WithUrl(rawUrl string)(*ItemAddpasswordAddPasswordRequestBuilder) {
    return NewItemAddpasswordAddPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
