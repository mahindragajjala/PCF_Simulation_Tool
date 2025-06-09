# \ParameterUpdateInTheNWDAFRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNwdafRegistration**](ParameterUpdateInTheNWDAFRegistrationAPI.md#UpdateNwdafRegistration) | **Patch** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | Update a parameter in the NWDAF registration



## UpdateNwdafRegistration

> UpdateNwdafRegistration200Response UpdateNwdafRegistration(ctx, ueId, nwdafRegistrationId).NwdafRegistrationModification(nwdafRegistrationModification).SupportedFeatures(supportedFeatures).Execute()

Update a parameter in the NWDAF registration

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
	nwdafRegistrationId := "nwdafRegistrationId_example" // string | NWDAF registration identifier
	nwdafRegistrationModification := *openapiclient.NewNwdafRegistrationModification("NwdafInstanceId_example") // NwdafRegistrationModification | 
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParameterUpdateInTheNWDAFRegistrationAPI.UpdateNwdafRegistration(context.Background(), ueId, nwdafRegistrationId).NwdafRegistrationModification(nwdafRegistrationModification).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParameterUpdateInTheNWDAFRegistrationAPI.UpdateNwdafRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNwdafRegistration`: UpdateNwdafRegistration200Response
	fmt.Fprintf(os.Stdout, "Response from `ParameterUpdateInTheNWDAFRegistrationAPI.UpdateNwdafRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**nwdafRegistrationId** | **string** | NWDAF registration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNwdafRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nwdafRegistrationModification** | [**NwdafRegistrationModification**](NwdafRegistrationModification.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**UpdateNwdafRegistration200Response**](UpdateNwdafRegistration200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

