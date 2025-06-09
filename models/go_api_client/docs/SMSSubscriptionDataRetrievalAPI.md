# \SMSSubscriptionDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmsData**](SMSSubscriptionDataRetrievalAPI.md#GetSmsData) | **Get** /{supi}/sms-data | retrieve a UE&#39;s SMS Subscription Data



## GetSmsData

> SmsSubscriptionData GetSmsData(ctx, supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's SMS Subscription Data

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
	supi := "supi_example" // string | Identifier of the UE
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
	plmnId := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId |  (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
	ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSSubscriptionDataRetrievalAPI.GetSmsData(context.Background(), supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSSubscriptionDataRetrievalAPI.GetSmsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmsData`: SmsSubscriptionData
	fmt.Fprintf(os.Stdout, "Response from `SMSSubscriptionDataRetrievalAPI.GetSmsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **plmnId** | [**PlmnId**](PlmnId.md) |  | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SmsSubscriptionData**](SmsSubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

