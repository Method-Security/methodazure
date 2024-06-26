package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemPlaypromptPlayPromptRequestBuilder provides operations to call the playPrompt method.
type CallsItemPlaypromptPlayPromptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemPlaypromptPlayPromptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemPlaypromptPlayPromptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemPlaypromptPlayPromptRequestBuilderInternal instantiates a new CallsItemPlaypromptPlayPromptRequestBuilder and sets the default values.
func NewCallsItemPlaypromptPlayPromptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemPlaypromptPlayPromptRequestBuilder) {
    m := &CallsItemPlaypromptPlayPromptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/playPrompt", pathParameters),
    }
    return m
}
// NewCallsItemPlaypromptPlayPromptRequestBuilder instantiates a new CallsItemPlaypromptPlayPromptRequestBuilder and sets the default values.
func NewCallsItemPlaypromptPlayPromptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemPlaypromptPlayPromptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemPlaypromptPlayPromptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post play a prompt in the call. For more information about how to handle operations, see commsOperation
// returns a PlayPromptOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-playprompt?view=graph-rest-1.0
func (m *CallsItemPlaypromptPlayPromptRequestBuilder) Post(ctx context.Context, body CallsItemPlaypromptPlayPromptPostRequestBodyable, requestConfiguration *CallsItemPlaypromptPlayPromptRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlayPromptOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlayPromptOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlayPromptOperationable), nil
}
// ToPostRequestInformation play a prompt in the call. For more information about how to handle operations, see commsOperation
// returns a *RequestInformation when successful
func (m *CallsItemPlaypromptPlayPromptRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemPlaypromptPlayPromptPostRequestBodyable, requestConfiguration *CallsItemPlaypromptPlayPromptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemPlaypromptPlayPromptRequestBuilder when successful
func (m *CallsItemPlaypromptPlayPromptRequestBuilder) WithUrl(rawUrl string)(*CallsItemPlaypromptPlayPromptRequestBuilder) {
    return NewCallsItemPlaypromptPlayPromptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
