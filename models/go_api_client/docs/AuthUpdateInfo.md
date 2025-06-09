# AuthUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationData** | [**ServiceSpecificAuthorizationData**](ServiceSpecificAuthorizationData.md) |  | 
**InvalidityInd** | Pointer to **bool** |  | [optional] 
**InvalidCause** | Pointer to [**InvalidCause**](InvalidCause.md) |  | [optional] 

## Methods

### NewAuthUpdateInfo

`func NewAuthUpdateInfo(authorizationData ServiceSpecificAuthorizationData, ) *AuthUpdateInfo`

NewAuthUpdateInfo instantiates a new AuthUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUpdateInfoWithDefaults

`func NewAuthUpdateInfoWithDefaults() *AuthUpdateInfo`

NewAuthUpdateInfoWithDefaults instantiates a new AuthUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationData

`func (o *AuthUpdateInfo) GetAuthorizationData() ServiceSpecificAuthorizationData`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *AuthUpdateInfo) GetAuthorizationDataOk() (*ServiceSpecificAuthorizationData, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *AuthUpdateInfo) SetAuthorizationData(v ServiceSpecificAuthorizationData)`

SetAuthorizationData sets AuthorizationData field to given value.


### GetInvalidityInd

`func (o *AuthUpdateInfo) GetInvalidityInd() bool`

GetInvalidityInd returns the InvalidityInd field if non-nil, zero value otherwise.

### GetInvalidityIndOk

`func (o *AuthUpdateInfo) GetInvalidityIndOk() (*bool, bool)`

GetInvalidityIndOk returns a tuple with the InvalidityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidityInd

`func (o *AuthUpdateInfo) SetInvalidityInd(v bool)`

SetInvalidityInd sets InvalidityInd field to given value.

### HasInvalidityInd

`func (o *AuthUpdateInfo) HasInvalidityInd() bool`

HasInvalidityInd returns a boolean if a field has been set.

### GetInvalidCause

`func (o *AuthUpdateInfo) GetInvalidCause() InvalidCause`

GetInvalidCause returns the InvalidCause field if non-nil, zero value otherwise.

### GetInvalidCauseOk

`func (o *AuthUpdateInfo) GetInvalidCauseOk() (*InvalidCause, bool)`

GetInvalidCauseOk returns a tuple with the InvalidCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCause

`func (o *AuthUpdateInfo) SetInvalidCause(v InvalidCause)`

SetInvalidCause sets InvalidCause field to given value.

### HasInvalidCause

`func (o *AuthUpdateInfo) HasInvalidCause() bool`

HasInvalidCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


