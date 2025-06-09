# RoamingInfoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roaming** | Pointer to **bool** |  | [optional] 
**ServingPlmn** | [**PlmnId**](PlmnId.md) |  | 

## Methods

### NewRoamingInfoUpdate

`func NewRoamingInfoUpdate(servingPlmn PlmnId, ) *RoamingInfoUpdate`

NewRoamingInfoUpdate instantiates a new RoamingInfoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingInfoUpdateWithDefaults

`func NewRoamingInfoUpdateWithDefaults() *RoamingInfoUpdate`

NewRoamingInfoUpdateWithDefaults instantiates a new RoamingInfoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoaming

`func (o *RoamingInfoUpdate) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *RoamingInfoUpdate) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *RoamingInfoUpdate) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *RoamingInfoUpdate) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetServingPlmn

`func (o *RoamingInfoUpdate) GetServingPlmn() PlmnId`

GetServingPlmn returns the ServingPlmn field if non-nil, zero value otherwise.

### GetServingPlmnOk

`func (o *RoamingInfoUpdate) GetServingPlmnOk() (*PlmnId, bool)`

GetServingPlmnOk returns a tuple with the ServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPlmn

`func (o *RoamingInfoUpdate) SetServingPlmn(v PlmnId)`

SetServingPlmn sets ServingPlmn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


