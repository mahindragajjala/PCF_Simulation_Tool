# \DeleteEESubscriptionAPI

All URIs are relative to *https://example.com/nudm-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEeSubscription**](DeleteEESubscriptionAPI.md#DeleteEeSubscription) | **Delete** /{ueIdentity}/ee-subscriptions/{subscriptionId} | Unsubscribe



## DeleteEeSubscription

> DeleteEeSubscription(ctx, ueIdentity, subscriptionId).Execute()

Unsubscribe

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
	subscriptionId := "subscriptionId_example" // string | Id of the EE Subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeleteEESubscriptionAPI.DeleteEeSubscription(context.Background(), ueIdentity, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteEESubscriptionAPI.DeleteEeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the subscription is applied. Contains the GPSI of the user or the external group ID or any UE. | 
**subscriptionId** | **string** | Id of the EE Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEeSubscriptionRequest struct via the builder pattern


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

