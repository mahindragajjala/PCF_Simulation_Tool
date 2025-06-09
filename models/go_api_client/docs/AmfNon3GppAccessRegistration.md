# AmfNon3GppAccessRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PurgeFlag** | Pointer to **bool** | This flag indicates whether an AMF is deregistered. This attribute may be included in notifications sent by the UDR to the UDM if purgeFlag is also set to true in the same notification.  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**ImsVoPs** | [**ImsVoPs**](ImsVoPs.md) |  | 
**DeregCallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AmfServiceNameDereg** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**PcscfRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AmfServiceNamePcscfRest** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**Guami** | [**Guami**](Guami.md) |  | 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**RatType** | [**RatType**](RatType.md) |  | 
**UrrpIndicator** | Pointer to **bool** |  | [optional] 
**AmfEeSubscriptionId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**RegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**VgmlcAddress** | Pointer to [**VgmlcAddress**](VgmlcAddress.md) |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**NoEeSubscriptionInd** | Pointer to **bool** |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**ReRegistrationRequired** | Pointer to **bool** |  | [optional] 
**AdminDeregSubWithdrawn** | Pointer to **bool** |  | [optional] 
**DataRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]
**SorSnpnSiSupported** | Pointer to **bool** |  | [optional] [default to false]
**UdrRestartInd** | Pointer to **bool** |  | [optional] [default to false]
**LastSynchronizationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAmfNon3GppAccessRegistration

`func NewAmfNon3GppAccessRegistration(amfInstanceId string, imsVoPs ImsVoPs, deregCallbackUri string, guami Guami, ratType RatType, ) *AmfNon3GppAccessRegistration`

NewAmfNon3GppAccessRegistration instantiates a new AmfNon3GppAccessRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfNon3GppAccessRegistrationWithDefaults

`func NewAmfNon3GppAccessRegistrationWithDefaults() *AmfNon3GppAccessRegistration`

NewAmfNon3GppAccessRegistrationWithDefaults instantiates a new AmfNon3GppAccessRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *AmfNon3GppAccessRegistration) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *AmfNon3GppAccessRegistration) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *AmfNon3GppAccessRegistration) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AmfNon3GppAccessRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPurgeFlag

`func (o *AmfNon3GppAccessRegistration) GetPurgeFlag() bool`

GetPurgeFlag returns the PurgeFlag field if non-nil, zero value otherwise.

### GetPurgeFlagOk

`func (o *AmfNon3GppAccessRegistration) GetPurgeFlagOk() (*bool, bool)`

GetPurgeFlagOk returns a tuple with the PurgeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeFlag

`func (o *AmfNon3GppAccessRegistration) SetPurgeFlag(v bool)`

SetPurgeFlag sets PurgeFlag field to given value.

### HasPurgeFlag

`func (o *AmfNon3GppAccessRegistration) HasPurgeFlag() bool`

HasPurgeFlag returns a boolean if a field has been set.

### GetPei

`func (o *AmfNon3GppAccessRegistration) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AmfNon3GppAccessRegistration) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AmfNon3GppAccessRegistration) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AmfNon3GppAccessRegistration) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetImsVoPs

`func (o *AmfNon3GppAccessRegistration) GetImsVoPs() ImsVoPs`

GetImsVoPs returns the ImsVoPs field if non-nil, zero value otherwise.

### GetImsVoPsOk

`func (o *AmfNon3GppAccessRegistration) GetImsVoPsOk() (*ImsVoPs, bool)`

GetImsVoPsOk returns a tuple with the ImsVoPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVoPs

`func (o *AmfNon3GppAccessRegistration) SetImsVoPs(v ImsVoPs)`

SetImsVoPs sets ImsVoPs field to given value.


### GetDeregCallbackUri

`func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *AmfNon3GppAccessRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.


### GetAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDereg() ServiceName`

GetAmfServiceNameDereg returns the AmfServiceNameDereg field if non-nil, zero value otherwise.

### GetAmfServiceNameDeregOk

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDeregOk() (*ServiceName, bool)`

GetAmfServiceNameDeregOk returns a tuple with the AmfServiceNameDereg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) SetAmfServiceNameDereg(v ServiceName)`

SetAmfServiceNameDereg sets AmfServiceNameDereg field to given value.

### HasAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) HasAmfServiceNameDereg() bool`

HasAmfServiceNameDereg returns a boolean if a field has been set.

### GetPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUri() string`

GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field if non-nil, zero value otherwise.

### GetPcscfRestorationCallbackUriOk

