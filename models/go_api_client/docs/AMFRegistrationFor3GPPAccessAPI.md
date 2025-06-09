# \AMFRegistrationFor3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call3GppRegistration**](AMFRegistrationFor3GPPAccessAPI.md#Call3GppRegistration) | **Put** /{ueId}/registrations/amf-3gpp-access | register as AMF for 3GPP access



## Call3GppRegistration

> Amf3GppAccessRegistration Call3GppRegistration(ctx, ueId).Amf3GppAccessRegistration(amf3GppAccessRegistration).Execute()

register as AMF for 3GPP access

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
	amf3GppAccessRegistration := *openapiclient.NewAmf3GppAccessRegistration("AmfInstanceId_example", "DeregCallbackUri_example", *openapiclient.NewGuami(*openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example"), "AmfId_example"), *openapiclient.NewRatType()) // Amf3GppAccessRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AMFRegistrationFor3GPPAccessAPI.Call3GppRegistration(context.Background(), ueId).Amf3GppAccessRegistration(amf3GppAccessRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AMFRegistrationFor3GPPAccessAPI.Call3GppRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Call3GppRegistration`: Amf3GppAccessRegistration
	fmt.Fprintf(os.Stdout, "Response from `AMFRegistrationFor3GPPAccessAPI.Call3GppRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall3GppRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amf3GppAccessRegistration** | [**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md) |  | 

### Return type

[**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

