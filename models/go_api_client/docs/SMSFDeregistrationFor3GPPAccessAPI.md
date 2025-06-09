# \SMSFDeregistrationFor3GPPAccessAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call3GppSmsfDeregistration**](SMSFDeregistrationFor3GPPAccessAPI.md#Call3GppSmsfDeregistration) | **Delete** /{ueId}/registrations/smsf-3gpp-access | delete the SMSF registration for 3GPP access



## Call3GppSmsfDeregistration

> Call3GppSmsfDeregistration(ctx, ueId).SmsfSetId(smsfSetId).IfMatch(ifMatch).Execute()

delete the SMSF registration for 3GPP access

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
	smsfSetId := "smsfSetId_example" // string |  (optional)
	ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in IETF RFC 9110, 3.1 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SMSFDeregistrationFor3GPPAccessAPI.Call3GppSmsfDeregistration(context.Background(), ueId).SmsfSetId(smsfSetId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSFDeregistrationFor3GPPAccessAPI.Call3GppSmsfDeregistration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCall3GppSmsfDeregistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfSetId** | **string** |  | 
 **ifMatch** | **string** | Validator for conditional requests, as described in IETF RFC 9110, 3.1 | 

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

