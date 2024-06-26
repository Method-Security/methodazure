package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder provides operations to call the ecma_Ceiling method.
type ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/ecma_Ceiling", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action ecma_Ceiling
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action ecma_Ceiling
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsEcma_ceilingEcma_CeilingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
