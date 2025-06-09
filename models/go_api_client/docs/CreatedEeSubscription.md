# CreatedEeSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 
**NumberOfUes** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**EventReports** | Pointer to [**[]MonitoringReport**](MonitoringReport.md) |  | [optional] 
**EpcStatusInd** | Pointer to **bool** |  | [optional] 
**Var5gOnlyInd** | Pointer to **bool** |  | [optional] 
**FailedMonitoringConfigs** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration  | [optional] 
**FailedMoniConfigsEPC** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration, the key value \&quot;ALL\&quot; may  be used to identify a map entry which contains the failed cause of the EE subscription  was not successful in EPC domain.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**CurrentStatusNotAvailableList** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 

## Methods

### NewCreatedEeSubscription

`func NewCreatedEeSubscription(eeSubscription EeSubscription, ) *CreatedEeSubscription`

NewCreatedEeSubscription instantiates a new CreatedEeSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedEeSubscriptionWithDefaults

`func NewCreatedEeSubscriptionWithDefaults() *CreatedEeSubscription`

NewCreatedEeSubscriptionWithDefaults instantiates a new CreatedEeSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEeSubscription

`func (o *CreatedEeSubscription) GetEeSubscription() EeSubscription`

GetEeSubscription returns the EeSubscription field if non-nil, zero value otherwise.

### GetEeSubscriptionOk

`func (o *CreatedEeSubscription) GetEeSubscriptionOk() (*EeSubscription, bool)`

GetEeSubscriptionOk returns a tuple with the EeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeSubscription

`func (o *CreatedEeSubscription) SetEeSubscription(v EeSubscription)`

SetEeSubscription sets EeSubscription field to given value.


### GetNumberOfUes

`func (o *CreatedEeSubscription) GetNumberOfUes() int32`

GetNumberOfUes returns the NumberOfUes field if non-nil, zero value otherwise.

### GetNumberOfUesOk

`func (o *CreatedEeSubscription) GetNumberOfUesOk() (*int32, bool)`

GetNumberOfUesOk returns a tuple with the NumberOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUes

`func (o *CreatedEeSubscription) SetNumberOfUes(v int32)`

SetNumberOfUes sets NumberOfUes field to given value.

### HasNumberOfUes

`func (o *CreatedEeSubscription) HasNumberOfUes() bool`

HasNumberOfUes returns a boolean if a field has been set.

### GetEventReports

`func (o *CreatedEeSubscription) GetEventReports() []MonitoringReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *CreatedEeSubscription) GetEventReportsOk() (*[]MonitoringReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *CreatedEeSubscription) SetEventReports(v []MonitoringReport)`

SetEventReports sets EventReports field to given value.

### HasEventReports

`func (o *CreatedEeSubscription) HasEventReports() bool`

HasEventReports returns a boolean if a field has been set.

### GetEpcStatusInd

`func (o *CreatedEeSubscription) GetEpcStatusInd() bool`

GetEpcStatusInd returns the EpcStatusInd field if non-nil, zero value otherwise.

### GetEpcStatusIndOk

`func (o *CreatedEeSubscription) GetEpcStatusIndOk() (*bool, bool)`

GetEpcStatusIndOk returns a tuple with the EpcStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpcStatusInd

`func (o *CreatedEeSubscription) SetEpcStatusInd(v bool)`

SetEpcStatusInd sets EpcStatusInd field to given value.

### HasEpcStatusInd

`func (o *CreatedEeSubscription) HasEpcStatusInd() bool`

HasEpcStatusInd returns a boolean if a field has been set.

### GetVar5gOnlyInd

`func (o *CreatedEeSubscription) GetVar5gOnlyInd() bool`

GetVar5gOnlyInd returns the Var5gOnlyInd field if non-nil, zero value otherwise.

### GetVar5gOnlyIndOk

`func (o *CreatedEeSubscription) GetVar5gOnlyIndOk() (*bool, bool)`

GetVar5gOnlyIndOk returns a tuple with the Var5gOnlyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gOnlyInd

`func (o *CreatedEeSubscription) SetVar5gOnlyInd(v bool)`

SetVar5gOnlyInd sets Var5gOnlyInd field to given value.

### HasVar5gOnlyInd

`func (o *CreatedEeSubscription) HasVar5gOnlyInd() bool`

HasVar5gOnlyInd returns a boolean if a field has been set.

### GetFailedMonitoringConfigs

`func (o *CreatedEeSubscription) GetFailedMonitoringConfigs() map[string]FailedMonitoringConfiguration`

GetFailedMonitoringConfigs returns the FailedMonitoringConfigs field if non-nil, zero value otherwise.

### GetFailedMonitoringConfigsOk

`func (o *CreatedEeSubscription) GetFailedMonitoringConfigsOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMonitoringConfigsOk returns a tuple with the FailedMonitoringConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMonitoringConfigs

`func (o *CreatedEeSubscription) SetFailedMonitoringConfigs(v map[string]FailedMonitoringConfiguration)`

SetFailedMonitoringConfigs sets FailedMonitoringConfigs field to given value.

### HasFailedMonitoringConfigs

`func (o *CreatedEeSubscription) HasFailedMonitoringConfigs() bool`

HasFailedMonitoringConfigs returns a boolean if a field has been set.

### GetFailedMoniConfigsEPC

`func (o *CreatedEeSubscription) GetFailedMoniConfigsEPC() map[string]FailedMonitoringConfiguration`

GetFailedMoniConfigsEPC returns the FailedMoniConfigsEPC field if non-nil, zero value otherwise.

### GetFailedMoniConfigsEPCOk

`func (o *CreatedEeSubscription) GetFailedMoniConfigsEPCOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMoniConfigsEPCOk returns a tuple with the FailedMoniConfigsEPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMoniConfigsEPC

`func (o *CreatedEeSubscription) SetFailedMoniConfigsEPC(v map[string]FailedMonitoringConfiguration)`

SetFailedMoniConfigsEPC sets FailedMoniConfigsEPC field to given value.

### HasFailedMoniConfigsEPC

`func (o *CreatedEeSubscription) HasFailedMoniConfigsEPC() bool`

HasFailedMoniConfigsEPC returns a boolean if a field has been set.

### GetResetIds

`func (o *CreatedEeSubscription) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *CreatedEeSubscription) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *CreatedEeSubscription) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *CreatedEeSubscription) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetCurrentStatusNotAvailableList

`func (o *CreatedEeSubscription) GetCurrentStatusNotAvailableList() []EventType`

GetCurrentStatusNotAvailableList returns the CurrentStatusNotAvailableList field if non-nil, zero value otherwise.

### GetCurrentStatusNotAvailableListOk

`func (o *CreatedEeSubscription) GetCurrentStatusNotAvailableListOk() (*[]EventType, bool)`

GetCurrentStatusNotAvailableListOk returns a tuple with the CurrentStatusNotAvailableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatusNotAvailableList

`func (o *CreatedEeSubscription) SetCurrentStatusNotAvailableList(v []EventType)`

SetCurrentStatusNotAvailableList sets CurrentStatusNotAvailableList field to given value.

### HasCurrentStatusNotAvailableList

`func (o *CreatedEeSubscription) HasCurrentStatusNotAvailableList() bool`

HasCurrentStatusNotAvailableList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


