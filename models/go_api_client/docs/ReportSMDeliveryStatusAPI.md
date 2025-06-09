# \ReportSMDeliveryStatusAPI

All URIs are relative to *https://example.com/nudm-rsds/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportSMDeliveryStatus**](ReportSMDeliveryStatusAPI.md#ReportSMDeliveryStatus) | **Post** /{ueIdentity}/sm-delivery-status | Report the SM Delivery Status



## ReportSMDeliveryStatus

> ReportSMDeliveryStatus(ctx, ueIdentity).SmDeliveryStatus(smDeliveryStatus).Execute()

Report the SM Delivery Status

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
	ueIdentity := "ueIdentity_example" // string | Represents the scope of the UE for which the Service Specific Parameters are authorized. Contains the GPSI of the user or the external group ID.
	smDeliveryStatus := *openapiclient.NewSmDeliveryStatus("Gpsi_example", "SmStatusReport_example") // SmDeliveryStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportSMDeliveryStatusAPI.ReportSMDeliveryStatus(context.Background(), ueIdentity).SmDeliveryStatus(smDeliveryStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSMDeliveryStatusAPI.ReportSMDeliveryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueIdentity** | **string** | Represents the scope of the UE for which the Service Specific Parameters are authorized. Contains the GPSI of the user or the external group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportSMDeliveryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smDeliveryStatus** | [**SmDeliveryStatus**](SmDeliveryStatus.md) |  | 

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

