# SmsfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SmsfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SmsfMAPAddress** | Pointer to **string** | This data type mentions International E.164 number of the SMSF; shall be present if the SMSF supports MAP.  | [optional] 
**SmsfDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**DataRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**SmsfSbiSupInd** | Pointer to **bool** |  | [optional] [default to false]
**UdrRestartInd** | Pointer to **bool** |  | [optional] [default to false]
**LastSynchronizationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**UeMemoryAvailableInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmsfRegistration

`func NewSmsfRegistration(smsfInstanceId string, plmnId PlmnId, ) *SmsfRegistration`

NewSmsfRegistration instantiates a new SmsfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsfRegistrationWithDefaults

`func NewSmsfRegistrationWithDefaults() *SmsfRegistration`

NewSmsfRegistrationWithDefaults instantiates a new SmsfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsfInstanceId

`func (o *SmsfRegistration) GetSmsfInstanceId() string`

GetSmsfInstanceId returns the SmsfInstanceId field if non-nil, zero value otherwise.

### GetSmsfInstanceIdOk

`func (o *SmsfRegistration) GetSmsfInstanceIdOk() (*string, bool)`

GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInstanceId

`func (o *SmsfRegistration) SetSmsfInstanceId(v string)`

SetSmsfInstanceId sets SmsfInstanceId field to given value.


### GetSmsfSetId

`func (o *SmsfRegistration) GetSmsfSetId() string`

GetSmsfSetId returns the SmsfSetId field if non-nil, zero value otherwise.

### GetSmsfSetIdOk

`func (o *SmsfRegistration) GetSmsfSetIdOk() (*string, bool)`

GetSmsfSetIdOk returns a tuple with the SmsfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSetId

`func (o *SmsfRegistration) SetSmsfSetId(v string)`

SetSmsfSetId sets SmsfSetId field to given value.

### HasSmsfSetId

`func (o *SmsfRegistration) HasSmsfSetId() bool`

HasSmsfSetId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmsfRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmsfRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmsfRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmsfRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPlmnId

`func (o *SmsfRegistration) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SmsfRegistration) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SmsfRegistration) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSmsfMAPAddress

`func (o *SmsfRegistration) GetSmsfMAPAddress() string`

GetSmsfMAPAddress returns the SmsfMAPAddress field if non-nil, zero value otherwise.

### GetSmsfMAPAddressOk

`func (o *SmsfRegistration) GetSmsfMAPAddressOk() (*string, bool)`

GetSmsfMAPAddressOk returns a tuple with the SmsfMAPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfMAPAddress

`func (o *SmsfRegistration) SetSmsfMAPAddress(v string)`

SetSmsfMAPAddress sets SmsfMAPAddress field to given value.

### HasSmsfMAPAddress

`func (o *SmsfRegistration) HasSmsfMAPAddress() bool`

HasSmsfMAPAddress returns a boolean if a field has been set.

### GetSmsfDiameterAddress

`func (o *SmsfRegistration) GetSmsfDiameterAddress() NetworkNodeDiameterAddress`

GetSmsfDiameterAddress returns the SmsfDiameterAddress field if non-nil, zero value otherwise.

### GetSmsfDiameterAddressOk

`func (o *SmsfRegistration) GetSmsfDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetSmsfDiameterAddressOk returns a tuple with the SmsfDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfDiameterAddress

`func (o *SmsfRegistration) SetSmsfDiameterAddress(v NetworkNodeDiameterAddress)`

SetSmsfDiameterAddress sets SmsfDiameterAddress field to given value.

### HasSmsfDiameterAddress

`func (o *SmsfRegistration) HasSmsfDiameterAddress() bool`

HasSmsfDiameterAddress returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *SmsfRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *SmsfRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *SmsfRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *SmsfRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetContextInfo

`func (o *SmsfRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SmsfRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SmsfRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SmsfRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetDataRestorationCallbackUri

`func (o *SmsfRegistration) GetDataRestorationCallbackUri() string`

GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field if non-nil, zero value otherwise.

### GetDataRestorationCallbackUriOk

`func (o *SmsfRegistration) GetDataRestorationCallbackUriOk() (*string, bool)`

GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRestorationCallbackUri

`func (o *SmsfRegistration) SetDataRestorationCallbackUri(v string)`

SetDataRestorationCallbackUri sets DataRestorationCallbackUri field to given value.

### HasDataRestorationCallbackUri

`func (o *SmsfRegistration) HasDataRestorationCallbackUri() bool`

HasDataRestorationCallbackUri returns a boolean if a field has been set.

### GetResetIds

`func (o *SmsfRegistration) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *SmsfRegistration) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *SmsfRegistration) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *SmsfRegistration) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetSmsfSbiSupInd

`func (o *SmsfRegistration) GetSmsfSbiSupInd() bool`

GetSmsfSbiSupInd returns the SmsfSbiSupInd field if non-nil, zero value otherwise.

### GetSmsfSbiSupIndOk

`func (o *SmsfRegistration) GetSmsfSbiSupIndOk() (*bool, bool)`

GetSmsfSbiSupIndOk returns a tuple with the SmsfSbiSupInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSbiSupInd

`func (o *SmsfRegistration) SetSmsfSbiSupInd(v bool)`

SetSmsfSbiSupInd sets SmsfSbiSupInd field to given value.

### HasSmsfSbiSupInd

`func (o *SmsfRegistration) HasSmsfSbiSupInd() bool`

HasSmsfSbiSupInd returns a boolean if a field has been set.

### GetUdrRestartInd

`func (o *SmsfRegistration) GetUdrRestartInd() bool`

GetUdrRestartInd returns the UdrRestartInd field if non-nil, zero value otherwise.

### GetUdrRestartIndOk

`func (o *SmsfRegistration) GetUdrRestartIndOk() (*bool, bool)`

GetUdrRestartIndOk returns a tuple with the UdrRestartInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrRestartInd

`func (o *SmsfRegistration) SetUdrRestartInd(v bool)`

SetUdrRestartInd sets UdrRestartInd field to given value.

### HasUdrRestartInd

`func (o *SmsfRegistration) HasUdrRestartInd() bool`

HasUdrRestartInd returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *SmsfRegistration) GetLastSynchronizationTime() time.Time`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *SmsfRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *SmsfRegistration) SetLastSynchronizationTime(v time.Time)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *SmsfRegistration) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.

### GetUeMemoryAvailableInd

`func (o *SmsfRegistration) GetUeMemoryAvailableInd() bool`

GetUeMemoryAvailableInd returns the UeMemoryAvailableInd field if non-nil, zero value otherwise.

### GetUeMemoryAvailableIndOk

`func (o *SmsfRegistration) GetUeMemoryAvailableIndOk() (*bool, bool)`

GetUeMemoryAvailableIndOk returns a tuple with the UeMemoryAvailableInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMemoryAvailableInd

`func (o *SmsfRegistration) SetUeMemoryAvailableInd(v bool)`

SetUeMemoryAvailableInd sets UeMemoryAvailableInd field to given value.

### HasUeMemoryAvailableInd

`func (o *SmsfRegistration) HasUeMemoryAvailableInd() bool`

HasUeMemoryAvailableInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


