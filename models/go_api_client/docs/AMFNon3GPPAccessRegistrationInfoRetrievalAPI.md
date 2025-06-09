# \AMFNon3GPPAccessRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNon3GppRegistration**](AMFNon3GPPAccessRegistrationInfoRetrievalAPI.md#GetNon3GppRegistration) | **Get** /{ueId}/registrations/amf-non-3gpp-access | retrieve the AMF registration for non-3GPP access information



## GetNon3GppRegistration

> AmfNon3GppAccessRegistration GetNon3GppRegistration(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

retrieve the AMF registration for non-3GPP access information

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
	ueId := "ueId_example" // string | Identifier of the UE
	supportedFeatures := "supportedFeatures_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AMFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppRegistration(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AMFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNon3GppRegistration`: AmfNon3GppAccessRegistration
	fmt.Fprintf(os.Stdout, "Response from `AMFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNon3GppRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** |  | 

### Return type

[**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

