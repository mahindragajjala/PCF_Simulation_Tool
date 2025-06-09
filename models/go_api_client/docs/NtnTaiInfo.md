# NtnTaiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**TacList** | **[]string** |  | 
**DerivedTac** | Pointer to **string** | 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.   | [optional] 

## Methods

### NewNtnTaiInfo

`func NewNtnTaiInfo(plmnId PlmnIdNid, tacList []string, ) *NtnTaiInfo`

NewNtnTaiInfo instantiates a new NtnTaiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtnTaiInfoWithDefaults

`func NewNtnTaiInfoWithDefaults() *NtnTaiInfo`

NewNtnTaiInfoWithDefaults instantiates a new NtnTaiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *NtnTaiInfo) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *NtnTaiInfo) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *NtnTaiInfo) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.


### GetTacList

`func (o *NtnTaiInfo) GetTacList() []string`

GetTacList returns the TacList field if non-nil, zero value otherwise.

### GetTacListOk

`func (o *NtnTaiInfo) GetTacListOk() (*[]string, bool)`

GetTacListOk returns a tuple with the TacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacList

`func (o *NtnTaiInfo) SetTacList(v []string)`

SetTacList sets TacList field to given value.


### GetDerivedTac

`func (o *NtnTaiInfo) GetDerivedTac() string`

GetDerivedTac returns the DerivedTac field if non-nil, zero value otherwise.

### GetDerivedTacOk

`func (o *NtnTaiInfo) GetDerivedTacOk() (*string, bool)`

GetDerivedTacOk returns a tuple with the DerivedTac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedTac

`func (o *NtnTaiInfo) SetDerivedTac(v string)`

SetDerivedTac sets DerivedTac field to given value.

### HasDerivedTac

`func (o *NtnTaiInfo) HasDerivedTac() bool`

HasDerivedTac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


