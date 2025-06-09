# Amf3GppAccessRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PurgeFlag** | Pointer to **bool** | This flag indicates whether an AMF is deregistered. This attribute may be included in notifications sent by the UDR to the UDM if purgeFlag is also set to true in the same notification.  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**ImsVoPs** | Pointer to [**ImsVoPs**](ImsVoPs.md) |  | [optional] 
**DeregCallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AmfServiceNameDereg** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**PcscfRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AmfServiceNamePcscfRest** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**InitialRegistrationInd** | Pointer to **bool** |  | [optional] 
**EmergencyRegistrationInd** | Pointer to **bool** |  | [optional] 
**Guami** | [**Guami**](Guami.md) |  | 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**DrFlag** | Pointer to **bool** | This data type indicates that the UDM+HSS is requested not to send S6a-CLR to the registered MME/SGSN (if any).  | [optional] 
**RatType** | [**RatType**](RatType.md) |  | 
**UrrpIndicator** | Pointer to **bool** |  | [optional] 
**AmfEeSubscriptionId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**EpsInterworkingInfo** | Pointer to [**EpsInterworkingInfo**](EpsInterworkingInfo.md) |  | [optional] 
**UeSrvccCapability** | Pointer to **bool** |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**VgmlcAddress** | Pointer to [**VgmlcAddress**](VgmlcAddress.md) |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**NoEeSubscriptionInd** | Pointer to **bool** |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UeReachableInd** | Pointer to [**UeReachableInd**](UeReachableInd.md) |  | [optional] 
**ReRegistrationRequired** | Pointer to **bool** |  | [optional] 
**AdminDeregSubWithdrawn** | Pointer to **bool** |  | [optional] 
**DataRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]
**UeMINTCapability** | Pointer to **bool** |  | [optional] 
**SorSnpnSiSupported** | Pointer to **bool** |  | [optional] [default to false]
**UdrRestartInd** | Pointer to **bool** |  | [optional] [default to false]
**LastSynchronizationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAmf3GppAccessRegistration

`func NewAmf3GppAccessRegistration(amfInstanceId string, deregCallbackUri string, guami Guami, ratType RatType, ) *Amf3GppAccessRegistration`

NewAmf3GppAccessRegistration instantiates a new Amf3GppAccessRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmf3GppAccessRegistrationWithDefaults

`func NewAmf3GppAccessRegistrationWithDefaults() *Amf3GppAccessRegistration`

NewAmf3GppAccessRegistrationWithDefaults instantiates a new Amf3GppAccessRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *Amf3GppAccessRegistration) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *Amf3GppAccessRegistration) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *Amf3GppAccessRegistration) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetSupportedFeatures

`func (o *Amf3GppAccessRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Amf3GppAccessRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Amf3GppAccessRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Amf3GppAccessRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPurgeFlag

`func (o *Amf3GppAccessRegistration) GetPurgeFlag() bool`

GetPurgeFlag returns the PurgeFlag field if non-nil, zero value otherwise.

### GetPurgeFlagOk

`func (o *Amf3GppAccessRegistration) GetPurgeFlagOk() (*bool, bool)`

GetPurgeFlagOk returns a tuple with the PurgeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeFlag

`func (o *Amf3GppAccessRegistration) SetPurgeFlag(v bool)`

SetPurgeFlag sets PurgeFlag field to given value.

### HasPurgeFlag

`func (o *Amf3GppAccessRegistration) HasPurgeFlag() bool`

HasPurgeFlag returns a boolean if a field has been set.

### GetPei

`func (o *Amf3GppAccessRegistration) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *Amf3GppAccessRegistration) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *Amf3GppAccessRegistration) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *Amf3GppAccessRegistration) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetImsVoPs

`func (o *Amf3GppAccessRegistration) GetImsVoPs() ImsVoPs`

GetImsVoPs returns the ImsVoPs field if non-nil, zero value otherwise.

### GetImsVoPsOk

`func (o *Amf3GppAccessRegistration) GetImsVoPsOk() (*ImsVoPs, bool)`

GetImsVoPsOk returns a tuple with the ImsVoPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVoPs

`func (o *Amf3GppAccessRegistration) SetImsVoPs(v ImsVoPs)`

SetImsVoPs sets ImsVoPs field to given value.

### HasImsVoPs

`func (o *Amf3GppAccessRegistration) HasImsVoPs() bool`

