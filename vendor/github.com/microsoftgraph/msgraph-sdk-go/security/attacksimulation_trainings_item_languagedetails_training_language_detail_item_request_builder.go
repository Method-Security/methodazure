package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder provides operations to manage the languageDetails property of the microsoft.graph.training entity.
type AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters language specific details on a training.
type AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters
}
// AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderInternal instantiates a new AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder and sets the default values.
func NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) {
    m := &AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/trainings/{training%2Did}/languageDetails/{trainingLanguageDetail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder instantiates a new AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder and sets the default values.
func NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property languageDetails for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get language specific details on a training.
// returns a TrainingLanguageDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTrainingLanguageDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable), nil
}
// Patch update the navigation property languageDetails in security
// returns a TrainingLanguageDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTrainingLanguageDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable), nil
}
// ToDeleteRequestInformation delete navigation property languageDetails for security
// returns a *RequestInformation when successful
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation language specific details on a training.
// returns a *RequestInformation when successful
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property languageDetails in security
// returns a *RequestInformation when successful
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TrainingLanguageDetailable, requestConfiguration *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder when successful
func (m *AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder) {
    return NewAttacksimulationTrainingsItemLanguagedetailsTrainingLanguageDetailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
