# \SendRoutingInfoSMCustomOperationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendRoutingInfoSm**](SendRoutingInfoSMCustomOperationAPI.md#SendRoutingInfoSm) | **Post** /{ueId}/registrations/send-routing-info-sm | Retreive addressing information for SMS delivery



## SendRoutingInfoSm

> RoutingInfoSmResponse SendRoutingInfoSm(ctx, ueId).RoutingInfoSmRequest(routingInfoSmRequest).Execute()

Retreive addressing information for SMS delivery

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
	routingInfoSmRequest := *openapiclient.NewRoutingInfoSmRequest() // RoutingInfoSmRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendRoutingInfoSMCustomOperationAPI.SendRoutingInfoSm(context.Background(), ueId).RoutingInfoSmRequest(routingInfoSmRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendRoutingInfoSMCustomOperationAPI.SendRoutingInfoSm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendRoutingInfoSm`: RoutingInfoSmResponse
	fmt.Fprintf(os.Stdout, "Response from `SendRoutingInfoSMCustomOperationAPI.SendRoutingInfoSm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRoutingInfoSmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingInfoSmRequest** | [**RoutingInfoSmRequest**](RoutingInfoSmRequest.md) |  | 

### Return type

[**RoutingInfoSmResponse**](RoutingInfoSmResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