HasImsVoPs returns a boolean if a field has been set.

### GetDeregCallbackUri

`func (o *Amf3GppAccessRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *Amf3GppAccessRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *Amf3GppAccessRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.


### GetAmfServiceNameDereg

`func (o *Amf3GppAccessRegistration) GetAmfServiceNameDereg() ServiceName`

GetAmfServiceNameDereg returns the AmfServiceNameDereg field if non-nil, zero value otherwise.

### GetAmfServiceNameDeregOk

`func (o *Amf3GppAccessRegistration) GetAmfServiceNameDeregOk() (*ServiceName, bool)`

GetAmfServiceNameDeregOk returns a tuple with the AmfServiceNameDereg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNameDereg

`func (o *Amf3GppAccessRegistration) SetAmfServiceNameDereg(v ServiceName)`

SetAmfServiceNameDereg sets AmfServiceNameDereg field to given value.

### HasAmfServiceNameDereg

`func (o *Amf3GppAccessRegistration) HasAmfServiceNameDereg() bool`

HasAmfServiceNameDereg returns a boolean if a field has been set.

### GetPcscfRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) GetPcscfRestorationCallbackUri() string`

GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field if non-nil, zero value otherwise.

### GetPcscfRestorationCallbackUriOk

`func (o *Amf3GppAccessRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool)`

GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) SetPcscfRestorationCallbackUri(v string)`

SetPcscfRestorationCallbackUri sets PcscfRestorationCallbackUri field to given value.

### HasPcscfRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) HasPcscfRestorationCallbackUri() bool`

HasPcscfRestorationCallbackUri returns a boolean if a field has been set.

### GetAmfServiceNamePcscfRest

`func (o *Amf3GppAccessRegistration) GetAmfServiceNamePcscfRest() ServiceName`

GetAmfServiceNamePcscfRest returns the AmfServiceNamePcscfRest field if non-nil, zero value otherwise.

### GetAmfServiceNamePcscfRestOk

`func (o *Amf3GppAccessRegistration) GetAmfServiceNamePcscfRestOk() (*ServiceName, bool)`

GetAmfServiceNamePcscfRestOk returns a tuple with the AmfServiceNamePcscfRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNamePcscfRest

`func (o *Amf3GppAccessRegistration) SetAmfServiceNamePcscfRest(v ServiceName)`

SetAmfServiceNamePcscfRest sets AmfServiceNamePcscfRest field to given value.

### HasAmfServiceNamePcscfRest

`func (o *Amf3GppAccessRegistration) HasAmfServiceNamePcscfRest() bool`

HasAmfServiceNamePcscfRest returns a boolean if a field has been set.

### GetInitialRegistrationInd

`func (o *Amf3GppAccessRegistration) GetInitialRegistrationInd() bool`

GetInitialRegistrationInd returns the InitialRegistrationInd field if non-nil, zero value otherwise.

### GetInitialRegistrationIndOk

`func (o *Amf3GppAccessRegistration) GetInitialRegistrationIndOk() (*bool, bool)`

GetInitialRegistrationIndOk returns a tuple with the InitialRegistrationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialRegistrationInd

`func (o *Amf3GppAccessRegistration) SetInitialRegistrationInd(v bool)`

SetInitialRegistrationInd sets InitialRegistrationInd field to given value.

### HasInitialRegistrationInd

`func (o *Amf3GppAccessRegistration) HasInitialRegistrationInd() bool`

HasInitialRegistrationInd returns a boolean if a field has been set.

### GetEmergencyRegistrationInd

`func (o *Amf3GppAccessRegistration) GetEmergencyRegistrationInd() bool`

GetEmergencyRegistrationInd returns the EmergencyRegistrationInd field if non-nil, zero value otherwise.

### GetEmergencyRegistrationIndOk

`func (o *Amf3GppAccessRegistration) GetEmergencyRegistrationIndOk() (*bool, bool)`

GetEmergencyRegistrationIndOk returns a tuple with the EmergencyRegistrationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyRegistrationInd

`func (o *Amf3GppAccessRegistration) SetEmergencyRegistrationInd(v bool)`

SetEmergencyRegistrationInd sets EmergencyRegistrationInd field to given value.

### HasEmergencyRegistrationInd

`func (o *Amf3GppAccessRegistration) HasEmergencyRegistrationInd() bool`

