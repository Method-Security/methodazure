package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderInternal instantiates a new ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder and sets the default values.
func NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) {
    m := &ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/completeMigration", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder instantiates a new ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder and sets the default values.
func NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete the message migration process by removing migration mode from a team. Migration mode is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a completeMigration request is made, you can't import additional messages into the team. You can add members to the team after the request returns a successful response.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-completemigration?view=graph-rest-1.0
func (m *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation complete the message migration process by removing migration mode from a team. Migration mode is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a completeMigration request is made, you can't import additional messages into the team. You can add members to the team after the request returns a successful response.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
