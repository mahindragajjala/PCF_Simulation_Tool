# SmfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**SingleNssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**EmergencyServices** | Pointer to **bool** |  | [optional] 
**PcscfRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**PgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PgwIpAddr** | Pointer to [**NullableIpAddress**](IpAddress.md) |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]
**DeregCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**RegistrationReason** | Pointer to [**RegistrationReason**](RegistrationReason.md) |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**DataRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**UdrRestartInd** | Pointer to **bool** |  | [optional] [default to false]
**LastSynchronizationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PduSessionReActivationRequired** | Pointer to **bool** |  | [optional] [default to false]
**StaleCheckCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**UdmStaleCheckCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**WildcardInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmfRegistration

`func NewSmfRegistration(smfInstanceId string, pduSessionId int32, singleNssai Snssai, plmnId PlmnId, ) *SmfRegistration`

NewSmfRegistration instantiates a new SmfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfRegistrationWithDefaults

`func NewSmfRegistrationWithDefaults() *SmfRegistration`

NewSmfRegistrationWithDefaults instantiates a new SmfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmfInstanceId

`func (o *SmfRegistration) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *SmfRegistration) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *SmfRegistration) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.


### GetSmfSetId

`func (o *SmfRegistration) GetSmfSetId() string`

GetSmfSetId returns the SmfSetId field if non-nil, zero value otherwise.

### GetSmfSetIdOk

`func (o *SmfRegistration) GetSmfSetIdOk() (*string, bool)`

GetSmfSetIdOk returns a tuple with the SmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSetId

`func (o *SmfRegistration) SetSmfSetId(v string)`

SetSmfSetId sets SmfSetId field to given value.

### HasSmfSetId

`func (o *SmfRegistration) HasSmfSetId() bool`

HasSmfSetId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmfRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmfRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmfRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmfRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmfRegistration) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmfRegistration) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmfRegistration) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetSingleNssai

`func (o *SmfRegistration) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *SmfRegistration) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *SmfRegistration) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.


### GetDnn

`func (o *SmfRegistration) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmfRegistration) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmfRegistration) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmfRegistration) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetEmergencyServices

`func (o *SmfRegistration) GetEmergencyServices() bool`

GetEmergencyServices returns the EmergencyServices field if non-nil, zero value otherwise.

### GetEmergencyServicesOk

`func (o *SmfRegistration) GetEmergencyServicesOk() (*bool, bool)`

GetEmergencyServicesOk returns a tuple with the EmergencyServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyServices

`func (o *SmfRegistration) SetEmergencyServices(v bool)`

SetEmergencyServices sets EmergencyServices field to given value.

### HasEmergencyServices

`func (o *SmfRegistration) HasEmergencyServices() bool`

HasEmergencyServices returns a boolean if a field has been set.

### GetPcscfRestorationCallbackUri

`func (o *SmfRegistration) GetPcscfRestorationCallbackUri() string`

GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field if non-nil, zero value otherwise.

### GetPcscfRestorationCallbackUriOk

`func (o *SmfRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool)`

GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationCallbackUri

`func (o *SmfRegistration) SetPcscfRestorationCallbackUri(v string)`

SetPcscfRestorationCallbackUri sets PcscfRestorationCallbackUri field to given value.

### HasPcscfRestorationCallbackUri

`func (o *SmfRegistration) HasPcscfRestorationCallbackUri() bool`

HasPcscfRestorationCallbackUri returns a boolean if a field has been set.

### GetPlmnId

`func (o *SmfRegistration) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SmfRegistration) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SmfRegistration) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetPgwFqdn

`func (o *SmfRegistration) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *SmfRegistration) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *SmfRegistration) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *SmfRegistration) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetPgwIpAddr

`func (o *SmfRegistration) GetPgwIpAddr() IpAddress`

GetPgwIpAddr returns the PgwIpAddr field if non-nil, zero value otherwise.

### GetPgwIpAddrOk

`func (o *SmfRegistration) GetPgwIpAddrOk() (*IpAddress, bool)`

GetPgwIpAddrOk returns a tuple with the PgwIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddr

`func (o *SmfRegistration) SetPgwIpAddr(v IpAddress)`

SetPgwIpAddr sets PgwIpAddr field to given value.

### HasPgwIpAddr

`func (o *SmfRegistration) HasPgwIpAddr() bool`

HasPgwIpAddr returns a boolean if a field has been set.

### SetPgwIpAddrNil

`func (o *SmfRegistration) SetPgwIpAddrNil(b bool)`

 SetPgwIpAddrNil sets the value for PgwIpAddr to be an explicit nil

### UnsetPgwIpAddr
`func (o *SmfRegistration) UnsetPgwIpAddr()`

UnsetPgwIpAddr ensures that no value is present for PgwIpAddr, not even an explicit nil
### GetEpdgInd

`func (o *SmfRegistration) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *SmfRegistration) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *SmfRegistration) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *SmfRegistration) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.

### GetDeregCallbackUri

