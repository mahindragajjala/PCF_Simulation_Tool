# MulticastMbsGroupMemb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MulticastGroupMemb** | **[]string** |  | 
**AfInstanceId** | Pointer to **string** |  | [optional] 
**InternalGroupIdentifier** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 

## Methods

### NewMulticastMbsGroupMemb

`func NewMulticastMbsGroupMemb(multicastGroupMemb []string, ) *MulticastMbsGroupMemb`

NewMulticastMbsGroupMemb instantiates a new MulticastMbsGroupMemb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticastMbsGroupMembWithDefaults

`func NewMulticastMbsGroupMembWithDefaults() *MulticastMbsGroupMemb`

NewMulticastMbsGroupMembWithDefaults instantiates a new MulticastMbsGroupMemb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMulticastGroupMemb

`func (o *MulticastMbsGroupMemb) GetMulticastGroupMemb() []string`

GetMulticastGroupMemb returns the MulticastGroupMemb field if non-nil, zero value otherwise.

### GetMulticastGroupMembOk

`func (o *MulticastMbsGroupMemb) GetMulticastGroupMembOk() (*[]string, bool)`

GetMulticastGroupMembOk returns a tuple with the MulticastGroupMemb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastGroupMemb

`func (o *MulticastMbsGroupMemb) SetMulticastGroupMemb(v []string)`

SetMulticastGroupMemb sets MulticastGroupMemb field to given value.


### GetAfInstanceId

`func (o *MulticastMbsGroupMemb) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *MulticastMbsGroupMemb) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *MulticastMbsGroupMemb) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.

### HasAfInstanceId

`func (o *MulticastMbsGroupMemb) HasAfInstanceId() bool`

HasAfInstanceId returns a boolean if a field has been set.

### GetInternalGroupIdentifier

`func (o *MulticastMbsGroupMemb) GetInternalGroupIdentifier() string`

GetInternalGroupIdentifier returns the InternalGroupIdentifier field if non-nil, zero value otherwise.

### GetInternalGroupIdentifierOk

`func (o *MulticastMbsGroupMemb) GetInternalGroupIdentifierOk() (*string, bool)`

GetInternalGroupIdentifierOk returns a tuple with the InternalGroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIdentifier

`func (o *MulticastMbsGroupMemb) SetInternalGroupIdentifier(v string)`

SetInternalGroupIdentifier sets InternalGroupIdentifier field to given value.

### HasInternalGroupIdentifier

`func (o *MulticastMbsGroupMemb) HasInternalGroupIdentifier() bool`

HasInternalGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


