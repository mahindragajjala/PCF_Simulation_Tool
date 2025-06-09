# RoutingInfoSmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSmGwInd** | Pointer to **bool** |  | [optional] [default to false]
**CorrelationId** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewRoutingInfoSmRequest

`func NewRoutingInfoSmRequest() *RoutingInfoSmRequest`

NewRoutingInfoSmRequest instantiates a new RoutingInfoSmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingInfoSmRequestWithDefaults

`func NewRoutingInfoSmRequestWithDefaults() *RoutingInfoSmRequest`

NewRoutingInfoSmRequestWithDefaults instantiates a new RoutingInfoSmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSmGwInd

`func (o *RoutingInfoSmRequest) GetIpSmGwInd() bool`

GetIpSmGwInd returns the IpSmGwInd field if non-nil, zero value otherwise.

### GetIpSmGwIndOk

`func (o *RoutingInfoSmRequest) GetIpSmGwIndOk() (*bool, bool)`

GetIpSmGwIndOk returns a tuple with the IpSmGwInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwInd

`func (o *RoutingInfoSmRequest) SetIpSmGwInd(v bool)`

SetIpSmGwInd sets IpSmGwInd field to given value.

### HasIpSmGwInd

`func (o *RoutingInfoSmRequest) HasIpSmGwInd() bool`

HasIpSmGwInd returns a boolean if a field has been set.

### GetCorrelationId

`func (o *RoutingInfoSmRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *RoutingInfoSmRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *RoutingInfoSmRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *RoutingInfoSmRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *RoutingInfoSmRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RoutingInfoSmRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RoutingInfoSmRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RoutingInfoSmRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


