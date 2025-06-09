# MbsrTimeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsrTimeWindow** | Pointer to [**[]ValidTimePeriod**](ValidTimePeriod.md) |  | [optional] 
**MbsrRecurTime** | Pointer to [**[]RecurTime**](RecurTime.md) |  | [optional] 

## Methods

### NewMbsrTimeInfo

`func NewMbsrTimeInfo() *MbsrTimeInfo`

NewMbsrTimeInfo instantiates a new MbsrTimeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsrTimeInfoWithDefaults

`func NewMbsrTimeInfoWithDefaults() *MbsrTimeInfo`

NewMbsrTimeInfoWithDefaults instantiates a new MbsrTimeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsrTimeWindow

`func (o *MbsrTimeInfo) GetMbsrTimeWindow() []ValidTimePeriod`

GetMbsrTimeWindow returns the MbsrTimeWindow field if non-nil, zero value otherwise.

### GetMbsrTimeWindowOk

`func (o *MbsrTimeInfo) GetMbsrTimeWindowOk() (*[]ValidTimePeriod, bool)`

GetMbsrTimeWindowOk returns a tuple with the MbsrTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsrTimeWindow

`func (o *MbsrTimeInfo) SetMbsrTimeWindow(v []ValidTimePeriod)`

SetMbsrTimeWindow sets MbsrTimeWindow field to given value.

### HasMbsrTimeWindow

`func (o *MbsrTimeInfo) HasMbsrTimeWindow() bool`

HasMbsrTimeWindow returns a boolean if a field has been set.

### GetMbsrRecurTime

`func (o *MbsrTimeInfo) GetMbsrRecurTime() []RecurTime`

GetMbsrRecurTime returns the MbsrRecurTime field if non-nil, zero value otherwise.

### GetMbsrRecurTimeOk

`func (o *MbsrTimeInfo) GetMbsrRecurTimeOk() (*[]RecurTime, bool)`

GetMbsrRecurTimeOk returns a tuple with the MbsrRecurTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsrRecurTime

`func (o *MbsrTimeInfo) SetMbsrRecurTime(v []RecurTime)`

SetMbsrRecurTime sets MbsrRecurTime field to given value.

### HasMbsrRecurTime

`func (o *MbsrTimeInfo) HasMbsrRecurTime() bool`

HasMbsrRecurTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


