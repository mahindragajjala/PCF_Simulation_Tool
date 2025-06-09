# \PEIUpdateAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeiUpdate**](PEIUpdateAPI.md#PeiUpdate) | **Post** /{ueId}/registrations/amf-3gpp-access/pei-update | Updates the PEI in the 3GPP access registration context



## PeiUpdate

> PeiUpdate(ctx, ueId).PeiUpdateInfo(peiUpdateInfo).Execute()

Updates the PEI in the 3GPP access registration context

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
	peiUpdateInfo := *openapiclient.NewPeiUpdateInfo("Pei_example") // PeiUpdateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PEIUpdateAPI.PeiUpdate(context.Background(), ueId).PeiUpdateInfo(peiUpdateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PEIUpdateAPI.PeiUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiPeiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **peiUpdateInfo** | [**PeiUpdateInfo**](PeiUpdateInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

