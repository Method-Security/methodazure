package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder provides operations to call the unsetReaction method.
type ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal instantiates a new ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    m := &ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/unsetReaction", pathParameters),
    }
    return m
}
// NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder instantiates a new ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unsetReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) Post(ctx context.Context, body ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action unsetReaction
// returns a *RequestInformation when successful
func (m *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemChatsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
