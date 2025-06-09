# IpSmGwRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSmGwMapAddress** | Pointer to **string** | This data type mentions International E.164 number of the SMSF; shall be present if the SMSF supports MAP.  | [optional] 
**IpSmGwDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**IpsmgwIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**IpsmgwIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**IpsmgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**NfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**UnriIndicator** | Pointer to **bool** |  | [optional] [default to false]
**ResetIds** | Pointer to **[]string** |  | [optional] 
**IpSmGwSbiSupInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIpSmGwRegistration

`func NewIpSmGwRegistration() *IpSmGwRegistration`

NewIpSmGwRegistration instantiates a new IpSmGwRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSmGwRegistrationWithDefaults

`func NewIpSmGwRegistrationWithDefaults() *IpSmGwRegistration`

NewIpSmGwRegistrationWithDefaults instantiates a new IpSmGwRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSmGwMapAddress

`func (o *IpSmGwRegistration) GetIpSmGwMapAddress() string`

GetIpSmGwMapAddress returns the IpSmGwMapAddress field if non-nil, zero value otherwise.

### GetIpSmGwMapAddressOk

`func (o *IpSmGwRegistration) GetIpSmGwMapAddressOk() (*string, bool)`

GetIpSmGwMapAddressOk returns a tuple with the IpSmGwMapAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwMapAddress

`func (o *IpSmGwRegistration) SetIpSmGwMapAddress(v string)`

SetIpSmGwMapAddress sets IpSmGwMapAddress field to given value.

### HasIpSmGwMapAddress

`func (o *IpSmGwRegistration) HasIpSmGwMapAddress() bool`

HasIpSmGwMapAddress returns a boolean if a field has been set.

### GetIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) GetIpSmGwDiameterAddress() NetworkNodeDiameterAddress`

GetIpSmGwDiameterAddress returns the IpSmGwDiameterAddress field if non-nil, zero value otherwise.

### GetIpSmGwDiameterAddressOk

`func (o *IpSmGwRegistration) GetIpSmGwDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetIpSmGwDiameterAddressOk returns a tuple with the IpSmGwDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) SetIpSmGwDiameterAddress(v NetworkNodeDiameterAddress)`

SetIpSmGwDiameterAddress sets IpSmGwDiameterAddress field to given value.

### HasIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) HasIpSmGwDiameterAddress() bool`

HasIpSmGwDiameterAddress returns a boolean if a field has been set.

### GetIpsmgwIpv4

`func (o *IpSmGwRegistration) GetIpsmgwIpv4() string`

GetIpsmgwIpv4 returns the IpsmgwIpv4 field if non-nil, zero value otherwise.

### GetIpsmgwIpv4Ok

`func (o *IpSmGwRegistration) GetIpsmgwIpv4Ok() (*string, bool)`

GetIpsmgwIpv4Ok returns a tuple with the IpsmgwIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwIpv4

`func (o *IpSmGwRegistration) SetIpsmgwIpv4(v string)`

SetIpsmgwIpv4 sets IpsmgwIpv4 field to given value.

### HasIpsmgwIpv4

`func (o *IpSmGwRegistration) HasIpsmgwIpv4() bool`

HasIpsmgwIpv4 returns a boolean if a field has been set.

### GetIpsmgwIpv6

`func (o *IpSmGwRegistration) GetIpsmgwIpv6() Ipv6Addr`

GetIpsmgwIpv6 returns the IpsmgwIpv6 field if non-nil, zero value otherwise.

### GetIpsmgwIpv6Ok

`func (o *IpSmGwRegistration) GetIpsmgwIpv6Ok() (*Ipv6Addr, bool)`

GetIpsmgwIpv6Ok returns a tuple with the IpsmgwIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwIpv6

`func (o *IpSmGwRegistration) SetIpsmgwIpv6(v Ipv6Addr)`

SetIpsmgwIpv6 sets IpsmgwIpv6 field to given value.

### HasIpsmgwIpv6

`func (o *IpSmGwRegistration) HasIpsmgwIpv6() bool`

HasIpsmgwIpv6 returns a boolean if a field has been set.

### GetIpsmgwFqdn

`func (o *IpSmGwRegistration) GetIpsmgwFqdn() string`

GetIpsmgwFqdn returns the IpsmgwFqdn field if non-nil, zero value otherwise.

### GetIpsmgwFqdnOk

`func (o *IpSmGwRegistration) GetIpsmgwFqdnOk() (*string, bool)`

GetIpsmgwFqdnOk returns a tuple with the IpsmgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwFqdn

`func (o *IpSmGwRegistration) SetIpsmgwFqdn(v string)`

SetIpsmgwFqdn sets IpsmgwFqdn field to given value.

### HasIpsmgwFqdn

`func (o *IpSmGwRegistration) HasIpsmgwFqdn() bool`

HasIpsmgwFqdn returns a boolean if a field has been set.

### GetNfInstanceId

`func (o *IpSmGwRegistration) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *IpSmGwRegistration) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *IpSmGwRegistration) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.

### HasNfInstanceId

`func (o *IpSmGwRegistration) HasNfInstanceId() bool`

HasNfInstanceId returns a boolean if a field has been set.

### GetUnriIndicator

`func (o *IpSmGwRegistration) GetUnriIndicator() bool`

GetUnriIndicator returns the UnriIndicator field if non-nil, zero value otherwise.

### GetUnriIndicatorOk

`func (o *IpSmGwRegistration) GetUnriIndicatorOk() (*bool, bool)`

GetUnriIndicatorOk returns a tuple with the UnriIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnriIndicator

`func (o *IpSmGwRegistration) SetUnriIndicator(v bool)`

SetUnriIndicator sets UnriIndicator field to given value.

### HasUnriIndicator

`func (o *IpSmGwRegistration) HasUnriIndicator() bool`

HasUnriIndicator returns a boolean if a field has been set.

### GetResetIds

`func (o *IpSmGwRegistration) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *IpSmGwRegistration) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *IpSmGwRegistration) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *IpSmGwRegistration) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetIpSmGwSbiSupInd

`func (o *IpSmGwRegistration) GetIpSmGwSbiSupInd() bool`

GetIpSmGwSbiSupInd returns the IpSmGwSbiSupInd field if non-nil, zero value otherwise.

### GetIpSmGwSbiSupIndOk

`func (o *IpSmGwRegistration) GetIpSmGwSbiSupIndOk() (*bool, bool)`

GetIpSmGwSbiSupIndOk returns a tuple with the IpSmGwSbiSupInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwSbiSupInd

`func (o *IpSmGwRegistration) SetIpSmGwSbiSupInd(v bool)`

SetIpSmGwSbiSupInd sets IpSmGwSbiSupInd field to given value.

### HasIpSmGwSbiSupInd

`func (o *IpSmGwRegistration) HasIpSmGwSbiSupInd() bool`

HasIpSmGwSbiSupInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


