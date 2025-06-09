# PpData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationCharacteristics** | Pointer to [**NullableCommunicationCharacteristics**](CommunicationCharacteristics.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ExpectedUeBehaviourParameters** | Pointer to [**NullableExpectedUeBehaviour**](ExpectedUeBehaviour.md) |  | [optional] 
**ExpectedUeBehaviourExtension** | Pointer to [**NullableExpectedUeBehaviourExtension**](ExpectedUeBehaviourExtension.md) |  | [optional] 
**EcRestriction** | Pointer to [**NullableEcRestriction**](EcRestriction.md) |  | [optional] 
**AcsInfo** | Pointer to [**NullableAcsInfo**](AcsInfo.md) |  | [optional] 
**StnSr** | Pointer to **NullableString** | String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003 with the OpenAPI &#39;nullable: true&#39; property.   | [optional] 
**LcsPrivacy** | Pointer to [**NullableLcsPrivacy**](LcsPrivacy.md) |  | [optional] 
**SorInfo** | Pointer to [**SorInfo**](SorInfo.md) |  | [optional] 
**Var5mbsAuthorizationInfo** | Pointer to [**NullableModel5MbsAuthorizationInfo**](Model5MbsAuthorizationInfo.md) |  | [optional] 
**DnnSnssaiSpecificGroup** | Pointer to [**NullableDnnSnssaiSpecificGroup**](DnnSnssaiSpecificGroup.md) |  | [optional] 
**MbsAssistanceInfo** | Pointer to [**MbsAssistanceInfo**](MbsAssistanceInfo.md) |  | [optional] 
**AppSpecificExpectedUeBehaviour** | Pointer to [**NullableAppSpecificExpectedUeBehaviour**](AppSpecificExpectedUeBehaviour.md) |  | [optional] 
**SliceUsageControlInfos** | Pointer to [**[]SliceUsageControlInfo**](SliceUsageControlInfo.md) |  | [optional] 

## Methods

### NewPpData

`func NewPpData() *PpData`

NewPpData instantiates a new PpData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpDataWithDefaults

`func NewPpDataWithDefaults() *PpData`

NewPpDataWithDefaults instantiates a new PpData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationCharacteristics

`func (o *PpData) GetCommunicationCharacteristics() CommunicationCharacteristics`

GetCommunicationCharacteristics returns the CommunicationCharacteristics field if non-nil, zero value otherwise.

### GetCommunicationCharacteristicsOk

`func (o *PpData) GetCommunicationCharacteristicsOk() (*CommunicationCharacteristics, bool)`

GetCommunicationCharacteristicsOk returns a tuple with the CommunicationCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationCharacteristics

`func (o *PpData) SetCommunicationCharacteristics(v CommunicationCharacteristics)`

SetCommunicationCharacteristics sets CommunicationCharacteristics field to given value.

### HasCommunicationCharacteristics

`func (o *PpData) HasCommunicationCharacteristics() bool`

HasCommunicationCharacteristics returns a boolean if a field has been set.

### SetCommunicationCharacteristicsNil

`func (o *PpData) SetCommunicationCharacteristicsNil(b bool)`

 SetCommunicationCharacteristicsNil sets the value for CommunicationCharacteristics to be an explicit nil

### UnsetCommunicationCharacteristics
`func (o *PpData) UnsetCommunicationCharacteristics()`

UnsetCommunicationCharacteristics ensures that no value is present for CommunicationCharacteristics, not even an explicit nil
### GetSupportedFeatures

