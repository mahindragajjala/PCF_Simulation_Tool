# \ServiceSpecificAuthorizationRemoveAPI

All URIs are relative to *https://example.com/nudm-ssau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceSpecificAuthorizationRemoval**](ServiceSpecificAuthorizationRemoveAPI.md#ServiceSpecificAuthorizationRemoval) | **Post** /{ueIdentity}/{serviceType}/remove | Remove the authorization of specific service&#39;s configuration.



## ServiceSpecificAuthorizationRemoval

> ServiceSpecificAuthorizationRemoval(ctx, ueIdentity, serviceType).ServiceSpecificAuthorizationRemoveData(serviceSpecificAuthorizationRemoveData).Execute()

Remove the authorization of specific service's configuration.

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
	ueIdentity := "ueIdentity_example" // string | Represents the scope of the UE for which the Service Specific configuration authorization to be removed. Contains the GPSI of the user or the external group ID.
	serviceType := *openapiclient.NewServiceType() // ServiceType | Represents the specific service for which the Service Specific configuration authorization to be removed.
	serviceSpecificAuthorizationRemoveData := *openapiclient.NewServiceSpecificAuthorizationRemoveData("AuthId_example") // ServiceSpecificAuthorizationRemoveData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceSpecificAuthorizationRemoveAPI.ServiceSpecificAuthorizationRemoval(context.Background(), ueIdentity, serviceType).ServiceSpecificAuthorizationRemoveData(serviceSpecificAuthorizationRemoveData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationRemoveAPI.ServiceSpecificAuthorizationRemoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the Service Specific configuration authorization to be removed. Contains the GPSI of the user or the external group ID. | 
**serviceType** | [**ServiceType**](.md) | Represents the specific service for which the Service Specific configuration authorization to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSpecificAuthorizationRemovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceSpecificAuthorizationRemoveData** | [**ServiceSpecificAuthorizationRemoveData**](ServiceSpecificAuthorizationRemoveData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

