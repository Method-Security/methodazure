package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters another way to input protected apps through xml files
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}/protectedAppLockerFiles/{windowsInformationProtectionAppLockerFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property protectedAppLockerFiles for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get another way to input protected apps through xml files
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property protectedAppLockerFiles in deviceAppManagement
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property protectedAppLockerFiles for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation another way to input protected apps through xml files
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property protectedAppLockerFiles in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
