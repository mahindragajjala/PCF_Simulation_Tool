# ExpectedUeBehaviourExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | 
**ExpectedUeBehaviourData** | Pointer to [**map[string]ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) | A map(list of key-value pairs) where a valid JSON pointer serves as key | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 

## Methods

### NewExpectedUeBehaviourExtension

`func NewExpectedUeBehaviourExtension(afInstanceId string, referenceId int32, ) *ExpectedUeBehaviourExtension`

NewExpectedUeBehaviourExtension instantiates a new ExpectedUeBehaviourExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedUeBehaviourExtensionWithDefaults

`func NewExpectedUeBehaviourExtensionWithDefaults() *ExpectedUeBehaviourExtension`

NewExpectedUeBehaviourExtensionWithDefaults instantiates a new ExpectedUeBehaviourExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *ExpectedUeBehaviourExtension) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *ExpectedUeBehaviourExtension) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *ExpectedUeBehaviourExtension) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *ExpectedUeBehaviourExtension) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ExpectedUeBehaviourExtension) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ExpectedUeBehaviourExtension) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetExpectedUeBehaviourData

`func (o *ExpectedUeBehaviourExtension) GetExpectedUeBehaviourData() map[string]ExpectedUeBehaviourData`

GetExpectedUeBehaviourData returns the ExpectedUeBehaviourData field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourDataOk

`func (o *ExpectedUeBehaviourExtension) GetExpectedUeBehaviourDataOk() (*map[string]ExpectedUeBehaviourData, bool)`

GetExpectedUeBehaviourDataOk returns a tuple with the ExpectedUeBehaviourData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourData

`func (o *ExpectedUeBehaviourExtension) SetExpectedUeBehaviourData(v map[string]ExpectedUeBehaviourData)`

SetExpectedUeBehaviourData sets ExpectedUeBehaviourData field to given value.

### HasExpectedUeBehaviourData

`func (o *ExpectedUeBehaviourExtension) HasExpectedUeBehaviourData() bool`

HasExpectedUeBehaviourData returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *ExpectedUeBehaviourExtension) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *ExpectedUeBehaviourExtension) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *ExpectedUeBehaviourExtension) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *ExpectedUeBehaviourExtension) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


