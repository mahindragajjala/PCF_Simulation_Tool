# \GetAuthDataForFNRGAPI

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRgAuthData**](GetAuthDataForFNRGAPI.md#GetRgAuthData) | **Get** /{supiOrSuci}/security-information-rg | Get authentication data for the FN-RG



## GetRgAuthData

> RgAuthCtx GetRgAuthData(ctx, supiOrSuci).AuthenticatedInd(authenticatedInd).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Get authentication data for the FN-RG

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
	supiOrSuci := "supiOrSuci_example" // string | SUPI or SUCI of the user
	authenticatedInd := true // bool | Authenticated indication
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
	plmnId := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId | serving PLMN ID (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
	ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetAuthDataForFNRGAPI.GetRgAuthData(context.Background(), supiOrSuci).AuthenticatedInd(authenticatedInd).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetAuthDataForFNRGAPI.GetRgAuthData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRgAuthData`: RgAuthCtx
	fmt.Fprintf(os.Stdout, "Response from `GetAuthDataForFNRGAPI.GetRgAuthData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supiOrSuci** | **string** | SUPI or SUCI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRgAuthDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatedInd** | **bool** | Authenticated indication | 
 **supportedFeatures** | **string** | Supported Features | 
 **plmnId** | [**PlmnId**](PlmnId.md) | serving PLMN ID | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**RgAuthCtx**](RgAuthCtx.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

