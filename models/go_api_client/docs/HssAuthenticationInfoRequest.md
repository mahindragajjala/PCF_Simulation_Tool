# HssAuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**HssAuthType** | [**HssAuthType**](HssAuthType.md) |  | 
**NumOfRequestedVectors** | **int32** |  | 
**ServingNetworkId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**AnId** | Pointer to [**AccessNetworkId**](AccessNetworkId.md) |  | [optional] 

## Methods

### NewHssAuthenticationInfoRequest

`func NewHssAuthenticationInfoRequest(hssAuthType HssAuthType, numOfRequestedVectors int32, ) *HssAuthenticationInfoRequest`

NewHssAuthenticationInfoRequest instantiates a new HssAuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHssAuthenticationInfoRequestWithDefaults

`func NewHssAuthenticationInfoRequestWithDefaults() *HssAuthenticationInfoRequest`

NewHssAuthenticationInfoRequestWithDefaults instantiates a new HssAuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *HssAuthenticationInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *HssAuthenticationInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *HssAuthenticationInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *HssAuthenticationInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetHssAuthType

`func (o *HssAuthenticationInfoRequest) GetHssAuthType() HssAuthType`

GetHssAuthType returns the HssAuthType field if non-nil, zero value otherwise.

### GetHssAuthTypeOk

`func (o *HssAuthenticationInfoRequest) GetHssAuthTypeOk() (*HssAuthType, bool)`

GetHssAuthTypeOk returns a tuple with the HssAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssAuthType

`func (o *HssAuthenticationInfoRequest) SetHssAuthType(v HssAuthType)`

SetHssAuthType sets HssAuthType field to given value.


### GetNumOfRequestedVectors

`func (o *HssAuthenticationInfoRequest) GetNumOfRequestedVectors() int32`

GetNumOfRequestedVectors returns the NumOfRequestedVectors field if non-nil, zero value otherwise.

### GetNumOfRequestedVectorsOk

`func (o *HssAuthenticationInfoRequest) GetNumOfRequestedVectorsOk() (*int32, bool)`

GetNumOfRequestedVectorsOk returns a tuple with the NumOfRequestedVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfRequestedVectors

`func (o *HssAuthenticationInfoRequest) SetNumOfRequestedVectors(v int32)`

SetNumOfRequestedVectors sets NumOfRequestedVectors field to given value.


### GetServingNetworkId

`func (o *HssAuthenticationInfoRequest) GetServingNetworkId() PlmnId`

GetServingNetworkId returns the ServingNetworkId field if non-nil, zero value otherwise.

### GetServingNetworkIdOk

`func (o *HssAuthenticationInfoRequest) GetServingNetworkIdOk() (*PlmnId, bool)`

GetServingNetworkIdOk returns a tuple with the ServingNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkId

`func (o *HssAuthenticationInfoRequest) SetServingNetworkId(v PlmnId)`

SetServingNetworkId sets ServingNetworkId field to given value.

### HasServingNetworkId

`func (o *HssAuthenticationInfoRequest) HasServingNetworkId() bool`

HasServingNetworkId returns a boolean if a field has been set.

### GetResynchronizationInfo

`func (o *HssAuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *HssAuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *HssAuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *HssAuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetAnId

`func (o *HssAuthenticationInfoRequest) GetAnId() AccessNetworkId`

GetAnId returns the AnId field if non-nil, zero value otherwise.

### GetAnIdOk

`func (o *HssAuthenticationInfoRequest) GetAnIdOk() (*AccessNetworkId, bool)`

GetAnIdOk returns a tuple with the AnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnId

`func (o *HssAuthenticationInfoRequest) SetAnId(v AccessNetworkId)`

SetAnId sets AnId field to given value.

### HasAnId

`func (o *HssAuthenticationInfoRequest) HasAnId() bool`

HasAnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


