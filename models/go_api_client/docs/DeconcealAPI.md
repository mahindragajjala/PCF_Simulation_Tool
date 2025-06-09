# \DeconcealAPI

All URIs are relative to *https://example.com/nudm-ueid/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deconceal**](DeconcealAPI.md#Deconceal) | **Post** /deconceal | Deconceal the SUCI to the SUPI



## Deconceal

> DeconcealRspData Deconceal(ctx).DeconcealReqData(deconcealReqData).Execute()

Deconceal the SUCI to the SUPI

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
	deconcealReqData := *openapiclient.NewDeconcealReqData("Suci_example") // DeconcealReqData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeconcealAPI.Deconceal(context.Background()).DeconcealReqData(deconcealReqData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeconcealAPI.Deconceal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Deconceal`: DeconcealRspData
	fmt.Fprintf(os.Stdout, "Response from `DeconcealAPI.Deconceal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeconcealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deconcealReqData** | [**DeconcealReqData**](DeconcealReqData.md) |  | 

### Return type

[**DeconcealRspData**](DeconcealRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

