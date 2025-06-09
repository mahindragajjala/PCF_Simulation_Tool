# \GenerateAuthDataAPI

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAuthData**](GenerateAuthDataAPI.md#GenerateAuthData) | **Post** /{supiOrSuci}/security-information/generate-auth-data | Generate authentication data for the UE



## GenerateAuthData

> AuthenticationInfoResult GenerateAuthData(ctx, supiOrSuci).AuthenticationInfoRequest(authenticationInfoRequest).Execute()

Generate authentication data for the UE

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
	authenticationInfoRequest := *openapiclient.NewAuthenticationInfoRequest("ServingNetworkName_example", "AusfInstanceId_example") // AuthenticationInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenerateAuthDataAPI.GenerateAuthData(context.Background(), supiOrSuci).AuthenticationInfoRequest(authenticationInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenerateAuthDataAPI.GenerateAuthData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAuthData`: AuthenticationInfoResult
	fmt.Fprintf(os.Stdout, "Response from `GenerateAuthDataAPI.GenerateAuthData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supiOrSuci** | **string** | SUPI or SUCI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAuthDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationInfoRequest** | [**AuthenticationInfoRequest**](AuthenticationInfoRequest.md) |  | 

### Return type

[**AuthenticationInfoResult**](AuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

