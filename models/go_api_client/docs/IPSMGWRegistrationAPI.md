# \IPSMGWRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpSmGwRegistration**](IPSMGWRegistrationAPI.md#IpSmGwRegistration) | **Put** /{ueId}/registrations/ip-sm-gw | Register an IP-SM-GW



## IpSmGwRegistration

> IpSmGwRegistration IpSmGwRegistration(ctx, ueId).IpSmGwRegistration(ipSmGwRegistration).Execute()

Register an IP-SM-GW

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
	ipSmGwRegistration := *openapiclient.NewIpSmGwRegistration() // IpSmGwRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPSMGWRegistrationAPI.IpSmGwRegistration(context.Background(), ueId).IpSmGwRegistration(ipSmGwRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationAPI.IpSmGwRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpSmGwRegistration`: IpSmGwRegistration
	fmt.Fprintf(os.Stdout, "Response from `IPSMGWRegistrationAPI.IpSmGwRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpSmGwRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipSmGwRegistration** | [**IpSmGwRegistration**](IpSmGwRegistration.md) |  | 

### Return type

[**IpSmGwRegistration**](IpSmGwRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

