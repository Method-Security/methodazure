package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder provides operations to manage the citationTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
type LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetQueryParameters the specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
type LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetQueryParameters
}
// NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderInternal instantiates a new LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) {
    m := &LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/retentionLabels/{retentionLabel%2Did}/descriptors/citationTemplate{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder instantiates a new LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
// returns a CitationTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CitationTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateCitationTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CitationTemplateable), nil
}
// ToGetRequestInformation the specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) WithUrl(rawUrl string)(*LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
