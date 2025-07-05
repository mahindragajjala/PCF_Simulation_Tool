https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.3hupsnwi18k9

This procedure is used to set up an AF application session context for the service as defined in 3GPP TS 23.501 [2], 3GPP TS 23.502 [3] and 3GPP TS 23.503 [4].




When a new AF application session context is being established and media information for this application session context is available at the AF and the related media requires PCC control, the AF shall invoke the Npcf_PolicyAuthorization_Create service operation by sending the HTTP POST request to the resource URI representing the "Application Sessions" collection resource of the PCF, as shown in figure 4.2.2.2-1, step 1.

The AF shall include in the "AppSessionContext" data type in the payload body of the HTTP POST request a partial representation of the "Individual Application Session Context" resource by providing the "AppSessionContextReqData" data type. 

The "Individual Application Session Context" resource and the "Events Subscription" sub-resource are created as described below.

The AF shall provide in the body of the HTTP POST request:

- for IP type PDU sessions, the IP address (IPv4 or IPv6) of the UE in the "ueIpv4" or "ueIpv6" attribute; and 

- for Ethernet type PDU sessions, the MAC address of the UE in the "ueMac" attribute.
- if the "TimeSensitiveNetworking" feature is supported, the "ueMac" attribute containing the MAC address of the DS-TT port as received from the PCF during the reporting of bridge information as defined in subclause 4.2.5.16.

NOTE: The determination of the DS-TT MAC address is specified in subclause 5.28.2 of 3GPP TS 23.501 [2].


The AF shall provide the corresponding service information in the "medComponents" attribute if available. 

The AF shall indicate to the PCF as part of the "medComponents" attribute whether the service data flow(s) (IP or Ethernet) should be enabled or disabled with the "fStatus" attribute.
If the "AuthorizationWithRequiredQoS" feature as defined in subclause 5.8 is supported, the AF may provide within the MediaComponent data structure required QoS information as specified in subclause 4.2.2.32.
Initial provisioning of required QoS information - 4.2.2.32
This procedure is used by an AF to request that a data session to a UE is set up with a specific QoS (e.g. low latency or jitter) and priority handling when the "AuthorizationWithRequiredQoS" feature is supported. The AF may provide within one or more entries of the "medComponents" attribute included in the "ascReqData"
attribute of the HTTP POST request message described in subclause 4.2.2.2 a reference to pre-defined QoS information within the "qosReference" attribute.
Additionally, if the AF supports adjustment to different QoS parameter combinations, the AF may provide a prioritized list of one or more QoS references within the "altSerReqs" attribute, where the lower the index of the array for a given entry, the higher the priority.
When the AF provides the "altSerReqs" attribute, the AF shall also subscribe to receive notifications when the GBR QoS targets for one or more service data flows can no longer (or can again) be guaranteed, as described in subclause 4.2.2.6.
Due to the received QoS information, the PCF may need to provision or modify the related PCC rules as specified in 3GPP TS 29.513 [7] and provide the related information towards the SMF following the corresponding procedures specified in 3GPP TS 29.512 [8].
The PCF shall reply to the AF as described in subclause 4.2.2.2
The AF may include the AF application identifier in the "afAppId" attribute into the body of the HTTP POST request in order to indicate the particular service that the AF session belongs to.
The AF application identifier may be provided at both "AppSessionContextReqData" data type level, and "MediaComponent" data type level. 
When provided at both levels, the AF application identifier provided at "MediaComponent" data type level shall have precedence.
The AF application identifier at the "AppSessionContextReqData" data type level may be used to trigger the PCF to indicate to the SMF/UPF to perform the application detection based on the operator's policy as defined in 3GPP TS 29.512 [8].
If the "IMS_SBI" feature is supported, the AF may include the AF charging identifier in the "afChargId" attribute for charging correlation purposes.
If the "TimeSensitiveNetworking" feature is supported the AF may provide TSN information as specified in subclauses 4.2.2.24 and 4.2.2.25.
The AF may also include the "evSubsc" attribute of "EventsSubscReqData" data type to request the notification of certain user plane events. The AF shall include the events to subscribe to in the "events" attribute, and the notification URI where to address the Npcf_PolicyAuthorization_Notify service operation in the "notifUri" attribute. 
The events subscription is provisioned in the "Events Subscription" sub-resource.
The AF shall also include the "notifUri" attribute in the "AppSessionContextReqData" data type to indicate the URI where the PCF can request to the AF the deletion of the "Individual Application Session Context" resource.
If the PCF cannot successfully fulfil the received HTTP POST request due to the internal PCF error or due to the error in the HTTP POST request, the PCF shall send the HTTP error response as specified in subclause 5.7.
Otherwise, when the PCF receives the HTTP POST request from the AF, the PCF shall apply session binding as described in 3GPP TS 29.513 [7]. 
To allow the PCF to identify the PDU session for which the HTTP POST request
applies, the AF shall provide in the body of the HTTP POST request:
for IP type PDU session, either the "ueIpv4" attribute or "ueIpv6" attribute containing the IPv4 or the IPv6 address applicable to an IP flow or IP flows towards the UE; and
for Ethernet type PDU session, the "ueMac" attribute containing the UE MAC address applicable to an Ethernet flow or Ethernet flows towards the UE.
The AF may provide DNN in the "dnn" attribute, SUPI in the "supi" attribute, GPSI in the "gpsi" attribute, the S-NSSAI in the "sliceInfo" attribute if available for session binding.
The AF may also provide the domain identity in the "ipDomain" attribute.

