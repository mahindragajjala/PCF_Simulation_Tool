# EeSubscriptionErrorAddInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubType** | Pointer to [**SubscriptionType**](SubscriptionType.md) |  | [optional] 
**FailedMonitoringConfigs** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration  | [optional] 
**FailedMoniConfigsEPC** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration, the key value \&quot;ALL\&quot; may  be used to identify a map entry which contains the failed cause of the EE subscription  was not successful in EPC domain.  | [optional] 

## Methods

### NewEeSubscriptionErrorAddInfo

`func NewEeSubscriptionErrorAddInfo() *EeSubscriptionErrorAddInfo`

NewEeSubscriptionErrorAddInfo instantiates a new EeSubscriptionErrorAddInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeSubscriptionErrorAddInfoWithDefaults

`func NewEeSubscriptionErrorAddInfoWithDefaults() *EeSubscriptionErrorAddInfo`

NewEeSubscriptionErrorAddInfoWithDefaults instantiates a new EeSubscriptionErrorAddInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubType

`func (o *EeSubscriptionErrorAddInfo) GetSubType() SubscriptionType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *EeSubscriptionErrorAddInfo) GetSubTypeOk() (*SubscriptionType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *EeSubscriptionErrorAddInfo) SetSubType(v SubscriptionType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *EeSubscriptionErrorAddInfo) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetFailedMonitoringConfigs

`func (o *EeSubscriptionErrorAddInfo) GetFailedMonitoringConfigs() map[string]FailedMonitoringConfiguration`

GetFailedMonitoringConfigs returns the FailedMonitoringConfigs field if non-nil, zero value otherwise.

### GetFailedMonitoringConfigsOk

`func (o *EeSubscriptionErrorAddInfo) GetFailedMonitoringConfigsOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMonitoringConfigsOk returns a tuple with the FailedMonitoringConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMonitoringConfigs

`func (o *EeSubscriptionErrorAddInfo) SetFailedMonitoringConfigs(v map[string]FailedMonitoringConfiguration)`

SetFailedMonitoringConfigs sets FailedMonitoringConfigs field to given value.

### HasFailedMonitoringConfigs

`func (o *EeSubscriptionErrorAddInfo) HasFailedMonitoringConfigs() bool`

HasFailedMonitoringConfigs returns a boolean if a field has been set.

### GetFailedMoniConfigsEPC

`func (o *EeSubscriptionErrorAddInfo) GetFailedMoniConfigsEPC() map[string]FailedMonitoringConfiguration`

GetFailedMoniConfigsEPC returns the FailedMoniConfigsEPC field if non-nil, zero value otherwise.

### GetFailedMoniConfigsEPCOk

`func (o *EeSubscriptionErrorAddInfo) GetFailedMoniConfigsEPCOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMoniConfigsEPCOk returns a tuple with the FailedMoniConfigsEPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMoniConfigsEPC

`func (o *EeSubscriptionErrorAddInfo) SetFailedMoniConfigsEPC(v map[string]FailedMonitoringConfiguration)`

SetFailedMoniConfigsEPC sets FailedMoniConfigsEPC field to given value.

### HasFailedMoniConfigsEPC

`func (o *EeSubscriptionErrorAddInfo) HasFailedMoniConfigsEPC() bool`

HasFailedMoniConfigsEPC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


