# \ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSmsfNon3GppRegistration**](ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessAPI.md#UpdateSmsfNon3GppRegistration) | **Patch** /{ueId}/registrations/smsf-non-3gpp-access | update a parameter in the SMSF registration for non-3GPP access



## UpdateSmsfNon3GppRegistration

> PatchResult UpdateSmsfNon3GppRegistration(ctx, ueId).SmsfRegistrationModification(smsfRegistrationModification).SupportedFeatures(supportedFeatures).Execute()

update a parameter in the SMSF registration for non-3GPP access

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
	smsfRegistrationModification := *openapiclient.NewSmsfRegistrationModification("SmsfInstanceId_example") // SmsfRegistrationModification | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessAPI.UpdateSmsfNon3GppRegistration(context.Background(), ueId).SmsfRegistrationModification(smsfRegistrationModification).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessAPI.UpdateSmsfNon3GppRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSmsfNon3GppRegistration`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessAPI.UpdateSmsfNon3GppRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmsfNon3GppRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfRegistrationModification** | [**SmsfRegistrationModification**](SmsfRegistrationModification.md) |  | 
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