NotifUri: "http://127.0.0.100:8000/nnef-callback/v1/applications/edge",
			IpDomain: "edgeIPDomain",

NOTE 1: 

The "ipDomain" attribute is helpful in the following scenario: Within a network slice instance, there are several separate IP address domains, with SMF/UPF(s) that allocate Ipv4 IP addresses out of the same private address range to UE PDU Sessions. 

The same IP address can thus be allocated to UE PDU sessions served by SMF/UPF(s) in different address domains. 

If one PCF controls several SMF/UPF(s) in different IP address domains, the UE IP address is thus not sufficient for the session binding. 

An AF can serve UEs in different IP address domains, either by having direct IP interfaces to those domains, or by having interconnections via NATs in the user plane between the UPF and the AF. 

If a NAT is used, the AF obtains the IP address allocated to the UE PDU session via application level signalling and supplies it for the session binding to the PCF in the "ueIpv4" attribute. 

The AF supplies an "ipDomain" attribute denoting the IP address domain behind the NAT in addition. The AF can derive the appropriate value from the source address (allocated by the NAT) of incoming user plane packets. The value provided in the "ipDomain" attribute is operator configurable.


NOTE 2: 

The "sliceInfo" attribute is helpful in the scenario where multiple network slice instances are deployed in the same DNN, and the same IPv4 address may be allocated to UE PDU sessions in different network slice instances. 

If one PCF controls several network slices, the UE IP address is not sufficient for the
session binding. The AF supplies "sliceInfo" attribute denoting the network slice instance that allocated the IPv4 address of the UE PDU session. How the AF derives S-NSSAI is out of the scope of this specification.


NOTE 3: 

When the scenario described in NOTE 1 applies and the AF is a P-CSCF it is assumed that the P-CSCF has direct IP interfaces to the different IP address domains and that no NAT is located between the UPF and P-CSCF. How a non-IMS AF obtains the UE private IP address to be provided to the PCF is out of scope of the present release; it is unspecified how to support applications that use a protocol that does not retain the original UE’s private IP address.

