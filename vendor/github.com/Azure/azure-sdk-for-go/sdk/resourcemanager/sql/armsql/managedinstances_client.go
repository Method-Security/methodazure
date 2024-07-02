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
	"strconv"
	"strings"
)

// ManagedInstancesClient contains the methods for the ManagedInstances group.
// Don't use this type directly, use NewManagedInstancesClient() instead.
type ManagedInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstancesClient creates a new instance of ManagedInstancesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - parameters - The requested managed instance resource state.
//   - options - ManagedInstancesClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedInstancesClient.BeginCreateOrUpdate
//     method.
func (client *ManagedInstancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstancesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancesClientBeginDeleteOptions contains the optional parameters for the ManagedInstancesClient.BeginDelete
//     method.
func (client *ManagedInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*runtime.Poller[ManagedInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
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
func (client *ManagedInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Failovers a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance to failover.
//   - options - ManagedInstancesClientBeginFailoverOptions contains the optional parameters for the ManagedInstancesClient.BeginFailover
//     method.
func (client *ManagedInstancesClient) BeginFailover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*runtime.Poller[ManagedInstancesClientFailoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.failover(ctx, resourceGroupName, managedInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstancesClientFailoverResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstancesClientFailoverResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Failover - Failovers a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) failover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstancesClient.BeginFailover"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
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

// failoverCreateRequest creates the Failover request.
func (client *ManagedInstancesClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ReplicaType != nil {
		reqQP.Set("replicaType", string(*options.ReplicaType))
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancesClientGetOptions contains the optional parameters for the ManagedInstancesClient.Get method.
func (client *ManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientGetOptions) (ManagedInstancesClientGetResponse, error) {
	var err error
	const operationName = "ManagedInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstancesClient) getHandleResponse(resp *http.Response) (ManagedInstancesClientGetResponse, error) {
	result := ManagedInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstance); err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all managed instances in the subscription.
//
// Generated from API version 2021-05-01-preview
//   - options - ManagedInstancesClientListOptions contains the optional parameters for the ManagedInstancesClient.NewListPager
//     method.
func (client *ManagedInstancesClient) NewListPager(options *ManagedInstancesClientListOptions) *runtime.Pager[ManagedInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListResponse]{
		More: func(page ManagedInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListResponse) (ManagedInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ManagedInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedInstancesClient) listCreateRequest(ctx context.Context, options *ManagedInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedInstancesClient) listHandleResponse(resp *http.Response) (ManagedInstancesClientListResponse, error) {
	result := ManagedInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByInstancePoolPager - Gets a list of all managed instances in an instance pool.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - instancePoolName - The instance pool name.
//   - options - ManagedInstancesClientListByInstancePoolOptions contains the optional parameters for the ManagedInstancesClient.NewListByInstancePoolPager
//     method.
func (client *ManagedInstancesClient) NewListByInstancePoolPager(resourceGroupName string, instancePoolName string, options *ManagedInstancesClientListByInstancePoolOptions) *runtime.Pager[ManagedInstancesClientListByInstancePoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByInstancePoolResponse]{
		More: func(page ManagedInstancesClientListByInstancePoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByInstancePoolResponse) (ManagedInstancesClientListByInstancePoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstancesClient.NewListByInstancePoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstancePoolCreateRequest(ctx, resourceGroupName, instancePoolName, options)
			}, nil)
			if err != nil {
				return ManagedInstancesClientListByInstancePoolResponse{}, err
			}
			return client.listByInstancePoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstancePoolCreateRequest creates the ListByInstancePool request.
func (client *ManagedInstancesClient) listByInstancePoolCreateRequest(ctx context.Context, resourceGroupName string, instancePoolName string, options *ManagedInstancesClientListByInstancePoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/managedInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instancePoolName == "" {
		return nil, errors.New("parameter instancePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instancePoolName}", url.PathEscape(instancePoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstancePoolHandleResponse handles the ListByInstancePool response.
func (client *ManagedInstancesClient) listByInstancePoolHandleResponse(resp *http.Response) (ManagedInstancesClientListByInstancePoolResponse, error) {
	result := ManagedInstancesClientListByInstancePoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListByInstancePoolResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Get top resource consuming queries of a managed instance.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancesClientListByManagedInstanceOptions contains the optional parameters for the ManagedInstancesClient.NewListByManagedInstancePager
//     method.
func (client *ManagedInstancesClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstancesClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByManagedInstanceResponse]{
		More: func(page ManagedInstancesClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByManagedInstanceResponse) (ManagedInstancesClientListByManagedInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstancesClient.NewListByManagedInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return ManagedInstancesClientListByManagedInstanceResponse{}, err
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstancesClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/topqueries"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.NumberOfQueries != nil {
		reqQP.Set("numberOfQueries", strconv.FormatInt(int64(*options.NumberOfQueries), 10))
	}
	if options != nil && options.Databases != nil {
		reqQP.Set("databases", *options.Databases)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", string(*options.Interval))
	}
	if options != nil && options.AggregationFunction != nil {
		reqQP.Set("aggregationFunction", string(*options.AggregationFunction))
	}
	if options != nil && options.ObservationMetric != nil {
		reqQP.Set("observationMetric", string(*options.ObservationMetric))
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstancesClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstancesClientListByManagedInstanceResponse, error) {
	result := ManagedInstancesClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopQueriesListResult); err != nil {
		return ManagedInstancesClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of managed instances in a resource group.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - options - ManagedInstancesClientListByResourceGroupOptions contains the optional parameters for the ManagedInstancesClient.NewListByResourceGroupPager
//     method.
func (client *ManagedInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedInstancesClientListByResourceGroupOptions) *runtime.Pager[ManagedInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByResourceGroupResponse]{
		More: func(page ManagedInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByResourceGroupResponse) (ManagedInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ManagedInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedInstancesClientListByResourceGroupResponse, error) {
	result := ManagedInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - parameters - The requested managed instance resource state.
//   - options - ManagedInstancesClientBeginUpdateOptions contains the optional parameters for the ManagedInstancesClient.BeginUpdate
//     method.
func (client *ManagedInstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*runtime.Poller[ManagedInstancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedInstanceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstancesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstancesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstancesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ManagedInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}