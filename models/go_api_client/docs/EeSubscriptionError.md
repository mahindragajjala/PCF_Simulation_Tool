# EeSubscriptionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.  | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AccessTokenError** | Pointer to [**AccessTokenErr**](AccessTokenErr.md) |  | [optional] 
**AccessTokenRequest** | Pointer to [**AccessTokenReq**](AccessTokenReq.md) |  | [optional] 
**NrfId** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**SupportedApiVersions** | Pointer to **[]string** |  | [optional] 
**NoProfileMatchInfo** | Pointer to [**NoProfileMatchInfo**](NoProfileMatchInfo.md) |  | [optional] 
**SubType** | Pointer to [**SubscriptionType**](SubscriptionType.md) |  | [optional] 
**FailedMonitoringConfigs** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration  | [optional] 
**FailedMoniConfigsEPC** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration, the key value \&quot;ALL\&quot; may  be used to identify a map entry which contains the failed cause of the EE subscription  was not successful in EPC domain.  | [optional] 

## Methods

### NewEeSubscriptionError

`func NewEeSubscriptionError() *EeSubscriptionError`

NewEeSubscriptionError instantiates a new EeSubscriptionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeSubscriptionErrorWithDefaults

`func NewEeSubscriptionErrorWithDefaults() *EeSubscriptionError`

NewEeSubscriptionErrorWithDefaults instantiates a new EeSubscriptionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EeSubscriptionError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EeSubscriptionError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EeSubscriptionError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EeSubscriptionError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *EeSubscriptionError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EeSubscriptionError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EeSubscriptionError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EeSubscriptionError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *EeSubscriptionError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EeSubscriptionError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EeSubscriptionError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EeSubscriptionError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *EeSubscriptionError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EeSubscriptionError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EeSubscriptionError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EeSubscriptionError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *EeSubscriptionError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *EeSubscriptionError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *EeSubscriptionError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *EeSubscriptionError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *EeSubscriptionError) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *EeSubscriptionError) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *EeSubscriptionError) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *EeSubscriptionError) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *EeSubscriptionError) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *EeSubscriptionError) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *EeSubscriptionError) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *EeSubscriptionError) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeSubscriptionError) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeSubscriptionError) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeSubscriptionError) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeSubscriptionError) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *EeSubscriptionError) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *EeSubscriptionError) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *EeSubscriptionError) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *EeSubscriptionError) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *EeSubscriptionError) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *EeSubscriptionError) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *EeSubscriptionError) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *EeSubscriptionError) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *EeSubscriptionError) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *EeSubscriptionError) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *EeSubscriptionError) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *EeSubscriptionError) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetSupportedApiVersions

`func (o *EeSubscriptionError) GetSupportedApiVersions() []string`

GetSupportedApiVersions returns the SupportedApiVersions field if non-nil, zero value otherwise.

### GetSupportedApiVersionsOk

`func (o *EeSubscriptionError) GetSupportedApiVersionsOk() (*[]string, bool)`

GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedApiVersions

`func (o *EeSubscriptionError) SetSupportedApiVersions(v []string)`

SetSupportedApiVersions sets SupportedApiVersions field to given value.

### HasSupportedApiVersions

`func (o *EeSubscriptionError) HasSupportedApiVersions() bool`

HasSupportedApiVersions returns a boolean if a field has been set.

### GetNoProfileMatchInfo

`func (o *EeSubscriptionError) GetNoProfileMatchInfo() NoProfileMatchInfo`

GetNoProfileMatchInfo returns the NoProfileMatchInfo field if non-nil, zero value otherwise.

### GetNoProfileMatchInfoOk

`func (o *EeSubscriptionError) GetNoProfileMatchInfoOk() (*NoProfileMatchInfo, bool)`

GetNoProfileMatchInfoOk returns a tuple with the NoProfileMatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProfileMatchInfo

`func (o *EeSubscriptionError) SetNoProfileMatchInfo(v NoProfileMatchInfo)`

SetNoProfileMatchInfo sets NoProfileMatchInfo field to given value.

### HasNoProfileMatchInfo

`func (o *EeSubscriptionError) HasNoProfileMatchInfo() bool`

HasNoProfileMatchInfo returns a boolean if a field has been set.

### GetSubType

`func (o *EeSubscriptionError) GetSubType() SubscriptionType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *EeSubscriptionError) GetSubTypeOk() (*SubscriptionType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *EeSubscriptionError) SetSubType(v SubscriptionType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *EeSubscriptionError) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetFailedMonitoringConfigs

`func (o *EeSubscriptionError) GetFailedMonitoringConfigs() map[string]FailedMonitoringConfiguration`

GetFailedMonitoringConfigs returns the FailedMonitoringConfigs field if non-nil, zero value otherwise.

### GetFailedMonitoringConfigsOk

`func (o *EeSubscriptionError) GetFailedMonitoringConfigsOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMonitoringConfigsOk returns a tuple with the FailedMonitoringConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMonitoringConfigs

`func (o *EeSubscriptionError) SetFailedMonitoringConfigs(v map[string]FailedMonitoringConfiguration)`

SetFailedMonitoringConfigs sets FailedMonitoringConfigs field to given value.

### HasFailedMonitoringConfigs

`func (o *EeSubscriptionError) HasFailedMonitoringConfigs() bool`

HasFailedMonitoringConfigs returns a boolean if a field has been set.

### GetFailedMoniConfigsEPC

`func (o *EeSubscriptionError) GetFailedMoniConfigsEPC() map[string]FailedMonitoringConfiguration`

GetFailedMoniConfigsEPC returns the FailedMoniConfigsEPC field if non-nil, zero value otherwise.

### GetFailedMoniConfigsEPCOk

`func (o *EeSubscriptionError) GetFailedMoniConfigsEPCOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMoniConfigsEPCOk returns a tuple with the FailedMoniConfigsEPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMoniConfigsEPC

`func (o *EeSubscriptionError) SetFailedMoniConfigsEPC(v map[string]FailedMonitoringConfiguration)`

SetFailedMoniConfigsEPC sets FailedMoniConfigsEPC field to given value.

### HasFailedMoniConfigsEPC

`func (o *EeSubscriptionError) HasFailedMoniConfigsEPC() bool`

HasFailedMoniConfigsEPC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