If the PCF fails in executing session binding, the PCF shall reject the Npcf_PolicyAuthorization_Create service operation with an HTTP "500 Internal Server Error" response including the "cause" attribute set to "PDU_SESSION_NOT_AVAILABLE".
If the request contains the "medComponents" attribute the PCF shall store the received service information. 
The PCF shall process the received service information according to the operator policy and may decide whether the request is accepted or not. 
The PCF may take the priority information within the "resPrio" attribute into account when making this decision.
If the service information provided in the body of the HTTP POST request is rejected (e.g. the subscribed guaranteed bandwidth for a particular user is exceeded), the PCF shall indicate in an HTTP "403 Forbidden" response message the cause for the rejection including the "cause" attribute set to "REQUESTED_SERVICE_NOT_AUTHORIZED".
If the service information provided in the HTTP POST request is rejected due to a temporary condition in the network (e.g. the NWDAF reported the network slice selected for the PDU session is congested), the PCF may include in the "403 Forbidden" response the "cause" attribute set to "REQUESTED_SERVICE_TEMPORARILY_NOT_AUTHORIZED". 
The PCF may also provide a retry interval within the "Retry-After" HTTP header field. 
When the AF receives the retry interval within the "Retry-After" HTTP header field, the AF shall not send the same service information to the PCF again (for the same application session context) until the retry interval has elapsed. 
The "Retry-After" HTTP header is described in 3GPP TS 29.500 [5] subclause 5.2.2.2.
The PCF may additionally provide the acceptable bandwidth within the attribute "acceptableServInfo" included in the "ExtendedProblemDetails" data structure returned in the rejection response message.
To allow the PCF and SMF/UPF to perform PCC rule authorization and QoS flow binding for the described service data flows, the AF shall supply:
for IP type PDU session, both source and destination IP addresses and port numbers in the "fDescs" attribute within the "medSubComps" attribute, if such information is available; and
for Ethernet type PDU session, the Ethernet Packet filters in the "ethfDescs" attribute within the "medSubComps" attribute, if such information is available.
The AF may specify the ToS traffic class within the "tosTrCl" attribute for the described service data flows together with the "fDescs" attribute.
The AF may include the "resPrio" attribute at the "AppSessionContextReqData" data type level to assign a priority to the AF Session as well as include the "resPrio" attribute at the "MediaComponent" data type level to assign a priority to the service data flow. 
The presence of the "resPrio" attribute in both levels does not constitute a conflict as they each represent different types of priority. The reservation priority at the "AppSessionContextReqData" data type level provides the relative priority for an AF session while the reservation priority at the "MediaComponent" data type level provides the relative priority for a service data flow within a session. If the "resPrio" attribute is not specified, the requested priority is PRIO_1.
The PCF shall check whether the received service information requires PCC rules to be created and provisioned as specified in 3GPP TS 29.513 [7]. Provisioning of PCC rules to the SMF shall be carried out as specified at 3GPP TS 29.512 [8].
Based on the received subscription information from the AF, the PCF may create a subscription to event notifications for a related PDU session from the SMF, as described in 3GPP TS 29.512 [8].
If the PCF created an "Individual Application Session Context" resource, the PCF shall send to the AF a "201 Created" response to the HTTP POST request, as shown in figure 4.2.2.2-1, step 2. The PCF shall include in the "201 Created" response:

- a Location header field; and
- an "AppSessionContext" data type in the payload body.

The Location header field shall contain the URI of the created individual application session context resource i.e. "{apiRoot}/npcf-policyauthorization/v1/app-sessions/{appSessionId}".

When "Events Subscription" sub-resource is created in this procedure, the AF shall build the sub-resource URI by adding the path segment "/events-subscription" at the end of the URI path received in the Location header field.

The "AppSessionContext" data type payload body shall contain the representation of the created "Individual Application Session Context" resource and may include the "Events Subscription" sub-resource.

The acknowledgement towards the AF should take place before or in parallel with any required PCC rule provisioning towards the SMF.

NOTE 5: 

The behaviour when the AF does not receive the HTTP response message, or when it arrives after the internal timer waiting for it has expired, or when it arrives with an indication different than a success indication, are outside the scope of this specification and based on operator policy. 

