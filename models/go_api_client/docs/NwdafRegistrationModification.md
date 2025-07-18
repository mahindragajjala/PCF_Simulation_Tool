# NwdafRegistrationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwdafInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NwdafSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**AnalyticsIds** | Pointer to [**[]EventId**](EventId.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewNwdafRegistrationModification

`func NewNwdafRegistrationModification(nwdafInstanceId string, ) *NwdafRegistrationModification`

NewNwdafRegistrationModification instantiates a new NwdafRegistrationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafRegistrationModificationWithDefaults

`func NewNwdafRegistrationModificationWithDefaults() *NwdafRegistrationModification`

NewNwdafRegistrationModificationWithDefaults instantiates a new NwdafRegistrationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwdafInstanceId

`func (o *NwdafRegistrationModification) GetNwdafInstanceId() string`

GetNwdafInstanceId returns the NwdafInstanceId field if non-nil, zero value otherwise.

### GetNwdafInstanceIdOk

`func (o *NwdafRegistrationModification) GetNwdafInstanceIdOk() (*string, bool)`

GetNwdafInstanceIdOk returns a tuple with the NwdafInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInstanceId

`func (o *NwdafRegistrationModification) SetNwdafInstanceId(v string)`

SetNwdafInstanceId sets NwdafInstanceId field to given value.


### GetNwdafSetId

`func (o *NwdafRegistrationModification) GetNwdafSetId() string`

GetNwdafSetId returns the NwdafSetId field if non-nil, zero value otherwise.

### GetNwdafSetIdOk

`func (o *NwdafRegistrationModification) GetNwdafSetIdOk() (*string, bool)`

GetNwdafSetIdOk returns a tuple with the NwdafSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafSetId

`func (o *NwdafRegistrationModification) SetNwdafSetId(v string)`

SetNwdafSetId sets NwdafSetId field to given value.

### HasNwdafSetId

`func (o *NwdafRegistrationModification) HasNwdafSetId() bool`

HasNwdafSetId returns a boolean if a field has been set.

### GetAnalyticsIds

`func (o *NwdafRegistrationModification) GetAnalyticsIds() []EventId`

GetAnalyticsIds returns the AnalyticsIds field if non-nil, zero value otherwise.

### GetAnalyticsIdsOk

`func (o *NwdafRegistrationModification) GetAnalyticsIdsOk() (*[]EventId, bool)`

GetAnalyticsIdsOk returns a tuple with the AnalyticsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsIds

`func (o *NwdafRegistrationModification) SetAnalyticsIds(v []EventId)`

SetAnalyticsIds sets AnalyticsIds field to given value.

### HasAnalyticsIds

`func (o *NwdafRegistrationModification) HasAnalyticsIds() bool`

HasAnalyticsIds returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NwdafRegistrationModification) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NwdafRegistrationModification) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NwdafRegistrationModification) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NwdafRegistrationModification) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