`func (o *PpData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PpData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PpData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PpData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetExpectedUeBehaviourParameters

`func (o *PpData) GetExpectedUeBehaviourParameters() ExpectedUeBehaviour`

GetExpectedUeBehaviourParameters returns the ExpectedUeBehaviourParameters field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourParametersOk

`func (o *PpData) GetExpectedUeBehaviourParametersOk() (*ExpectedUeBehaviour, bool)`

GetExpectedUeBehaviourParametersOk returns a tuple with the ExpectedUeBehaviourParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourParameters

`func (o *PpData) SetExpectedUeBehaviourParameters(v ExpectedUeBehaviour)`

SetExpectedUeBehaviourParameters sets ExpectedUeBehaviourParameters field to given value.

### HasExpectedUeBehaviourParameters

`func (o *PpData) HasExpectedUeBehaviourParameters() bool`

HasExpectedUeBehaviourParameters returns a boolean if a field has been set.

### SetExpectedUeBehaviourParametersNil

`func (o *PpData) SetExpectedUeBehaviourParametersNil(b bool)`

 SetExpectedUeBehaviourParametersNil sets the value for ExpectedUeBehaviourParameters to be an explicit nil

### UnsetExpectedUeBehaviourParameters
`func (o *PpData) UnsetExpectedUeBehaviourParameters()`

UnsetExpectedUeBehaviourParameters ensures that no value is present for ExpectedUeBehaviourParameters, not even an explicit nil
### GetExpectedUeBehaviourExtension

`func (o *PpData) GetExpectedUeBehaviourExtension() ExpectedUeBehaviourExtension`

GetExpectedUeBehaviourExtension returns the ExpectedUeBehaviourExtension field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourExtensionOk

`func (o *PpData) GetExpectedUeBehaviourExtensionOk() (*ExpectedUeBehaviourExtension, bool)`

GetExpectedUeBehaviourExtensionOk returns a tuple with the ExpectedUeBehaviourExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourExtension

`func (o *PpData) SetExpectedUeBehaviourExtension(v ExpectedUeBehaviourExtension)`

SetExpectedUeBehaviourExtension sets ExpectedUeBehaviourExtension field to given value.

### HasExpectedUeBehaviourExtension

`func (o *PpData) HasExpectedUeBehaviourExtension() bool`

HasExpectedUeBehaviourExtension returns a boolean if a field has been set.

### SetExpectedUeBehaviourExtensionNil

`func (o *PpData) SetExpectedUeBehaviourExtensionNil(b bool)`

 SetExpectedUeBehaviourExtensionNil sets the value for ExpectedUeBehaviourExtension to be an explicit nil

### UnsetExpectedUeBehaviourExtension
`func (o *PpData) UnsetExpectedUeBehaviourExtension()`

UnsetExpectedUeBehaviourExtension ensures that no value is present for ExpectedUeBehaviourExtension, not even an explicit nil
### GetEcRestriction

`func (o *PpData) GetEcRestriction() EcRestriction`

GetEcRestriction returns the EcRestriction field if non-nil, zero value otherwise.

### GetEcRestrictionOk

`func (o *PpData) GetEcRestrictionOk() (*EcRestriction, bool)`

GetEcRestrictionOk returns a tuple with the EcRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestriction

`func (o *PpData) SetEcRestriction(v EcRestriction)`

SetEcRestriction sets EcRestriction field to given value.

### HasEcRestriction

`func (o *PpData) HasEcRestriction() bool`

HasEcRestriction returns a boolean if a field has been set.

### SetEcRestrictionNil

`func (o *PpData) SetEcRestrictionNil(b bool)`

 SetEcRestrictionNil sets the value for EcRestriction to be an explicit nil

### UnsetEcRestriction
`func (o *PpData) UnsetEcRestriction()`

UnsetEcRestriction ensures that no value is present for EcRestriction, not even an explicit nil
### GetAcsInfo

`func (o *PpData) GetAcsInfo() AcsInfo`

GetAcsInfo returns the AcsInfo field if non-nil, zero value otherwise.

### GetAcsInfoOk

`func (o *PpData) GetAcsInfoOk() (*AcsInfo, bool)`

GetAcsInfoOk returns a tuple with the AcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfo

`func (o *PpData) SetAcsInfo(v AcsInfo)`

SetAcsInfo sets AcsInfo field to given value.

### HasAcsInfo

`func (o *PpData) HasAcsInfo() bool`

HasAcsInfo returns a boolean if a field has been set.

### SetAcsInfoNil

`func (o *PpData) SetAcsInfoNil(b bool)`

 SetAcsInfoNil sets the value for AcsInfo to be an explicit nil

### UnsetAcsInfo
`func (o *PpData) UnsetAcsInfo()`

UnsetAcsInfo ensures that no value is present for AcsInfo, not even an explicit nil
### GetStnSr

`func (o *PpData) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *PpData) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *PpData) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *PpData) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### SetStnSrNil

