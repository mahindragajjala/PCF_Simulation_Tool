https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.johs8gq4msgg

This procedure is used by an AF to indicate sponsored data connectivity when "SponsoredConnectivity" feature is supported.

The AF shall provide in the "AppSessionContext" data type of the HTTP POST request message described in subclause 4.2.2.2 an application service provider identity and a sponsor identity within the "aspId" attribute and "sponId" attribute within the "ascReqData" attribute. 

Additionally, the AF may provide an indication to the PCF of sponsored data connectivity not enabled by including the "sponStatus" attribute set to "SPONSOR_DISABLED".

To support the usage monitoring of sponsored data connectivity, the AF may subscribe with the PCF to the notification of usage threshold reached. The AF shall include:

- an entry of the "AfEventSubscription" data type in the "events" attribute with the "event" attribute set to "USAGE_REPORT"; and
- the "usgThres" attribute of "UsageThreshold" data type in the "EventsSubscReqData" data type with:
a) the total volume in the "totalVolume" attribute; or
b) the uplink volume only in the "uplinkVolume" attribute; or
c) the downlink volume only in the "downlinkVolume"; and/or
d) the time in the "duration" attribute.

NOTE 1: If the AF is in the user plane, the AF can handle the usage monitoring and therefore it is not required to provide a usage threshold to the PCF as part of the sponsored connectivity functionality.

When the AF indicated to enable sponsored data connectivity, and the UE is roaming in a VPLMN, the following procedures apply:

- If the AF is located in the HPLMN, for home routed roaming case and when the operator policies do not allow accessing the sponsored data connectivity with this roaming case, the H-PCF shall reject the service request and shall include in the HTTP "403 Forbidden" response message the "cause" attribute set to "UNAUTHORIZED_SPONSORED_DATA_CONNECTIVITY".

- If the AF is located in the VPLMN, the V-PCF shall reject the service request and shall include in the HTTP "403 Forbidden" response message the "cause" attribute set to "UNAUTHORIZED_SPONSORED_DATA_CONNECTIVITY".

When the AF indicated to enable sponsored data connectivity, and the UE is non-roaming or roaming with the home routed case and the operator policies allow accessing the sponsored data connectivity with this roaming case, the following procedures apply:

If the SMF does not support sponsored connectivity and the required reporting level for that service indicates a sponsored connectivity level according to 3GPP TS 29.512 [8], then the PCF shall reject the request and shall include in the HTTP "403 Forbidden" response message the "cause" attribute set to "REQUESTED_SERVICE_NOT_AUTHORIZED".

- If the SMF supports sponsored data connectivity feature or the required reporting level is different from sponsored connectivity level as described in 3GPP TS 29.512 [8], then the PCF, based on operator policies, shall check whether it is required to validate the sponsored connectivity data. If it is required, it shall perform the authorizations based on sponsored data connectivity profiles. If the authorization fails, the PCF shall include in the HTTP "403 Forbidden" response message the "cause" attribute set to
"UNAUTHORIZED_SPONSORED_DATA_CONNECTIVITY".

NOTE 2: The PCF is not required to verify that a trust relationship exists between the operator and the sponsors.

The PCF shall reply to the AF as described in subclause 4.2.2.2.