`func (o *SmfRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *SmfRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *SmfRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.

### HasDeregCallbackUri

`func (o *SmfRegistration) HasDeregCallbackUri() bool`

HasDeregCallbackUri returns a boolean if a field has been set.

### GetRegistrationReason

`func (o *SmfRegistration) GetRegistrationReason() RegistrationReason`

GetRegistrationReason returns the RegistrationReason field if non-nil, zero value otherwise.

### GetRegistrationReasonOk

`func (o *SmfRegistration) GetRegistrationReasonOk() (*RegistrationReason, bool)`

GetRegistrationReasonOk returns a tuple with the RegistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationReason

`func (o *SmfRegistration) SetRegistrationReason(v RegistrationReason)`

SetRegistrationReason sets RegistrationReason field to given value.

### HasRegistrationReason

`func (o *SmfRegistration) HasRegistrationReason() bool`

HasRegistrationReason returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *SmfRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *SmfRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *SmfRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *SmfRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetContextInfo

`func (o *SmfRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SmfRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SmfRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SmfRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetPcfId

`func (o *SmfRegistration) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *SmfRegistration) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *SmfRegistration) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *SmfRegistration) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetDataRestorationCallbackUri

`func (o *SmfRegistration) GetDataRestorationCallbackUri() string`

GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field if non-nil, zero value otherwise.

### GetDataRestorationCallbackUriOk

`func (o *SmfRegistration) GetDataRestorationCallbackUriOk() (*string, bool)`

GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRestorationCallbackUri

`func (o *SmfRegistration) SetDataRestorationCallbackUri(v string)`

SetDataRestorationCallbackUri sets DataRestorationCallbackUri field to given value.

### HasDataRestorationCallbackUri

`func (o *SmfRegistration) HasDataRestorationCallbackUri() bool`

HasDataRestorationCallbackUri returns a boolean if a field has been set.

### GetResetIds

`func (o *SmfRegistration) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *SmfRegistration) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *SmfRegistration) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *SmfRegistration) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetUdrRestartInd

`func (o *SmfRegistration) GetUdrRestartInd() bool`

GetUdrRestartInd returns the UdrRestartInd field if non-nil, zero value otherwise.

### GetUdrRestartIndOk

`func (o *SmfRegistration) GetUdrRestartIndOk() (*bool, bool)`

GetUdrRestartIndOk returns a tuple with the UdrRestartInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrRestartInd

`func (o *SmfRegistration) SetUdrRestartInd(v bool)`

SetUdrRestartInd sets UdrRestartInd field to given value.

### HasUdrRestartInd

`func (o *SmfRegistration) HasUdrRestartInd() bool`

HasUdrRestartInd returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *SmfRegistration) GetLastSynchronizationTime() time.Time`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *SmfRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *SmfRegistration) SetLastSynchronizationTime(v time.Time)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *SmfRegistration) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.

### GetPduSessionReActivationRequired

`func (o *SmfRegistration) GetPduSessionReActivationRequired() bool`

GetPduSessionReActivationRequired returns the PduSessionReActivationRequired field if non-nil, zero value otherwise.

### GetPduSessionReActivationRequiredOk

`func (o *SmfRegistration) GetPduSessionReActivationRequiredOk() (*bool, bool)`

GetPduSessionReActivationRequiredOk returns a tuple with the PduSessionReActivationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionReActivationRequired

`func (o *SmfRegistration) SetPduSessionReActivationRequired(v bool)`

SetPduSessionReActivationRequired sets PduSessionReActivationRequired field to given value.

### HasPduSessionReActivationRequired

`func (o *SmfRegistration) HasPduSessionReActivationRequired() bool`

HasPduSessionReActivationRequired returns a boolean if a field has been set.

### GetStaleCheckCallbackUri

`func (o *SmfRegistration) GetStaleCheckCallbackUri() string`

GetStaleCheckCallbackUri returns the StaleCheckCallbackUri field if non-nil, zero value otherwise.

### GetStaleCheckCallbackUriOk

`func (o *SmfRegistration) GetStaleCheckCallbackUriOk() (*string, bool)`

GetStaleCheckCallbackUriOk returns a tuple with the StaleCheckCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCheckCallbackUri

`func (o *SmfRegistration) SetStaleCheckCallbackUri(v string)`

SetStaleCheckCallbackUri sets StaleCheckCallbackUri field to given value.

### HasStaleCheckCallbackUri

`func (o *SmfRegistration) HasStaleCheckCallbackUri() bool`

HasStaleCheckCallbackUri returns a boolean if a field has been set.

### GetUdmStaleCheckCallbackUri

`func (o *SmfRegistration) GetUdmStaleCheckCallbackUri() string`

GetUdmStaleCheckCallbackUri returns the UdmStaleCheckCallbackUri field if non-nil, zero value otherwise.

### GetUdmStaleCheckCallbackUriOk

`func (o *SmfRegistration) GetUdmStaleCheckCallbackUriOk() (*string, bool)`

GetUdmStaleCheckCallbackUriOk returns a tuple with the UdmStaleCheckCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmStaleCheckCallbackUri

`func (o *SmfRegistration) SetUdmStaleCheckCallbackUri(v string)`

SetUdmStaleCheckCallbackUri sets UdmStaleCheckCallbackUri field to given value.

### HasUdmStaleCheckCallbackUri

`func (o *SmfRegistration) HasUdmStaleCheckCallbackUri() bool`

HasUdmStaleCheckCallbackUri returns a boolean if a field has been set.

### GetWildcardInd

`func (o *SmfRegistration) GetWildcardInd() bool`

GetWildcardInd returns the WildcardInd field if non-nil, zero value otherwise.

### GetWildcardIndOk

`func (o *SmfRegistration) GetWildcardIndOk() (*bool, bool)`

GetWildcardIndOk returns a tuple with the WildcardInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardInd

`func (o *SmfRegistration) SetWildcardInd(v bool)`

SetWildcardInd sets WildcardInd field to given value.

### HasWildcardInd

`func (o *SmfRegistration) HasWildcardInd() bool`

HasWildcardInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


