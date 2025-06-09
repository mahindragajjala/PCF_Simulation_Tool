# AuthUpdateNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | [**ServiceType**](ServiceType.md) |  | 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**AuthUpdateInfoList** | [**[]AuthUpdateInfo**](AuthUpdateInfo.md) |  | 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**AfId** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthUpdateNotification

`func NewAuthUpdateNotification(serviceType ServiceType, authUpdateInfoList []AuthUpdateInfo, ) *AuthUpdateNotification`

NewAuthUpdateNotification instantiates a new AuthUpdateNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUpdateNotificationWithDefaults

`func NewAuthUpdateNotificationWithDefaults() *AuthUpdateNotification`

NewAuthUpdateNotificationWithDefaults instantiates a new AuthUpdateNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *AuthUpdateNotification) GetServiceType() ServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AuthUpdateNotification) GetServiceTypeOk() (*ServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AuthUpdateNotification) SetServiceType(v ServiceType)`

SetServiceType sets ServiceType field to given value.


### GetSnssai

`func (o *AuthUpdateNotification) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AuthUpdateNotification) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AuthUpdateNotification) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *AuthUpdateNotification) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetDnn

`func (o *AuthUpdateNotification) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AuthUpdateNotification) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AuthUpdateNotification) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AuthUpdateNotification) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetAuthUpdateInfoList

`func (o *AuthUpdateNotification) GetAuthUpdateInfoList() []AuthUpdateInfo`

GetAuthUpdateInfoList returns the AuthUpdateInfoList field if non-nil, zero value otherwise.

### GetAuthUpdateInfoListOk

`func (o *AuthUpdateNotification) GetAuthUpdateInfoListOk() (*[]AuthUpdateInfo, bool)`

GetAuthUpdateInfoListOk returns a tuple with the AuthUpdateInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUpdateInfoList

`func (o *AuthUpdateNotification) SetAuthUpdateInfoList(v []AuthUpdateInfo)`

SetAuthUpdateInfoList sets AuthUpdateInfoList field to given value.


### GetMtcProviderInformation

`func (o *AuthUpdateNotification) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *AuthUpdateNotification) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *AuthUpdateNotification) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *AuthUpdateNotification) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetAfId

`func (o *AuthUpdateNotification) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *AuthUpdateNotification) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *AuthUpdateNotification) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *AuthUpdateNotification) HasAfId() bool`

HasAfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


