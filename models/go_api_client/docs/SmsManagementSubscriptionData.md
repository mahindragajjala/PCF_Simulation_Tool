# SmsManagementSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MtSmsSubscribed** | Pointer to **bool** |  | [optional] 
**MtSmsBarringAll** | Pointer to **bool** |  | [optional] 
**MtSmsBarringRoaming** | Pointer to **bool** |  | [optional] 
**MoSmsSubscribed** | Pointer to **bool** |  | [optional] 
**MoSmsBarringAll** | Pointer to **bool** |  | [optional] 
**MoSmsBarringRoaming** | Pointer to **bool** |  | [optional] 
**SharedSmsMngDataIds** | Pointer to **[]string** |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SharedTraceDataId** | Pointer to **string** |  | [optional] 

## Methods

### NewSmsManagementSubscriptionData

`func NewSmsManagementSubscriptionData() *SmsManagementSubscriptionData`

NewSmsManagementSubscriptionData instantiates a new SmsManagementSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsManagementSubscriptionDataWithDefaults

`func NewSmsManagementSubscriptionDataWithDefaults() *SmsManagementSubscriptionData`

NewSmsManagementSubscriptionDataWithDefaults instantiates a new SmsManagementSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *SmsManagementSubscriptionData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmsManagementSubscriptionData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmsManagementSubscriptionData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmsManagementSubscriptionData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) GetMtSmsSubscribed() bool`

GetMtSmsSubscribed returns the MtSmsSubscribed field if non-nil, zero value otherwise.

### GetMtSmsSubscribedOk

`func (o *SmsManagementSubscriptionData) GetMtSmsSubscribedOk() (*bool, bool)`

GetMtSmsSubscribedOk returns a tuple with the MtSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) SetMtSmsSubscribed(v bool)`

SetMtSmsSubscribed sets MtSmsSubscribed field to given value.

### HasMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) HasMtSmsSubscribed() bool`

HasMtSmsSubscribed returns a boolean if a field has been set.

### GetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringAll() bool`

GetMtSmsBarringAll returns the MtSmsBarringAll field if non-nil, zero value otherwise.

### GetMtSmsBarringAllOk

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringAllOk() (*bool, bool)`

GetMtSmsBarringAllOk returns a tuple with the MtSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) SetMtSmsBarringAll(v bool)`

SetMtSmsBarringAll sets MtSmsBarringAll field to given value.

### HasMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) HasMtSmsBarringAll() bool`

HasMtSmsBarringAll returns a boolean if a field has been set.

### GetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoaming() bool`

GetMtSmsBarringRoaming returns the MtSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMtSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoamingOk() (*bool, bool)`

GetMtSmsBarringRoamingOk returns a tuple with the MtSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) SetMtSmsBarringRoaming(v bool)`

SetMtSmsBarringRoaming sets MtSmsBarringRoaming field to given value.

### HasMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) HasMtSmsBarringRoaming() bool`

HasMtSmsBarringRoaming returns a boolean if a field has been set.

### GetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) GetMoSmsSubscribed() bool`

GetMoSmsSubscribed returns the MoSmsSubscribed field if non-nil, zero value otherwise.

### GetMoSmsSubscribedOk

`func (o *SmsManagementSubscriptionData) GetMoSmsSubscribedOk() (*bool, bool)`

GetMoSmsSubscribedOk returns a tuple with the MoSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) SetMoSmsSubscribed(v bool)`

SetMoSmsSubscribed sets MoSmsSubscribed field to given value.

### HasMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) HasMoSmsSubscribed() bool`

HasMoSmsSubscribed returns a boolean if a field has been set.

### GetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringAll() bool`

GetMoSmsBarringAll returns the MoSmsBarringAll field if non-nil, zero value otherwise.

### GetMoSmsBarringAllOk

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringAllOk() (*bool, bool)`

GetMoSmsBarringAllOk returns a tuple with the MoSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) SetMoSmsBarringAll(v bool)`

SetMoSmsBarringAll sets MoSmsBarringAll field to given value.

### HasMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) HasMoSmsBarringAll() bool`

HasMoSmsBarringAll returns a boolean if a field has been set.

### GetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoaming() bool`

GetMoSmsBarringRoaming returns the MoSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMoSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoamingOk() (*bool, bool)`

GetMoSmsBarringRoamingOk returns a tuple with the MoSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) SetMoSmsBarringRoaming(v bool)`

SetMoSmsBarringRoaming sets MoSmsBarringRoaming field to given value.

### HasMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) HasMoSmsBarringRoaming() bool`

HasMoSmsBarringRoaming returns a boolean if a field has been set.

### GetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIds() []string`

GetSharedSmsMngDataIds returns the SharedSmsMngDataIds field if non-nil, zero value otherwise.

### GetSharedSmsMngDataIdsOk

`func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIdsOk() (*[]string, bool)`

GetSharedSmsMngDataIdsOk returns a tuple with the SharedSmsMngDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) SetSharedSmsMngDataIds(v []string)`

SetSharedSmsMngDataIds sets SharedSmsMngDataIds field to given value.

### HasSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) HasSharedSmsMngDataIds() bool`

HasSharedSmsMngDataIds returns a boolean if a field has been set.

### GetTraceData

`func (o *SmsManagementSubscriptionData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmsManagementSubscriptionData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmsManagementSubscriptionData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmsManagementSubscriptionData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmsManagementSubscriptionData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmsManagementSubscriptionData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSharedTraceDataId

`func (o *SmsManagementSubscriptionData) GetSharedTraceDataId() string`

GetSharedTraceDataId returns the SharedTraceDataId field if non-nil, zero value otherwise.

### GetSharedTraceDataIdOk

`func (o *SmsManagementSubscriptionData) GetSharedTraceDataIdOk() (*string, bool)`

GetSharedTraceDataIdOk returns a tuple with the SharedTraceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTraceDataId

`func (o *SmsManagementSubscriptionData) SetSharedTraceDataId(v string)`

SetSharedTraceDataId sets SharedTraceDataId field to given value.

### HasSharedTraceDataId

`func (o *SmsManagementSubscriptionData) HasSharedTraceDataId() bool`

HasSharedTraceDataId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


