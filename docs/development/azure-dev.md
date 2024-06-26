# Azure Development

The Azure and Microsoft SDKs and APIs can be difficult to deal with. Compared to AWS for example, they are not clearly documented, data is less organized, and there are different patterns used for different resources.

Below are some tips for dealing with the different SDKs.

## Using the Azure SDK

The [API Reference](https://learn.microsoft.com/en-us/rest/api/azure/) and Go documentation pages [example](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions), are both great resources for understanding query patterns and data structures.

## Using the Microsoft Graph SDK

The Microsoft Graph SDK is the preferred and up to date method when dealing with Entra ID. For more information on the transition from Azure AD SDKs, view this [article](https://learn.microsoft.com/en-us/graph/migrate-azure-ad-graph-overview).

The Microsoft Graph SDK is particularly difficult to deal with. Similar to the Azure SDK, the [API Reference](https://learn.microsoft.com/en-us/graph/api/resources/serviceprincipal?view=graph-rest-1.0) is very useful because you need to specify fields to query.

Additionally, the [Graph Explorer](https://developer.microsoft.com/en-us/graph/graph-explorer) is also useful for testing requests over HTTPS before brining them to the SDK.