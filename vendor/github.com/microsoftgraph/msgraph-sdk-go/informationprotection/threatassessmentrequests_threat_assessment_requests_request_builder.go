package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
type ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters retrieve a list of threatAssessmentRequest objects. A threat assessment request can be one of the following types:
type ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters struct {
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
// ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters
}
// ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByThreatAssessmentRequestId provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
// returns a *ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ByThreatAssessmentRequestId(threatAssessmentRequestId string)(*ThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if threatAssessmentRequestId != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = threatAssessmentRequestId
    }
    return NewThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal instantiates a new ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder and sets the default values.
func NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    m := &ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/threatAssessmentRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder instantiates a new ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder and sets the default values.
func NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatassessmentrequestsCountRequestBuilder when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Count()(*ThreatassessmentrequestsCountRequestBuilder) {
    return NewThreatassessmentrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of threatAssessmentRequest objects. A threat assessment request can be one of the following types:
// returns a ThreatAssessmentRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotection-list-threatassessmentrequests?view=graph-rest-1.0
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateThreatAssessmentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestCollectionResponseable), nil
}
// Post create a new threat assessment request. A threat assessment request can be one of the following types:
// returns a ThreatAssessmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotection-post-threatassessmentrequests?view=graph-rest-1.0
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable), nil
}
// ToGetRequestInformation retrieve a list of threatAssessmentRequest objects. A threat assessment request can be one of the following types:
// returns a *RequestInformation when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new threat assessment request. A threat assessment request can be one of the following types:
// returns a *RequestInformation when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ThreatAssessmentRequestable, requestConfiguration *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder when successful
func (m *ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) WithUrl(rawUrl string)(*ThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    return NewThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
