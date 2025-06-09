# NiddAuthUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationData** | [**AuthorizationData**](AuthorizationData.md) |  | 
**InvalidityInd** | Pointer to **bool** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**NiddCause** | Pointer to [**NiddCause**](NiddCause.md) |  | [optional] 

## Methods

### NewNiddAuthUpdateInfo

`func NewNiddAuthUpdateInfo(authorizationData AuthorizationData, ) *NiddAuthUpdateInfo`

NewNiddAuthUpdateInfo instantiates a new NiddAuthUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddAuthUpdateInfoWithDefaults

`func NewNiddAuthUpdateInfoWithDefaults() *NiddAuthUpdateInfo`

NewNiddAuthUpdateInfoWithDefaults instantiates a new NiddAuthUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationData

`func (o *NiddAuthUpdateInfo) GetAuthorizationData() AuthorizationData`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *NiddAuthUpdateInfo) GetAuthorizationDataOk() (*AuthorizationData, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *NiddAuthUpdateInfo) SetAuthorizationData(v AuthorizationData)`

SetAuthorizationData sets AuthorizationData field to given value.


### GetInvalidityInd

`func (o *NiddAuthUpdateInfo) GetInvalidityInd() bool`

GetInvalidityInd returns the InvalidityInd field if non-nil, zero value otherwise.

### GetInvalidityIndOk

`func (o *NiddAuthUpdateInfo) GetInvalidityIndOk() (*bool, bool)`

GetInvalidityIndOk returns a tuple with the InvalidityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidityInd

`func (o *NiddAuthUpdateInfo) SetInvalidityInd(v bool)`

SetInvalidityInd sets InvalidityInd field to given value.

### HasInvalidityInd

`func (o *NiddAuthUpdateInfo) HasInvalidityInd() bool`

HasInvalidityInd returns a boolean if a field has been set.

### GetSnssai

`func (o *NiddAuthUpdateInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NiddAuthUpdateInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NiddAuthUpdateInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *NiddAuthUpdateInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetDnn

`func (o *NiddAuthUpdateInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *NiddAuthUpdateInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *NiddAuthUpdateInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *NiddAuthUpdateInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetNiddCause

`func (o *NiddAuthUpdateInfo) GetNiddCause() NiddCause`

GetNiddCause returns the NiddCause field if non-nil, zero value otherwise.

### GetNiddCauseOk

`func (o *NiddAuthUpdateInfo) GetNiddCauseOk() (*NiddCause, bool)`

GetNiddCauseOk returns a tuple with the NiddCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddCause

`func (o *NiddAuthUpdateInfo) SetNiddCause(v NiddCause)`

SetNiddCause sets NiddCause field to given value.

### HasNiddCause

`func (o *NiddAuthUpdateInfo) HasNiddCause() bool`

HasNiddCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


