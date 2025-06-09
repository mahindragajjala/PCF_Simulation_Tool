# \SMFDeregistrationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmfDeregistration**](SMFDeregistrationAPI.md#SmfDeregistration) | **Delete** /{ueId}/registrations/smf-registrations/{pduSessionId} | delete an SMF registration



## SmfDeregistration

> SmfDeregistration(ctx, ueId, pduSessionId).SmfSetId(smfSetId).SmfInstanceId(smfInstanceId).SmfEventsImplicitlyUnsubscribed(smfEventsImplicitlyUnsubscribed).Execute()

delete an SMF registration

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
	smfSetId := "smfSetId_example" // string |  (optional)
	smfInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	smfEventsImplicitlyUnsubscribed := true // bool | Indication on SMF event subscriptions implicitly unsubscribed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SMFDeregistrationAPI.SmfDeregistration(context.Background(), ueId, pduSessionId).SmfSetId(smfSetId).SmfInstanceId(smfInstanceId).SmfEventsImplicitlyUnsubscribed(smfEventsImplicitlyUnsubscribed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMFDeregistrationAPI.SmfDeregistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**pduSessionId** | **int32** | Identifier of the PDU session | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmfDeregistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfSetId** | **string** |  | 
 **smfInstanceId** | **string** |  | 
 **smfEventsImplicitlyUnsubscribed** | **bool** | Indication on SMF event subscriptions implicitly unsubscribed. | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

