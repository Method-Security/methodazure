package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
type ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetQueryParameters get the properties of a certificateBasedAuthConfiguration object.
type ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetQueryParameters
}
// NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal instantiates a new ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder and sets the default values.
func NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    m := &ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder instantiates a new ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder and sets the default values.
func NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a certificateBasedAuthConfiguration object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificatebasedauthconfiguration-delete?view=graph-rest-1.0
func (m *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of a certificateBasedAuthConfiguration object.
// returns a CertificateBasedAuthConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificatebasedauthconfiguration-get?view=graph-rest-1.0
func (m *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CertificateBasedAuthConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCertificateBasedAuthConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CertificateBasedAuthConfigurationable), nil
}
// ToDeleteRequestInformation delete a certificateBasedAuthConfiguration object.
// returns a *RequestInformation when successful
func (m *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties of a certificateBasedAuthConfiguration object.
// returns a *RequestInformation when successful
func (m *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder when successful
func (m *ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*ItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    return NewItemCertificatebasedauthconfigurationCertificateBasedAuthConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
