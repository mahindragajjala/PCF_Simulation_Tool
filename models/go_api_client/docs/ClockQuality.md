# ClockQuality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceabilityToGnss** | Pointer to **bool** |  | [optional] 
**TraceabilityToUtc** | Pointer to **bool** |  | [optional] 
**FrequencyStability** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.  | [optional] 
**ClockAccuracyIndex** | Pointer to **string** |  | [optional] 
**ClockAccuracyValue** | Pointer to **int32** |  | [optional] 

## Methods

### NewClockQuality

`func NewClockQuality() *ClockQuality`

NewClockQuality instantiates a new ClockQuality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClockQualityWithDefaults

`func NewClockQualityWithDefaults() *ClockQuality`

NewClockQualityWithDefaults instantiates a new ClockQuality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceabilityToGnss

`func (o *ClockQuality) GetTraceabilityToGnss() bool`

GetTraceabilityToGnss returns the TraceabilityToGnss field if non-nil, zero value otherwise.

### GetTraceabilityToGnssOk

`func (o *ClockQuality) GetTraceabilityToGnssOk() (*bool, bool)`

GetTraceabilityToGnssOk returns a tuple with the TraceabilityToGnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceabilityToGnss

`func (o *ClockQuality) SetTraceabilityToGnss(v bool)`

SetTraceabilityToGnss sets TraceabilityToGnss field to given value.

### HasTraceabilityToGnss

`func (o *ClockQuality) HasTraceabilityToGnss() bool`

HasTraceabilityToGnss returns a boolean if a field has been set.

### GetTraceabilityToUtc

`func (o *ClockQuality) GetTraceabilityToUtc() bool`

GetTraceabilityToUtc returns the TraceabilityToUtc field if non-nil, zero value otherwise.

### GetTraceabilityToUtcOk

`func (o *ClockQuality) GetTraceabilityToUtcOk() (*bool, bool)`

GetTraceabilityToUtcOk returns a tuple with the TraceabilityToUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceabilityToUtc

`func (o *ClockQuality) SetTraceabilityToUtc(v bool)`

SetTraceabilityToUtc sets TraceabilityToUtc field to given value.

### HasTraceabilityToUtc

`func (o *ClockQuality) HasTraceabilityToUtc() bool`

HasTraceabilityToUtc returns a boolean if a field has been set.

### GetFrequencyStability

`func (o *ClockQuality) GetFrequencyStability() int32`

GetFrequencyStability returns the FrequencyStability field if non-nil, zero value otherwise.

### GetFrequencyStabilityOk

`func (o *ClockQuality) GetFrequencyStabilityOk() (*int32, bool)`

GetFrequencyStabilityOk returns a tuple with the FrequencyStability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyStability

`func (o *ClockQuality) SetFrequencyStability(v int32)`

SetFrequencyStability sets FrequencyStability field to given value.

### HasFrequencyStability

`func (o *ClockQuality) HasFrequencyStability() bool`

HasFrequencyStability returns a boolean if a field has been set.

### GetClockAccuracyIndex

`func (o *ClockQuality) GetClockAccuracyIndex() string`

GetClockAccuracyIndex returns the ClockAccuracyIndex field if non-nil, zero value otherwise.

### GetClockAccuracyIndexOk

`func (o *ClockQuality) GetClockAccuracyIndexOk() (*string, bool)`

GetClockAccuracyIndexOk returns a tuple with the ClockAccuracyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockAccuracyIndex

`func (o *ClockQuality) SetClockAccuracyIndex(v string)`

SetClockAccuracyIndex sets ClockAccuracyIndex field to given value.

### HasClockAccuracyIndex

`func (o *ClockQuality) HasClockAccuracyIndex() bool`

HasClockAccuracyIndex returns a boolean if a field has been set.

### GetClockAccuracyValue

`func (o *ClockQuality) GetClockAccuracyValue() int32`

GetClockAccuracyValue returns the ClockAccuracyValue field if non-nil, zero value otherwise.

### GetClockAccuracyValueOk

`func (o *ClockQuality) GetClockAccuracyValueOk() (*int32, bool)`

GetClockAccuracyValueOk returns a tuple with the ClockAccuracyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockAccuracyValue

`func (o *ClockQuality) SetClockAccuracyValue(v int32)`

SetClockAccuracyValue sets ClockAccuracyValue field to given value.

### HasClockAccuracyValue

`func (o *ClockQuality) HasClockAccuracyValue() bool`

HasClockAccuracyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


