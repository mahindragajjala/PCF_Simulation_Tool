# \QueryUEInfoAPI

All URIs are relative to *https://example.com/nudm-mt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryUeInfo**](QueryUEInfoAPI.md#QueryUeInfo) | **Get** /{supi} | Query Information for the UE



## QueryUeInfo

> UeInfo QueryUeInfo(ctx, supi).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Query Information for the UE

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
	supi := "supi_example" // string | Identifier of the UE
	fields := []string{"Inner_example"} // []string | attributes to be retrieved
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryUEInfoAPI.QueryUeInfo(context.Background(), supi).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryUEInfoAPI.QueryUeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryUeInfo`: UeInfo
	fmt.Fprintf(os.Stdout, "Response from `QueryUEInfoAPI.QueryUeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryUeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**UeInfo**](UeInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

