# ReportingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportMode** | Pointer to [**EventReportMode**](EventReportMode.md) |  | [optional] 
**MaxNumOfReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SamplingRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**GuardTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ReportPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifFlag** | Pointer to [**NotificationFlag**](NotificationFlag.md) |  | [optional] 
**MutingExcInstructions** | Pointer to [**MutingExceptionInstructions**](MutingExceptionInstructions.md) |  | [optional] 
**MutingNotSettings** | Pointer to [**MutingNotificationsSettings**](MutingNotificationsSettings.md) |  | [optional] [readonly] 
**VarRepPeriodInfo** | Pointer to [**[]VarRepPeriod**](VarRepPeriod.md) |  | [optional] 

## Methods

### NewReportingOptions

`func NewReportingOptions() *ReportingOptions`

NewReportingOptions instantiates a new ReportingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingOptionsWithDefaults

`func NewReportingOptionsWithDefaults() *ReportingOptions`

NewReportingOptionsWithDefaults instantiates a new ReportingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportMode

`func (o *ReportingOptions) GetReportMode() EventReportMode`

GetReportMode returns the ReportMode field if non-nil, zero value otherwise.

### GetReportModeOk

`func (o *ReportingOptions) GetReportModeOk() (*EventReportMode, bool)`

GetReportModeOk returns a tuple with the ReportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportMode

`func (o *ReportingOptions) SetReportMode(v EventReportMode)`

SetReportMode sets ReportMode field to given value.

### HasReportMode

`func (o *ReportingOptions) HasReportMode() bool`

HasReportMode returns a boolean if a field has been set.

### GetMaxNumOfReports

`func (o *ReportingOptions) GetMaxNumOfReports() int32`

GetMaxNumOfReports returns the MaxNumOfReports field if non-nil, zero value otherwise.

### GetMaxNumOfReportsOk

`func (o *ReportingOptions) GetMaxNumOfReportsOk() (*int32, bool)`

GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfReports

`func (o *ReportingOptions) SetMaxNumOfReports(v int32)`

SetMaxNumOfReports sets MaxNumOfReports field to given value.

### HasMaxNumOfReports

`func (o *ReportingOptions) HasMaxNumOfReports() bool`

HasMaxNumOfReports returns a boolean if a field has been set.

### GetExpiry

`func (o *ReportingOptions) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ReportingOptions) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ReportingOptions) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ReportingOptions) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSamplingRatio

`func (o *ReportingOptions) GetSamplingRatio() int32`

GetSamplingRatio returns the SamplingRatio field if non-nil, zero value otherwise.

### GetSamplingRatioOk

`func (o *ReportingOptions) GetSamplingRatioOk() (*int32, bool)`

GetSamplingRatioOk returns a tuple with the SamplingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRatio

`func (o *ReportingOptions) SetSamplingRatio(v int32)`

SetSamplingRatio sets SamplingRatio field to given value.

### HasSamplingRatio

`func (o *ReportingOptions) HasSamplingRatio() bool`

HasSamplingRatio returns a boolean if a field has been set.

### GetGuardTime

`func (o *ReportingOptions) GetGuardTime() int32`

GetGuardTime returns the GuardTime field if non-nil, zero value otherwise.

### GetGuardTimeOk

`func (o *ReportingOptions) GetGuardTimeOk() (*int32, bool)`

GetGuardTimeOk returns a tuple with the GuardTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardTime

`func (o *ReportingOptions) SetGuardTime(v int32)`

SetGuardTime sets GuardTime field to given value.

### HasGuardTime

`func (o *ReportingOptions) HasGuardTime() bool`

HasGuardTime returns a boolean if a field has been set.

### GetReportPeriod

`func (o *ReportingOptions) GetReportPeriod() int32`

GetReportPeriod returns the ReportPeriod field if non-nil, zero value otherwise.

### GetReportPeriodOk

`func (o *ReportingOptions) GetReportPeriodOk() (*int32, bool)`

GetReportPeriodOk returns a tuple with the ReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPeriod

`func (o *ReportingOptions) SetReportPeriod(v int32)`

SetReportPeriod sets ReportPeriod field to given value.

### HasReportPeriod

`func (o *ReportingOptions) HasReportPeriod() bool`

HasReportPeriod returns a boolean if a field has been set.

### GetNotifFlag

`func (o *ReportingOptions) GetNotifFlag() NotificationFlag`

GetNotifFlag returns the NotifFlag field if non-nil, zero value otherwise.

### GetNotifFlagOk

`func (o *ReportingOptions) GetNotifFlagOk() (*NotificationFlag, bool)`

GetNotifFlagOk returns a tuple with the NotifFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifFlag

`func (o *ReportingOptions) SetNotifFlag(v NotificationFlag)`

SetNotifFlag sets NotifFlag field to given value.

### HasNotifFlag

`func (o *ReportingOptions) HasNotifFlag() bool`

HasNotifFlag returns a boolean if a field has been set.

### GetMutingExcInstructions

`func (o *ReportingOptions) GetMutingExcInstructions() MutingExceptionInstructions`

GetMutingExcInstructions returns the MutingExcInstructions field if non-nil, zero value otherwise.

### GetMutingExcInstructionsOk

`func (o *ReportingOptions) GetMutingExcInstructionsOk() (*MutingExceptionInstructions, bool)`

GetMutingExcInstructionsOk returns a tuple with the MutingExcInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutingExcInstructions

`func (o *ReportingOptions) SetMutingExcInstructions(v MutingExceptionInstructions)`

SetMutingExcInstructions sets MutingExcInstructions field to given value.

### HasMutingExcInstructions

`func (o *ReportingOptions) HasMutingExcInstructions() bool`

HasMutingExcInstructions returns a boolean if a field has been set.

### GetMutingNotSettings

`func (o *ReportingOptions) GetMutingNotSettings() MutingNotificationsSettings`

GetMutingNotSettings returns the MutingNotSettings field if non-nil, zero value otherwise.

### GetMutingNotSettingsOk

`func (o *ReportingOptions) GetMutingNotSettingsOk() (*MutingNotificationsSettings, bool)`

GetMutingNotSettingsOk returns a tuple with the MutingNotSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutingNotSettings

`func (o *ReportingOptions) SetMutingNotSettings(v MutingNotificationsSettings)`

SetMutingNotSettings sets MutingNotSettings field to given value.

### HasMutingNotSettings

`func (o *ReportingOptions) HasMutingNotSettings() bool`

HasMutingNotSettings returns a boolean if a field has been set.

### GetVarRepPeriodInfo

`func (o *ReportingOptions) GetVarRepPeriodInfo() []VarRepPeriod`

GetVarRepPeriodInfo returns the VarRepPeriodInfo field if non-nil, zero value otherwise.

### GetVarRepPeriodInfoOk

`func (o *ReportingOptions) GetVarRepPeriodInfoOk() (*[]VarRepPeriod, bool)`

GetVarRepPeriodInfoOk returns a tuple with the VarRepPeriodInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarRepPeriodInfo

`func (o *ReportingOptions) SetVarRepPeriodInfo(v []VarRepPeriod)`

SetVarRepPeriodInfo sets VarRepPeriodInfo field to given value.

### HasVarRepPeriodInfo

`func (o *ReportingOptions) HasVarRepPeriodInfo() bool`

HasVarRepPeriodInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


