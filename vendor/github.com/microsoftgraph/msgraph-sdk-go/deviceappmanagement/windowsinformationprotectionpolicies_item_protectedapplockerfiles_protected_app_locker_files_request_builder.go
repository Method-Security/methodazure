package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters another way to input protected apps through xml files
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters struct {
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
// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsInformationProtectionAppLockerFileId provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ByWindowsInformationProtectionAppLockerFileId(windowsInformationProtectionAppLockerFileId string)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionAppLockerFileId != "" {
        urlTplParams["windowsInformationProtectionAppLockerFile%2Did"] = windowsInformationProtectionAppLockerFileId
    }
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal instantiates a new WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    m := &WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}/protectedAppLockerFiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder instantiates a new WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Count()(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get another way to input protected apps through xml files
// returns a WindowsInformationProtectionAppLockerFileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLockerFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileCollectionResponseable), nil
}
// Post create new navigation property to protectedAppLockerFiles for deviceAppManagement
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable), nil
}
// ToGetRequestInformation another way to input protected apps through xml files
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to protectedAppLockerFiles for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
