# \GenerateHSSAuthenticationVectorsAPI

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAv**](GenerateHSSAuthenticationVectorsAPI.md#GenerateAv) | **Post** /{supi}/hss-security-information/{hssAuthType}/generate-av | Generate authentication data for the UE in EPS or IMS domain



## GenerateAv

> HssAuthenticationInfoResult GenerateAv(ctx, supi, hssAuthType).HssAuthenticationInfoRequest(hssAuthenticationInfoRequest).Execute()

Generate authentication data for the UE in EPS or IMS domain

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
	supi := "supi_example" // string | SUPI of the user
	hssAuthType := *openapiclient.NewHssAuthTypeInUri() // HssAuthTypeInUri | Type of AV requested by HSS
	hssAuthenticationInfoRequest := *openapiclient.NewHssAuthenticationInfoRequest(*openapiclient.NewHssAuthType(), int32(123)) // HssAuthenticationInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenerateHSSAuthenticationVectorsAPI.GenerateAv(context.Background(), supi, hssAuthType).HssAuthenticationInfoRequest(hssAuthenticationInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenerateHSSAuthenticationVectorsAPI.GenerateAv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAv`: HssAuthenticationInfoResult
	fmt.Fprintf(os.Stdout, "Response from `GenerateHSSAuthenticationVectorsAPI.GenerateAv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | SUPI of the user | 
**hssAuthType** | [**HssAuthTypeInUri**](.md) | Type of AV requested by HSS | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hssAuthenticationInfoRequest** | [**HssAuthenticationInfoRequest**](HssAuthenticationInfoRequest.md) |  | 

### Return type

[**HssAuthenticationInfoResult**](HssAuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

