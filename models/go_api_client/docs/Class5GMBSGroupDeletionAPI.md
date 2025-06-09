# \Class5GMBSGroupDeletionAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete5GMBSGroup**](Class5GMBSGroupDeletionAPI.md#Delete5GMBSGroup) | **Delete** /mbs-group-membership/{extGroupId} | delete a 5G MBS Group



## Delete5GMBSGroup

> Delete5GMBSGroup(ctx, extGroupId).Execute()

delete a 5G MBS Group

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class5GMBSGroupDeletionAPI.Delete5GMBSGroup(context.Background(), extGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GMBSGroupDeletionAPI.Delete5GMBSGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDelete5GMBSGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

