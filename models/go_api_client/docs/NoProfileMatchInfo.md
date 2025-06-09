# NoProfileMatchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**NoProfileMatchReason**](NoProfileMatchReason.md) |  | 
**QueryParamCombinationList** | Pointer to [**[]QueryParamCombination**](QueryParamCombination.md) |  | [optional] 

## Methods

### NewNoProfileMatchInfo

`func NewNoProfileMatchInfo(reason NoProfileMatchReason, ) *NoProfileMatchInfo`

NewNoProfileMatchInfo instantiates a new NoProfileMatchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoProfileMatchInfoWithDefaults

`func NewNoProfileMatchInfoWithDefaults() *NoProfileMatchInfo`

NewNoProfileMatchInfoWithDefaults instantiates a new NoProfileMatchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *NoProfileMatchInfo) GetReason() NoProfileMatchReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NoProfileMatchInfo) GetReasonOk() (*NoProfileMatchReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NoProfileMatchInfo) SetReason(v NoProfileMatchReason)`

SetReason sets Reason field to given value.


### GetQueryParamCombinationList

`func (o *NoProfileMatchInfo) GetQueryParamCombinationList() []QueryParamCombination`

GetQueryParamCombinationList returns the QueryParamCombinationList field if non-nil, zero value otherwise.

### GetQueryParamCombinationListOk

`func (o *NoProfileMatchInfo) GetQueryParamCombinationListOk() (*[]QueryParamCombination, bool)`

GetQueryParamCombinationListOk returns a tuple with the QueryParamCombinationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParamCombinationList

`func (o *NoProfileMatchInfo) SetQueryParamCombinationList(v []QueryParamCombination)`

SetQueryParamCombinationList sets QueryParamCombinationList field to given value.

### HasQueryParamCombinationList

`func (o *NoProfileMatchInfo) HasQueryParamCombinationList() bool`

HasQueryParamCombinationList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


