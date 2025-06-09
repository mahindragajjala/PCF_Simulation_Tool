# \NWDAFDeregistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NwdafDeregistration**](NWDAFDeregistrationAPI.md#NwdafDeregistration) | **Delete** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | delete an NWDAF registration



## NwdafDeregistration

> NwdafDeregistration(ctx, ueId, nwdafRegistrationId).Execute()

delete an NWDAF registration

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NWDAFDeregistrationAPI.NwdafDeregistration(context.Background(), ueId, nwdafRegistrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NWDAFDeregistrationAPI.NwdafDeregistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**nwdafRegistrationId** | **string** | NWDAF registration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiNwdafDeregistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

