# PpDlPacketCountExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 

## Methods

### NewPpDlPacketCountExt

`func NewPpDlPacketCountExt(afInstanceId string, referenceId int32, ) *PpDlPacketCountExt`

NewPpDlPacketCountExt instantiates a new PpDlPacketCountExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpDlPacketCountExtWithDefaults

`func NewPpDlPacketCountExtWithDefaults() *PpDlPacketCountExt`

NewPpDlPacketCountExtWithDefaults instantiates a new PpDlPacketCountExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *PpDlPacketCountExt) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *PpDlPacketCountExt) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *PpDlPacketCountExt) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *PpDlPacketCountExt) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PpDlPacketCountExt) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PpDlPacketCountExt) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetDnn

`func (o *PpDlPacketCountExt) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PpDlPacketCountExt) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PpDlPacketCountExt) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PpDlPacketCountExt) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSingleNssai

`func (o *PpDlPacketCountExt) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *PpDlPacketCountExt) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *PpDlPacketCountExt) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *PpDlPacketCountExt) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetValidityTime

`func (o *PpDlPacketCountExt) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *PpDlPacketCountExt) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *PpDlPacketCountExt) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *PpDlPacketCountExt) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *PpDlPacketCountExt) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *PpDlPacketCountExt) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *PpDlPacketCountExt) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *PpDlPacketCountExt) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


