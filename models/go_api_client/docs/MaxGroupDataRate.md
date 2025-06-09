# MaxGroupDataRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Downlink** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewMaxGroupDataRate

`func NewMaxGroupDataRate() *MaxGroupDataRate`

NewMaxGroupDataRate instantiates a new MaxGroupDataRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxGroupDataRateWithDefaults

`func NewMaxGroupDataRateWithDefaults() *MaxGroupDataRate`

NewMaxGroupDataRateWithDefaults instantiates a new MaxGroupDataRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *MaxGroupDataRate) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *MaxGroupDataRate) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *MaxGroupDataRate) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *MaxGroupDataRate) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetDownlink

`func (o *MaxGroupDataRate) GetDownlink() string`

GetDownlink returns the Downlink field if non-nil, zero value otherwise.

### GetDownlinkOk

`func (o *MaxGroupDataRate) GetDownlinkOk() (*string, bool)`

GetDownlinkOk returns a tuple with the Downlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlink

`func (o *MaxGroupDataRate) SetDownlink(v string)`

SetDownlink sets Downlink field to given value.

### HasDownlink

`func (o *MaxGroupDataRate) HasDownlink() bool`

HasDownlink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


