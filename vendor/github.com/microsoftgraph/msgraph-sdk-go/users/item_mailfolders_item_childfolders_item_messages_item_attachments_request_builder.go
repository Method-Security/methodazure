package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder provides operations to manage the attachments property of the microsoft.graph.message entity.
type ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters the fileAttachment and itemAttachment attachments for the message.
type ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters struct {
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
// ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters
}
// ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttachmentId provides operations to manage the attachments property of the microsoft.graph.message entity.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) ByAttachmentId(attachmentId string)(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attachmentId != "" {
        urlTplParams["attachment%2Did"] = attachmentId
    }
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderInternal instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) {
    m := &ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/attachments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCountRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) Count()(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCountRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) CreateUploadSession()(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the fileAttachment and itemAttachment attachments for the message.
// returns a AttachmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentCollectionResponseable), nil
}
// Post create new navigation property to attachments for users
// returns a Attachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Attachmentable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Attachmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Attachmentable), nil
}
// ToGetRequestInformation the fileAttachment and itemAttachment attachments for the message.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attachments for users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Attachmentable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemAttachmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
