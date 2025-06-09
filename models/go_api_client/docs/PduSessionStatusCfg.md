# PduSessionStatusCfg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 

## Methods

### NewPduSessionStatusCfg

`func NewPduSessionStatusCfg() *PduSessionStatusCfg`

NewPduSessionStatusCfg instantiates a new PduSessionStatusCfg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionStatusCfgWithDefaults

`func NewPduSessionStatusCfgWithDefaults() *PduSessionStatusCfg`

NewPduSessionStatusCfgWithDefaults instantiates a new PduSessionStatusCfg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PduSessionStatusCfg) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionStatusCfg) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionStatusCfg) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PduSessionStatusCfg) HasDnn() bool`

HasDnn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


