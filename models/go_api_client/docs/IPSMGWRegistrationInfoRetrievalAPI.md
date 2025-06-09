# \IPSMGWRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpSmGwRegistration**](IPSMGWRegistrationInfoRetrievalAPI.md#GetIpSmGwRegistration) | **Get** /{ueId}/registrations/ip-sm-gw | Retrieve the IP-SM-GW registration information



## GetIpSmGwRegistration

> IpSmGwRegistration GetIpSmGwRegistration(ctx, ueId).Execute()

Retrieve the IP-SM-GW registration information

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPSMGWRegistrationInfoRetrievalAPI.GetIpSmGwRegistration(context.Background(), ueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationInfoRetrievalAPI.GetIpSmGwRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpSmGwRegistration`: IpSmGwRegistration
	fmt.Fprintf(os.Stdout, "Response from `IPSMGWRegistrationInfoRetrievalAPI.GetIpSmGwRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpSmGwRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpSmGwRegistration**](IpSmGwRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

