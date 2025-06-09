# DnnSnssaiSpecificGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**DefQos** | Pointer to [**AfReqDefaultQoS**](AfReqDefaultQoS.md) |  | [optional] 
**AfReqServArea** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 

## Methods

### NewDnnSnssaiSpecificGroup

`func NewDnnSnssaiSpecificGroup(dnn string, snssai Snssai, ) *DnnSnssaiSpecificGroup`

NewDnnSnssaiSpecificGroup instantiates a new DnnSnssaiSpecificGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnSnssaiSpecificGroupWithDefaults

`func NewDnnSnssaiSpecificGroupWithDefaults() *DnnSnssaiSpecificGroup`

NewDnnSnssaiSpecificGroupWithDefaults instantiates a new DnnSnssaiSpecificGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnSnssaiSpecificGroup) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnSnssaiSpecificGroup) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnSnssaiSpecificGroup) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSnssai

`func (o *DnnSnssaiSpecificGroup) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *DnnSnssaiSpecificGroup) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *DnnSnssaiSpecificGroup) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDefQos

`func (o *DnnSnssaiSpecificGroup) GetDefQos() AfReqDefaultQoS`

GetDefQos returns the DefQos field if non-nil, zero value otherwise.

### GetDefQosOk

`func (o *DnnSnssaiSpecificGroup) GetDefQosOk() (*AfReqDefaultQoS, bool)`

GetDefQosOk returns a tuple with the DefQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefQos

`func (o *DnnSnssaiSpecificGroup) SetDefQos(v AfReqDefaultQoS)`

SetDefQos sets DefQos field to given value.

### HasDefQos

`func (o *DnnSnssaiSpecificGroup) HasDefQos() bool`

HasDefQos returns a boolean if a field has been set.

### GetAfReqServArea

`func (o *DnnSnssaiSpecificGroup) GetAfReqServArea() []Tai`

GetAfReqServArea returns the AfReqServArea field if non-nil, zero value otherwise.

### GetAfReqServAreaOk

`func (o *DnnSnssaiSpecificGroup) GetAfReqServAreaOk() (*[]Tai, bool)`

GetAfReqServAreaOk returns a tuple with the AfReqServArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfReqServArea

`func (o *DnnSnssaiSpecificGroup) SetAfReqServArea(v []Tai)`

SetAfReqServArea sets AfReqServArea field to given value.

### HasAfReqServArea

`func (o *DnnSnssaiSpecificGroup) HasAfReqServArea() bool`

HasAfReqServArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


