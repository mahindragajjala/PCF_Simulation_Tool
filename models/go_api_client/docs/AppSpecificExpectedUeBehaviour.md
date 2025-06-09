# AppSpecificExpectedUeBehaviour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | 
**AppSpecificExpectedUeBehaviourData** | [**map[string]AppSpecificExpectedUeBehaviourData**](AppSpecificExpectedUeBehaviourData.md) | A map(list of key-value pairs) where a valid JSON pointer serves as key | 

## Methods

### NewAppSpecificExpectedUeBehaviour

`func NewAppSpecificExpectedUeBehaviour(afInstanceId string, referenceId int32, appSpecificExpectedUeBehaviourData map[string]AppSpecificExpectedUeBehaviourData, ) *AppSpecificExpectedUeBehaviour`

NewAppSpecificExpectedUeBehaviour instantiates a new AppSpecificExpectedUeBehaviour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSpecificExpectedUeBehaviourWithDefaults

`func NewAppSpecificExpectedUeBehaviourWithDefaults() *AppSpecificExpectedUeBehaviour`

NewAppSpecificExpectedUeBehaviourWithDefaults instantiates a new AppSpecificExpectedUeBehaviour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *AppSpecificExpectedUeBehaviour) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *AppSpecificExpectedUeBehaviour) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *AppSpecificExpectedUeBehaviour) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *AppSpecificExpectedUeBehaviour) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AppSpecificExpectedUeBehaviour) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AppSpecificExpectedUeBehaviour) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetAppSpecificExpectedUeBehaviourData

`func (o *AppSpecificExpectedUeBehaviour) GetAppSpecificExpectedUeBehaviourData() map[string]AppSpecificExpectedUeBehaviourData`

GetAppSpecificExpectedUeBehaviourData returns the AppSpecificExpectedUeBehaviourData field if non-nil, zero value otherwise.

### GetAppSpecificExpectedUeBehaviourDataOk

`func (o *AppSpecificExpectedUeBehaviour) GetAppSpecificExpectedUeBehaviourDataOk() (*map[string]AppSpecificExpectedUeBehaviourData, bool)`

GetAppSpecificExpectedUeBehaviourDataOk returns a tuple with the AppSpecificExpectedUeBehaviourData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpecificExpectedUeBehaviourData

`func (o *AppSpecificExpectedUeBehaviour) SetAppSpecificExpectedUeBehaviourData(v map[string]AppSpecificExpectedUeBehaviourData)`

SetAppSpecificExpectedUeBehaviourData sets AppSpecificExpectedUeBehaviourData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