HasEmergencyRegistrationInd returns a boolean if a field has been set.

### GetGuami

`func (o *Amf3GppAccessRegistration) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *Amf3GppAccessRegistration) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *Amf3GppAccessRegistration) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetBackupAmfInfo

`func (o *Amf3GppAccessRegistration) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *Amf3GppAccessRegistration) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *Amf3GppAccessRegistration) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *Amf3GppAccessRegistration) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetDrFlag

`func (o *Amf3GppAccessRegistration) GetDrFlag() bool`

GetDrFlag returns the DrFlag field if non-nil, zero value otherwise.

### GetDrFlagOk

`func (o *Amf3GppAccessRegistration) GetDrFlagOk() (*bool, bool)`

GetDrFlagOk returns a tuple with the DrFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrFlag

`func (o *Amf3GppAccessRegistration) SetDrFlag(v bool)`

SetDrFlag sets DrFlag field to given value.

### HasDrFlag

`func (o *Amf3GppAccessRegistration) HasDrFlag() bool`

HasDrFlag returns a boolean if a field has been set.

### GetRatType

`func (o *Amf3GppAccessRegistration) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *Amf3GppAccessRegistration) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *Amf3GppAccessRegistration) SetRatType(v RatType)`

SetRatType sets RatType field to given value.


### GetUrrpIndicator

`func (o *Amf3GppAccessRegistration) GetUrrpIndicator() bool`

GetUrrpIndicator returns the UrrpIndicator field if non-nil, zero value otherwise.

### GetUrrpIndicatorOk

`func (o *Amf3GppAccessRegistration) GetUrrpIndicatorOk() (*bool, bool)`

GetUrrpIndicatorOk returns a tuple with the UrrpIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrrpIndicator

`func (o *Amf3GppAccessRegistration) SetUrrpIndicator(v bool)`

SetUrrpIndicator sets UrrpIndicator field to given value.

### HasUrrpIndicator

`func (o *Amf3GppAccessRegistration) HasUrrpIndicator() bool`

HasUrrpIndicator returns a boolean if a field has been set.

### GetAmfEeSubscriptionId

`func (o *Amf3GppAccessRegistration) GetAmfEeSubscriptionId() string`

GetAmfEeSubscriptionId returns the AmfEeSubscriptionId field if non-nil, zero value otherwise.

### GetAmfEeSubscriptionIdOk

`func (o *Amf3GppAccessRegistration) GetAmfEeSubscriptionIdOk() (*string, bool)`

GetAmfEeSubscriptionIdOk returns a tuple with the AmfEeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfEeSubscriptionId

`func (o *Amf3GppAccessRegistration) SetAmfEeSubscriptionId(v string)`

SetAmfEeSubscriptionId sets AmfEeSubscriptionId field to given value.

### HasAmfEeSubscriptionId

`func (o *Amf3GppAccessRegistration) HasAmfEeSubscriptionId() bool`

HasAmfEeSubscriptionId returns a boolean if a field has been set.

### GetEpsInterworkingInfo

`func (o *Amf3GppAccessRegistration) GetEpsInterworkingInfo() EpsInterworkingInfo`

GetEpsInterworkingInfo returns the EpsInterworkingInfo field if non-nil, zero value otherwise.

### GetEpsInterworkingInfoOk

`func (o *Amf3GppAccessRegistration) GetEpsInterworkingInfoOk() (*EpsInterworkingInfo, bool)`

GetEpsInterworkingInfoOk returns a tuple with the EpsInterworkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInfo

`func (o *Amf3GppAccessRegistration) SetEpsInterworkingInfo(v EpsInterworkingInfo)`

SetEpsInterworkingInfo sets EpsInterworkingInfo field to given value.

### HasEpsInterworkingInfo

`func (o *Amf3GppAccessRegistration) HasEpsInterworkingInfo() bool`

HasEpsInterworkingInfo returns a boolean if a field has been set.

### GetUeSrvccCapability

`func (o *Amf3GppAccessRegistration) GetUeSrvccCapability() bool`

GetUeSrvccCapability returns the UeSrvccCapability field if non-nil, zero value otherwise.

### GetUeSrvccCapabilityOk

`func (o *Amf3GppAccessRegistration) GetUeSrvccCapabilityOk() (*bool, bool)`

GetUeSrvccCapabilityOk returns a tuple with the UeSrvccCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSrvccCapability

