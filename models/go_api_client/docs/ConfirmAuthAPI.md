# \ConfirmAuthAPI

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmAuth**](ConfirmAuthAPI.md#ConfirmAuth) | **Post** /{supi}/auth-events | Create a new confirmation event



## ConfirmAuth

> AuthEvent ConfirmAuth(ctx, supi).AuthEvent(authEvent).Execute()

Create a new confirmation event

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
	authEvent := *openapiclient.NewAuthEvent("NfInstanceId_example", false, time.Now(), *openapiclient.NewAuthType(), "ServingNetworkName_example") // AuthEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfirmAuthAPI.ConfirmAuth(context.Background(), supi).AuthEvent(authEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfirmAuthAPI.ConfirmAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmAuth`: AuthEvent
	fmt.Fprintf(os.Stdout, "Response from `ConfirmAuthAPI.ConfirmAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | SUPI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authEvent** | [**AuthEvent**](AuthEvent.md) |  | 

### Return type

[**AuthEvent**](AuthEvent.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

