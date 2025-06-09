# \SMFSelectionSubscriptionDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmfSelData**](SMFSelectionSubscriptionDataRetrievalAPI.md#GetSmfSelData) | **Get** /{supi}/smf-select-data | retrieve a UE&#39;s SMF Selection Subscription Data



## GetSmfSelData

> SmfSelectionSubscriptionData GetSmfSelData(ctx, supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's SMF Selection Subscription Data

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
	plmnId := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId | serving PLMN ID (optional)
	disasterRoamingInd := true // bool | Indication whether Disaster Roaming service is applied or not (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
	ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMFSelectionSubscriptionDataRetrievalAPI.GetSmfSelData(context.Background(), supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMFSelectionSubscriptionDataRetrievalAPI.GetSmfSelData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmfSelData`: SmfSelectionSubscriptionData
	fmt.Fprintf(os.Stdout, "Response from `SMFSelectionSubscriptionDataRetrievalAPI.GetSmfSelData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmfSelDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **plmnId** | [**PlmnId**](PlmnId.md) | serving PLMN ID | 
 **disasterRoamingInd** | **bool** | Indication whether Disaster Roaming service is applied or not | [default to false]
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

