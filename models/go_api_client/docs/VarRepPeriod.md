# VarRepPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepPeriod** | **int32** | indicating a time in seconds. | 
**PercValueNfLoad** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewVarRepPeriod

`func NewVarRepPeriod(repPeriod int32, ) *VarRepPeriod`

NewVarRepPeriod instantiates a new VarRepPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVarRepPeriodWithDefaults

`func NewVarRepPeriodWithDefaults() *VarRepPeriod`

NewVarRepPeriodWithDefaults instantiates a new VarRepPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepPeriod

`func (o *VarRepPeriod) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *VarRepPeriod) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *VarRepPeriod) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.


### GetPercValueNfLoad

`func (o *VarRepPeriod) GetPercValueNfLoad() int32`

GetPercValueNfLoad returns the PercValueNfLoad field if non-nil, zero value otherwise.

### GetPercValueNfLoadOk

`func (o *VarRepPeriod) GetPercValueNfLoadOk() (*int32, bool)`

GetPercValueNfLoadOk returns a tuple with the PercValueNfLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercValueNfLoad

`func (o *VarRepPeriod) SetPercValueNfLoad(v int32)`

SetPercValueNfLoad sets PercValueNfLoad field to given value.

### HasPercValueNfLoad

`func (o *VarRepPeriod) HasPercValueNfLoad() bool`

HasPercValueNfLoad returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


