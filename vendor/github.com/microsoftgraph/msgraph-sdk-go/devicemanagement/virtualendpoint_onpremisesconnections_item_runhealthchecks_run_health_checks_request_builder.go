package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder provides operations to call the runHealthChecks method.
type VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderInternal instantiates a new VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) {
    m := &VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/runHealthChecks", pathParameters),
    }
    return m
}
// NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder instantiates a new VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run health checks on the cloudPcOnPremisesConnection object. It triggers a new health check for the cloudPcOnPremisesConnection (../resources/cloudpconpremisesconnection.md) object and changes the healthCheckStatus and [healthCheckStatusDetail properties when check finished.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-runhealthcheck?view=graph-rest-1.0
func (m *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation run health checks on the cloudPcOnPremisesConnection object. It triggers a new health check for the cloudPcOnPremisesConnection (../resources/cloudpconpremisesconnection.md) object and changes the healthCheckStatus and [healthCheckStatusDetail properties when check finished.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsItemRunhealthchecksRunHealthChecksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
