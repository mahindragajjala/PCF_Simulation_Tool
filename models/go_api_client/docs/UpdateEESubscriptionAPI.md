# \UpdateEESubscriptionAPI

All URIs are relative to *https://example.com/nudm-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateEeSubscription**](UpdateEESubscriptionAPI.md#UpdateEeSubscription) | **Patch** /{ueIdentity}/ee-subscriptions/{subscriptionId} | Patch



## UpdateEeSubscription

> PatchResult UpdateEeSubscription(ctx, ueIdentity, subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Patch

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
	patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateEESubscriptionAPI.UpdateEeSubscription(context.Background(), ueIdentity, subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateEESubscriptionAPI.UpdateEeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEeSubscription`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `UpdateEESubscriptionAPI.UpdateEeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the subscription is applied. Contains the GPSI of the user or the external group ID or any UE. | 
**subscriptionId** | **string** | Id of the EE Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

