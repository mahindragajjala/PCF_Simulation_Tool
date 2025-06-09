# ServiceSpecificAuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**AuthUpdateCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AfId** | Pointer to **string** |  | [optional] 
**NefId** | Pointer to **string** | Identity of the NEF | [optional] 

## Methods

### NewServiceSpecificAuthorizationInfo

`func NewServiceSpecificAuthorizationInfo() *ServiceSpecificAuthorizationInfo`

NewServiceSpecificAuthorizationInfo instantiates a new ServiceSpecificAuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificAuthorizationInfoWithDefaults

`func NewServiceSpecificAuthorizationInfoWithDefaults() *ServiceSpecificAuthorizationInfo`

NewServiceSpecificAuthorizationInfoWithDefaults instantiates a new ServiceSpecificAuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *ServiceSpecificAuthorizationInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ServiceSpecificAuthorizationInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ServiceSpecificAuthorizationInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ServiceSpecificAuthorizationInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetDnn

`func (o *ServiceSpecificAuthorizationInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ServiceSpecificAuthorizationInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ServiceSpecificAuthorizationInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ServiceSpecificAuthorizationInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *ServiceSpecificAuthorizationInfo) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *ServiceSpecificAuthorizationInfo) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *ServiceSpecificAuthorizationInfo) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *ServiceSpecificAuthorizationInfo) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetAuthUpdateCallbackUri

`func (o *ServiceSpecificAuthorizationInfo) GetAuthUpdateCallbackUri() string`

GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field if non-nil, zero value otherwise.

### GetAuthUpdateCallbackUriOk

`func (o *ServiceSpecificAuthorizationInfo) GetAuthUpdateCallbackUriOk() (*string, bool)`

GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUpdateCallbackUri

`func (o *ServiceSpecificAuthorizationInfo) SetAuthUpdateCallbackUri(v string)`

SetAuthUpdateCallbackUri sets AuthUpdateCallbackUri field to given value.

### HasAuthUpdateCallbackUri

`func (o *ServiceSpecificAuthorizationInfo) HasAuthUpdateCallbackUri() bool`

HasAuthUpdateCallbackUri returns a boolean if a field has been set.

### GetAfId

`func (o *ServiceSpecificAuthorizationInfo) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *ServiceSpecificAuthorizationInfo) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *ServiceSpecificAuthorizationInfo) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *ServiceSpecificAuthorizationInfo) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetNefId

`func (o *ServiceSpecificAuthorizationInfo) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *ServiceSpecificAuthorizationInfo) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *ServiceSpecificAuthorizationInfo) SetNefId(v string)`

SetNefId sets NefId field to given value.

### HasNefId

`func (o *ServiceSpecificAuthorizationInfo) HasNefId() bool`

HasNefId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


