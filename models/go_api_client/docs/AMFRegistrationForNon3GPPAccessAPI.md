# \AMFRegistrationForNon3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Non3GppRegistration**](AMFRegistrationForNon3GPPAccessAPI.md#Non3GppRegistration) | **Put** /{ueId}/registrations/amf-non-3gpp-access | register as AMF for non-3GPP access



## Non3GppRegistration

> AmfNon3GppAccessRegistration Non3GppRegistration(ctx, ueId).AmfNon3GppAccessRegistration(amfNon3GppAccessRegistration).Execute()

register as AMF for non-3GPP access

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
	amfNon3GppAccessRegistration := *openapiclient.NewAmfNon3GppAccessRegistration("AmfInstanceId_example", *openapiclient.NewImsVoPs(), "DeregCallbackUri_example", *openapiclient.NewGuami(*openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example"), "AmfId_example"), *openapiclient.NewRatType()) // AmfNon3GppAccessRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AMFRegistrationForNon3GPPAccessAPI.Non3GppRegistration(context.Background(), ueId).AmfNon3GppAccessRegistration(amfNon3GppAccessRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AMFRegistrationForNon3GPPAccessAPI.Non3GppRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Non3GppRegistration`: AmfNon3GppAccessRegistration
	fmt.Fprintf(os.Stdout, "Response from `AMFRegistrationForNon3GPPAccessAPI.Non3GppRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiNon3GppRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amfNon3GppAccessRegistration** | [**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md) |  | 

### Return type

[**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

