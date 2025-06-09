# \AuthorizeTheNIDDConfigurationRequestAPI

All URIs are relative to *https://example.com/nudm-niddau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeNiddData**](AuthorizeTheNIDDConfigurationRequestAPI.md#AuthorizeNiddData) | **Post** /{ueIdentity}/authorize | Authorize the NIDD configuration request.



## AuthorizeNiddData

> AuthorizationData AuthorizeNiddData(ctx, ueIdentity).AuthorizationInfo(authorizationInfo).Execute()

Authorize the NIDD configuration request.

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
	ueIdentity := "ueIdentity_example" // string | Represents the scope of the UE for which the NIDD configuration are authorized. Contains the GPSI of the user or the external group ID.
	authorizationInfo := *openapiclient.NewAuthorizationInfo(*openapiclient.NewSnssai(int32(123)), "Dnn_example", "MtcProviderInformation_example", "AuthUpdateCallbackUri_example") // AuthorizationInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeTheNIDDConfigurationRequestAPI.AuthorizeNiddData(context.Background(), ueIdentity).AuthorizationInfo(authorizationInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeTheNIDDConfigurationRequestAPI.AuthorizeNiddData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeNiddData`: AuthorizationData
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeTheNIDDConfigurationRequestAPI.AuthorizeNiddData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the NIDD configuration are authorized. Contains the GPSI of the user or the external group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeNiddDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationInfo** | [**AuthorizationInfo**](AuthorizationInfo.md) |  | 

### Return type

[**AuthorizationData**](AuthorizationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

