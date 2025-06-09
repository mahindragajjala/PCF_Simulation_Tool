# \NWDAFRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNwdafRegistration**](NWDAFRegistrationInfoRetrievalAPI.md#GetNwdafRegistration) | **Get** /{ueId}/registrations/nwdaf-registrations | retrieve the NWDAF registration



## GetNwdafRegistration

> []NwdafRegistration GetNwdafRegistration(ctx, ueId).AnalyticsIds(analyticsIds).SupportedFeatures(supportedFeatures).Execute()

retrieve the NWDAF registration

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
	analyticsIds := []openapiclient.EventId{*openapiclient.NewEventId()} // []EventId | List of analytics Id(s) provided by the consumers of NWDAF. (optional)
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NWDAFRegistrationInfoRetrievalAPI.GetNwdafRegistration(context.Background(), ueId).AnalyticsIds(analyticsIds).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NWDAFRegistrationInfoRetrievalAPI.GetNwdafRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNwdafRegistration`: []NwdafRegistration
	fmt.Fprintf(os.Stdout, "Response from `NWDAFRegistrationInfoRetrievalAPI.GetNwdafRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNwdafRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsIds** | [**[]EventId**](EventId.md) | List of analytics Id(s) provided by the consumers of NWDAF. | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]NwdafRegistration**](NwdafRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

