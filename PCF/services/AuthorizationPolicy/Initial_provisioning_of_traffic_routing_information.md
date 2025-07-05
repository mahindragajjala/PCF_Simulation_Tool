https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.s823bcf7isk7

Initial provisioning of traffic routing information

This procedure is used by an AF to:

Influence SMF traffic routing decisions to a local access to a Data Network identified by a DNAI; and/or
Request subscriptions to notifications about UP path management events related to the PDU session, when "InfluenceOnTrafficRouting" feature is supported.


NOTE 1: The AF uses the Npcf_PolicyAuthorization service for requests targeting specific on-going PDU sessions of individual UE(s). 

The AF requests that target existing or future PDU Sessions of multiple UE(s) or any UE are sent via the NEF and may target multiple PCF(s), as described in 3GPP TS 29.513 [7].
The AF shall include in the HTTP POST request message described in subclause 4.2.2.2 the "afRoutReq" attribute of "AfRoutingRequirement" data type with specific routing requirements for the application traffic flows either within "AppSessionContextReqData" data type for the service indicated in the "afAppId" attribute, or within the "medComponents" attribute. 
When provided at both levels, the "afRoutReq" attribute value in the "medComponents" attribute shall have precedence over the "afRoutReq" attribute included in the "AppSessionContextReqData" data type.
The AF may include traffic routing requirements together with service information.
The AF may request to influence SMF traffic routing decisions to a DNAI. The AF shall include in the "afRoutReq" attribute:
a) A list of routes to locations of applications in the "routeToLocs" attribute. Each element of the list shall contain:
a DNAI in the "dnai" attribute to indicate the location of the application towards which the traffic routing is applied; and
either a routing profile identifier in the "routeProfId" attribute, or the explicit routing information in the "routeInfo" attribute.
The AF may include in the "afRoutReq" attribute:
a) Indication of application relocation possibility in the "appReloc" attribute.
b) Temporal validity during which the AF request is valid shall be indicated with the "startTime" and "stopTime" attributes.
c) Spatial validity during which the AF request is valid shall be indicated in terms of validity areas encoded in the "spVal" attribute of "SpatialValidity" data type. The "SpatialValidity" data type consists of a list of presence areas included in the "presenceInfoList" attribute, where each element shall include the presence reporting area identifier in the "praId" attribute and may include the elements composing a presence area encoded in the attributes: "trackingAreaList", "ecgList", "ncgList", "globalRanNodeIdList".
d) Indication of UE IP address preservation in the "addrPreserInd" attribute if the URLLC feature is supported.



AfRoutingRequirement{
				AppReloc: false,
				UpPathChgSub: &models.UpPathChgEvent{
					DnaiChgType:     models.DnaiChangeType_LATE,
					NotificationUri: "http://127.0.0.100:8000/nnef-callback/v1/traffic-influence/edge",
					NotifCorreId:    "1234",
				},
				RouteToLocs: []models.RouteToLocation{
					{
						Dnai:        "edge",
						RouteProfId: "MEC1",
					},
				},
			},
{
					Dnai: "Dnai",
					RouteInfo: &models.RouteInformation{
						Ipv4Addr:   "111.11.11.1",
						Ipv6Addr:   "222.22.22.2",
						PortNumber: 9999,
					},
					RouteProfId: "RouteProfId",
				},
AfRoutReq: &models.AfRoutingRequirement{
				AppReloc: false,
				UpPathChgSub: &models.UpPathChgEvent{
					DnaiChgType:     models.DnaiChangeType_LATE,
					NotificationUri: "http://127.0.0.100:8000/nnef-callback/v1/traffic-influence/edge",
					NotifCorreId:    "1234",
				},
				RouteToLocs: []models.RouteToLocation{
					{
						Dnai:        "edge",
						RouteProfId: "MEC1",
					},
				},
			},
The AF may also subscribe to notifications about UP path management events. The AF shall include in the "upPathChgSub" attribute:
notifications of early and/or late DNAI change, using the attribute "dnaiChgType" indicating whether the subscription is for "EARLY", "LATE" or "EARLY_LATE";
the notification URI where the AF is receiving the Nsmf_EventExposure_Notify service operation in the "notificationUri" attribute; and
the notification correlation identifier assigned by the AF in the "notifCorrelId" attribute.If the URLLC feature is supported, the AF may include an indication of AF acknowledgement to be expected as an "afAckInd" attribute within the "upPathChgSub" attribute.
The PCF shall reply to the AF as described in subclause 4.2.2.2.
The PCF shall store the routing requirements included in the "afRoutReq" attribute.

The PCF shall check whether the received routing requirements requires PCC rules to be created or provisioned to include or modify traffic steering policies, the AF transaction identifier and the application relocation possibility as specified in 3GPP TS 29.513 [7]. 
Provisioning of PCC rules to the SMF shall be carried out as specified in 3GPP TS 29.512 [8].


NOTE 2: The AF receives the notification about UP path management events by the Nsmf_EventExposure_Notify service operation as defined in subclause 4.2.2.2 of 3GPP TS 29.508 [13].
