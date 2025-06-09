# MonitoringEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**RevokedCause** | Pointer to [**RevokedCause**](RevokedCause.md) |  | [optional] 

## Methods

### NewMonitoringEvent

`func NewMonitoringEvent(eventType EventType, ) *MonitoringEvent`

NewMonitoringEvent instantiates a new MonitoringEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringEventWithDefaults

`func NewMonitoringEventWithDefaults() *MonitoringEvent`

NewMonitoringEventWithDefaults instantiates a new MonitoringEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MonitoringEvent) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitoringEvent) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitoringEvent) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetRevokedCause

`func (o *MonitoringEvent) GetRevokedCause() RevokedCause`

GetRevokedCause returns the RevokedCause field if non-nil, zero value otherwise.

### GetRevokedCauseOk

`func (o *MonitoringEvent) GetRevokedCauseOk() (*RevokedCause, bool)`

GetRevokedCauseOk returns a tuple with the RevokedCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCause

`func (o *MonitoringEvent) SetRevokedCause(v RevokedCause)`

SetRevokedCause sets RevokedCause field to given value.

### HasRevokedCause

`func (o *MonitoringEvent) HasRevokedCause() bool`

HasRevokedCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


