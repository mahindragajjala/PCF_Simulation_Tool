# \UECMRegistrationInfoRetrievalAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistrations**](UECMRegistrationInfoRetrievalAPI.md#GetRegistrations) | **Get** /{ueId}/registrations | retrieve UE registration data sets



## GetRegistrations

> RegistrationDataSets GetRegistrations(ctx, ueId).RegistrationDatasetNames(registrationDatasetNames).SupportedFeatures(supportedFeatures).SingleNssai(singleNssai).Dnn(dnn).Execute()

retrieve UE registration data sets

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
	registrationDatasetNames := []openapiclient.RegistrationDataSetName{*openapiclient.NewRegistrationDataSetName()} // []RegistrationDataSetName | List of UECM registration dataset names
	supportedFeatures := "supportedFeatures_example" // string |  (optional)
	singleNssai := *openapiclient.NewSnssai(int32(123)) // Snssai |  (optional)
	dnn := "dnn_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UECMRegistrationInfoRetrievalAPI.GetRegistrations(context.Background(), ueId).RegistrationDatasetNames(registrationDatasetNames).SupportedFeatures(supportedFeatures).SingleNssai(singleNssai).Dnn(dnn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UECMRegistrationInfoRetrievalAPI.GetRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistrations`: RegistrationDataSets
	fmt.Fprintf(os.Stdout, "Response from `UECMRegistrationInfoRetrievalAPI.GetRegistrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registrationDatasetNames** | [**[]RegistrationDataSetName**](RegistrationDataSetName.md) | List of UECM registration dataset names | 
 **supportedFeatures** | **string** |  | 
 **singleNssai** | [**Snssai**](Snssai.md) |  | 
 **dnn** | **string** |  | 

### Return type

[**RegistrationDataSets**](RegistrationDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

