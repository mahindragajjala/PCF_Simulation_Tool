# \SMFSmfRegistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmfRegistration**](SMFSmfRegistrationAPI.md#GetSmfRegistration) | **Get** /{ueId}/registrations/smf-registrations | retrieve the SMF registration information
[**Registration**](SMFSmfRegistrationAPI.md#Registration) | **Put** /{ueId}/registrations/smf-registrations/{pduSessionId} | register as SMF



## GetSmfRegistration

> SmfRegistrationInfo GetSmfRegistration(ctx, ueId).SingleNssai(singleNssai).Dnn(dnn).SupportedFeatures(supportedFeatures).Execute()

retrieve the SMF registration information

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
	singleNssai := *openapiclient.NewSnssai(int32(123)) // Snssai |  (optional)
	dnn := "dnn_example" // string |  (optional)
	supportedFeatures := "supportedFeatures_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMFSmfRegistrationAPI.GetSmfRegistration(context.Background(), ueId).SingleNssai(singleNssai).Dnn(dnn).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMFSmfRegistrationAPI.GetSmfRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmfRegistration`: SmfRegistrationInfo
	fmt.Fprintf(os.Stdout, "Response from `SMFSmfRegistrationAPI.GetSmfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleNssai** | [**Snssai**](Snssai.md) |  | 
 **dnn** | **string** |  | 
 **supportedFeatures** | **string** |  | 

### Return type

[**SmfRegistrationInfo**](SmfRegistrationInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Registration

> SmfRegistration Registration(ctx, ueId, pduSessionId).SmfRegistration(smfRegistration).Execute()

register as SMF

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
	smfRegistration := *openapiclient.NewSmfRegistration("SmfInstanceId_example", int32(123), *openapiclient.NewSnssai(int32(123)), *openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // SmfRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMFSmfRegistrationAPI.Registration(context.Background(), ueId, pduSessionId).SmfRegistration(smfRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMFSmfRegistrationAPI.Registration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Registration`: SmfRegistration
	fmt.Fprintf(os.Stdout, "Response from `SMFSmfRegistrationAPI.Registration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**pduSessionId** | **int32** | Identifier of the PDU session | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfRegistration** | [**SmfRegistration**](SmfRegistration.md) |  | 

### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

