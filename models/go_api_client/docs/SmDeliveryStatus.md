# SmDeliveryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**SmStatusReport** | **string** |  | 

## Methods

### NewSmDeliveryStatus

`func NewSmDeliveryStatus(gpsi string, smStatusReport string, ) *SmDeliveryStatus`

NewSmDeliveryStatus instantiates a new SmDeliveryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmDeliveryStatusWithDefaults

`func NewSmDeliveryStatusWithDefaults() *SmDeliveryStatus`

NewSmDeliveryStatusWithDefaults instantiates a new SmDeliveryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *SmDeliveryStatus) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmDeliveryStatus) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmDeliveryStatus) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSmStatusReport

`func (o *SmDeliveryStatus) GetSmStatusReport() string`

GetSmStatusReport returns the SmStatusReport field if non-nil, zero value otherwise.

### GetSmStatusReportOk

`func (o *SmDeliveryStatus) GetSmStatusReportOk() (*string, bool)`

GetSmStatusReportOk returns a tuple with the SmStatusReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmStatusReport

`func (o *SmDeliveryStatus) SetSmStatusReport(v string)`

SetSmStatusReport sets SmStatusReport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


