# \CreateEESubscriptionAPI

All URIs are relative to *https://example.com/nudm-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEeSubscription**](CreateEESubscriptionAPI.md#CreateEeSubscription) | **Post** /{ueIdentity}/ee-subscriptions | Subscribe



## CreateEeSubscription

> CreatedEeSubscription CreateEeSubscription(ctx, ueIdentity).EeSubscription(eeSubscription).Execute()

Subscribe

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
	ueIdentity := "ueIdentity_example" // string | Represents the scope of the UE for which the subscription is applied. Contains the GPSI of the user or the external group ID or any UE.
	eeSubscription := *openapiclient.NewEeSubscription("CallbackReference_example", map[string]MonitoringConfiguration{"key": *openapiclient.NewMonitoringConfiguration(*openapiclient.NewEventType())}) // EeSubscription | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateEESubscriptionAPI.CreateEeSubscription(context.Background(), ueIdentity).EeSubscription(eeSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateEESubscriptionAPI.CreateEeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEeSubscription`: CreatedEeSubscription
	fmt.Fprintf(os.Stdout, "Response from `CreateEESubscriptionAPI.CreateEeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the subscription is applied. Contains the GPSI of the user or the external group ID or any UE. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 

### Return type

[**CreatedEeSubscription**](CreatedEeSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

