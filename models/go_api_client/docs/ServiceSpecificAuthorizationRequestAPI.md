# \ServiceSpecificAuthorizationRequestAPI

All URIs are relative to *https://example.com/nudm-ssau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceSpecificAuthorization**](ServiceSpecificAuthorizationRequestAPI.md#ServiceSpecificAuthorization) | **Post** /{ueIdentity}/{serviceType}/authorize | Authorization for the Service specific parameters in the request.



## ServiceSpecificAuthorization

> ServiceSpecificAuthorizationData ServiceSpecificAuthorization(ctx, ueIdentity, serviceType).ServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationInfo).Execute()

Authorization for the Service specific parameters in the request.

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
	ueIdentity := "ueIdentity_example" // string | Represents the scope of the UE for which the Service Specific Parameters are authorized. Contains the GPSI of the user or the external group ID.
	serviceType := *openapiclient.NewServiceType() // ServiceType | Represents the specific service for which the Service Specific Parameters are authorized.
	serviceSpecificAuthorizationInfo := *openapiclient.NewServiceSpecificAuthorizationInfo() // ServiceSpecificAuthorizationInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSpecificAuthorizationRequestAPI.ServiceSpecificAuthorization(context.Background(), ueIdentity, serviceType).ServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationRequestAPI.ServiceSpecificAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSpecificAuthorization`: ServiceSpecificAuthorizationData
	fmt.Fprintf(os.Stdout, "Response from `ServiceSpecificAuthorizationRequestAPI.ServiceSpecificAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the Service Specific Parameters are authorized. Contains the GPSI of the user or the external group ID. | 
**serviceType** | [**ServiceType**](.md) | Represents the specific service for which the Service Specific Parameters are authorized. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSpecificAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceSpecificAuthorizationInfo** | [**ServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfo.md) |  | 

### Return type

[**ServiceSpecificAuthorizationData**](ServiceSpecificAuthorizationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

