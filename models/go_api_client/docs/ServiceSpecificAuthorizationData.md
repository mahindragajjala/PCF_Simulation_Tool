# ServiceSpecificAuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUeId** | Pointer to [**AuthorizationUeId**](AuthorizationUeId.md) |  | [optional] 
**ExtGroupId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**IntGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceSpecificAuthorizationData

`func NewServiceSpecificAuthorizationData() *ServiceSpecificAuthorizationData`

NewServiceSpecificAuthorizationData instantiates a new ServiceSpecificAuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificAuthorizationDataWithDefaults

`func NewServiceSpecificAuthorizationDataWithDefaults() *ServiceSpecificAuthorizationData`

NewServiceSpecificAuthorizationDataWithDefaults instantiates a new ServiceSpecificAuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUeId

`func (o *ServiceSpecificAuthorizationData) GetAuthorizationUeId() AuthorizationUeId`

GetAuthorizationUeId returns the AuthorizationUeId field if non-nil, zero value otherwise.

### GetAuthorizationUeIdOk

`func (o *ServiceSpecificAuthorizationData) GetAuthorizationUeIdOk() (*AuthorizationUeId, bool)`

GetAuthorizationUeIdOk returns a tuple with the AuthorizationUeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUeId

`func (o *ServiceSpecificAuthorizationData) SetAuthorizationUeId(v AuthorizationUeId)`

SetAuthorizationUeId sets AuthorizationUeId field to given value.

### HasAuthorizationUeId

`func (o *ServiceSpecificAuthorizationData) HasAuthorizationUeId() bool`

HasAuthorizationUeId returns a boolean if a field has been set.

### GetExtGroupId

`func (o *ServiceSpecificAuthorizationData) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *ServiceSpecificAuthorizationData) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *ServiceSpecificAuthorizationData) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.

### HasExtGroupId

`func (o *ServiceSpecificAuthorizationData) HasExtGroupId() bool`

HasExtGroupId returns a boolean if a field has been set.

### GetIntGroupId

`func (o *ServiceSpecificAuthorizationData) GetIntGroupId() string`

GetIntGroupId returns the IntGroupId field if non-nil, zero value otherwise.

### GetIntGroupIdOk

`func (o *ServiceSpecificAuthorizationData) GetIntGroupIdOk() (*string, bool)`

GetIntGroupIdOk returns a tuple with the IntGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGroupId

`func (o *ServiceSpecificAuthorizationData) SetIntGroupId(v string)`

SetIntGroupId sets IntGroupId field to given value.

### HasIntGroupId

`func (o *ServiceSpecificAuthorizationData) HasIntGroupId() bool`

HasIntGroupId returns a boolean if a field has been set.

### GetAuthId

`func (o *ServiceSpecificAuthorizationData) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *ServiceSpecificAuthorizationData) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *ServiceSpecificAuthorizationData) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *ServiceSpecificAuthorizationData) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


