package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingcurrenciesBookingCurrenciesRequestBuilder provides operations to manage the bookingCurrencies property of the microsoft.graph.solutionsRoot entity.
type BookingcurrenciesBookingCurrenciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingcurrenciesBookingCurrenciesRequestBuilderGetQueryParameters get a list of bookingCurrency objects available to a Microsoft Bookings business.
type BookingcurrenciesBookingCurrenciesRequestBuilderGetQueryParameters struct {
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
// BookingcurrenciesBookingCurrenciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingcurrenciesBookingCurrenciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingcurrenciesBookingCurrenciesRequestBuilderGetQueryParameters
}
// BookingcurrenciesBookingCurrenciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingcurrenciesBookingCurrenciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBookingCurrencyId provides operations to manage the bookingCurrencies property of the microsoft.graph.solutionsRoot entity.
// returns a *BookingcurrenciesBookingCurrencyItemRequestBuilder when successful
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) ByBookingCurrencyId(bookingCurrencyId string)(*BookingcurrenciesBookingCurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bookingCurrencyId != "" {
        urlTplParams["bookingCurrency%2Did"] = bookingCurrencyId
    }
    return NewBookingcurrenciesBookingCurrencyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBookingcurrenciesBookingCurrenciesRequestBuilderInternal instantiates a new BookingcurrenciesBookingCurrenciesRequestBuilder and sets the default values.
func NewBookingcurrenciesBookingCurrenciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingcurrenciesBookingCurrenciesRequestBuilder) {
    m := &BookingcurrenciesBookingCurrenciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingCurrencies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBookingcurrenciesBookingCurrenciesRequestBuilder instantiates a new BookingcurrenciesBookingCurrenciesRequestBuilder and sets the default values.
func NewBookingcurrenciesBookingCurrenciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingcurrenciesBookingCurrenciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingcurrenciesBookingCurrenciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BookingcurrenciesCountRequestBuilder when successful
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) Count()(*BookingcurrenciesCountRequestBuilder) {
    return NewBookingcurrenciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of bookingCurrency objects available to a Microsoft Bookings business.
// returns a BookingCurrencyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingcurrency-list?view=graph-rest-1.0
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingcurrenciesBookingCurrenciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingCurrencyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyCollectionResponseable), nil
}
// Post create new navigation property to bookingCurrencies for solutions
// returns a BookingCurrencyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyable, requestConfiguration *BookingcurrenciesBookingCurrenciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingCurrencyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyable), nil
}
// ToGetRequestInformation get a list of bookingCurrency objects available to a Microsoft Bookings business.
// returns a *RequestInformation when successful
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingcurrenciesBookingCurrenciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to bookingCurrencies for solutions
// returns a *RequestInformation when successful
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingCurrencyable, requestConfiguration *BookingcurrenciesBookingCurrenciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BookingcurrenciesBookingCurrenciesRequestBuilder when successful
func (m *BookingcurrenciesBookingCurrenciesRequestBuilder) WithUrl(rawUrl string)(*BookingcurrenciesBookingCurrenciesRequestBuilder) {
    return NewBookingcurrenciesBookingCurrenciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
