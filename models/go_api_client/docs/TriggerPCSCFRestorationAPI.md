# \TriggerPCSCFRestorationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TriggerPCSCFRestoration**](TriggerPCSCFRestorationAPI.md#TriggerPCSCFRestoration) | **Post** /restore-pcscf | Trigger the Restoration of the P-CSCF



## TriggerPCSCFRestoration

> TriggerPCSCFRestoration(ctx).TriggerRequest(triggerRequest).Execute()

Trigger the Restoration of the P-CSCF

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
	triggerRequest := *openapiclient.NewTriggerRequest("Supi_example") // TriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerPCSCFRestorationAPI.TriggerPCSCFRestoration(context.Background()).TriggerRequest(triggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerPCSCFRestorationAPI.TriggerPCSCFRestoration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerPCSCFRestorationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerRequest** | [**TriggerRequest**](TriggerRequest.md) |  | 

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