`func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool)`

GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) SetPcscfRestorationCallbackUri(v string)`

SetPcscfRestorationCallbackUri sets PcscfRestorationCallbackUri field to given value.

### HasPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) HasPcscfRestorationCallbackUri() bool`

HasPcscfRestorationCallbackUri returns a boolean if a field has been set.

### GetAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRest() ServiceName`

GetAmfServiceNamePcscfRest returns the AmfServiceNamePcscfRest field if non-nil, zero value otherwise.

### GetAmfServiceNamePcscfRestOk

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRestOk() (*ServiceName, bool)`

GetAmfServiceNamePcscfRestOk returns a tuple with the AmfServiceNamePcscfRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) SetAmfServiceNamePcscfRest(v ServiceName)`

SetAmfServiceNamePcscfRest sets AmfServiceNamePcscfRest field to given value.

### HasAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) HasAmfServiceNamePcscfRest() bool`

HasAmfServiceNamePcscfRest returns a boolean if a field has been set.

### GetGuami

`func (o *AmfNon3GppAccessRegistration) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *AmfNon3GppAccessRegistration) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *AmfNon3GppAccessRegistration) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetRatType

`func (o *AmfNon3GppAccessRegistration) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *AmfNon3GppAccessRegistration) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *AmfNon3GppAccessRegistration) SetRatType(v RatType)`

SetRatType sets RatType field to given value.


### GetUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) GetUrrpIndicator() bool`

GetUrrpIndicator returns the UrrpIndicator field if non-nil, zero value otherwise.

### GetUrrpIndicatorOk

`func (o *AmfNon3GppAccessRegistration) GetUrrpIndicatorOk() (*bool, bool)`

GetUrrpIndicatorOk returns a tuple with the UrrpIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) SetUrrpIndicator(v bool)`

SetUrrpIndicator sets UrrpIndicator field to given value.

### HasUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) HasUrrpIndicator() bool`

HasUrrpIndicator returns a boolean if a field has been set.

### GetAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionId() string`

GetAmfEeSubscriptionId returns the AmfEeSubscriptionId field if non-nil, zero value otherwise.

### GetAmfEeSubscriptionIdOk

`func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionIdOk() (*string, bool)`

GetAmfEeSubscriptionIdOk returns a tuple with the AmfEeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) SetAmfEeSubscriptionId(v string)`

SetAmfEeSubscriptionId sets AmfEeSubscriptionId field to given value.

### HasAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) HasAmfEeSubscriptionId() bool`

HasAmfEeSubscriptionId returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *AmfNon3GppAccessRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *AmfNon3GppAccessRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *AmfNon3GppAccessRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *AmfNon3GppAccessRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) GetVgmlcAddress() VgmlcAddress`

GetVgmlcAddress returns the VgmlcAddress field if non-nil, zero value otherwise.

### GetVgmlcAddressOk

`func (o *AmfNon3GppAccessRegistration) GetVgmlcAddressOk() (*VgmlcAddress, bool)`

GetVgmlcAddressOk returns a tuple with the VgmlcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) SetVgmlcAddress(v VgmlcAddress)`

SetVgmlcAddress sets VgmlcAddress field to given value.

### HasVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) HasVgmlcAddress() bool`

HasVgmlcAddress returns a boolean if a field has been set.

### GetContextInfo

`func (o *AmfNon3GppAccessRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *AmfNon3GppAccessRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *AmfNon3GppAccessRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *AmfNon3GppAccessRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionInd() bool`

GetNoEeSubscriptionInd returns the NoEeSubscriptionInd field if non-nil, zero value otherwise.

### GetNoEeSubscriptionIndOk

`func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionIndOk() (*bool, bool)`

GetNoEeSubscriptionIndOk returns a tuple with the NoEeSubscriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) SetNoEeSubscriptionInd(v bool)`

SetNoEeSubscriptionInd sets NoEeSubscriptionInd field to given value.

### HasNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) HasNoEeSubscriptionInd() bool`

HasNoEeSubscriptionInd returns a boolean if a field has been set.

### GetSupi

