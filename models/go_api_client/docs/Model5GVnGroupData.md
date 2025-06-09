# Model5GVnGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AppDescriptors** | Pointer to [**[]AppDescriptor**](AppDescriptor.md) |  | [optional] 
**SecondaryAuth** | Pointer to **bool** |  | [optional] 
**DnAaaIpAddressAllocation** | Pointer to **bool** |  | [optional] 
**DnAaaAddress** | Pointer to [**NullableIpAddress**](IpAddress.md) |  | [optional] 
**AdditionalDnAaaAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**DnAaaFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**Var5gVnGroupCommunicationInd** | Pointer to **bool** |  | [optional] 
**MaxGroupDataRate** | Pointer to [**MaxGroupDataRate**](MaxGroupDataRate.md) |  | [optional] 

## Methods

### NewModel5GVnGroupData

`func NewModel5GVnGroupData(dnn string, sNssai Snssai, ) *Model5GVnGroupData`

NewModel5GVnGroupData instantiates a new Model5GVnGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GVnGroupDataWithDefaults

`func NewModel5GVnGroupDataWithDefaults() *Model5GVnGroupData`

NewModel5GVnGroupDataWithDefaults instantiates a new Model5GVnGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *Model5GVnGroupData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Model5GVnGroupData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Model5GVnGroupData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSNssai

`func (o *Model5GVnGroupData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *Model5GVnGroupData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *Model5GVnGroupData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetPduSessionTypes

`func (o *Model5GVnGroupData) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *Model5GVnGroupData) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *Model5GVnGroupData) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *Model5GVnGroupData) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAppDescriptors

`func (o *Model5GVnGroupData) GetAppDescriptors() []AppDescriptor`

GetAppDescriptors returns the AppDescriptors field if non-nil, zero value otherwise.

### GetAppDescriptorsOk

`func (o *Model5GVnGroupData) GetAppDescriptorsOk() (*[]AppDescriptor, bool)`

GetAppDescriptorsOk returns a tuple with the AppDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptors

`func (o *Model5GVnGroupData) SetAppDescriptors(v []AppDescriptor)`

SetAppDescriptors sets AppDescriptors field to given value.

### HasAppDescriptors

`func (o *Model5GVnGroupData) HasAppDescriptors() bool`

HasAppDescriptors returns a boolean if a field has been set.

### GetSecondaryAuth

`func (o *Model5GVnGroupData) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *Model5GVnGroupData) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *Model5GVnGroupData) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *Model5GVnGroupData) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### GetDnAaaIpAddressAllocation

`func (o *Model5GVnGroupData) GetDnAaaIpAddressAllocation() bool`

GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field if non-nil, zero value otherwise.

### GetDnAaaIpAddressAllocationOk

`func (o *Model5GVnGroupData) GetDnAaaIpAddressAllocationOk() (*bool, bool)`

GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaIpAddressAllocation

`func (o *Model5GVnGroupData) SetDnAaaIpAddressAllocation(v bool)`

SetDnAaaIpAddressAllocation sets DnAaaIpAddressAllocation field to given value.

### HasDnAaaIpAddressAllocation

`func (o *Model5GVnGroupData) HasDnAaaIpAddressAllocation() bool`

HasDnAaaIpAddressAllocation returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *Model5GVnGroupData) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *Model5GVnGroupData) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *Model5GVnGroupData) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *Model5GVnGroupData) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### SetDnAaaAddressNil

`func (o *Model5GVnGroupData) SetDnAaaAddressNil(b bool)`

 SetDnAaaAddressNil sets the value for DnAaaAddress to be an explicit nil

### UnsetDnAaaAddress
`func (o *Model5GVnGroupData) UnsetDnAaaAddress()`

UnsetDnAaaAddress ensures that no value is present for DnAaaAddress, not even an explicit nil
### GetAdditionalDnAaaAddresses

`func (o *Model5GVnGroupData) GetAdditionalDnAaaAddresses() []IpAddress`

GetAdditionalDnAaaAddresses returns the AdditionalDnAaaAddresses field if non-nil, zero value otherwise.

### GetAdditionalDnAaaAddressesOk

`func (o *Model5GVnGroupData) GetAdditionalDnAaaAddressesOk() (*[]IpAddress, bool)`

GetAdditionalDnAaaAddressesOk returns a tuple with the AdditionalDnAaaAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDnAaaAddresses

`func (o *Model5GVnGroupData) SetAdditionalDnAaaAddresses(v []IpAddress)`

SetAdditionalDnAaaAddresses sets AdditionalDnAaaAddresses field to given value.

### HasAdditionalDnAaaAddresses

`func (o *Model5GVnGroupData) HasAdditionalDnAaaAddresses() bool`

HasAdditionalDnAaaAddresses returns a boolean if a field has been set.

### GetDnAaaFqdn

`func (o *Model5GVnGroupData) GetDnAaaFqdn() string`

GetDnAaaFqdn returns the DnAaaFqdn field if non-nil, zero value otherwise.

### GetDnAaaFqdnOk

`func (o *Model5GVnGroupData) GetDnAaaFqdnOk() (*string, bool)`

GetDnAaaFqdnOk returns a tuple with the DnAaaFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaFqdn

`func (o *Model5GVnGroupData) SetDnAaaFqdn(v string)`

SetDnAaaFqdn sets DnAaaFqdn field to given value.

### HasDnAaaFqdn

`func (o *Model5GVnGroupData) HasDnAaaFqdn() bool`

HasDnAaaFqdn returns a boolean if a field has been set.

### GetVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupData) GetVar5gVnGroupCommunicationInd() bool`

GetVar5gVnGroupCommunicationInd returns the Var5gVnGroupCommunicationInd field if non-nil, zero value otherwise.

### GetVar5gVnGroupCommunicationIndOk

`func (o *Model5GVnGroupData) GetVar5gVnGroupCommunicationIndOk() (*bool, bool)`

GetVar5gVnGroupCommunicationIndOk returns a tuple with the Var5gVnGroupCommunicationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupData) SetVar5gVnGroupCommunicationInd(v bool)`

SetVar5gVnGroupCommunicationInd sets Var5gVnGroupCommunicationInd field to given value.

### HasVar5gVnGroupCommunicationInd

`func (o *Model5GVnGroupData) HasVar5gVnGroupCommunicationInd() bool`

HasVar5gVnGroupCommunicationInd returns a boolean if a field has been set.

### GetMaxGroupDataRate

`func (o *Model5GVnGroupData) GetMaxGroupDataRate() MaxGroupDataRate`

GetMaxGroupDataRate returns the MaxGroupDataRate field if non-nil, zero value otherwise.

### GetMaxGroupDataRateOk

`func (o *Model5GVnGroupData) GetMaxGroupDataRateOk() (*MaxGroupDataRate, bool)`

GetMaxGroupDataRateOk returns a tuple with the MaxGroupDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupDataRate

`func (o *Model5GVnGroupData) SetMaxGroupDataRate(v MaxGroupDataRate)`

SetMaxGroupDataRate sets MaxGroupDataRate field to given value.

### HasMaxGroupDataRate

`func (o *Model5GVnGroupData) HasMaxGroupDataRate() bool`

HasMaxGroupDataRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


