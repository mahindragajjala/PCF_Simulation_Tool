# SliceUsageControlInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssai** | [**Snssai**](Snssai.md) |  | 
**DeregInactTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SessInactTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewSliceUsageControlInfo

`func NewSliceUsageControlInfo(sNssai Snssai, ) *SliceUsageControlInfo`

NewSliceUsageControlInfo instantiates a new SliceUsageControlInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceUsageControlInfoWithDefaults

`func NewSliceUsageControlInfoWithDefaults() *SliceUsageControlInfo`

NewSliceUsageControlInfoWithDefaults instantiates a new SliceUsageControlInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssai

`func (o *SliceUsageControlInfo) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SliceUsageControlInfo) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SliceUsageControlInfo) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetDeregInactTimer

`func (o *SliceUsageControlInfo) GetDeregInactTimer() int32`

GetDeregInactTimer returns the DeregInactTimer field if non-nil, zero value otherwise.

### GetDeregInactTimerOk

`func (o *SliceUsageControlInfo) GetDeregInactTimerOk() (*int32, bool)`

GetDeregInactTimerOk returns a tuple with the DeregInactTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregInactTimer

`func (o *SliceUsageControlInfo) SetDeregInactTimer(v int32)`

SetDeregInactTimer sets DeregInactTimer field to given value.

### HasDeregInactTimer

`func (o *SliceUsageControlInfo) HasDeregInactTimer() bool`

HasDeregInactTimer returns a boolean if a field has been set.

### GetSessInactTimer

`func (o *SliceUsageControlInfo) GetSessInactTimer() int32`

GetSessInactTimer returns the SessInactTimer field if non-nil, zero value otherwise.

### GetSessInactTimerOk

`func (o *SliceUsageControlInfo) GetSessInactTimerOk() (*int32, bool)`

GetSessInactTimerOk returns a tuple with the SessInactTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInactTimer

`func (o *SliceUsageControlInfo) SetSessInactTimer(v int32)`

SetSessInactTimer sets SessInactTimer field to given value.

### HasSessInactTimer

`func (o *SliceUsageControlInfo) HasSessInactTimer() bool`

HasSessInactTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