`func (o *Amf3GppAccessRegistration) SetUeSrvccCapability(v bool)`

SetUeSrvccCapability sets UeSrvccCapability field to given value.

### HasUeSrvccCapability

`func (o *Amf3GppAccessRegistration) HasUeSrvccCapability() bool`

HasUeSrvccCapability returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *Amf3GppAccessRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *Amf3GppAccessRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *Amf3GppAccessRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *Amf3GppAccessRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetVgmlcAddress

`func (o *Amf3GppAccessRegistration) GetVgmlcAddress() VgmlcAddress`

GetVgmlcAddress returns the VgmlcAddress field if non-nil, zero value otherwise.

### GetVgmlcAddressOk

`func (o *Amf3GppAccessRegistration) GetVgmlcAddressOk() (*VgmlcAddress, bool)`

GetVgmlcAddressOk returns a tuple with the VgmlcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddress

`func (o *Amf3GppAccessRegistration) SetVgmlcAddress(v VgmlcAddress)`

SetVgmlcAddress sets VgmlcAddress field to given value.

### HasVgmlcAddress

`func (o *Amf3GppAccessRegistration) HasVgmlcAddress() bool`

HasVgmlcAddress returns a boolean if a field has been set.

### GetContextInfo

`func (o *Amf3GppAccessRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *Amf3GppAccessRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *Amf3GppAccessRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *Amf3GppAccessRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetNoEeSubscriptionInd

`func (o *Amf3GppAccessRegistration) GetNoEeSubscriptionInd() bool`

GetNoEeSubscriptionInd returns the NoEeSubscriptionInd field if non-nil, zero value otherwise.

### GetNoEeSubscriptionIndOk

`func (o *Amf3GppAccessRegistration) GetNoEeSubscriptionIndOk() (*bool, bool)`

GetNoEeSubscriptionIndOk returns a tuple with the NoEeSubscriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEeSubscriptionInd

`func (o *Amf3GppAccessRegistration) SetNoEeSubscriptionInd(v bool)`

SetNoEeSubscriptionInd sets NoEeSubscriptionInd field to given value.

### HasNoEeSubscriptionInd

`func (o *Amf3GppAccessRegistration) HasNoEeSubscriptionInd() bool`

HasNoEeSubscriptionInd returns a boolean if a field has been set.

### GetSupi

`func (o *Amf3GppAccessRegistration) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *Amf3GppAccessRegistration) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *Amf3GppAccessRegistration) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *Amf3GppAccessRegistration) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeReachableInd

`func (o *Amf3GppAccessRegistration) GetUeReachableInd() UeReachableInd`

GetUeReachableInd returns the UeReachableInd field if non-nil, zero value otherwise.

### GetUeReachableIndOk

`func (o *Amf3GppAccessRegistration) GetUeReachableIndOk() (*UeReachableInd, bool)`

GetUeReachableIndOk returns a tuple with the UeReachableInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeReachableInd

`func (o *Amf3GppAccessRegistration) SetUeReachableInd(v UeReachableInd)`

SetUeReachableInd sets UeReachableInd field to given value.

### HasUeReachableInd

`func (o *Amf3GppAccessRegistration) HasUeReachableInd() bool`

HasUeReachableInd returns a boolean if a field has been set.

### GetReRegistrationRequired

`func (o *Amf3GppAccessRegistration) GetReRegistrationRequired() bool`

GetReRegistrationRequired returns the ReRegistrationRequired field if non-nil, zero value otherwise.

### GetReRegistrationRequiredOk

`func (o *Amf3GppAccessRegistration) GetReRegistrationRequiredOk() (*bool, bool)`

GetReRegistrationRequiredOk returns a tuple with the ReRegistrationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegistrationRequired

`func (o *Amf3GppAccessRegistration) SetReRegistrationRequired(v bool)`

SetReRegistrationRequired sets ReRegistrationRequired field to given value.

### HasReRegistrationRequired

`func (o *Amf3GppAccessRegistration) HasReRegistrationRequired() bool`

HasReRegistrationRequired returns a boolean if a field has been set.

### GetAdminDeregSubWithdrawn

`func (o *Amf3GppAccessRegistration) GetAdminDeregSubWithdrawn() bool`

GetAdminDeregSubWithdrawn returns the AdminDeregSubWithdrawn field if non-nil, zero value otherwise.

