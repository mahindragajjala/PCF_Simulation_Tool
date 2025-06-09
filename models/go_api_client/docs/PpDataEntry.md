# PpDataEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationCharacteristics** | Pointer to [**NullableCommunicationCharacteristicsAF**](CommunicationCharacteristicsAF.md) |  | [optional] 
**ReferenceId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**EcsAddrConfigInfo** | Pointer to [**NullableEcsAddrConfigInfo**](EcsAddrConfigInfo.md) |  | [optional] 
**AdditionalEcsAddrConfigInfos** | Pointer to [**[]EcsAddrConfigInfo**](EcsAddrConfigInfo.md) |  | [optional] 
**EcsAddrConfigInfoPerPlmn** | Pointer to [**map[string][]EcsAddrConfigInfo**](array.md) | A map (list of key-value pairs) where the key of the map is the serving PLMN id; and the value is an array of EcsAddrConfigInfo for that serving PLMN.  | [optional] 
**EcRestriction** | Pointer to [**NullableEcRestriction**](EcRestriction.md) |  | [optional] 
**SliceUsageControlInfos** | Pointer to [**[]SliceUsageControlInfo**](SliceUsageControlInfo.md) |  | [optional] 

## Methods

### NewPpDataEntry

`func NewPpDataEntry() *PpDataEntry`

NewPpDataEntry instantiates a new PpDataEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpDataEntryWithDefaults

`func NewPpDataEntryWithDefaults() *PpDataEntry`

NewPpDataEntryWithDefaults instantiates a new PpDataEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationCharacteristics

`func (o *PpDataEntry) GetCommunicationCharacteristics() CommunicationCharacteristicsAF`

GetCommunicationCharacteristics returns the CommunicationCharacteristics field if non-nil, zero value otherwise.

### GetCommunicationCharacteristicsOk

`func (o *PpDataEntry) GetCommunicationCharacteristicsOk() (*CommunicationCharacteristicsAF, bool)`

GetCommunicationCharacteristicsOk returns a tuple with the CommunicationCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationCharacteristics

`func (o *PpDataEntry) SetCommunicationCharacteristics(v CommunicationCharacteristicsAF)`

SetCommunicationCharacteristics sets CommunicationCharacteristics field to given value.

### HasCommunicationCharacteristics

`func (o *PpDataEntry) HasCommunicationCharacteristics() bool`

HasCommunicationCharacteristics returns a boolean if a field has been set.

### SetCommunicationCharacteristicsNil

`func (o *PpDataEntry) SetCommunicationCharacteristicsNil(b bool)`

 SetCommunicationCharacteristicsNil sets the value for CommunicationCharacteristics to be an explicit nil

### UnsetCommunicationCharacteristics
`func (o *PpDataEntry) UnsetCommunicationCharacteristics()`

UnsetCommunicationCharacteristics ensures that no value is present for CommunicationCharacteristics, not even an explicit nil
### GetReferenceId

`func (o *PpDataEntry) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PpDataEntry) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PpDataEntry) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PpDataEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetValidityTime

`func (o *PpDataEntry) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *PpDataEntry) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *PpDataEntry) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *PpDataEntry) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *PpDataEntry) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *PpDataEntry) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *PpDataEntry) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *PpDataEntry) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PpDataEntry) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PpDataEntry) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PpDataEntry) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PpDataEntry) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetEcsAddrConfigInfo

`func (o *PpDataEntry) GetEcsAddrConfigInfo() EcsAddrConfigInfo`

GetEcsAddrConfigInfo returns the EcsAddrConfigInfo field if non-nil, zero value otherwise.

### GetEcsAddrConfigInfoOk

`func (o *PpDataEntry) GetEcsAddrConfigInfoOk() (*EcsAddrConfigInfo, bool)`

GetEcsAddrConfigInfoOk returns a tuple with the EcsAddrConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsAddrConfigInfo

`func (o *PpDataEntry) SetEcsAddrConfigInfo(v EcsAddrConfigInfo)`

SetEcsAddrConfigInfo sets EcsAddrConfigInfo field to given value.

### HasEcsAddrConfigInfo

`func (o *PpDataEntry) HasEcsAddrConfigInfo() bool`

HasEcsAddrConfigInfo returns a boolean if a field has been set.

### SetEcsAddrConfigInfoNil

`func (o *PpDataEntry) SetEcsAddrConfigInfoNil(b bool)`

 SetEcsAddrConfigInfoNil sets the value for EcsAddrConfigInfo to be an explicit nil

