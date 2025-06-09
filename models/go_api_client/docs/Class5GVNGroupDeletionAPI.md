# \Class5GVNGroupDeletionAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete5GVNGroup**](Class5GVNGroupDeletionAPI.md#Delete5GVNGroup) | **Delete** /5g-vn-groups/{extGroupId} | delete a 5G VN Group



## Delete5GVNGroup

> Delete5GVNGroup(ctx, extGroupId).MtcProviderInfo(mtcProviderInfo).AfId(afId).Execute()

delete a 5G VN Group

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
	extGroupId := "extGroupId_example" // string | External Identifier of the Group
	mtcProviderInfo := "mtcProviderInfo_example" // string | MTC Provider Information that originated the service operation (optional)
	afId := "afId_example" // string | AF ID that originated the service operation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class5GVNGroupDeletionAPI.Delete5GVNGroup(context.Background(), extGroupId).MtcProviderInfo(mtcProviderInfo).AfId(afId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupDeletionAPI.Delete5GVNGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extGroupId** | **string** | External Identifier of the Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete5GVNGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mtcProviderInfo** | **string** | MTC Provider Information that originated the service operation | 
 **afId** | **string** | AF ID that originated the service operation | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

