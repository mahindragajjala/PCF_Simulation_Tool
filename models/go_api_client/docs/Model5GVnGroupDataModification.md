# Model5GVnGroupDataModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDescriptors** | Pointer to [**[]AppDescriptor**](AppDescriptor.md) |  | [optional] 
**SecondaryAuth** | Pointer to **NullableBool** |  | [optional] 
**DnAaaIpAddressAllocation** | Pointer to **NullableBool** |  | [optional] 
**DnAaaAddress** | Pointer to [**NullableIpAddress**](IpAddress.md) |  | [optional] 
**AdditionalDnAaaAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**DnAaaFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**Var5gVnGroupCommunicationInd** | Pointer to **NullableBool** |  | [optional] 
**MaxGroupDataRate** | Pointer to [**NullableMaxGroupDataRate**](MaxGroupDataRate.md) |  | [optional] 

## Methods

### NewModel5GVnGroupDataModification

`func NewModel5GVnGroupDataModification() *Model5GVnGroupDataModification`

NewModel5GVnGroupDataModification instantiates a new Model5GVnGroupDataModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GVnGroupDataModificationWithDefaults

`func NewModel5GVnGroupDataModificationWithDefaults() *Model5GVnGroupDataModification`

NewModel5GVnGroupDataModificationWithDefaults instantiates a new Model5GVnGroupDataModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDescriptors

`func (o *Model5GVnGroupDataModification) GetAppDescriptors() []AppDescriptor`

GetAppDescriptors returns the AppDescriptors field if non-nil, zero value otherwise.

### GetAppDescriptorsOk

`func (o *Model5GVnGroupDataModification) GetAppDescriptorsOk() (*[]AppDescriptor, bool)`

GetAppDescriptorsOk returns a tuple with the AppDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptors

`func (o *Model5GVnGroupDataModification) SetAppDescriptors(v []AppDescriptor)`

SetAppDescriptors sets AppDescriptors field to given value.

### HasAppDescriptors

`func (o *Model5GVnGroupDataModification) HasAppDescriptors() bool`

HasAppDescriptors returns a boolean if a field has been set.

### SetAppDescriptorsNil

`func (o *Model5GVnGroupDataModification) SetAppDescriptorsNil(b bool)`

 SetAppDescriptorsNil sets the value for AppDescriptors to be an explicit nil

### UnsetAppDescriptors
`func (o *Model5GVnGroupDataModification) UnsetAppDescriptors()`

UnsetAppDescriptors ensures that no value is present for AppDescriptors, not even an explicit nil
### GetSecondaryAuth

`func (o *Model5GVnGroupDataModification) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *Model5GVnGroupDataModification) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *Model5GVnGroupDataModification) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *Model5GVnGroupDataModification) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### SetSecondaryAuthNil

`func (o *Model5GVnGroupDataModification) SetSecondaryAuthNil(b bool)`

 SetSecondaryAuthNil sets the value for SecondaryAuth to be an explicit nil

### UnsetSecondaryAuth
`func (o *Model5GVnGroupDataModification) UnsetSecondaryAuth()`

UnsetSecondaryAuth ensures that no value is present for SecondaryAuth, not even an explicit nil
### GetDnAaaIpAddressAllocation

`func (o *Model5GVnGroupDataModification) GetDnAaaIpAddressAllocation() bool`

GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field if non-nil, zero value otherwise.

### GetDnAaaIpAddressAllocationOk

`func (o *Model5GVnGroupDataModification) GetDnAaaIpAddressAllocationOk() (*bool, bool)`

GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaIpAddressAllocation

`func (o *Model5GVnGroupDataModification) SetDnAaaIpAddressAllocation(v bool)`

SetDnAaaIpAddressAllocation sets DnAaaIpAddressAllocation field to given value.

### HasDnAaaIpAddressAllocation

`func (o *Model5GVnGroupDataModification) HasDnAaaIpAddressAllocation() bool`

HasDnAaaIpAddressAllocation returns a boolean if a field has been set.

### SetDnAaaIpAddressAllocationNil

`func (o *Model5GVnGroupDataModification) SetDnAaaIpAddressAllocationNil(b bool)`

 SetDnAaaIpAddressAllocationNil sets the value for DnAaaIpAddressAllocation to be an explicit nil

### UnsetDnAaaIpAddressAllocation
`func (o *Model5GVnGroupDataModification) UnsetDnAaaIpAddressAllocation()`

UnsetDnAaaIpAddressAllocation ensures that no value is present for DnAaaIpAddressAllocation, not even an explicit nil
### GetDnAaaAddress

`func (o *Model5GVnGroupDataModification) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *Model5GVnGroupDataModification) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *Model5GVnGroupDataModification) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *Model5GVnGroupDataModification) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### SetDnAaaAddressNil

`func (o *Model5GVnGroupDataModification) SetDnAaaAddressNil(b bool)`

 SetDnAaaAddressNil sets the value for DnAaaAddress to be an explicit nil

### UnsetDnAaaAddress
`func (o *Model5GVnGroupDataModification) UnsetDnAaaAddress()`

UnsetDnAaaAddress ensures that no value is present for DnAaaAddress, not even an explicit nil
### GetAdditionalDnAaaAddresses