`func (o *AmfNon3GppAccessRegistration) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AmfNon3GppAccessRegistration) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AmfNon3GppAccessRegistration) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AmfNon3GppAccessRegistration) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetReRegistrationRequired

`func (o *AmfNon3GppAccessRegistration) GetReRegistrationRequired() bool`

GetReRegistrationRequired returns the ReRegistrationRequired field if non-nil, zero value otherwise.

### GetReRegistrationRequiredOk

`func (o *AmfNon3GppAccessRegistration) GetReRegistrationRequiredOk() (*bool, bool)`

GetReRegistrationRequiredOk returns a tuple with the ReRegistrationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegistrationRequired

`func (o *AmfNon3GppAccessRegistration) SetReRegistrationRequired(v bool)`

SetReRegistrationRequired sets ReRegistrationRequired field to given value.

### HasReRegistrationRequired

`func (o *AmfNon3GppAccessRegistration) HasReRegistrationRequired() bool`

HasReRegistrationRequired returns a boolean if a field has been set.

### GetAdminDeregSubWithdrawn

`func (o *AmfNon3GppAccessRegistration) GetAdminDeregSubWithdrawn() bool`

GetAdminDeregSubWithdrawn returns the AdminDeregSubWithdrawn field if non-nil, zero value otherwise.

### GetAdminDeregSubWithdrawnOk

`func (o *AmfNon3GppAccessRegistration) GetAdminDeregSubWithdrawnOk() (*bool, bool)`

GetAdminDeregSubWithdrawnOk returns a tuple with the AdminDeregSubWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDeregSubWithdrawn

`func (o *AmfNon3GppAccessRegistration) SetAdminDeregSubWithdrawn(v bool)`

SetAdminDeregSubWithdrawn sets AdminDeregSubWithdrawn field to given value.

### HasAdminDeregSubWithdrawn

`func (o *AmfNon3GppAccessRegistration) HasAdminDeregSubWithdrawn() bool`

HasAdminDeregSubWithdrawn returns a boolean if a field has been set.

### GetDataRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) GetDataRestorationCallbackUri() string`

GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field if non-nil, zero value otherwise.

### GetDataRestorationCallbackUriOk

`func (o *AmfNon3GppAccessRegistration) GetDataRestorationCallbackUriOk() (*string, bool)`

GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) SetDataRestorationCallbackUri(v string)`

SetDataRestorationCallbackUri sets DataRestorationCallbackUri field to given value.

### HasDataRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) HasDataRestorationCallbackUri() bool`

HasDataRestorationCallbackUri returns a boolean if a field has been set.

### GetResetIds

`func (o *AmfNon3GppAccessRegistration) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *AmfNon3GppAccessRegistration) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *AmfNon3GppAccessRegistration) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *AmfNon3GppAccessRegistration) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *AmfNon3GppAccessRegistration) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *AmfNon3GppAccessRegistration) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *AmfNon3GppAccessRegistration) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *AmfNon3GppAccessRegistration) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.

### GetSorSnpnSiSupported

`func (o *AmfNon3GppAccessRegistration) GetSorSnpnSiSupported() bool`

GetSorSnpnSiSupported returns the SorSnpnSiSupported field if non-nil, zero value otherwise.

### GetSorSnpnSiSupportedOk

`func (o *AmfNon3GppAccessRegistration) GetSorSnpnSiSupportedOk() (*bool, bool)`

GetSorSnpnSiSupportedOk returns a tuple with the SorSnpnSiSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSnpnSiSupported

`func (o *AmfNon3GppAccessRegistration) SetSorSnpnSiSupported(v bool)`

SetSorSnpnSiSupported sets SorSnpnSiSupported field to given value.

### HasSorSnpnSiSupported

`func (o *AmfNon3GppAccessRegistration) HasSorSnpnSiSupported() bool`

HasSorSnpnSiSupported returns a boolean if a field has been set.

### GetUdrRestartInd

`func (o *AmfNon3GppAccessRegistration) GetUdrRestartInd() bool`

GetUdrRestartInd returns the UdrRestartInd field if non-nil, zero value otherwise.

### GetUdrRestartIndOk

`func (o *AmfNon3GppAccessRegistration) GetUdrRestartIndOk() (*bool, bool)`

GetUdrRestartIndOk returns a tuple with the UdrRestartInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrRestartInd

`func (o *AmfNon3GppAccessRegistration) SetUdrRestartInd(v bool)`

SetUdrRestartInd sets UdrRestartInd field to given value.

### HasUdrRestartInd

`func (o *AmfNon3GppAccessRegistration) HasUdrRestartInd() bool`

HasUdrRestartInd returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *AmfNon3GppAccessRegistration) GetLastSynchronizationTime() time.Time`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *AmfNon3GppAccessRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *AmfNon3GppAccessRegistration) SetLastSynchronizationTime(v time.Time)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *AmfNon3GppAccessRegistration) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


