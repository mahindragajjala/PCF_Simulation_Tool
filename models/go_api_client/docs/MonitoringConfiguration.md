# MonitoringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**ImmediateFlag** | Pointer to **bool** |  | [optional] 
**LocationReportingConfiguration** | Pointer to [**LocationReportingConfiguration**](LocationReportingConfiguration.md) |  | [optional] 
**AssociationType** | Pointer to [**AssociationType**](AssociationType.md) |  | [optional] 
**DatalinkReportCfg** | Pointer to [**DatalinkReportingConfiguration**](DatalinkReportingConfiguration.md) |  | [optional] 
**LossConnectivityCfg** | Pointer to [**LossConnectivityCfg**](LossConnectivityCfg.md) |  | [optional] 
**MaximumLatency** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SuggestedPacketNumDl** | Pointer to **int32** |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**PduSessionStatusCfg** | Pointer to [**PduSessionStatusCfg**](PduSessionStatusCfg.md) |  | [optional] 
**ReachabilityForSmsCfg** | Pointer to [**ReachabilityForSmsConfiguration**](ReachabilityForSmsConfiguration.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**AfId** | Pointer to **string** |  | [optional] 
**ReachabilityForDataCfg** | Pointer to [**ReachabilityForDataConfiguration**](ReachabilityForDataConfiguration.md) |  | [optional] 
**IdleStatusInd** | Pointer to **bool** |  | [optional] [default to false]
**MonitoringSuspension** | Pointer to [**MonitoringSuspension**](MonitoringSuspension.md) |  | [optional] 

## Methods

### NewMonitoringConfiguration

`func NewMonitoringConfiguration(eventType EventType, ) *MonitoringConfiguration`

NewMonitoringConfiguration instantiates a new MonitoringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigurationWithDefaults

`func NewMonitoringConfigurationWithDefaults() *MonitoringConfiguration`

NewMonitoringConfigurationWithDefaults instantiates a new MonitoringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MonitoringConfiguration) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitoringConfiguration) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitoringConfiguration) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetImmediateFlag

`func (o *MonitoringConfiguration) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *MonitoringConfiguration) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *MonitoringConfiguration) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *MonitoringConfiguration) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.

### GetLocationReportingConfiguration

`func (o *MonitoringConfiguration) GetLocationReportingConfiguration() LocationReportingConfiguration`

GetLocationReportingConfiguration returns the LocationReportingConfiguration field if non-nil, zero value otherwise.

### GetLocationReportingConfigurationOk

`func (o *MonitoringConfiguration) GetLocationReportingConfigurationOk() (*LocationReportingConfiguration, bool)`

GetLocationReportingConfigurationOk returns a tuple with the LocationReportingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingConfiguration

`func (o *MonitoringConfiguration) SetLocationReportingConfiguration(v LocationReportingConfiguration)`

SetLocationReportingConfiguration sets LocationReportingConfiguration field to given value.

### HasLocationReportingConfiguration

`func (o *MonitoringConfiguration) HasLocationReportingConfiguration() bool`

HasLocationReportingConfiguration returns a boolean if a field has been set.

### GetAssociationType

`func (o *MonitoringConfiguration) GetAssociationType() AssociationType`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *MonitoringConfiguration) GetAssociationTypeOk() (*AssociationType, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *MonitoringConfiguration) SetAssociationType(v AssociationType)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *MonitoringConfiguration) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetDatalinkReportCfg

`func (o *MonitoringConfiguration) GetDatalinkReportCfg() DatalinkReportingConfiguration`

GetDatalinkReportCfg returns the DatalinkReportCfg field if non-nil, zero value otherwise.

### GetDatalinkReportCfgOk

`func (o *MonitoringConfiguration) GetDatalinkReportCfgOk() (*DatalinkReportingConfiguration, bool)`

GetDatalinkReportCfgOk returns a tuple with the DatalinkReportCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalinkReportCfg

`func (o *MonitoringConfiguration) SetDatalinkReportCfg(v DatalinkReportingConfiguration)`

SetDatalinkReportCfg sets DatalinkReportCfg field to given value.

### HasDatalinkReportCfg

`func (o *MonitoringConfiguration) HasDatalinkReportCfg() bool`

HasDatalinkReportCfg returns a boolean if a field has been set.

### GetLossConnectivityCfg

`func (o *MonitoringConfiguration) GetLossConnectivityCfg() LossConnectivityCfg`

GetLossConnectivityCfg returns the LossConnectivityCfg field if non-nil, zero value otherwise.

### GetLossConnectivityCfgOk

`func (o *MonitoringConfiguration) GetLossConnectivityCfgOk() (*LossConnectivityCfg, bool)`

GetLossConnectivityCfgOk returns a tuple with the LossConnectivityCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossConnectivityCfg

`func (o *MonitoringConfiguration) SetLossConnectivityCfg(v LossConnectivityCfg)`

SetLossConnectivityCfg sets LossConnectivityCfg field to given value.

### HasLossConnectivityCfg

`func (o *MonitoringConfiguration) HasLossConnectivityCfg() bool`

HasLossConnectivityCfg returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *MonitoringConfiguration) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *MonitoringConfiguration) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *MonitoringConfiguration) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *MonitoringConfiguration) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *MonitoringConfiguration) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *MonitoringConfiguration) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *MonitoringConfiguration) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *MonitoringConfiguration) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetSuggestedPacketNumDl

`func (o *MonitoringConfiguration) GetSuggestedPacketNumDl() int32`

GetSuggestedPacketNumDl returns the SuggestedPacketNumDl field if non-nil, zero value otherwise.

### GetSuggestedPacketNumDlOk

`func (o *MonitoringConfiguration) GetSuggestedPacketNumDlOk() (*int32, bool)`

GetSuggestedPacketNumDlOk returns a tuple with the SuggestedPacketNumDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPacketNumDl

`func (o *MonitoringConfiguration) SetSuggestedPacketNumDl(v int32)`

SetSuggestedPacketNumDl sets SuggestedPacketNumDl field to given value.

### HasSuggestedPacketNumDl

`func (o *MonitoringConfiguration) HasSuggestedPacketNumDl() bool`

HasSuggestedPacketNumDl returns a boolean if a field has been set.

### GetDnn

`func (o *MonitoringConfiguration) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *MonitoringConfiguration) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *MonitoringConfiguration) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *MonitoringConfiguration) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSingleNssai

