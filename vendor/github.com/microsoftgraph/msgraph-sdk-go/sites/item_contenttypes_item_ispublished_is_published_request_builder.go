package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemContenttypesItemIspublishedIsPublishedRequestBuilder provides operations to call the isPublished method.
type ItemContenttypesItemIspublishedIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal instantiates a new ItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    m := &ItemContenttypesItemIspublishedIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewItemContenttypesItemIspublishedIsPublishedRequestBuilder instantiates a new ItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the publishing status of a contentType in a content type hub site.
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a ItemContenttypesItemIspublishedIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemContenttypesItemIspublishedIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemContenttypesItemIspublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContenttypesItemIspublishedIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse check the publishing status of a contentType in a content type hub site.
// returns a ItemContenttypesItemIspublishedIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemContenttypesItemIspublishedIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *ItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemContenttypesItemIspublishedIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContenttypesItemIspublishedIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContenttypesItemIspublishedIsPublishedGetResponseable), nil
}
// ToGetRequestInformation check the publishing status of a contentType in a content type hub site.
// returns a *RequestInformation when successful
func (m *ItemContenttypesItemIspublishedIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *ItemContenttypesItemIspublishedIsPublishedRequestBuilder) WithUrl(rawUrl string)(*ItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}