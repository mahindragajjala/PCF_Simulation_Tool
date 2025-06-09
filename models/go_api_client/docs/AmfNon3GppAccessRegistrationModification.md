# AmfNon3GppAccessRegistrationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guami** | [**Guami**](Guami.md) |  | 
**PurgeFlag** | Pointer to **bool** | This flag indicates whether an AMF is deregistered. This attribute may be included in notifications sent by the UDR to the UDM if purgeFlag is also set to true in the same notification.  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**ImsVoPs** | Pointer to [**ImsVoPs**](ImsVoPs.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 

## Methods

### NewAmfNon3GppAccessRegistrationModification

`func NewAmfNon3GppAccessRegistrationModification(guami Guami, ) *AmfNon3GppAccessRegistrationModification`

NewAmfNon3GppAccessRegistrationModification instantiates a new AmfNon3GppAccessRegistrationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfNon3GppAccessRegistrationModificationWithDefaults

`func NewAmfNon3GppAccessRegistrationModificationWithDefaults() *AmfNon3GppAccessRegistrationModification`

NewAmfNon3GppAccessRegistrationModificationWithDefaults instantiates a new AmfNon3GppAccessRegistrationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuami

`func (o *AmfNon3GppAccessRegistrationModification) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *AmfNon3GppAccessRegistrationModification) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *AmfNon3GppAccessRegistrationModification) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetPurgeFlag

`func (o *AmfNon3GppAccessRegistrationModification) GetPurgeFlag() bool`

GetPurgeFlag returns the PurgeFlag field if non-nil, zero value otherwise.

### GetPurgeFlagOk

`func (o *AmfNon3GppAccessRegistrationModification) GetPurgeFlagOk() (*bool, bool)`

GetPurgeFlagOk returns a tuple with the PurgeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeFlag

`func (o *AmfNon3GppAccessRegistrationModification) SetPurgeFlag(v bool)`

SetPurgeFlag sets PurgeFlag field to given value.

### HasPurgeFlag

`func (o *AmfNon3GppAccessRegistrationModification) HasPurgeFlag() bool`

HasPurgeFlag returns a boolean if a field has been set.

### GetPei

`func (o *AmfNon3GppAccessRegistrationModification) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AmfNon3GppAccessRegistrationModification) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AmfNon3GppAccessRegistrationModification) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AmfNon3GppAccessRegistrationModification) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetImsVoPs

`func (o *AmfNon3GppAccessRegistrationModification) GetImsVoPs() ImsVoPs`

GetImsVoPs returns the ImsVoPs field if non-nil, zero value otherwise.

### GetImsVoPsOk

`func (o *AmfNon3GppAccessRegistrationModification) GetImsVoPsOk() (*ImsVoPs, bool)`

GetImsVoPsOk returns a tuple with the ImsVoPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVoPs

`func (o *AmfNon3GppAccessRegistrationModification) SetImsVoPs(v ImsVoPs)`

SetImsVoPs sets ImsVoPs field to given value.

### HasImsVoPs

`func (o *AmfNon3GppAccessRegistrationModification) HasImsVoPs() bool`

HasImsVoPs returns a boolean if a field has been set.

### GetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistrationModification) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *AmfNon3GppAccessRegistrationModification) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistrationModification) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *AmfNon3GppAccessRegistrationModification) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


