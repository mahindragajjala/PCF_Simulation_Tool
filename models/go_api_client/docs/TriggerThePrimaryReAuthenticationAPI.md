# \TriggerThePrimaryReAuthenticationAPI

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTrigger**](TriggerThePrimaryReAuthenticationAPI.md#AuthTrigger) | **Get** /{ueId}/registrations/auth-trigger | trigger the primary (re-)authentication



## AuthTrigger

> AuthTrigger(ctx, ueId).AuthTriggerInfo(authTriggerInfo).Execute()

trigger the primary (re-)authentication

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
	authTriggerInfo := *openapiclient.NewAuthTriggerInfo() // AuthTriggerInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerThePrimaryReAuthenticationAPI.AuthTrigger(context.Background(), ueId).AuthTriggerInfo(authTriggerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerThePrimaryReAuthenticationAPI.AuthTrigger``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAuthTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authTriggerInfo** | [**AuthTriggerInfo**](AuthTriggerInfo.md) |  | 

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

