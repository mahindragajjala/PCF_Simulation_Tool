# Model5GVnGroupConfigurationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5gVnGroupData** | Pointer to [**NullableModel5GVnGroupDataModification**](Model5GVnGroupDataModification.md) |  | [optional] 
**AfInstanceId** | Pointer to **string** |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 
**MembersData** | Pointer to **map[string]map[string]interface{}** | Contains the list of 5G VN Group members, each member is identified by GPSI. A map (list of key-value pairs where Gpsi serves as key) of GpsiInfo.  The value in each entries of the map shall be an empty JSON object.  | [optional] 

## Methods

### NewModel5GVnGroupConfigurationModification

`func NewModel5GVnGroupConfigurationModification() *Model5GVnGroupConfigurationModification`

NewModel5GVnGroupConfigurationModification instantiates a new Model5GVnGroupConfigurationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GVnGroupConfigurationModificationWithDefaults

`func NewModel5GVnGroupConfigurationModificationWithDefaults() *Model5GVnGroupConfigurationModification`

NewModel5GVnGroupConfigurationModificationWithDefaults instantiates a new Model5GVnGroupConfigurationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5gVnGroupData

`func (o *Model5GVnGroupConfigurationModification) GetVar5gVnGroupData() Model5GVnGroupDataModification`

GetVar5gVnGroupData returns the Var5gVnGroupData field if non-nil, zero value otherwise.

### GetVar5gVnGroupDataOk

`func (o *Model5GVnGroupConfigurationModification) GetVar5gVnGroupDataOk() (*Model5GVnGroupDataModification, bool)`

GetVar5gVnGroupDataOk returns a tuple with the Var5gVnGroupData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gVnGroupData

`func (o *Model5GVnGroupConfigurationModification) SetVar5gVnGroupData(v Model5GVnGroupDataModification)`

SetVar5gVnGroupData sets Var5gVnGroupData field to given value.

### HasVar5gVnGroupData

`func (o *Model5GVnGroupConfigurationModification) HasVar5gVnGroupData() bool`

HasVar5gVnGroupData returns a boolean if a field has been set.

### SetVar5gVnGroupDataNil

`func (o *Model5GVnGroupConfigurationModification) SetVar5gVnGroupDataNil(b bool)`

 SetVar5gVnGroupDataNil sets the value for Var5gVnGroupData to be an explicit nil

### UnsetVar5gVnGroupData
`func (o *Model5GVnGroupConfigurationModification) UnsetVar5gVnGroupData()`

UnsetVar5gVnGroupData ensures that no value is present for Var5gVnGroupData, not even an explicit nil
### GetAfInstanceId

`func (o *Model5GVnGroupConfigurationModification) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *Model5GVnGroupConfigurationModification) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *Model5GVnGroupConfigurationModification) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.

### HasAfInstanceId

`func (o *Model5GVnGroupConfigurationModification) HasAfInstanceId() bool`

HasAfInstanceId returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *Model5GVnGroupConfigurationModification) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *Model5GVnGroupConfigurationModification) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *Model5GVnGroupConfigurationModification) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *Model5GVnGroupConfigurationModification) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetMembers

`func (o *Model5GVnGroupConfigurationModification) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Model5GVnGroupConfigurationModification) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Model5GVnGroupConfigurationModification) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Model5GVnGroupConfigurationModification) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *Model5GVnGroupConfigurationModification) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *Model5GVnGroupConfigurationModification) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetMembersData

`func (o *Model5GVnGroupConfigurationModification) GetMembersData() map[string]map[string]interface{}`

GetMembersData returns the MembersData field if non-nil, zero value otherwise.

### GetMembersDataOk

`func (o *Model5GVnGroupConfigurationModification) GetMembersDataOk() (*map[string]map[string]interface{}, bool)`

GetMembersDataOk returns a tuple with the MembersData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersData

`func (o *Model5GVnGroupConfigurationModification) SetMembersData(v map[string]map[string]interface{})`

SetMembersData sets MembersData field to given value.

### HasMembersData

`func (o *Model5GVnGroupConfigurationModification) HasMembersData() bool`

HasMembersData returns a boolean if a field has been set.

### SetMembersDataNil

`func (o *Model5GVnGroupConfigurationModification) SetMembersDataNil(b bool)`

 SetMembersDataNil sets the value for MembersData to be an explicit nil

### UnsetMembersData
`func (o *Model5GVnGroupConfigurationModification) UnsetMembersData()`

UnsetMembersData ensures that no value is present for MembersData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


