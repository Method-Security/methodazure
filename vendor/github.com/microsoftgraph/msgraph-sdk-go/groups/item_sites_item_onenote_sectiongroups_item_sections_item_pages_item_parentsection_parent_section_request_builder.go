package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetQueryParameters the section that contains the page. Read-only.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetQueryParameters
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderInternal instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) {
    m := &ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/parentSection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the section that contains the page. Read-only.
// returns a OnenoteSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable), nil
}
// ToGetRequestInformation the section that contains the page. Read-only.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesItemParentsectionParentSectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
