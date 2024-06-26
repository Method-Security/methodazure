package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder builds and executes requests for operations under \identity\authenticationEventsFlows\{authenticationEventsFlow-id}\graph.externalUsersSelfServiceSignUpEventsFlow\onAuthenticationMethodLoadStart
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetQueryParameters required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetQueryParameters
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Get required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.
// returns a OnAuthenticationMethodLoadStartHandlerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnAuthenticationMethodLoadStartHandlerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnAuthenticationMethodLoadStartHandlerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnAuthenticationMethodLoadStartHandlerable), nil
}
// GraphOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp casts the previous resource to onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupGraphOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) GraphOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupGraphOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupGraphOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartOnAuthenticationMethodLoadStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}