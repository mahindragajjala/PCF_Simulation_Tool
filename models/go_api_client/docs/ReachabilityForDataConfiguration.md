# ReachabilityForDataConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportCfg** | [**ReachabilityForDataReportConfig**](ReachabilityForDataReportConfig.md) |  | 
**MinInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewReachabilityForDataConfiguration

`func NewReachabilityForDataConfiguration(reportCfg ReachabilityForDataReportConfig, ) *ReachabilityForDataConfiguration`

NewReachabilityForDataConfiguration instantiates a new ReachabilityForDataConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityForDataConfigurationWithDefaults

`func NewReachabilityForDataConfigurationWithDefaults() *ReachabilityForDataConfiguration`

NewReachabilityForDataConfigurationWithDefaults instantiates a new ReachabilityForDataConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportCfg

`func (o *ReachabilityForDataConfiguration) GetReportCfg() ReachabilityForDataReportConfig`

GetReportCfg returns the ReportCfg field if non-nil, zero value otherwise.

### GetReportCfgOk

`func (o *ReachabilityForDataConfiguration) GetReportCfgOk() (*ReachabilityForDataReportConfig, bool)`

GetReportCfgOk returns a tuple with the ReportCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCfg

`func (o *ReachabilityForDataConfiguration) SetReportCfg(v ReachabilityForDataReportConfig)`

SetReportCfg sets ReportCfg field to given value.


### GetMinInterval

`func (o *ReachabilityForDataConfiguration) GetMinInterval() int32`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *ReachabilityForDataConfiguration) GetMinIntervalOk() (*int32, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *ReachabilityForDataConfiguration) SetMinInterval(v int32)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *ReachabilityForDataConfiguration) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


