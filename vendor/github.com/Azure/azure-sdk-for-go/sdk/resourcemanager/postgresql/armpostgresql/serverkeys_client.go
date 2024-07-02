//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// ServerKeysClient contains the methods for the ServerKeys group.
// Don't use this type directly, use NewServerKeysClient() instead.
type ServerKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerKeysClient creates a new instance of ServerKeysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerKeysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a PostgreSQL Server key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - serverName - The name of the server.
//   - keyName - The name of the PostgreSQL Server key to be operated on (updated or created).
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - The requested PostgreSQL Server key resource state.
//   - options - ServerKeysClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerKeysClient.BeginCreateOrUpdate
//     method.
func (client *ServerKeysClient) BeginCreateOrUpdate(ctx context.Context, serverName string, keyName string, resourceGroupName string, parameters ServerKey, options *ServerKeysClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerKeysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, serverName, keyName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerKeysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerKeysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a PostgreSQL Server key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
func (client *ServerKeysClient) createOrUpdate(ctx context.Context, serverName string, keyName string, resourceGroupName string, parameters ServerKey, options *ServerKeysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerKeysClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, serverName, keyName, resourceGroupName, parameters, options)
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
func (client *ServerKeysClient) createOrUpdateCreateRequest(ctx context.Context, serverName string, keyName string, resourceGroupName string, parameters ServerKey, options *ServerKeysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}"
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the PostgreSQL Server key with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - serverName - The name of the server.
//   - keyName - The name of the PostgreSQL Server key to be deleted.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ServerKeysClientBeginDeleteOptions contains the optional parameters for the ServerKeysClient.BeginDelete method.
func (client *ServerKeysClient) BeginDelete(ctx context.Context, serverName string, keyName string, resourceGroupName string, options *ServerKeysClientBeginDeleteOptions) (*runtime.Poller[ServerKeysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, serverName, keyName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerKeysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerKeysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the PostgreSQL Server key with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
func (client *ServerKeysClient) deleteOperation(ctx context.Context, serverName string, keyName string, resourceGroupName string, options *ServerKeysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerKeysClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, serverName, keyName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerKeysClient) deleteCreateRequest(ctx context.Context, serverName string, keyName string, resourceGroupName string, options *ServerKeysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}"
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a PostgreSQL Server key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - keyName - The name of the PostgreSQL Server key to be retrieved.
//   - options - ServerKeysClientGetOptions contains the optional parameters for the ServerKeysClient.Get method.
func (client *ServerKeysClient) Get(ctx context.Context, resourceGroupName string, serverName string, keyName string, options *ServerKeysClientGetOptions) (ServerKeysClientGetResponse, error) {
	var err error
	const operationName = "ServerKeysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, keyName, options)
	if err != nil {
		return ServerKeysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerKeysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, keyName string, options *ServerKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerKeysClient) getHandleResponse(resp *http.Response) (ServerKeysClientGetResponse, error) {
	result := ServerKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerKey); err != nil {
		return ServerKeysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Server keys.
//
// Generated from API version 2020-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ServerKeysClientListOptions contains the optional parameters for the ServerKeysClient.NewListPager method.
func (client *ServerKeysClient) NewListPager(resourceGroupName string, serverName string, options *ServerKeysClientListOptions) *runtime.Pager[ServerKeysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerKeysClientListResponse]{
		More: func(page ServerKeysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerKeysClientListResponse) (ServerKeysClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerKeysClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerKeysClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ServerKeysClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerKeysClient) listHandleResponse(resp *http.Response) (ServerKeysClientListResponse, error) {
	result := ServerKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerKeyListResult); err != nil {
		return ServerKeysClientListResponse{}, err
	}
	return result, nil
}