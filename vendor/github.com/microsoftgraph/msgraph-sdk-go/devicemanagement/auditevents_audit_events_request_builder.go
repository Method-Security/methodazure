package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuditeventsAuditEventsRequestBuilder provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
type AuditeventsAuditEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuditeventsAuditEventsRequestBuilderGetQueryParameters list properties and relationships of the auditEvent objects.
type AuditeventsAuditEventsRequestBuilderGetQueryParameters struct {
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
// AuditeventsAuditEventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditeventsAuditEventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditeventsAuditEventsRequestBuilderGetQueryParameters
}
// AuditeventsAuditEventsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditeventsAuditEventsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuditEventId provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AuditeventsAuditEventItemRequestBuilder when successful
func (m *AuditeventsAuditEventsRequestBuilder) ByAuditEventId(auditEventId string)(*AuditeventsAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if auditEventId != "" {
        urlTplParams["auditEvent%2Did"] = auditEventId
    }
    return NewAuditeventsAuditEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuditeventsAuditEventsRequestBuilderInternal instantiates a new AuditeventsAuditEventsRequestBuilder and sets the default values.
func NewAuditeventsAuditEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditeventsAuditEventsRequestBuilder) {
    m := &AuditeventsAuditEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/auditEvents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuditeventsAuditEventsRequestBuilder instantiates a new AuditeventsAuditEventsRequestBuilder and sets the default values.
func NewAuditeventsAuditEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditeventsAuditEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditeventsAuditEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuditeventsCountRequestBuilder when successful
func (m *AuditeventsAuditEventsRequestBuilder) Count()(*AuditeventsCountRequestBuilder) {
    return NewAuditeventsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the auditEvent objects.
// returns a AuditEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-auditing-auditevent-list?view=graph-rest-1.0
func (m *AuditeventsAuditEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditeventsAuditEventsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuditEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventCollectionResponseable), nil
}
// GetAuditActivityTypesWithCategory provides operations to call the getAuditActivityTypes method.
// returns a *AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder when successful
func (m *AuditeventsAuditEventsRequestBuilder) GetAuditActivityTypesWithCategory(category *string)(*AuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilder) {
    return NewAuditeventsGetauditactivitytypeswithcategoryGetAuditActivityTypesWithCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, category)
}
// GetAuditCategories provides operations to call the getAuditCategories method.
// returns a *AuditeventsGetauditcategoriesGetAuditCategoriesRequestBuilder when successful
func (m *AuditeventsAuditEventsRequestBuilder) GetAuditCategories()(*AuditeventsGetauditcategoriesGetAuditCategoriesRequestBuilder) {
    return NewAuditeventsGetauditcategoriesGetAuditCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new auditEvent object.
// returns a AuditEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-auditing-auditevent-create?view=graph-rest-1.0
func (m *AuditeventsAuditEventsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventable, requestConfiguration *AuditeventsAuditEventsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuditEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventable), nil
}
// ToGetRequestInformation list properties and relationships of the auditEvent objects.
// returns a *RequestInformation when successful
func (m *AuditeventsAuditEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditeventsAuditEventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new auditEvent object.
// returns a *RequestInformation when successful
func (m *AuditeventsAuditEventsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuditEventable, requestConfiguration *AuditeventsAuditEventsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuditeventsAuditEventsRequestBuilder when successful
func (m *AuditeventsAuditEventsRequestBuilder) WithUrl(rawUrl string)(*AuditeventsAuditEventsRequestBuilder) {
    return NewAuditeventsAuditEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
