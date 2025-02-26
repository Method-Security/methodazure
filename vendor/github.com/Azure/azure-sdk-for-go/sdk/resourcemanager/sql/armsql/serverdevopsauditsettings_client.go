//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerDevOpsAuditSettingsClient contains the methods for the ServerDevOpsAuditSettings group.
// Don't use this type directly, use NewServerDevOpsAuditSettingsClient() instead.
type ServerDevOpsAuditSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerDevOpsAuditSettingsClient creates a new instance of ServerDevOpsAuditSettingsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerDevOpsAuditSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerDevOpsAuditSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerDevOpsAuditSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - devOpsAuditingSettingsName - The name of the devops audit settings. This should always be 'default'.
//   - parameters - Properties of DevOps audit settings
//   - options - ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.BeginCreateOrUpdate
//     method.
func (client *ServerDevOpsAuditSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ServerDevOpsAuditSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerDevOpsAuditSettingsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - devOpsAuditingSettingsName - The name of the devops audit settings. This should always be 'default'.
//   - options - ServerDevOpsAuditSettingsClientGetOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.Get
//     method.
func (client *ServerDevOpsAuditSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsClientGetOptions) (ServerDevOpsAuditSettingsClientGetResponse, error) {
	var err error
	const operationName = "ServerDevOpsAuditSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, options)
	if err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerDevOpsAuditSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerDevOpsAuditSettingsClient) getHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsClientGetResponse, error) {
	result := ServerDevOpsAuditSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditingSettings); err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Lists DevOps audit settings of a server.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerDevOpsAuditSettingsClientListByServerOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.NewListByServerPager
//     method.
func (client *ServerDevOpsAuditSettingsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsClientListByServerOptions) *runtime.Pager[ServerDevOpsAuditSettingsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerDevOpsAuditSettingsClientListByServerResponse]{
		More: func(page ServerDevOpsAuditSettingsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerDevOpsAuditSettingsClientListByServerResponse) (ServerDevOpsAuditSettingsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerDevOpsAuditSettingsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerDevOpsAuditSettingsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerDevOpsAuditSettingsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerDevOpsAuditSettingsClient) listByServerHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsClientListByServerResponse, error) {
	result := ServerDevOpsAuditSettingsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditSettingsListResult); err != nil {
		return ServerDevOpsAuditSettingsClientListByServerResponse{}, err
	}
	return result, nil
}
