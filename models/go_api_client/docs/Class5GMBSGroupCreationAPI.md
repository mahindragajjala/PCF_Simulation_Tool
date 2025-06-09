# \Class5GMBSGroupCreationAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create5GMBSGroup**](Class5GMBSGroupCreationAPI.md#Create5GMBSGroup) | **Put** /mbs-group-membership/{extGroupId} | create a 5G MBS Group



## Create5GMBSGroup

> Create5GMBSGroup(ctx, extGroupId).MulticastMbsGroupMemb(multicastMbsGroupMemb).Execute()

create a 5G MBS Group

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
	multicastMbsGroupMemb := *openapiclient.NewMulticastMbsGroupMemb([]string{"MulticastGroupMemb_example"}) // MulticastMbsGroupMemb | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class5GMBSGroupCreationAPI.Create5GMBSGroup(context.Background(), extGroupId).MulticastMbsGroupMemb(multicastMbsGroupMemb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class5GMBSGroupCreationAPI.Create5GMBSGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreate5GMBSGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multicastMbsGroupMemb** | [**MulticastMbsGroupMemb**](MulticastMbsGroupMemb.md) |  | 

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

