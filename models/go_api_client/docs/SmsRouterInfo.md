# SmsRouterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**DiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**MapAddress** | Pointer to **string** | This data type mentions International E.164 number of the SMSF; shall be present if the SMSF supports MAP.  | [optional] 
**RouterIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**RouterIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**RouterFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewSmsRouterInfo

`func NewSmsRouterInfo() *SmsRouterInfo`

NewSmsRouterInfo instantiates a new SmsRouterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRouterInfoWithDefaults

`func NewSmsRouterInfoWithDefaults() *SmsRouterInfo`

NewSmsRouterInfoWithDefaults instantiates a new SmsRouterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SmsRouterInfo) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SmsRouterInfo) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SmsRouterInfo) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.

### HasNfInstanceId

`func (o *SmsRouterInfo) HasNfInstanceId() bool`

HasNfInstanceId returns a boolean if a field has been set.

### GetDiameterAddress

`func (o *SmsRouterInfo) GetDiameterAddress() NetworkNodeDiameterAddress`

GetDiameterAddress returns the DiameterAddress field if non-nil, zero value otherwise.

### GetDiameterAddressOk

`func (o *SmsRouterInfo) GetDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetDiameterAddressOk returns a tuple with the DiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterAddress

`func (o *SmsRouterInfo) SetDiameterAddress(v NetworkNodeDiameterAddress)`

SetDiameterAddress sets DiameterAddress field to given value.

### HasDiameterAddress

`func (o *SmsRouterInfo) HasDiameterAddress() bool`

HasDiameterAddress returns a boolean if a field has been set.

### GetMapAddress

`func (o *SmsRouterInfo) GetMapAddress() string`

GetMapAddress returns the MapAddress field if non-nil, zero value otherwise.

### GetMapAddressOk

`func (o *SmsRouterInfo) GetMapAddressOk() (*string, bool)`

GetMapAddressOk returns a tuple with the MapAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAddress

`func (o *SmsRouterInfo) SetMapAddress(v string)`

SetMapAddress sets MapAddress field to given value.

### HasMapAddress

`func (o *SmsRouterInfo) HasMapAddress() bool`

HasMapAddress returns a boolean if a field has been set.

### GetRouterIpv4

`func (o *SmsRouterInfo) GetRouterIpv4() string`

GetRouterIpv4 returns the RouterIpv4 field if non-nil, zero value otherwise.

### GetRouterIpv4Ok

`func (o *SmsRouterInfo) GetRouterIpv4Ok() (*string, bool)`

GetRouterIpv4Ok returns a tuple with the RouterIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIpv4

`func (o *SmsRouterInfo) SetRouterIpv4(v string)`

SetRouterIpv4 sets RouterIpv4 field to given value.

### HasRouterIpv4

`func (o *SmsRouterInfo) HasRouterIpv4() bool`

HasRouterIpv4 returns a boolean if a field has been set.

### GetRouterIpv6

`func (o *SmsRouterInfo) GetRouterIpv6() Ipv6Addr`

GetRouterIpv6 returns the RouterIpv6 field if non-nil, zero value otherwise.

### GetRouterIpv6Ok

`func (o *SmsRouterInfo) GetRouterIpv6Ok() (*Ipv6Addr, bool)`

GetRouterIpv6Ok returns a tuple with the RouterIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIpv6

`func (o *SmsRouterInfo) SetRouterIpv6(v Ipv6Addr)`

SetRouterIpv6 sets RouterIpv6 field to given value.

### HasRouterIpv6

`func (o *SmsRouterInfo) HasRouterIpv6() bool`

HasRouterIpv6 returns a boolean if a field has been set.

### GetRouterFqdn

`func (o *SmsRouterInfo) GetRouterFqdn() string`

GetRouterFqdn returns the RouterFqdn field if non-nil, zero value otherwise.

### GetRouterFqdnOk

`func (o *SmsRouterInfo) GetRouterFqdnOk() (*string, bool)`

GetRouterFqdnOk returns a tuple with the RouterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterFqdn

`func (o *SmsRouterInfo) SetRouterFqdn(v string)`

SetRouterFqdn sets RouterFqdn field to given value.

### HasRouterFqdn

`func (o *SmsRouterInfo) HasRouterFqdn() bool`

HasRouterFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


