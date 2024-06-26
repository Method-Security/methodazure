package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder provides operations to manage the serviceConfigurationRecords property of the microsoft.graph.domain entity.
type ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetQueryParameters retrieves a list of domainDnsRecord objects needed to enable services for the domain. Use the returned list to add records to the zone file of the domain. This can be done through the domain registrar or DNS server configuration.
type ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetQueryParameters struct {
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
// ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetQueryParameters
}
// ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDomainDnsRecordId provides operations to manage the serviceConfigurationRecords property of the microsoft.graph.domain entity.
// returns a *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder when successful
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) ByDomainDnsRecordId(domainDnsRecordId string)(*ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if domainDnsRecordId != "" {
        urlTplParams["domainDnsRecord%2Did"] = domainDnsRecordId
    }
    return NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderInternal instantiates a new ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder and sets the default values.
func NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) {
    m := &ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/serviceConfigurationRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder instantiates a new ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder and sets the default values.
func NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemServiceconfigurationrecordsCountRequestBuilder when successful
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) Count()(*ItemServiceconfigurationrecordsCountRequestBuilder) {
    return NewItemServiceconfigurationrecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieves a list of domainDnsRecord objects needed to enable services for the domain. Use the returned list to add records to the zone file of the domain. This can be done through the domain registrar or DNS server configuration.
// returns a DomainDnsRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/domain-list-serviceconfigurationrecords?view=graph-rest-1.0
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDomainDnsRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordCollectionResponseable), nil
}
// Post create new navigation property to serviceConfigurationRecords for domains
// returns a DomainDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, requestConfiguration *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDomainDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable), nil
}
// ToGetRequestInformation retrieves a list of domainDnsRecord objects needed to enable services for the domain. Use the returned list to add records to the zone file of the domain. This can be done through the domain registrar or DNS server configuration.
// returns a *RequestInformation when successful
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to serviceConfigurationRecords for domains
// returns a *RequestInformation when successful
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, requestConfiguration *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder when successful
func (m *ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) WithUrl(rawUrl string)(*ItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder) {
    return NewItemServiceconfigurationrecordsServiceConfigurationRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
