# RoutingInfoSmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Smsf3Gpp** | Pointer to [**SmsfRegistration**](SmsfRegistration.md) |  | [optional] 
**SmsfNon3Gpp** | Pointer to [**SmsfRegistration**](SmsfRegistration.md) |  | [optional] 
**IpSmGw** | Pointer to [**IpSmGwInfo**](IpSmGwInfo.md) |  | [optional] 
**SmsRouter** | Pointer to [**SmsRouterInfo**](SmsRouterInfo.md) |  | [optional] 

## Methods

### NewRoutingInfoSmResponse

`func NewRoutingInfoSmResponse() *RoutingInfoSmResponse`

NewRoutingInfoSmResponse instantiates a new RoutingInfoSmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingInfoSmResponseWithDefaults

`func NewRoutingInfoSmResponseWithDefaults() *RoutingInfoSmResponse`

NewRoutingInfoSmResponseWithDefaults instantiates a new RoutingInfoSmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *RoutingInfoSmResponse) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *RoutingInfoSmResponse) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *RoutingInfoSmResponse) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *RoutingInfoSmResponse) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetSmsf3Gpp

`func (o *RoutingInfoSmResponse) GetSmsf3Gpp() SmsfRegistration`

GetSmsf3Gpp returns the Smsf3Gpp field if non-nil, zero value otherwise.

### GetSmsf3GppOk

`func (o *RoutingInfoSmResponse) GetSmsf3GppOk() (*SmsfRegistration, bool)`

GetSmsf3GppOk returns a tuple with the Smsf3Gpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsf3Gpp

`func (o *RoutingInfoSmResponse) SetSmsf3Gpp(v SmsfRegistration)`

SetSmsf3Gpp sets Smsf3Gpp field to given value.

### HasSmsf3Gpp

`func (o *RoutingInfoSmResponse) HasSmsf3Gpp() bool`

HasSmsf3Gpp returns a boolean if a field has been set.

### GetSmsfNon3Gpp

`func (o *RoutingInfoSmResponse) GetSmsfNon3Gpp() SmsfRegistration`

GetSmsfNon3Gpp returns the SmsfNon3Gpp field if non-nil, zero value otherwise.

### GetSmsfNon3GppOk

`func (o *RoutingInfoSmResponse) GetSmsfNon3GppOk() (*SmsfRegistration, bool)`

GetSmsfNon3GppOk returns a tuple with the SmsfNon3Gpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfNon3Gpp

`func (o *RoutingInfoSmResponse) SetSmsfNon3Gpp(v SmsfRegistration)`

SetSmsfNon3Gpp sets SmsfNon3Gpp field to given value.

### HasSmsfNon3Gpp

`func (o *RoutingInfoSmResponse) HasSmsfNon3Gpp() bool`

HasSmsfNon3Gpp returns a boolean if a field has been set.

### GetIpSmGw

`func (o *RoutingInfoSmResponse) GetIpSmGw() IpSmGwInfo`

GetIpSmGw returns the IpSmGw field if non-nil, zero value otherwise.

### GetIpSmGwOk

`func (o *RoutingInfoSmResponse) GetIpSmGwOk() (*IpSmGwInfo, bool)`

GetIpSmGwOk returns a tuple with the IpSmGw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGw

`func (o *RoutingInfoSmResponse) SetIpSmGw(v IpSmGwInfo)`

SetIpSmGw sets IpSmGw field to given value.

### HasIpSmGw

`func (o *RoutingInfoSmResponse) HasIpSmGw() bool`

HasIpSmGw returns a boolean if a field has been set.

### GetSmsRouter

`func (o *RoutingInfoSmResponse) GetSmsRouter() SmsRouterInfo`

GetSmsRouter returns the SmsRouter field if non-nil, zero value otherwise.

### GetSmsRouterOk

`func (o *RoutingInfoSmResponse) GetSmsRouterOk() (*SmsRouterInfo, bool)`

GetSmsRouterOk returns a tuple with the SmsRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsRouter

`func (o *RoutingInfoSmResponse) SetSmsRouter(v SmsRouterInfo)`

SetSmsRouter sets SmsRouter field to given value.

### HasSmsRouter

`func (o *RoutingInfoSmResponse) HasSmsRouter() bool`

HasSmsRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


