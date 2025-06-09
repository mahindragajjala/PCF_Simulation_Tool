# \RoamingInformationUpdateAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRoamingInformation**](RoamingInformationUpdateAPI.md#UpdateRoamingInformation) | **Post** /{ueId}/registrations/amf-3gpp-access/roaming-info-update | Update the Roaming Information



## UpdateRoamingInformation

> RoamingInfoUpdate UpdateRoamingInformation(ctx, ueId).RoamingInfoUpdate(roamingInfoUpdate).Execute()

Update the Roaming Information

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
	roamingInfoUpdate := *openapiclient.NewRoamingInfoUpdate(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // RoamingInfoUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoamingInformationUpdateAPI.UpdateRoamingInformation(context.Background(), ueId).RoamingInfoUpdate(roamingInfoUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoamingInformationUpdateAPI.UpdateRoamingInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoamingInformation`: RoamingInfoUpdate
	fmt.Fprintf(os.Stdout, "Response from `RoamingInformationUpdateAPI.UpdateRoamingInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoamingInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roamingInfoUpdate** | [**RoamingInfoUpdate**](RoamingInfoUpdate.md) |  | 

### Return type

[**RoamingInfoUpdate**](RoamingInfoUpdate.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

