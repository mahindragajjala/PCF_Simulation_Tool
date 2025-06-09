# \Class5GMBSGroupModificationAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get5GMBSGroup**](Class5GMBSGroupModificationAPI.md#Get5GMBSGroup) | **Get** /mbs-group-membership/{extGroupId} | get 5G MBS Group
[**Modify5GMBSGroup**](Class5GMBSGroupModificationAPI.md#Modify5GMBSGroup) | **Patch** /mbs-group-membership/{extGroupId} | modify a 5G MBS Group



## Get5GMBSGroup

> MulticastMbsGroupMemb Get5GMBSGroup(ctx, extGroupId).Execute()

get 5G MBS Group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	extGroupId := "extGroupId_example" // string | External Identifier of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class5GMBSGroupModificationAPI.Get5GMBSGroup(context.Background(), extGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GMBSGroupModificationAPI.Get5GMBSGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get5GMBSGroup`: MulticastMbsGroupMemb
	fmt.Fprintf(os.Stdout, "Response from `Class5GMBSGroupModificationAPI.Get5GMBSGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extGroupId** | **string** | External Identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet5GMBSGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MulticastMbsGroupMemb**](MulticastMbsGroupMemb.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Modify5GMBSGroup

> PatchResult Modify5GMBSGroup(ctx, extGroupId).MulticastMbsGroupMemb(multicastMbsGroupMemb).SupportedFeatures(supportedFeatures).Execute()

modify a 5G MBS Group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	extGroupId := "extGroupId_example" // string | External Identifier of the group
	multicastMbsGroupMemb := *openapiclient.NewMulticastMbsGroupMemb([]string{"MulticastGroupMemb_example"}) // MulticastMbsGroupMemb | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class5GMBSGroupModificationAPI.Modify5GMBSGroup(context.Background(), extGroupId).MulticastMbsGroupMemb(multicastMbsGroupMemb).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GMBSGroupModificationAPI.Modify5GMBSGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Modify5GMBSGroup`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `Class5GMBSGroupModificationAPI.Modify5GMBSGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extGroupId** | **string** | External Identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiModify5GMBSGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multicastMbsGroupMemb** | [**MulticastMbsGroupMemb**](MulticastMbsGroupMemb.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

