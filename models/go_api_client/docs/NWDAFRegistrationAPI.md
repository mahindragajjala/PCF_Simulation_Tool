# \NWDAFRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NwdafRegistration**](NWDAFRegistrationAPI.md#NwdafRegistration) | **Put** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | register as NWDAF



## NwdafRegistration

> NwdafRegistration NwdafRegistration(ctx, ueId, nwdafRegistrationId).NwdafRegistration(nwdafRegistration).Execute()

register as NWDAF

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
	nwdafRegistrationId := "nwdafRegistrationId_example" // string | NWDAF registration identifier
	nwdafRegistration := *openapiclient.NewNwdafRegistration("NwdafInstanceId_example", []openapiclient.EventId{*openapiclient.NewEventId()}) // NwdafRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NWDAFRegistrationAPI.NwdafRegistration(context.Background(), ueId, nwdafRegistrationId).NwdafRegistration(nwdafRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NWDAFRegistrationAPI.NwdafRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NwdafRegistration`: NwdafRegistration
	fmt.Fprintf(os.Stdout, "Response from `NWDAFRegistrationAPI.NwdafRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**nwdafRegistrationId** | **string** | NWDAF registration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiNwdafRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nwdafRegistration** | [**NwdafRegistration**](NwdafRegistration.md) |  | 

### Return type

[**NwdafRegistration**](NwdafRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

