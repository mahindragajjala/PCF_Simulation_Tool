# ReachabilityReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AccessTypeList** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**Reachability** | Pointer to [**UeReachability**](UeReachability.md) |  | [optional] 
**MaxAvailabilityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**IdleStatusIndication** | Pointer to [**IdleStatusIndication**](IdleStatusIndication.md) |  | [optional] 

## Methods

### NewReachabilityReport

`func NewReachabilityReport() *ReachabilityReport`

NewReachabilityReport instantiates a new ReachabilityReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityReportWithDefaults

`func NewReachabilityReportWithDefaults() *ReachabilityReport`

NewReachabilityReportWithDefaults instantiates a new ReachabilityReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *ReachabilityReport) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *ReachabilityReport) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *ReachabilityReport) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.

### HasAmfInstanceId

`func (o *ReachabilityReport) HasAmfInstanceId() bool`

HasAmfInstanceId returns a boolean if a field has been set.

### GetAccessTypeList

`func (o *ReachabilityReport) GetAccessTypeList() []AccessType`

GetAccessTypeList returns the AccessTypeList field if non-nil, zero value otherwise.

### GetAccessTypeListOk

`func (o *ReachabilityReport) GetAccessTypeListOk() (*[]AccessType, bool)`

GetAccessTypeListOk returns a tuple with the AccessTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypeList

`func (o *ReachabilityReport) SetAccessTypeList(v []AccessType)`

SetAccessTypeList sets AccessTypeList field to given value.

### HasAccessTypeList

`func (o *ReachabilityReport) HasAccessTypeList() bool`

HasAccessTypeList returns a boolean if a field has been set.

### GetReachability

`func (o *ReachabilityReport) GetReachability() UeReachability`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ReachabilityReport) GetReachabilityOk() (*UeReachability, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ReachabilityReport) SetReachability(v UeReachability)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ReachabilityReport) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetMaxAvailabilityTime

`func (o *ReachabilityReport) GetMaxAvailabilityTime() time.Time`

GetMaxAvailabilityTime returns the MaxAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxAvailabilityTimeOk

`func (o *ReachabilityReport) GetMaxAvailabilityTimeOk() (*time.Time, bool)`

GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailabilityTime

`func (o *ReachabilityReport) SetMaxAvailabilityTime(v time.Time)`

SetMaxAvailabilityTime sets MaxAvailabilityTime field to given value.

### HasMaxAvailabilityTime

`func (o *ReachabilityReport) HasMaxAvailabilityTime() bool`

HasMaxAvailabilityTime returns a boolean if a field has been set.

### GetIdleStatusIndication

`func (o *ReachabilityReport) GetIdleStatusIndication() IdleStatusIndication`

GetIdleStatusIndication returns the IdleStatusIndication field if non-nil, zero value otherwise.

### GetIdleStatusIndicationOk

`func (o *ReachabilityReport) GetIdleStatusIndicationOk() (*IdleStatusIndication, bool)`

GetIdleStatusIndicationOk returns a tuple with the IdleStatusIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusIndication

`func (o *ReachabilityReport) SetIdleStatusIndication(v IdleStatusIndication)`

SetIdleStatusIndication sets IdleStatusIndication field to given value.

### HasIdleStatusIndication

`func (o *ReachabilityReport) HasIdleStatusIndication() bool`

HasIdleStatusIndication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


