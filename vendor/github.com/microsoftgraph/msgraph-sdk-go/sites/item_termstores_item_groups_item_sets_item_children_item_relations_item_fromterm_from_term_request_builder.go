package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
type ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetQueryParameters the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
type ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetQueryParameters
}
// NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderInternal instantiates a new ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder and sets the default values.
func NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) {
    m := &ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}/relations/{relation%2Did}/fromTerm{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder instantiates a new ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder and sets the default values.
func NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// ToGetRequestInformation the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
// returns a *RequestInformation when successful
func (m *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder when successful
func (m *ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) {
    return NewItemTermstoresItemGroupsItemSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}