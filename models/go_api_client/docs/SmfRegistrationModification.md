# SmfRegistrationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**PgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewSmfRegistrationModification

`func NewSmfRegistrationModification(smfInstanceId string, ) *SmfRegistrationModification`

NewSmfRegistrationModification instantiates a new SmfRegistrationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfRegistrationModificationWithDefaults

`func NewSmfRegistrationModificationWithDefaults() *SmfRegistrationModification`

NewSmfRegistrationModificationWithDefaults instantiates a new SmfRegistrationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmfInstanceId

`func (o *SmfRegistrationModification) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *SmfRegistrationModification) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *SmfRegistrationModification) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.


### GetSmfSetId

`func (o *SmfRegistrationModification) GetSmfSetId() string`

GetSmfSetId returns the SmfSetId field if non-nil, zero value otherwise.

### GetSmfSetIdOk

`func (o *SmfRegistrationModification) GetSmfSetIdOk() (*string, bool)`

GetSmfSetIdOk returns a tuple with the SmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSetId

`func (o *SmfRegistrationModification) SetSmfSetId(v string)`

SetSmfSetId sets SmfSetId field to given value.

### HasSmfSetId

`func (o *SmfRegistrationModification) HasSmfSetId() bool`

HasSmfSetId returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *SmfRegistrationModification) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *SmfRegistrationModification) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *SmfRegistrationModification) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *SmfRegistrationModification) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