`func (o *MonitoringConfiguration) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *MonitoringConfiguration) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *MonitoringConfiguration) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *MonitoringConfiguration) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetAppId

`func (o *MonitoringConfiguration) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MonitoringConfiguration) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MonitoringConfiguration) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MonitoringConfiguration) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetPduSessionStatusCfg

`func (o *MonitoringConfiguration) GetPduSessionStatusCfg() PduSessionStatusCfg`

GetPduSessionStatusCfg returns the PduSessionStatusCfg field if non-nil, zero value otherwise.

### GetPduSessionStatusCfgOk

`func (o *MonitoringConfiguration) GetPduSessionStatusCfgOk() (*PduSessionStatusCfg, bool)`

GetPduSessionStatusCfgOk returns a tuple with the PduSessionStatusCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionStatusCfg

`func (o *MonitoringConfiguration) SetPduSessionStatusCfg(v PduSessionStatusCfg)`

SetPduSessionStatusCfg sets PduSessionStatusCfg field to given value.

### HasPduSessionStatusCfg

`func (o *MonitoringConfiguration) HasPduSessionStatusCfg() bool`

HasPduSessionStatusCfg returns a boolean if a field has been set.

### GetReachabilityForSmsCfg

`func (o *MonitoringConfiguration) GetReachabilityForSmsCfg() ReachabilityForSmsConfiguration`

GetReachabilityForSmsCfg returns the ReachabilityForSmsCfg field if non-nil, zero value otherwise.

### GetReachabilityForSmsCfgOk

`func (o *MonitoringConfiguration) GetReachabilityForSmsCfgOk() (*ReachabilityForSmsConfiguration, bool)`

GetReachabilityForSmsCfgOk returns a tuple with the ReachabilityForSmsCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForSmsCfg

`func (o *MonitoringConfiguration) SetReachabilityForSmsCfg(v ReachabilityForSmsConfiguration)`

SetReachabilityForSmsCfg sets ReachabilityForSmsCfg field to given value.

### HasReachabilityForSmsCfg

`func (o *MonitoringConfiguration) HasReachabilityForSmsCfg() bool`

HasReachabilityForSmsCfg returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *MonitoringConfiguration) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *MonitoringConfiguration) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *MonitoringConfiguration) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *MonitoringConfiguration) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetAfId

`func (o *MonitoringConfiguration) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *MonitoringConfiguration) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *MonitoringConfiguration) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *MonitoringConfiguration) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetReachabilityForDataCfg

`func (o *MonitoringConfiguration) GetReachabilityForDataCfg() ReachabilityForDataConfiguration`

GetReachabilityForDataCfg returns the ReachabilityForDataCfg field if non-nil, zero value otherwise.

### GetReachabilityForDataCfgOk

`func (o *MonitoringConfiguration) GetReachabilityForDataCfgOk() (*ReachabilityForDataConfiguration, bool)`

GetReachabilityForDataCfgOk returns a tuple with the ReachabilityForDataCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForDataCfg

`func (o *MonitoringConfiguration) SetReachabilityForDataCfg(v ReachabilityForDataConfiguration)`

SetReachabilityForDataCfg sets ReachabilityForDataCfg field to given value.

### HasReachabilityForDataCfg

`func (o *MonitoringConfiguration) HasReachabilityForDataCfg() bool`

HasReachabilityForDataCfg returns a boolean if a field has been set.

### GetIdleStatusInd

`func (o *MonitoringConfiguration) GetIdleStatusInd() bool`

GetIdleStatusInd returns the IdleStatusInd field if non-nil, zero value otherwise.

### GetIdleStatusIndOk

`func (o *MonitoringConfiguration) GetIdleStatusIndOk() (*bool, bool)`

GetIdleStatusIndOk returns a tuple with the IdleStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInd

`func (o *MonitoringConfiguration) SetIdleStatusInd(v bool)`

SetIdleStatusInd sets IdleStatusInd field to given value.

### HasIdleStatusInd

`func (o *MonitoringConfiguration) HasIdleStatusInd() bool`

HasIdleStatusInd returns a boolean if a field has been set.

### GetMonitoringSuspension

`func (o *MonitoringConfiguration) GetMonitoringSuspension() MonitoringSuspension`

GetMonitoringSuspension returns the MonitoringSuspension field if non-nil, zero value otherwise.

### GetMonitoringSuspensionOk

`func (o *MonitoringConfiguration) GetMonitoringSuspensionOk() (*MonitoringSuspension, bool)`

GetMonitoringSuspensionOk returns a tuple with the MonitoringSuspension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringSuspension

`func (o *MonitoringConfiguration) SetMonitoringSuspension(v MonitoringSuspension)`

SetMonitoringSuspension sets MonitoringSuspension field to given value.

### HasMonitoringSuspension

`func (o *MonitoringConfiguration) HasMonitoringSuspension() bool`

HasMonitoringSuspension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


