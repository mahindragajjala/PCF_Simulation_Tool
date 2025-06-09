# \DeleteAuthAPI

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuth**](DeleteAuthAPI.md#DeleteAuth) | **Put** /{supi}/auth-events/{authEventId} | Deletes the authentication result in the UDM



## DeleteAuth

> DeleteAuth(ctx, supi, authEventId).AuthEvent(authEvent).Execute()

Deletes the authentication result in the UDM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	supi := "supi_example" // string | SUPI of the user
	authEventId := "authEventId_example" // string | authEvent Id
	authEvent := *openapiclient.NewAuthEvent("NfInstanceId_example", false, time.Now(), *openapiclient.NewAuthType(), "ServingNetworkName_example") // AuthEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeleteAuthAPI.DeleteAuth(context.Background(), supi, authEventId).AuthEvent(authEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteAuthAPI.DeleteAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | SUPI of the user | 
**authEventId** | **string** | authEvent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authEvent** | [**AuthEvent**](AuthEvent.md) |  | 

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

