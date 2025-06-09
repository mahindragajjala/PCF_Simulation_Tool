# \RetrieveSMFRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveSmfRegistration**](RetrieveSMFRegistrationAPI.md#RetrieveSmfRegistration) | **Get** /{ueId}/registrations/smf-registrations/{pduSessionId} | get an SMF registration



## RetrieveSmfRegistration

> SmfRegistration RetrieveSmfRegistration(ctx, ueId, pduSessionId).Execute()

get an SMF registration

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
	pduSessionId := int32(56) // int32 | Identifier of the PDU session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveSMFRegistrationAPI.RetrieveSmfRegistration(context.Background(), ueId, pduSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveSMFRegistrationAPI.RetrieveSmfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSmfRegistration`: SmfRegistration
	fmt.Fprintf(os.Stdout, "Response from `RetrieveSMFRegistrationAPI.RetrieveSmfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**pduSessionId** | **int32** | Identifier of the PDU session | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSmfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

