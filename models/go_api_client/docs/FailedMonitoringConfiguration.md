# FailedMonitoringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**FailedCause** | [**FailedCause**](FailedCause.md) |  | 

## Methods

### NewFailedMonitoringConfiguration

`func NewFailedMonitoringConfiguration(eventType EventType, failedCause FailedCause, ) *FailedMonitoringConfiguration`

NewFailedMonitoringConfiguration instantiates a new FailedMonitoringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedMonitoringConfigurationWithDefaults

`func NewFailedMonitoringConfigurationWithDefaults() *FailedMonitoringConfiguration`

NewFailedMonitoringConfigurationWithDefaults instantiates a new FailedMonitoringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *FailedMonitoringConfiguration) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *FailedMonitoringConfiguration) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *FailedMonitoringConfiguration) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetFailedCause

`func (o *FailedMonitoringConfiguration) GetFailedCause() FailedCause`

GetFailedCause returns the FailedCause field if non-nil, zero value otherwise.

### GetFailedCauseOk

`func (o *FailedMonitoringConfiguration) GetFailedCauseOk() (*FailedCause, bool)`

GetFailedCauseOk returns a tuple with the FailedCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCause

`func (o *FailedMonitoringConfiguration) SetFailedCause(v FailedCause)`

SetFailedCause sets FailedCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


