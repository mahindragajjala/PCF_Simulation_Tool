# EcsAddrConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsServerAddr** | Pointer to [**EcsServerAddr**](EcsServerAddr.md) |  | [optional] 
**SpatialValidityCond** | Pointer to [**SpatialValidityCond**](SpatialValidityCond.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewEcsAddrConfigInfo

`func NewEcsAddrConfigInfo() *EcsAddrConfigInfo`

NewEcsAddrConfigInfo instantiates a new EcsAddrConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcsAddrConfigInfoWithDefaults

`func NewEcsAddrConfigInfoWithDefaults() *EcsAddrConfigInfo`

NewEcsAddrConfigInfoWithDefaults instantiates a new EcsAddrConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsServerAddr

`func (o *EcsAddrConfigInfo) GetEcsServerAddr() EcsServerAddr`

GetEcsServerAddr returns the EcsServerAddr field if non-nil, zero value otherwise.

### GetEcsServerAddrOk

`func (o *EcsAddrConfigInfo) GetEcsServerAddrOk() (*EcsServerAddr, bool)`

GetEcsServerAddrOk returns a tuple with the EcsServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsServerAddr

`func (o *EcsAddrConfigInfo) SetEcsServerAddr(v EcsServerAddr)`

SetEcsServerAddr sets EcsServerAddr field to given value.

### HasEcsServerAddr

`func (o *EcsAddrConfigInfo) HasEcsServerAddr() bool`

HasEcsServerAddr returns a boolean if a field has been set.

### GetSpatialValidityCond

`func (o *EcsAddrConfigInfo) GetSpatialValidityCond() SpatialValidityCond`

GetSpatialValidityCond returns the SpatialValidityCond field if non-nil, zero value otherwise.

### GetSpatialValidityCondOk

`func (o *EcsAddrConfigInfo) GetSpatialValidityCondOk() (*SpatialValidityCond, bool)`

GetSpatialValidityCondOk returns a tuple with the SpatialValidityCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidityCond

`func (o *EcsAddrConfigInfo) SetSpatialValidityCond(v SpatialValidityCond)`

SetSpatialValidityCond sets SpatialValidityCond field to given value.

### HasSpatialValidityCond

`func (o *EcsAddrConfigInfo) HasSpatialValidityCond() bool`

HasSpatialValidityCond returns a boolean if a field has been set.

### GetDnn

`func (o *EcsAddrConfigInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EcsAddrConfigInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EcsAddrConfigInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EcsAddrConfigInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *EcsAddrConfigInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EcsAddrConfigInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EcsAddrConfigInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EcsAddrConfigInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


