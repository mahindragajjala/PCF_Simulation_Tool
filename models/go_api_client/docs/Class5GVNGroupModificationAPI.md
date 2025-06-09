# \Class5GVNGroupModificationAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get5GVNGroup**](Class5GVNGroupModificationAPI.md#Get5GVNGroup) | **Get** /5g-vn-groups/{extGroupId} | get 5G VN Group
[**Modify5GVNGroup**](Class5GVNGroupModificationAPI.md#Modify5GVNGroup) | **Patch** /5g-vn-groups/{extGroupId} | modify a 5G VN Group



## Get5GVNGroup

> Model5GVnGroupConfiguration Get5GVNGroup(ctx, extGroupId).Execute()

get 5G VN Group

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
	resp, r, err := apiClient.Class5GVNGroupModificationAPI.Get5GVNGroup(context.Background(), extGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupModificationAPI.Get5GVNGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get5GVNGroup`: Model5GVnGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `Class5GVNGroupModificationAPI.Get5GVNGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extGroupId** | **string** | External Identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet5GVNGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Modify5GVNGroup

> PatchResult Modify5GVNGroup(ctx, extGroupId).Model5GVnGroupConfigurationModification(model5GVnGroupConfigurationModification).SupportedFeatures(supportedFeatures).Execute()

modify a 5G VN Group

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
	model5GVnGroupConfigurationModification := *openapiclient.NewModel5GVnGroupConfigurationModification() // Model5GVnGroupConfigurationModification | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class5GVNGroupModificationAPI.Modify5GVNGroup(context.Background(), extGroupId).Model5GVnGroupConfigurationModification(model5GVnGroupConfigurationModification).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupModificationAPI.Modify5GVNGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Modify5GVNGroup`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `Class5GVNGroupModificationAPI.Modify5GVNGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extGroupId** | **string** | External Identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiModify5GVNGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model5GVnGroupConfigurationModification** | [**Model5GVnGroupConfigurationModification**](Model5GVnGroupConfigurationModification.md) |  | 
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