### UnsetEcsAddrConfigInfo
`func (o *PpDataEntry) UnsetEcsAddrConfigInfo()`

UnsetEcsAddrConfigInfo ensures that no value is present for EcsAddrConfigInfo, not even an explicit nil
### GetAdditionalEcsAddrConfigInfos

`func (o *PpDataEntry) GetAdditionalEcsAddrConfigInfos() []EcsAddrConfigInfo`

GetAdditionalEcsAddrConfigInfos returns the AdditionalEcsAddrConfigInfos field if non-nil, zero value otherwise.

### GetAdditionalEcsAddrConfigInfosOk

`func (o *PpDataEntry) GetAdditionalEcsAddrConfigInfosOk() (*[]EcsAddrConfigInfo, bool)`

GetAdditionalEcsAddrConfigInfosOk returns a tuple with the AdditionalEcsAddrConfigInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEcsAddrConfigInfos

`func (o *PpDataEntry) SetAdditionalEcsAddrConfigInfos(v []EcsAddrConfigInfo)`

SetAdditionalEcsAddrConfigInfos sets AdditionalEcsAddrConfigInfos field to given value.

### HasAdditionalEcsAddrConfigInfos

`func (o *PpDataEntry) HasAdditionalEcsAddrConfigInfos() bool`

HasAdditionalEcsAddrConfigInfos returns a boolean if a field has been set.

### GetEcsAddrConfigInfoPerPlmn

`func (o *PpDataEntry) GetEcsAddrConfigInfoPerPlmn() map[string][]EcsAddrConfigInfo`

GetEcsAddrConfigInfoPerPlmn returns the EcsAddrConfigInfoPerPlmn field if non-nil, zero value otherwise.

### GetEcsAddrConfigInfoPerPlmnOk

`func (o *PpDataEntry) GetEcsAddrConfigInfoPerPlmnOk() (*map[string][]EcsAddrConfigInfo, bool)`

GetEcsAddrConfigInfoPerPlmnOk returns a tuple with the EcsAddrConfigInfoPerPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsAddrConfigInfoPerPlmn

`func (o *PpDataEntry) SetEcsAddrConfigInfoPerPlmn(v map[string][]EcsAddrConfigInfo)`

SetEcsAddrConfigInfoPerPlmn sets EcsAddrConfigInfoPerPlmn field to given value.

### HasEcsAddrConfigInfoPerPlmn

`func (o *PpDataEntry) HasEcsAddrConfigInfoPerPlmn() bool`

HasEcsAddrConfigInfoPerPlmn returns a boolean if a field has been set.

### GetEcRestriction

`func (o *PpDataEntry) GetEcRestriction() EcRestriction`

GetEcRestriction returns the EcRestriction field if non-nil, zero value otherwise.

### GetEcRestrictionOk

`func (o *PpDataEntry) GetEcRestrictionOk() (*EcRestriction, bool)`

GetEcRestrictionOk returns a tuple with the EcRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestriction

`func (o *PpDataEntry) SetEcRestriction(v EcRestriction)`

SetEcRestriction sets EcRestriction field to given value.

### HasEcRestriction

`func (o *PpDataEntry) HasEcRestriction() bool`

HasEcRestriction returns a boolean if a field has been set.

### SetEcRestrictionNil

`func (o *PpDataEntry) SetEcRestrictionNil(b bool)`

 SetEcRestrictionNil sets the value for EcRestriction to be an explicit nil

### UnsetEcRestriction
`func (o *PpDataEntry) UnsetEcRestriction()`

UnsetEcRestriction ensures that no value is present for EcRestriction, not even an explicit nil
### GetSliceUsageControlInfos

`func (o *PpDataEntry) GetSliceUsageControlInfos() []SliceUsageControlInfo`

GetSliceUsageControlInfos returns the SliceUsageControlInfos field if non-nil, zero value otherwise.

### GetSliceUsageControlInfosOk

`func (o *PpDataEntry) GetSliceUsageControlInfosOk() (*[]SliceUsageControlInfo, bool)`

GetSliceUsageControlInfosOk returns a tuple with the SliceUsageControlInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceUsageControlInfos

`func (o *PpDataEntry) SetSliceUsageControlInfos(v []SliceUsageControlInfo)`

SetSliceUsageControlInfos sets SliceUsageControlInfos field to given value.

### HasSliceUsageControlInfos

`func (o *PpDataEntry) HasSliceUsageControlInfos() bool`

HasSliceUsageControlInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


