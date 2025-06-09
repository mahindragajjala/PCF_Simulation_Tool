# \SubscriptionDataUpdateAPI

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Update**](SubscriptionDataUpdateAPI.md#Update) | **Patch** /{ueId}/pp-data | provision parameters



## Update

> PatchResult Update(ctx, ueId).PpData(ppData).SupportedFeatures(supportedFeatures).Execute()

provision parameters

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
	ueId := *openapiclient.NewUpdateUeIdParameter() // UpdateUeIdParameter | Identifier of the UE
	ppData := *openapiclient.NewPpData() // PpData | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionDataUpdateAPI.Update(context.Background(), ueId).PpData(ppData).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionDataUpdateAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionDataUpdateAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UpdateUeIdParameter**](.md) | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ppData** | [**PpData**](PpData.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

