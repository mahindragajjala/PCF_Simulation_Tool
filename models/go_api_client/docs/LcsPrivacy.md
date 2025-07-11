# LcsPrivacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**Lpi** | Pointer to [**Lpi**](Lpi.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**EvtRptExpectedArea** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**AreaUsageInd** | Pointer to [**AreaUsageInd**](AreaUsageInd.md) |  | [optional] [default to POSITIVE_SENSE]
**UpLocRepIndAf** | Pointer to [**UpLocRepIndAf**](UpLocRepIndAf.md) |  | [optional] [default to USER_PLANE_REPORT_NOT_ALLOWED]

## Methods

### NewLcsPrivacy

`func NewLcsPrivacy() *LcsPrivacy`

NewLcsPrivacy instantiates a new LcsPrivacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLcsPrivacyWithDefaults

`func NewLcsPrivacyWithDefaults() *LcsPrivacy`

NewLcsPrivacyWithDefaults instantiates a new LcsPrivacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *LcsPrivacy) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *LcsPrivacy) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *LcsPrivacy) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.

### HasAfInstanceId

`func (o *LcsPrivacy) HasAfInstanceId() bool`

HasAfInstanceId returns a boolean if a field has been set.

### GetReferenceId

`func (o *LcsPrivacy) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LcsPrivacy) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *LcsPrivacy) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *LcsPrivacy) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetLpi

`func (o *LcsPrivacy) GetLpi() Lpi`

GetLpi returns the Lpi field if non-nil, zero value otherwise.

### GetLpiOk

`func (o *LcsPrivacy) GetLpiOk() (*Lpi, bool)`

GetLpiOk returns a tuple with the Lpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpi

`func (o *LcsPrivacy) SetLpi(v Lpi)`

SetLpi sets Lpi field to given value.

### HasLpi

`func (o *LcsPrivacy) HasLpi() bool`

HasLpi returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *LcsPrivacy) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *LcsPrivacy) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *LcsPrivacy) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *LcsPrivacy) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetEvtRptExpectedArea

`func (o *LcsPrivacy) GetEvtRptExpectedArea() GeographicArea`

GetEvtRptExpectedArea returns the EvtRptExpectedArea field if non-nil, zero value otherwise.

### GetEvtRptExpectedAreaOk

`func (o *LcsPrivacy) GetEvtRptExpectedAreaOk() (*GeographicArea, bool)`

GetEvtRptExpectedAreaOk returns a tuple with the EvtRptExpectedArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtRptExpectedArea

`func (o *LcsPrivacy) SetEvtRptExpectedArea(v GeographicArea)`

SetEvtRptExpectedArea sets EvtRptExpectedArea field to given value.

### HasEvtRptExpectedArea

`func (o *LcsPrivacy) HasEvtRptExpectedArea() bool`

HasEvtRptExpectedArea returns a boolean if a field has been set.

### GetAreaUsageInd

`func (o *LcsPrivacy) GetAreaUsageInd() AreaUsageInd`

GetAreaUsageInd returns the AreaUsageInd field if non-nil, zero value otherwise.

### GetAreaUsageIndOk

`func (o *LcsPrivacy) GetAreaUsageIndOk() (*AreaUsageInd, bool)`

GetAreaUsageIndOk returns a tuple with the AreaUsageInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaUsageInd

`func (o *LcsPrivacy) SetAreaUsageInd(v AreaUsageInd)`

SetAreaUsageInd sets AreaUsageInd field to given value.

### HasAreaUsageInd

`func (o *LcsPrivacy) HasAreaUsageInd() bool`

HasAreaUsageInd returns a boolean if a field has been set.

### GetUpLocRepIndAf

`func (o *LcsPrivacy) GetUpLocRepIndAf() UpLocRepIndAf`

GetUpLocRepIndAf returns the UpLocRepIndAf field if non-nil, zero value otherwise.

### GetUpLocRepIndAfOk

`func (o *LcsPrivacy) GetUpLocRepIndAfOk() (*UpLocRepIndAf, bool)`

GetUpLocRepIndAfOk returns a tuple with the UpLocRepIndAf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLocRepIndAf

`func (o *LcsPrivacy) SetUpLocRepIndAf(v UpLocRepIndAf)`

SetUpLocRepIndAf sets UpLocRepIndAf field to given value.

### HasUpLocRepIndAf

`func (o *LcsPrivacy) HasUpLocRepIndAf() bool`

HasUpLocRepIndAf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