`func (o *PpData) SetStnSrNil(b bool)`

 SetStnSrNil sets the value for StnSr to be an explicit nil

### UnsetStnSr
`func (o *PpData) UnsetStnSr()`

UnsetStnSr ensures that no value is present for StnSr, not even an explicit nil
### GetLcsPrivacy

`func (o *PpData) GetLcsPrivacy() LcsPrivacy`

GetLcsPrivacy returns the LcsPrivacy field if non-nil, zero value otherwise.

### GetLcsPrivacyOk

`func (o *PpData) GetLcsPrivacyOk() (*LcsPrivacy, bool)`

GetLcsPrivacyOk returns a tuple with the LcsPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacy

`func (o *PpData) SetLcsPrivacy(v LcsPrivacy)`

SetLcsPrivacy sets LcsPrivacy field to given value.

### HasLcsPrivacy

`func (o *PpData) HasLcsPrivacy() bool`

HasLcsPrivacy returns a boolean if a field has been set.

### SetLcsPrivacyNil

`func (o *PpData) SetLcsPrivacyNil(b bool)`

 SetLcsPrivacyNil sets the value for LcsPrivacy to be an explicit nil

### UnsetLcsPrivacy
`func (o *PpData) UnsetLcsPrivacy()`

UnsetLcsPrivacy ensures that no value is present for LcsPrivacy, not even an explicit nil
### GetSorInfo

`func (o *PpData) GetSorInfo() SorInfo`

GetSorInfo returns the SorInfo field if non-nil, zero value otherwise.

### GetSorInfoOk

`func (o *PpData) GetSorInfoOk() (*SorInfo, bool)`

GetSorInfoOk returns a tuple with the SorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfo

`func (o *PpData) SetSorInfo(v SorInfo)`

SetSorInfo sets SorInfo field to given value.

### HasSorInfo

`func (o *PpData) HasSorInfo() bool`

HasSorInfo returns a boolean if a field has been set.

### GetVar5mbsAuthorizationInfo

`func (o *PpData) GetVar5mbsAuthorizationInfo() Model5MbsAuthorizationInfo`

GetVar5mbsAuthorizationInfo returns the Var5mbsAuthorizationInfo field if non-nil, zero value otherwise.

### GetVar5mbsAuthorizationInfoOk

`func (o *PpData) GetVar5mbsAuthorizationInfoOk() (*Model5MbsAuthorizationInfo, bool)`

GetVar5mbsAuthorizationInfoOk returns a tuple with the Var5mbsAuthorizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5mbsAuthorizationInfo

`func (o *PpData) SetVar5mbsAuthorizationInfo(v Model5MbsAuthorizationInfo)`

SetVar5mbsAuthorizationInfo sets Var5mbsAuthorizationInfo field to given value.

### HasVar5mbsAuthorizationInfo

`func (o *PpData) HasVar5mbsAuthorizationInfo() bool`

HasVar5mbsAuthorizationInfo returns a boolean if a field has been set.

### SetVar5mbsAuthorizationInfoNil

`func (o *PpData) SetVar5mbsAuthorizationInfoNil(b bool)`

 SetVar5mbsAuthorizationInfoNil sets the value for Var5mbsAuthorizationInfo to be an explicit nil

### UnsetVar5mbsAuthorizationInfo
`func (o *PpData) UnsetVar5mbsAuthorizationInfo()`

UnsetVar5mbsAuthorizationInfo ensures that no value is present for Var5mbsAuthorizationInfo, not even an explicit nil
### GetDnnSnssaiSpecificGroup

`func (o *PpData) GetDnnSnssaiSpecificGroup() DnnSnssaiSpecificGroup`

GetDnnSnssaiSpecificGroup returns the DnnSnssaiSpecificGroup field if non-nil, zero value otherwise.

### GetDnnSnssaiSpecificGroupOk

`func (o *PpData) GetDnnSnssaiSpecificGroupOk() (*DnnSnssaiSpecificGroup, bool)`

GetDnnSnssaiSpecificGroupOk returns a tuple with the DnnSnssaiSpecificGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiSpecificGroup

`func (o *PpData) SetDnnSnssaiSpecificGroup(v DnnSnssaiSpecificGroup)`

