# \SMSFNon3GPPAccessRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNon3GppSmsfRegistration**](SMSFNon3GPPAccessRegistrationInfoRetrievalAPI.md#GetNon3GppSmsfRegistration) | **Get** /{ueId}/registrations/smsf-non-3gpp-access | retrieve the SMSF registration for non-3GPP access information



## GetNon3GppSmsfRegistration

> SmsfRegistration GetNon3GppSmsfRegistration(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

retrieve the SMSF registration for non-3GPP access information

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
	resp, r, err := apiClient.SMSFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppSmsfRegistration(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppSmsfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNon3GppSmsfRegistration`: SmsfRegistration
	fmt.Fprintf(os.Stdout, "Response from `SMSFNon3GPPAccessRegistrationInfoRetrievalAPI.GetNon3GppSmsfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNon3GppSmsfRegistrationRequest struct via the builder pattern


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

