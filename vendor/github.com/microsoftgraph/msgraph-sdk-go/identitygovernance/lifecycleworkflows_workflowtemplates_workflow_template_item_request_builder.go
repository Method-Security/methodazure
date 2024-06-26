package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetQueryParameters read the properties and relationships of a workflowTemplate object.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) {
    m := &LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder instantiates a new LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a workflowTemplate object.
// returns a WorkflowTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflowtemplate-get?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowTemplateable), nil
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowTemplate entity.
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) Tasks()(*LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a workflowTemplate object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
