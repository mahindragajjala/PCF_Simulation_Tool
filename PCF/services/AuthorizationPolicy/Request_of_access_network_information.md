Request of access network information 

In 5G, "access network information" refers to the details about the radio access network (RAN), which is the part of the cellular network that directly connects user devices like smartphones to the core network through radio signals, essentially providing the means to access the network by transmitting and receiving data over wireless waves; it includes information like the location of cell towers, signal strength, and available frequencies used for communication within a specific area. 


This procedure is used by an AF to request the PCF to report the access network information (i.e. user location and/or user timezone information) at the creation of the "Individual Application Session Context" resource, when the "NetLoc" feature is supported.

The AF shall include in the HTTP POST request message described in subclause 4.2.2.2:
- an entry of the "AfEventSubscription" data type in the "events" attribute with:
a) the "event" attribute set to the value "ANI_REPORT"; and
b) the "notifMethod" attribute set to the value "ONE_TIME"; and

the "reqAnis" attribute, with the required access network information, i.e. user location and/or user time zone information).
When the PCF determines that the access network does not support the access network information reporting because the SMF does not support the NetLoc feature, the PCF shall respond to the AF including in the "EventsNotification" data type the "noNetLocSupp" attribute set to "ANR_NOT_SUPPORTED" value. Otherwise, the PCF shall immediately configure the SMF to provide such access information, as specified in 3GPP TS 29.512 [8].
The PCF shall reply to the AF with an HTTP response message as described in subclause 4.2.2.2

{
  "nfInstanceId": "smf-12345",
  "subscriptionId": "imsi-208930000000001",
  "reqAnis": {
    "accessType": "3GPP_ACCESS",
    "cellId": "12345678",
    "tai": {
      "plmnId": {
        "mcc": "208",
        "mnc": "93"
      },
      "tac": "000001"
    }
  },
  "pduSessionId": 10,
  "supi": "208930000000001",
  "dnn": "internet",
  "snssai": {
    "sst": 1,
    "sd": "010203"
  },
  "ueLocation": {
    "latitude": 48.8566,
    "longitude": 2.3522
  }
}


