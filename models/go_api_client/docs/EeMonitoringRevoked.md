# EeMonitoringRevoked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokedMonitoringEventList** | [**map[string]MonitoringEvent**](MonitoringEvent.md) | A map (list of key-value pairs where ReferenceId serves as key) of MonitoringEvents | 
**RemovedGpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExcludeGpsiList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEeMonitoringRevoked

`func NewEeMonitoringRevoked(revokedMonitoringEventList map[string]MonitoringEvent, ) *EeMonitoringRevoked`

NewEeMonitoringRevoked instantiates a new EeMonitoringRevoked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeMonitoringRevokedWithDefaults

`func NewEeMonitoringRevokedWithDefaults() *EeMonitoringRevoked`

NewEeMonitoringRevokedWithDefaults instantiates a new EeMonitoringRevoked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokedMonitoringEventList

`func (o *EeMonitoringRevoked) GetRevokedMonitoringEventList() map[string]MonitoringEvent`

GetRevokedMonitoringEventList returns the RevokedMonitoringEventList field if non-nil, zero value otherwise.

### GetRevokedMonitoringEventListOk

`func (o *EeMonitoringRevoked) GetRevokedMonitoringEventListOk() (*map[string]MonitoringEvent, bool)`

GetRevokedMonitoringEventListOk returns a tuple with the RevokedMonitoringEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedMonitoringEventList

`func (o *EeMonitoringRevoked) SetRevokedMonitoringEventList(v map[string]MonitoringEvent)`

SetRevokedMonitoringEventList sets RevokedMonitoringEventList field to given value.


### GetRemovedGpsi

`func (o *EeMonitoringRevoked) GetRemovedGpsi() string`

GetRemovedGpsi returns the RemovedGpsi field if non-nil, zero value otherwise.

### GetRemovedGpsiOk

`func (o *EeMonitoringRevoked) GetRemovedGpsiOk() (*string, bool)`

GetRemovedGpsiOk returns a tuple with the RemovedGpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedGpsi

`func (o *EeMonitoringRevoked) SetRemovedGpsi(v string)`

SetRemovedGpsi sets RemovedGpsi field to given value.

### HasRemovedGpsi

`func (o *EeMonitoringRevoked) HasRemovedGpsi() bool`

HasRemovedGpsi returns a boolean if a field has been set.

### GetExcludeGpsiList

`func (o *EeMonitoringRevoked) GetExcludeGpsiList() []string`

GetExcludeGpsiList returns the ExcludeGpsiList field if non-nil, zero value otherwise.

### GetExcludeGpsiListOk

`func (o *EeMonitoringRevoked) GetExcludeGpsiListOk() (*[]string, bool)`

GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGpsiList

`func (o *EeMonitoringRevoked) SetExcludeGpsiList(v []string)`

SetExcludeGpsiList sets ExcludeGpsiList field to given value.

### HasExcludeGpsiList

`func (o *EeMonitoringRevoked) HasExcludeGpsiList() bool`

HasExcludeGpsiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


