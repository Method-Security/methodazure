package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder provides operations to call the accept method.
type ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/accept", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post accept the specified event in a user calendar.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-accept?view=graph-rest-1.0
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation accept the specified event in a user calendar.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}