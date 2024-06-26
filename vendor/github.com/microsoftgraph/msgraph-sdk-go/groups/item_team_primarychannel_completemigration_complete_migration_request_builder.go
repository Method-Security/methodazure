package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderInternal instantiates a new ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) {
    m := &ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/completeMigration", pathParameters),
    }
    return m
}
// NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder instantiates a new ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete the message migration process by removing migration mode from a channel in a team. Migration mode is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a completeMigration request is made, you can't import additional messages into the team. You can add members to the team after the request returns a successful response.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-completemigration?view=graph-rest-1.0
func (m *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation complete the message migration process by removing migration mode from a channel in a team. Migration mode is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a completeMigration request is made, you can't import additional messages into the team. You can add members to the team after the request returns a successful response.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemTeamPrimarychannelCompletemigrationCompleteMigrationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