`func (o *Model5GVnGroupDataModification) GetAdditionalDnAaaAddresses() []IpAddress`

GetAdditionalDnAaaAddresses returns the AdditionalDnAaaAddresses field if non-nil, zero value otherwise.

### GetAdditionalDnAaaAddressesOk

`func (o *Model5GVnGroupDataModification) GetAdditionalDnAaaAddressesOk() (*[]IpAddress, bool)`

GetAdditionalDnAaaAddressesOk returns a tuple with the AdditionalDnAaaAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDnAaaAddresses

`func (o *Model5GVnGroupDataModification) SetAdditionalDnAaaAddresses(v []IpAddress)`

SetAdditionalDnAaaAddresses sets AdditionalDnAaaAddresses field to given value.

### HasAdditionalDnAaaAddresses

`func (o *Model5GVnGroupDataModification) HasAdditionalDnAaaAddresses() bool`

HasAdditionalDnAaaAddresses returns a boolean if a field has been set.

### SetAdditionalDnAaaAddressesNil

`func (o *Model5GVnGroupDataModification) SetAdditionalDnAaaAddressesNil(b bool)`

 SetAdditionalDnAaaAddressesNil sets the value for AdditionalDnAaaAddresses to be an explicit nil

### UnsetAdditionalDnAaaAddresses
`func (o *Model5GVnGroupDataModification) UnsetAdditionalDnAaaAddresses()`

UnsetAdditionalDnAaaAddresses ensures that no value is present for AdditionalDnAaaAddresses, not even an explicit nil
### GetDnAaaFqdn

`func (o *Model5GVnGroupDataModification) GetDnAaaFqdn() string`

GetDnAaaFqdn returns the DnAaaFqdn field if non-nil, zero value otherwise.

### GetDnAaaFqdnOk

`func (o *Model5GVnGroupDataModification) GetDnAaaFqdnOk() (*string, bool)`

GetDnAaaFqdnOk returns a tuple with the DnAaaFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaFqdn

`func (o *Model5GVnGroupDataModification) SetDnAaaFqdn(v string)`

SetDnAaaFqdn sets DnAaaFqdn field to given value.

### HasDnAaaFqdn

`func (o *Model5GVnGroupDataModification) HasDnAaaFqdn() bool`

HasDnAaaFqdn returns a boolean if a field has been set.

### GetVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupDataModification) GetVar5gVnGroupCommunicationInd() bool`

GetVar5gVnGroupCommunicationInd returns the Var5gVnGroupCommunicationInd field if non-nil, zero value otherwise.

### GetVar5gVnGroupCommunicationIndOk

`func (o *Model5GVnGroupDataModification) GetVar5gVnGroupCommunicationIndOk() (*bool, bool)`

GetVar5gVnGroupCommunicationIndOk returns a tuple with the Var5gVnGroupCommunicationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupDataModification) SetVar5gVnGroupCommunicationInd(v bool)`

SetVar5gVnGroupCommunicationInd sets Var5gVnGroupCommunicationInd field to given value.

### HasVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupDataModification) HasVar5gVnGroupCommunicationInd() bool`

HasVar5gVnGroupCommunicationInd returns a boolean if a field has been set.

### SetVar5gVnGroupCommunicationIndNil

`func (o *Model5GVnGroupDataModification) SetVar5gVnGroupCommunicationIndNil(b bool)`

 SetVar5gVnGroupCommunicationIndNil sets the value for Var5gVnGroupCommunicationInd to be an explicit nil

### UnsetVar5gVnGroupCommunicationInd
`func (o *Model5GVnGroupDataModification) UnsetVar5gVnGroupCommunicationInd()`

UnsetVar5gVnGroupCommunicationInd ensures that no value is present for Var5gVnGroupCommunicationInd, not even an explicit nil
### GetMaxGroupDataRate

`func (o *Model5GVnGroupDataModification) GetMaxGroupDataRate() MaxGroupDataRate`

GetMaxGroupDataRate returns the MaxGroupDataRate field if non-nil, zero value otherwise.

### GetMaxGroupDataRateOk

`func (o *Model5GVnGroupDataModification) GetMaxGroupDataRateOk() (*MaxGroupDataRate, bool)`

GetMaxGroupDataRateOk returns a tuple with the MaxGroupDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupDataRate

`func (o *Model5GVnGroupDataModification) SetMaxGroupDataRate(v MaxGroupDataRate)`

SetMaxGroupDataRate sets MaxGroupDataRate field to given value.

### HasMaxGroupDataRate

`func (o *Model5GVnGroupDataModification) HasMaxGroupDataRate() bool`

HasMaxGroupDataRate returns a boolean if a field has been set.

### SetMaxGroupDataRateNil

`func (o *Model5GVnGroupDataModification) SetMaxGroupDataRateNil(b bool)`

 SetMaxGroupDataRateNil sets the value for MaxGroupDataRate to be an explicit nil

### UnsetMaxGroupDataRate
`func (o *Model5GVnGroupDataModification) UnsetMaxGroupDataRate()`

UnsetMaxGroupDataRate ensures that no value is present for MaxGroupDataRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


