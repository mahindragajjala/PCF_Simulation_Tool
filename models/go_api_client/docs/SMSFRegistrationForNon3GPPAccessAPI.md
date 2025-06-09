# \SMSFRegistrationForNon3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Non3GppSmsfRegistration**](SMSFRegistrationForNon3GPPAccessAPI.md#Non3GppSmsfRegistration) | **Put** /{ueId}/registrations/smsf-non-3gpp-access | register as SMSF for non-3GPP access



## Non3GppSmsfRegistration

> SmsfRegistration Non3GppSmsfRegistration(ctx, ueId).SmsfRegistration(smsfRegistration).Execute()

register as SMSF for non-3GPP access

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
	smsfRegistration := *openapiclient.NewSmsfRegistration("SmsfInstanceId_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // SmsfRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSFRegistrationForNon3GPPAccessAPI.Non3GppSmsfRegistration(context.Background(), ueId).SmsfRegistration(smsfRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSFRegistrationForNon3GPPAccessAPI.Non3GppSmsfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Non3GppSmsfRegistration`: SmsfRegistration
	fmt.Fprintf(os.Stdout, "Response from `SMSFRegistrationForNon3GPPAccessAPI.Non3GppSmsfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiNon3GppSmsfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfRegistration** | [**SmsfRegistration**](SmsfRegistration.md) |  | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

