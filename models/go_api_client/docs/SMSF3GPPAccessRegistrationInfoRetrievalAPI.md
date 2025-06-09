# \SMSF3GPPAccessRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get3GppSmsfRegistration**](SMSF3GPPAccessRegistrationInfoRetrievalAPI.md#Get3GppSmsfRegistration) | **Get** /{ueId}/registrations/smsf-3gpp-access | retrieve the SMSF registration for 3GPP access information



## Get3GppSmsfRegistration

> SmsfRegistration Get3GppSmsfRegistration(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

retrieve the SMSF registration for 3GPP access information

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
	resp, r, err := apiClient.SMSF3GPPAccessRegistrationInfoRetrievalAPI.Get3GppSmsfRegistration(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSF3GPPAccessRegistrationInfoRetrievalAPI.Get3GppSmsfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get3GppSmsfRegistration`: SmsfRegistration
	fmt.Fprintf(os.Stdout, "Response from `SMSF3GPPAccessRegistrationInfoRetrievalAPI.Get3GppSmsfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet3GppSmsfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** |  | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

