package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder provides operations to call the createLink method.
type ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderInternal instantiates a new ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder and sets the default values.
func NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) {
    m := &ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/createLink", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder instantiates a new ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder and sets the default values.
func NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createLink
// returns a Permissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) Post(ctx context.Context, body ItemListsItemItemsItemCreatelinkCreateLinkPostRequestBodyable, requestConfiguration *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable), nil
}
// ToPostRequestInformation invoke action createLink
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemListsItemItemsItemCreatelinkCreateLinkPostRequestBodyable, requestConfiguration *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
