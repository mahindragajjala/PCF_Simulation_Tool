# RecurTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurTimeWindow** | Pointer to [**ValidTimePeriod**](ValidTimePeriod.md) |  | [optional] 
**RecurType** | Pointer to [**RecurType**](RecurType.md) |  | [optional] 
**RecurMonth** | Pointer to **[]int32** |  | [optional] 
**RecurWeek** | Pointer to **[]int32** |  | [optional] 
**RecurDay** | Pointer to **[]int32** |  | [optional] 
**RecurDate** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewRecurTime

`func NewRecurTime() *RecurTime`

NewRecurTime instantiates a new RecurTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurTimeWithDefaults

`func NewRecurTimeWithDefaults() *RecurTime`

NewRecurTimeWithDefaults instantiates a new RecurTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurTimeWindow

`func (o *RecurTime) GetRecurTimeWindow() ValidTimePeriod`

GetRecurTimeWindow returns the RecurTimeWindow field if non-nil, zero value otherwise.

### GetRecurTimeWindowOk

`func (o *RecurTime) GetRecurTimeWindowOk() (*ValidTimePeriod, bool)`

GetRecurTimeWindowOk returns a tuple with the RecurTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurTimeWindow

`func (o *RecurTime) SetRecurTimeWindow(v ValidTimePeriod)`

SetRecurTimeWindow sets RecurTimeWindow field to given value.

### HasRecurTimeWindow

`func (o *RecurTime) HasRecurTimeWindow() bool`

HasRecurTimeWindow returns a boolean if a field has been set.

### GetRecurType

`func (o *RecurTime) GetRecurType() RecurType`

GetRecurType returns the RecurType field if non-nil, zero value otherwise.

### GetRecurTypeOk

`func (o *RecurTime) GetRecurTypeOk() (*RecurType, bool)`

GetRecurTypeOk returns a tuple with the RecurType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurType

`func (o *RecurTime) SetRecurType(v RecurType)`

SetRecurType sets RecurType field to given value.

### HasRecurType

`func (o *RecurTime) HasRecurType() bool`

HasRecurType returns a boolean if a field has been set.

### GetRecurMonth

`func (o *RecurTime) GetRecurMonth() []int32`

GetRecurMonth returns the RecurMonth field if non-nil, zero value otherwise.

### GetRecurMonthOk

`func (o *RecurTime) GetRecurMonthOk() (*[]int32, bool)`

GetRecurMonthOk returns a tuple with the RecurMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurMonth

`func (o *RecurTime) SetRecurMonth(v []int32)`

SetRecurMonth sets RecurMonth field to given value.

### HasRecurMonth

`func (o *RecurTime) HasRecurMonth() bool`

HasRecurMonth returns a boolean if a field has been set.

### GetRecurWeek

`func (o *RecurTime) GetRecurWeek() []int32`

GetRecurWeek returns the RecurWeek field if non-nil, zero value otherwise.

### GetRecurWeekOk

`func (o *RecurTime) GetRecurWeekOk() (*[]int32, bool)`

GetRecurWeekOk returns a tuple with the RecurWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurWeek

`func (o *RecurTime) SetRecurWeek(v []int32)`

SetRecurWeek sets RecurWeek field to given value.

### HasRecurWeek

`func (o *RecurTime) HasRecurWeek() bool`

HasRecurWeek returns a boolean if a field has been set.

### GetRecurDay

`func (o *RecurTime) GetRecurDay() []int32`

GetRecurDay returns the RecurDay field if non-nil, zero value otherwise.

### GetRecurDayOk

`func (o *RecurTime) GetRecurDayOk() (*[]int32, bool)`

GetRecurDayOk returns a tuple with the RecurDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurDay

`func (o *RecurTime) SetRecurDay(v []int32)`

SetRecurDay sets RecurDay field to given value.

### HasRecurDay

`func (o *RecurTime) HasRecurDay() bool`

HasRecurDay returns a boolean if a field has been set.

### GetRecurDate

`func (o *RecurTime) GetRecurDate() []time.Time`

GetRecurDate returns the RecurDate field if non-nil, zero value otherwise.

### GetRecurDateOk

`func (o *RecurTime) GetRecurDateOk() (*[]time.Time, bool)`

GetRecurDateOk returns a tuple with the RecurDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurDate

`func (o *RecurTime) SetRecurDate(v []time.Time)`

SetRecurDate sets RecurDate field to given value.

### HasRecurDate

`func (o *RecurTime) HasRecurDate() bool`

HasRecurDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


