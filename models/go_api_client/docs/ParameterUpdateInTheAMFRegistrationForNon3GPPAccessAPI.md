# \ParameterUpdateInTheAMFRegistrationForNon3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNon3GppRegistration**](ParameterUpdateInTheAMFRegistrationForNon3GPPAccessAPI.md#UpdateNon3GppRegistration) | **Patch** /{ueId}/registrations/amf-non-3gpp-access | update a parameter in the AMF registration for non-3GPP access



## UpdateNon3GppRegistration

> PatchResult UpdateNon3GppRegistration(ctx, ueId).AmfNon3GppAccessRegistrationModification(amfNon3GppAccessRegistrationModification).SupportedFeatures(supportedFeatures).Execute()

update a parameter in the AMF registration for non-3GPP access

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
	amfNon3GppAccessRegistrationModification := *openapiclient.NewAmfNon3GppAccessRegistrationModification(*openapiclient.NewGuami(*openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example"), "AmfId_example")) // AmfNon3GppAccessRegistrationModification | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParameterUpdateInTheAMFRegistrationForNon3GPPAccessAPI.UpdateNon3GppRegistration(context.Background(), ueId).AmfNon3GppAccessRegistrationModification(amfNon3GppAccessRegistrationModification).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParameterUpdateInTheAMFRegistrationForNon3GPPAccessAPI.UpdateNon3GppRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNon3GppRegistration`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `ParameterUpdateInTheAMFRegistrationForNon3GPPAccessAPI.UpdateNon3GppRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNon3GppRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amfNon3GppAccessRegistrationModification** | [**AmfNon3GppAccessRegistrationModification**](AmfNon3GppAccessRegistrationModification.md) |  | 
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

