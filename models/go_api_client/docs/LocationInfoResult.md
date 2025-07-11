# LocationInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AmfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SmsfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**CurrentLoc** | Pointer to **bool** |  | [optional] 
**GeoInfo** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**LocatoinAge** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**Timezone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewLocationInfoResult

`func NewLocationInfoResult() *LocationInfoResult`

NewLocationInfoResult instantiates a new LocationInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoResultWithDefaults

`func NewLocationInfoResultWithDefaults() *LocationInfoResult`

NewLocationInfoResultWithDefaults instantiates a new LocationInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVPlmnId

`func (o *LocationInfoResult) GetVPlmnId() PlmnId`

GetVPlmnId returns the VPlmnId field if non-nil, zero value otherwise.

### GetVPlmnIdOk

`func (o *LocationInfoResult) GetVPlmnIdOk() (*PlmnId, bool)`

GetVPlmnIdOk returns a tuple with the VPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPlmnId

`func (o *LocationInfoResult) SetVPlmnId(v PlmnId)`

SetVPlmnId sets VPlmnId field to given value.

### HasVPlmnId

`func (o *LocationInfoResult) HasVPlmnId() bool`

HasVPlmnId returns a boolean if a field has been set.

### GetAmfInstanceId

`func (o *LocationInfoResult) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *LocationInfoResult) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *LocationInfoResult) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.

### HasAmfInstanceId

`func (o *LocationInfoResult) HasAmfInstanceId() bool`

HasAmfInstanceId returns a boolean if a field has been set.

### GetSmsfInstanceId

`func (o *LocationInfoResult) GetSmsfInstanceId() string`

GetSmsfInstanceId returns the SmsfInstanceId field if non-nil, zero value otherwise.

### GetSmsfInstanceIdOk

`func (o *LocationInfoResult) GetSmsfInstanceIdOk() (*string, bool)`

GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInstanceId

`func (o *LocationInfoResult) SetSmsfInstanceId(v string)`

SetSmsfInstanceId sets SmsfInstanceId field to given value.

### HasSmsfInstanceId

`func (o *LocationInfoResult) HasSmsfInstanceId() bool`

HasSmsfInstanceId returns a boolean if a field has been set.

### GetNcgi

`func (o *LocationInfoResult) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *LocationInfoResult) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *LocationInfoResult) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *LocationInfoResult) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetEcgi

`func (o *LocationInfoResult) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *LocationInfoResult) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *LocationInfoResult) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *LocationInfoResult) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetTai

`func (o *LocationInfoResult) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *LocationInfoResult) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *LocationInfoResult) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *LocationInfoResult) HasTai() bool`

HasTai returns a boolean if a field has been set.

### GetCurrentLoc

`func (o *LocationInfoResult) GetCurrentLoc() bool`

GetCurrentLoc returns the CurrentLoc field if non-nil, zero value otherwise.

### GetCurrentLocOk

`func (o *LocationInfoResult) GetCurrentLocOk() (*bool, bool)`

GetCurrentLocOk returns a tuple with the CurrentLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLoc

`func (o *LocationInfoResult) SetCurrentLoc(v bool)`

SetCurrentLoc sets CurrentLoc field to given value.

### HasCurrentLoc

`func (o *LocationInfoResult) HasCurrentLoc() bool`

HasCurrentLoc returns a boolean if a field has been set.

### GetGeoInfo

`func (o *LocationInfoResult) GetGeoInfo() GeographicArea`

GetGeoInfo returns the GeoInfo field if non-nil, zero value otherwise.

### GetGeoInfoOk

`func (o *LocationInfoResult) GetGeoInfoOk() (*GeographicArea, bool)`

GetGeoInfoOk returns a tuple with the GeoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoInfo

`func (o *LocationInfoResult) SetGeoInfo(v GeographicArea)`

SetGeoInfo sets GeoInfo field to given value.

### HasGeoInfo

`func (o *LocationInfoResult) HasGeoInfo() bool`

HasGeoInfo returns a boolean if a field has been set.

### GetLocatoinAge

`func (o *LocationInfoResult) GetLocatoinAge() int32`

GetLocatoinAge returns the LocatoinAge field if non-nil, zero value otherwise.

### GetLocatoinAgeOk

`func (o *LocationInfoResult) GetLocatoinAgeOk() (*int32, bool)`

GetLocatoinAgeOk returns a tuple with the LocatoinAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatoinAge

`func (o *LocationInfoResult) SetLocatoinAge(v int32)`

SetLocatoinAge sets LocatoinAge field to given value.

### HasLocatoinAge

`func (o *LocationInfoResult) HasLocatoinAge() bool`

HasLocatoinAge returns a boolean if a field has been set.

### GetRatType

`func (o *LocationInfoResult) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *LocationInfoResult) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *LocationInfoResult) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *LocationInfoResult) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetTimezone

`func (o *LocationInfoResult) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LocationInfoResult) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *LocationInfoResult) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *LocationInfoResult) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *LocationInfoResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *LocationInfoResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *LocationInfoResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *LocationInfoResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


