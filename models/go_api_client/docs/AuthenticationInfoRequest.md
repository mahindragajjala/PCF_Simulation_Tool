# AuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ServingNetworkName** | **string** |  | 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**AusfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**CagId** | Pointer to **string** | String containing a Closed Access Group Identifier. | [optional] 
**N5gcInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthenticationInfoRequest

`func NewAuthenticationInfoRequest(servingNetworkName string, ausfInstanceId string, ) *AuthenticationInfoRequest`

NewAuthenticationInfoRequest instantiates a new AuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoRequestWithDefaults

`func NewAuthenticationInfoRequestWithDefaults() *AuthenticationInfoRequest`

NewAuthenticationInfoRequestWithDefaults instantiates a new AuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *AuthenticationInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServingNetworkName

`func (o *AuthenticationInfoRequest) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *AuthenticationInfoRequest) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *AuthenticationInfoRequest) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.


### GetResynchronizationInfo

`func (o *AuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *AuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *AuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *AuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetAusfInstanceId

`func (o *AuthenticationInfoRequest) GetAusfInstanceId() string`

GetAusfInstanceId returns the AusfInstanceId field if non-nil, zero value otherwise.

### GetAusfInstanceIdOk

`func (o *AuthenticationInfoRequest) GetAusfInstanceIdOk() (*string, bool)`

GetAusfInstanceIdOk returns a tuple with the AusfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInstanceId

`func (o *AuthenticationInfoRequest) SetAusfInstanceId(v string)`

SetAusfInstanceId sets AusfInstanceId field to given value.


### GetCagId

`func (o *AuthenticationInfoRequest) GetCagId() string`

GetCagId returns the CagId field if non-nil, zero value otherwise.

### GetCagIdOk

`func (o *AuthenticationInfoRequest) GetCagIdOk() (*string, bool)`

GetCagIdOk returns a tuple with the CagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagId

`func (o *AuthenticationInfoRequest) SetCagId(v string)`

SetCagId sets CagId field to given value.

### HasCagId

`func (o *AuthenticationInfoRequest) HasCagId() bool`

HasCagId returns a boolean if a field has been set.

### GetN5gcInd

`func (o *AuthenticationInfoRequest) GetN5gcInd() bool`

GetN5gcInd returns the N5gcInd field if non-nil, zero value otherwise.

### GetN5gcIndOk

`func (o *AuthenticationInfoRequest) GetN5gcIndOk() (*bool, bool)`

GetN5gcIndOk returns a tuple with the N5gcInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN5gcInd

`func (o *AuthenticationInfoRequest) SetN5gcInd(v bool)`

SetN5gcInd sets N5gcInd field to given value.

### HasN5gcInd

`func (o *AuthenticationInfoRequest) HasN5gcInd() bool`

HasN5gcInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


