# \Class5GVNGroupCreationAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create5GVNGroup**](Class5GVNGroupCreationAPI.md#Create5GVNGroup) | **Put** /5g-vn-groups/{extGroupId} | create a 5G VN Group



## Create5GVNGroup

> Create5GVNGroup(ctx, extGroupId).Model5GVnGroupConfiguration(model5GVnGroupConfiguration).Execute()

create a 5G VN Group

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
	model5GVnGroupConfiguration := *openapiclient.NewModel5GVnGroupConfiguration() // Model5GVnGroupConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class5GVNGroupCreationAPI.Create5GVNGroup(context.Background(), extGroupId).Model5GVnGroupConfiguration(model5GVnGroupConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupCreationAPI.Create5GVNGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreate5GVNGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model5GVnGroupConfiguration** | [**Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

