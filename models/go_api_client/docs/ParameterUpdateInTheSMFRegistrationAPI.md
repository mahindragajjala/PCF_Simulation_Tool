# \ParameterUpdateInTheSMFRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSmfRegistration**](ParameterUpdateInTheSMFRegistrationAPI.md#UpdateSmfRegistration) | **Patch** /{ueId}/registrations/smf-registrations/{pduSessionId} | update a parameter in the SMF registration



## UpdateSmfRegistration

> PatchResult UpdateSmfRegistration(ctx, ueId, pduSessionId).SmfRegistrationModification(smfRegistrationModification).SupportedFeatures(supportedFeatures).Execute()

update a parameter in the SMF registration

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
	pduSessionId := int32(56) // int32 | Identifier of the PDU session
	smfRegistrationModification := *openapiclient.NewSmfRegistrationModification("SmfInstanceId_example") // SmfRegistrationModification | 
	supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParameterUpdateInTheSMFRegistrationAPI.UpdateSmfRegistration(context.Background(), ueId, pduSessionId).SmfRegistrationModification(smfRegistrationModification).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParameterUpdateInTheSMFRegistrationAPI.UpdateSmfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSmfRegistration`: PatchResult
	fmt.Fprintf(os.Stdout, "Response from `ParameterUpdateInTheSMFRegistrationAPI.UpdateSmfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**pduSessionId** | **int32** | Identifier of the PDU session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfRegistrationModification** | [**SmfRegistrationModification**](SmfRegistrationModification.md) |  | 
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