SetDnnSnssaiSpecificGroup sets DnnSnssaiSpecificGroup field to given value.

### HasDnnSnssaiSpecificGroup

`func (o *PpData) HasDnnSnssaiSpecificGroup() bool`

HasDnnSnssaiSpecificGroup returns a boolean if a field has been set.

### SetDnnSnssaiSpecificGroupNil

`func (o *PpData) SetDnnSnssaiSpecificGroupNil(b bool)`

 SetDnnSnssaiSpecificGroupNil sets the value for DnnSnssaiSpecificGroup to be an explicit nil

### UnsetDnnSnssaiSpecificGroup
`func (o *PpData) UnsetDnnSnssaiSpecificGroup()`

UnsetDnnSnssaiSpecificGroup ensures that no value is present for DnnSnssaiSpecificGroup, not even an explicit nil
### GetMbsAssistanceInfo

`func (o *PpData) GetMbsAssistanceInfo() MbsAssistanceInfo`

GetMbsAssistanceInfo returns the MbsAssistanceInfo field if non-nil, zero value otherwise.

### GetMbsAssistanceInfoOk

`func (o *PpData) GetMbsAssistanceInfoOk() (*MbsAssistanceInfo, bool)`

GetMbsAssistanceInfoOk returns a tuple with the MbsAssistanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsAssistanceInfo

`func (o *PpData) SetMbsAssistanceInfo(v MbsAssistanceInfo)`

SetMbsAssistanceInfo sets MbsAssistanceInfo field to given value.

### HasMbsAssistanceInfo

`func (o *PpData) HasMbsAssistanceInfo() bool`

HasMbsAssistanceInfo returns a boolean if a field has been set.

### GetAppSpecificExpectedUeBehaviour

`func (o *PpData) GetAppSpecificExpectedUeBehaviour() AppSpecificExpectedUeBehaviour`

GetAppSpecificExpectedUeBehaviour returns the AppSpecificExpectedUeBehaviour field if non-nil, zero value otherwise.

### GetAppSpecificExpectedUeBehaviourOk

`func (o *PpData) GetAppSpecificExpectedUeBehaviourOk() (*AppSpecificExpectedUeBehaviour, bool)`

GetAppSpecificExpectedUeBehaviourOk returns a tuple with the AppSpecificExpectedUeBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpecificExpectedUeBehaviour

`func (o *PpData) SetAppSpecificExpectedUeBehaviour(v AppSpecificExpectedUeBehaviour)`

SetAppSpecificExpectedUeBehaviour sets AppSpecificExpectedUeBehaviour field to given value.

### HasAppSpecificExpectedUeBehaviour

`func (o *PpData) HasAppSpecificExpectedUeBehaviour() bool`

HasAppSpecificExpectedUeBehaviour returns a boolean if a field has been set.

### SetAppSpecificExpectedUeBehaviourNil

`func (o *PpData) SetAppSpecificExpectedUeBehaviourNil(b bool)`

 SetAppSpecificExpectedUeBehaviourNil sets the value for AppSpecificExpectedUeBehaviour to be an explicit nil

### UnsetAppSpecificExpectedUeBehaviour
`func (o *PpData) UnsetAppSpecificExpectedUeBehaviour()`

UnsetAppSpecificExpectedUeBehaviour ensures that no value is present for AppSpecificExpectedUeBehaviour, not even an explicit nil
### GetSliceUsageControlInfos

`func (o *PpData) GetSliceUsageControlInfos() []SliceUsageControlInfo`

GetSliceUsageControlInfos returns the SliceUsageControlInfos field if non-nil, zero value otherwise.

### GetSliceUsageControlInfosOk

`func (o *PpData) GetSliceUsageControlInfosOk() (*[]SliceUsageControlInfo, bool)`

GetSliceUsageControlInfosOk returns a tuple with the SliceUsageControlInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceUsageControlInfos

`func (o *PpData) SetSliceUsageControlInfos(v []SliceUsageControlInfo)`

SetSliceUsageControlInfos sets SliceUsageControlInfos field to given value.

### HasSliceUsageControlInfos

`func (o *PpData) HasSliceUsageControlInfos() bool`

HasSliceUsageControlInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


