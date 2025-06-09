# \ProvideUELocationAPI

All URIs are relative to *https://example.com/nudm-mt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvideLocationInfo**](ProvideUELocationAPI.md#ProvideLocationInfo) | **Post** /{supi}/loc-info/provide-loc-info | Provides the UE&#39;s 5GS location information



## ProvideLocationInfo

> LocationInfoResult ProvideLocationInfo(ctx, supi).LocationInfoRequest(locationInfoRequest).Execute()

Provides the UE's 5GS location information

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
	locationInfoRequest := *openapiclient.NewLocationInfoRequest() // LocationInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvideUELocationAPI.ProvideLocationInfo(context.Background(), supi).LocationInfoRequest(locationInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvideUELocationAPI.ProvideLocationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvideLocationInfo`: LocationInfoResult
	fmt.Fprintf(os.Stdout, "Response from `ProvideUELocationAPI.ProvideLocationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideLocationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationInfoRequest** | [**LocationInfoRequest**](LocationInfoRequest.md) |  | 

### Return type

[**LocationInfoResult**](LocationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

