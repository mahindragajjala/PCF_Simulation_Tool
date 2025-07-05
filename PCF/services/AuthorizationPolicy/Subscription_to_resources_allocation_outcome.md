https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.bvs8694sflnr

Subscription to resources allocation outcome

This procedure is used by an AF to subscribe to notifications when the resources associated to the corresponding service information have been allocated and/or cannot be allocated.

The AF shall use the "EventsSubscReqData" data type as described in subclause 4.2.2.2 and shall include in the HTTP POST request message:

- if the AF requests the PCF to provide a notification when the resources associated to the service information have been allocated, an event entry within the "events" attribute with the "event" attribute set to "SUCCESSFUL_RESOURCES_ALLOCATION"; and/or

- if the AF requests the PCF to provide a notification when the resources associated to the service information cannot be allocated, an event entry within the "events" attribute with the "event" attribute set to "FAILED_RESOURCES_ALLOCATION".

The PCF shall reply to the AF as described in subclause 4.2.2.2.

As a result of this action, the PCF shall set the appropriate subscription to notification of resources allocation outcome for the corresponding PCC Rule(s) as described in 3GPP TS 29.512 [8].


EventsSubscReqData {
	UpdateEventsSubsc201Data := models.EventsSubscReqData{
		Events: []models.AfEventSubscription{
			{
				Event:       models.AfEvent_ACCESS_TYPE_CHANGE,
				NotifMethod: models.AfNotifMethod_EVENT_DETECTION,
			},
			{
				Event:       models.AfEvent_PLMN_CHG,
				NotifMethod: models.AfNotifMethod_EVENT_DETECTION,
			},
		},
		NotifUri: "https://127.0.0.1:12345",
	}
	return UpdateEventsSubsc201Data

EventsSubscReqData {
	UpdateEventsSubsc200Data := models.EventsSubscReqData{
		Events: []models.AfEventSubscription{
			{
				Event:       models.AfEvent_PLMN_CHG,
				NotifMethod: models.AfNotifMethod_EVENT_DETECTION,
			},
		},
		NotifUri: "https://127.0.0.1:12345",
	}
	return UpdateEventsSubsc200Data

EventsSubscReqData {
	UpdateEventsSubsc204Data := models.EventsSubscReqData{
		Events: []models.AfEventSubscription{
			{
				Event:       models.AfEvent_SUCCESSFUL_RESOURCES_ALLOCATION,
				NotifMethod: models.AfNotifMethod_EVENT_DETECTION,
			},
		},
		NotifUri: "https://127.0.0.1:12345",
	}
	return UpdateEventsSubsc204Data

EventsSubscReqData {
	UpdateEventsSubsc400Data := models.EventsSubscReqData{
		UsgThres: &models.UsageThreshold{
			Duration:       0,
			TotalVolume:    0,
			DownlinkVolume: 0,
			UplinkVolume:   0},
	}
	return UpdateEventsSubsc400Data

EvSubsc: &models.EventsSubscReqDataRm{
			NotifUri: "EvSubsc_NotifUri",
			Events: []models.AfEventSubscription{
				{
					Event:       models.AfEvent_ACCESS_TYPE_CHANGE,
					NotifMethod: models.AfNotifMethod_EVENT_DETECTION,
				},
			},
			UsgThres: &models.UsageThresholdRm{
				Duration:    10,
				TotalVolume: 10,
			},

