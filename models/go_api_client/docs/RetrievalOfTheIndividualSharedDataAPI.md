# \RetrievalOfTheIndividualSharedDataAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndividualSharedData**](RetrievalOfTheIndividualSharedDataAPI.md#GetIndividualSharedData) | **Get** /shared-data/{sharedDataId} | retrieve the individual shared data



## GetIndividualSharedData

> SharedData GetIndividualSharedData(ctx, sharedDataId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve the individual shared data

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
	sharedDataId := []string{"Inner_example"} // []string | Id of the Shared data
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
	ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalOfTheIndividualSharedDataAPI.GetIndividualSharedData(context.Background(), sharedDataId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfTheIndividualSharedDataAPI.GetIndividualSharedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndividualSharedData`: SharedData
	fmt.Fprintf(os.Stdout, "Response from `RetrievalOfTheIndividualSharedDataAPI.GetIndividualSharedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDataId** | [**[]string**](string.md) | Id of the Shared data | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

