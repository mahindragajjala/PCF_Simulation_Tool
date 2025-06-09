# AfReqDefaultQoS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | 
**Arp** | [**Arp**](Arp.md) |  | 
**PriorityLevel** | Pointer to **int32** | Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.   | [optional] 

## Methods

### NewAfReqDefaultQoS

`func NewAfReqDefaultQoS(var5qi int32, arp Arp, ) *AfReqDefaultQoS`

NewAfReqDefaultQoS instantiates a new AfReqDefaultQoS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfReqDefaultQoSWithDefaults

`func NewAfReqDefaultQoSWithDefaults() *AfReqDefaultQoS`

NewAfReqDefaultQoSWithDefaults instantiates a new AfReqDefaultQoS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *AfReqDefaultQoS) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *AfReqDefaultQoS) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *AfReqDefaultQoS) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.


### GetArp

`func (o *AfReqDefaultQoS) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *AfReqDefaultQoS) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *AfReqDefaultQoS) SetArp(v Arp)`

SetArp sets Arp field to given value.


### GetPriorityLevel

`func (o *AfReqDefaultQoS) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *AfReqDefaultQoS) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *AfReqDefaultQoS) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *AfReqDefaultQoS) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


