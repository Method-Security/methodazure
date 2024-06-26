package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder provides operations to call the norm_S_Inv method.
type ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/norm_S_Inv", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action norm_S_Inv
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action norm_S_Inv
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsNorm_s_invNorm_S_InvRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
