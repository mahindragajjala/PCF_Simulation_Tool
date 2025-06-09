# \TriggerAMFFor3GPPAccessDeregistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregAMF**](TriggerAMFFor3GPPAccessDeregistrationAPI.md#DeregAMF) | **Post** /{ueId}/registrations/amf-3gpp-access/dereg-amf | trigger AMF for 3GPP access deregistration



## DeregAMF

> DeregAMF(ctx, ueId).AmfDeregInfo(amfDeregInfo).Execute()

trigger AMF for 3GPP access deregistration

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
	amfDeregInfo := *openapiclient.NewAmfDeregInfo(*openapiclient.NewDeregistrationReason()) // AmfDeregInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAMFFor3GPPAccessDeregistrationAPI.DeregAMF(context.Background(), ueId).AmfDeregInfo(amfDeregInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAMFFor3GPPAccessDeregistrationAPI.DeregAMF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeregAMFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amfDeregInfo** | [**AmfDeregInfo**](AmfDeregInfo.md) |  | 

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

