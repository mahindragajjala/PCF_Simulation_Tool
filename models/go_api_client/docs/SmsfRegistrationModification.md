# SmsfRegistrationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SmsfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**UeMemoryAvailableInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmsfRegistrationModification

`func NewSmsfRegistrationModification(smsfInstanceId string, ) *SmsfRegistrationModification`

NewSmsfRegistrationModification instantiates a new SmsfRegistrationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsfRegistrationModificationWithDefaults

`func NewSmsfRegistrationModificationWithDefaults() *SmsfRegistrationModification`

NewSmsfRegistrationModificationWithDefaults instantiates a new SmsfRegistrationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsfInstanceId

`func (o *SmsfRegistrationModification) GetSmsfInstanceId() string`

GetSmsfInstanceId returns the SmsfInstanceId field if non-nil, zero value otherwise.

### GetSmsfInstanceIdOk

`func (o *SmsfRegistrationModification) GetSmsfInstanceIdOk() (*string, bool)`

GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInstanceId

`func (o *SmsfRegistrationModification) SetSmsfInstanceId(v string)`

SetSmsfInstanceId sets SmsfInstanceId field to given value.


### GetSmsfSetId

`func (o *SmsfRegistrationModification) GetSmsfSetId() string`

GetSmsfSetId returns the SmsfSetId field if non-nil, zero value otherwise.

### GetSmsfSetIdOk

`func (o *SmsfRegistrationModification) GetSmsfSetIdOk() (*string, bool)`

GetSmsfSetIdOk returns a tuple with the SmsfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSetId

`func (o *SmsfRegistrationModification) SetSmsfSetId(v string)`

SetSmsfSetId sets SmsfSetId field to given value.

### HasSmsfSetId

`func (o *SmsfRegistrationModification) HasSmsfSetId() bool`

HasSmsfSetId returns a boolean if a field has been set.

### GetUeMemoryAvailableInd

`func (o *SmsfRegistrationModification) GetUeMemoryAvailableInd() bool`

GetUeMemoryAvailableInd returns the UeMemoryAvailableInd field if non-nil, zero value otherwise.

### GetUeMemoryAvailableIndOk

`func (o *SmsfRegistrationModification) GetUeMemoryAvailableIndOk() (*bool, bool)`

GetUeMemoryAvailableIndOk returns a tuple with the UeMemoryAvailableInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMemoryAvailableInd

`func (o *SmsfRegistrationModification) SetUeMemoryAvailableInd(v bool)`

SetUeMemoryAvailableInd sets UeMemoryAvailableInd field to given value.

### HasUeMemoryAvailableInd

`func (o *SmsfRegistrationModification) HasUeMemoryAvailableInd() bool`

HasUeMemoryAvailableInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