### GetAdminDeregSubWithdrawnOk

`func (o *Amf3GppAccessRegistration) GetAdminDeregSubWithdrawnOk() (*bool, bool)`

GetAdminDeregSubWithdrawnOk returns a tuple with the AdminDeregSubWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDeregSubWithdrawn

`func (o *Amf3GppAccessRegistration) SetAdminDeregSubWithdrawn(v bool)`

SetAdminDeregSubWithdrawn sets AdminDeregSubWithdrawn field to given value.

### HasAdminDeregSubWithdrawn

`func (o *Amf3GppAccessRegistration) HasAdminDeregSubWithdrawn() bool`

HasAdminDeregSubWithdrawn returns a boolean if a field has been set.

### GetDataRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) GetDataRestorationCallbackUri() string`

GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field if non-nil, zero value otherwise.

### GetDataRestorationCallbackUriOk

`func (o *Amf3GppAccessRegistration) GetDataRestorationCallbackUriOk() (*string, bool)`

GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) SetDataRestorationCallbackUri(v string)`

SetDataRestorationCallbackUri sets DataRestorationCallbackUri field to given value.

### HasDataRestorationCallbackUri

`func (o *Amf3GppAccessRegistration) HasDataRestorationCallbackUri() bool`

HasDataRestorationCallbackUri returns a boolean if a field has been set.

### GetResetIds

`func (o *Amf3GppAccessRegistration) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *Amf3GppAccessRegistration) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *Amf3GppAccessRegistration) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *Amf3GppAccessRegistration) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *Amf3GppAccessRegistration) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *Amf3GppAccessRegistration) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *Amf3GppAccessRegistration) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *Amf3GppAccessRegistration) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.

### GetUeMINTCapability

`func (o *Amf3GppAccessRegistration) GetUeMINTCapability() bool`

GetUeMINTCapability returns the UeMINTCapability field if non-nil, zero value otherwise.

### GetUeMINTCapabilityOk

`func (o *Amf3GppAccessRegistration) GetUeMINTCapabilityOk() (*bool, bool)`

GetUeMINTCapabilityOk returns a tuple with the UeMINTCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMINTCapability

`func (o *Amf3GppAccessRegistration) SetUeMINTCapability(v bool)`

SetUeMINTCapability sets UeMINTCapability field to given value.

### HasUeMINTCapability

`func (o *Amf3GppAccessRegistration) HasUeMINTCapability() bool`

HasUeMINTCapability returns a boolean if a field has been set.

### GetSorSnpnSiSupported

`func (o *Amf3GppAccessRegistration) GetSorSnpnSiSupported() bool`

GetSorSnpnSiSupported returns the SorSnpnSiSupported field if non-nil, zero value otherwise.

### GetSorSnpnSiSupportedOk

`func (o *Amf3GppAccessRegistration) GetSorSnpnSiSupportedOk() (*bool, bool)`

GetSorSnpnSiSupportedOk returns a tuple with the SorSnpnSiSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSnpnSiSupported

`func (o *Amf3GppAccessRegistration) SetSorSnpnSiSupported(v bool)`

SetSorSnpnSiSupported sets SorSnpnSiSupported field to given value.

### HasSorSnpnSiSupported

`func (o *Amf3GppAccessRegistration) HasSorSnpnSiSupported() bool`

HasSorSnpnSiSupported returns a boolean if a field has been set.

### GetUdrRestartInd

`func (o *Amf3GppAccessRegistration) GetUdrRestartInd() bool`

GetUdrRestartInd returns the UdrRestartInd field if non-nil, zero value otherwise.

### GetUdrRestartIndOk

`func (o *Amf3GppAccessRegistration) GetUdrRestartIndOk() (*bool, bool)`

GetUdrRestartIndOk returns a tuple with the UdrRestartInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrRestartInd

`func (o *Amf3GppAccessRegistration) SetUdrRestartInd(v bool)`

SetUdrRestartInd sets UdrRestartInd field to given value.

### HasUdrRestartInd

`func (o *Amf3GppAccessRegistration) HasUdrRestartInd() bool`

HasUdrRestartInd returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *Amf3GppAccessRegistration) GetLastSynchronizationTime() time.Time`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *Amf3GppAccessRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *Amf3GppAccessRegistration) SetLastSynchronizationTime(v time.Time)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *Amf3GppAccessRegistration) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


