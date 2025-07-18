# PcscfRestorationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**FailedPcscf** | Pointer to [**NullablePcscfAddress**](PcscfAddress.md) |  | [optional] 

## Methods

### NewPcscfRestorationNotification

`func NewPcscfRestorationNotification(supi string, ) *PcscfRestorationNotification`

NewPcscfRestorationNotification instantiates a new PcscfRestorationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcscfRestorationNotificationWithDefaults

`func NewPcscfRestorationNotificationWithDefaults() *PcscfRestorationNotification`

NewPcscfRestorationNotificationWithDefaults instantiates a new PcscfRestorationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PcscfRestorationNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PcscfRestorationNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PcscfRestorationNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetFailedPcscf

`func (o *PcscfRestorationNotification) GetFailedPcscf() PcscfAddress`

GetFailedPcscf returns the FailedPcscf field if non-nil, zero value otherwise.

### GetFailedPcscfOk

`func (o *PcscfRestorationNotification) GetFailedPcscfOk() (*PcscfAddress, bool)`

GetFailedPcscfOk returns a tuple with the FailedPcscf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPcscf

`func (o *PcscfRestorationNotification) SetFailedPcscf(v PcscfAddress)`

SetFailedPcscf sets FailedPcscf field to given value.

### HasFailedPcscf

`func (o *PcscfRestorationNotification) HasFailedPcscf() bool`

HasFailedPcscf returns a boolean if a field has been set.

### SetFailedPcscfNil

`func (o *PcscfRestorationNotification) SetFailedPcscfNil(b bool)`

 SetFailedPcscfNil sets the value for FailedPcscf to be an explicit nil

### UnsetFailedPcscf
`func (o *PcscfRestorationNotification) UnsetFailedPcscf()`

UnsetFailedPcscf ensures that no value is present for FailedPcscf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


