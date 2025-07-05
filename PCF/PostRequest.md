https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.ly2s4dba7l8z

The procedure in the present sub clause is applicable when the NF service consumer creates an AM policy association.

when the UE registers to the network, and when the AMF is relocated (between the different AMF sets) and the new AMF selects a new PCF. 

The procedure for the case where the AMF is relocated and the new AMF selects the old PCF is defined in subclause 4.2.3.1.

The creation of an AM policy association only applies for normally registered UEs, i.e., it does not apply for Emergency Registered UEs.

Figure 4.2.2.1-1 illustrates the creation of a policy association. 
When a UE registers and a UE context is being established, the AMF can obtain Service Area Restrictions, 
RFSP index,
subscribed UE-AMBR and GPSI from the UDM during the update location procedure 
allowed NSSAI from local configuration or from the NSSF during the slice selection procedure and shall decide based on local policies
whether to request policies from the PCF.

To request policies from the PCF, the NF service consumer (e.g. AMF) shall send an HTTP POST request with:
"{apiRoot}/npcf-am-policy-control/v1/policies" as Resource URI and the PolicyAssociationRequest data structure as
request body that shall include:

- Notification URI encoded as "notificationUri" attribute;
- SUPI encoded as "supi" attribute; and
- if the feature "SliceSupport" or the feature "DNNReplacementControl" is supported in the AMF and the UE is registered in the 3GPP access, the allowed NSSAI in the 3GPP access encoded in the "allowedSnssais" attribute;
and that shall include when available:
- GPSI encoded as "gpsi" attribute;
- if the feature "MultipleAccessTypes" is not supported, the access type encoded as "accessType" attribute; 
- Permanent Equipment Identifier (PEI) encoded as "pei" attribute; 
- User Location Information encoded as "userLoc" attribute; 
- UE Time Zone encoded as "timeZone" attribute; 
- Serving PLMN Identifier and for SNPN the NID encoded as "servingPlmn" attribute; 
- if the feature "MultipleAccessTypes" is not supported, the RAT type encoded as "ratType" attribute; 
- Service Area Restrictions (see subclause 4.2.2.3.1) derived from the Service Area Restrictions obtained from the UDM by mapping any service areas denoted by geographical information into Tracking Area Identities (TAIs) and encoded as "servAreaRes" attribute; 
- RFSP index (see subclause 4.2.2.3.2) as obtained from the UDM encoded as "rfsp" attribute; 
- A list of Internal Group Identifiers encoded as "groupIds" attribute; 
- if the NF service consumer is an AMF, the GUAMI encoded as "guami" attribute; 
- if the NF service consumer is an AMF, the name of a service produced by the AMF that expects to receive information within Npcf_AMPolicyControl_UpdateNotify service operation encoded as "serviceName" attribute; 
- Alternate or backup IPv4 Address(es) where to send Notifications encoded as "altNotifIpv4Addrs" attribute;
- Alternate or backup IPv6 Address(es) where to send Notifications encoded as "altNotifIpv6Addrs" attribute;
- Alternate or backup FQDN(s) where to send Notifications encoded as "altNotifFqdns" attribute;
- trace control and configuration parameters information encoded as "traceReq" attribute;
- if the feature "UE-AMBR_Authorization" is supported in the AMF, the subscribed UE-AMBR (see subclause 4.2.2.3.3) in the "ueAmbr" attribute; and
- if the feature"DNNReplacementControl" is supported, the mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN encoded in the "mappingSnssais" attribute.

Upon the reception of the HTTP POST request, the PCF:
- shall assign a policy association ID;
- shall determine ( పరిష్కరించు,కచ్చితంగా తెలుసుకొను )
 the applicable policy (taking into consideration and optionally modifying possibly received UEAMBR, Service Area Restrictions and/or RFSP index);

- for the successful case shall send a HTTP "201 Created" response with the URI for the created resource in the "Location" header field

NOTE: The assigned policy association ID is part of the URI for the created resource and is thus associated with the SUPI.and the PolicyAssociation data type as body including:

- conditionally AMF Access and Mobility Policy (see subclause 4.2.2.3), i.e.:
a) if the PCF received the "servAreaRes" attribute in the request, Service Area Restrictions encoded as
"servAreaRes" attribute; and/or
b) if the PCF received the "rfsp" attribute in the request, RAT Frequency Selection Priority (RFSP) Index
encoded as "rfsp" attribute; and/or
c) if the feature "UE-AMBR_Authorization" is supported and the PCF received the "ueAmbr" attribute in
the request, the authorized UE-AMBR encoded as "ueAmbr" attribute;

- optionally one or several of the following Policy Control Request Trigger(s) encoded as "triggers" attribute (see subclause 4.2.3.2):

a) Location change (tracking area); and
b) Change of UE presence in PRA; and
c) if the "SliceSupport" feature or the "DNNReplacementControl" feature is supported, change of allowed NSSAI; and
d) if the "DNNReplacementControl" feature is supported, change of SMF selection information; and

- if the Policy Control Request Trigger "Change of UE presence in PRA" is provided, the presence reporting areas for which reporting is required encoded as "pras" attribute;

- if the Policy Control Request Trigger "Change of SMF selection information" is provided, the SMF selection information representing the conditions upon which the AMF shall request a DNN replacement (see subclause 4.2.2.3.4) encoded as "smfSelInfo" attribute;

- if errors occur when processing the HTTP POST request, shall apply error handling procedures as specified in subclause 5.7 and according to the following provisions:
- if the user information received within the "supi" attribute is unknown, the PCF shall reject the request and
include in an HTTP "400 Bad Request" response message the "cause" attribute of the ProblemDetails data
structure set to "USER_UNKNOWN";
- if the PCF is, due to incomplete, erroneous or missing information in the request not able to provision an AM policy decision, the PCF may reject the request and include in an HTTP "400 Bad Request" response message the "cause" attribute of the ProblemDetails data structure set to
"ERROR_REQUEST_PARAMETERS".
- if the PCF rejects the AM policy association establishment, the NF service consumer shall apply the policy
retrieved from the UDM if available; otherwise, the NF service consumer shall apply the operator configured policy.

If the PCF received an GUAMI, the PCF may subscribe to GUAMI changes using the AMFStatusChange service operation of the Namf_Communication service specified in 3GPP TS 29.518 [14], and it may use the Nnrf_NFDiscovery Service specified in 3GPP TS 29.510 [13] (using the obtained GUAMI and possibly service name) to query the other AMFs within the AMF set.
If the PCF received a "traceReq" attribute, it shall perform trace procedures as defined in 3GPP TS 32.422 [18]. 

The provided information is a detailed description of how the Access and Mobility Management Function (AMF) in 5G communicates with the Policy Control Function (PCF) to obtain and establish policies for a User Equipment (UE) using an HTTP POST request. Let’s break it down and explain it step-by-step.

Explanation of the Procedure (Line by Line)
1. Initial Context Establishment
"When a UE registers and a UE context is being established..."
When a 5G device (UE) registers with the network, the AMF interacts with various network functions (NFs) to gather necessary information for creating a UE-specific context.
2. Retrieving Information
The AMF can obtain Service Area Restrictions, RFSP index, subscribed UE-AMBR, and GPSI from the UDM during the update location procedure..."


The AMF retrieves key details like:
Service Area Restrictions: Areas where the UE is allowed/disallowed.
RFSP Index: Priority for radio frequency resources.
UE-AMBR: Maximum bit rate for aggregated traffic.
GPSI: Generic Public Subscription Identifier, used to identify the UE.
These details come from the UDM (Unified Data Management) during a location update procedure.
"...and the allowed NSSAI from local configuration or from the NSSF during the slice selection procedure."


The AMF determines the allowed NSSAI (Network Slice Selection Assistance Information) from either its configuration or the NSSF (Network Slice Selection Function).
"...and shall decide based on local policies whether to request policies from the PCF."


Based on the information gathered and internal policies, the AMF determines whether it needs to interact with the PCF for policy decisions.

3. Requesting Policies from the PCF
"To request policies from the PCF, the NF service consumer (e.g., AMF) shall send an HTTP POST request..."
The AMF sends an HTTP POST request to the PCF at:
 {apiRoot}/npcf-am-policy-control/v1/policies
The request body includes a PolicyAssociationRequest structure.

4. Attributes in the Request Body
The request body includes the following:
Mandatory Attributes:


"notificationUri": URI for receiving notifications from the PCF.
"supi": The UE’s identifier.
Optional Attributes (when applicable):


"allowedSnssais": Allowed network slices for 3GPP access.
"gpsi": The UE’s public subscription identifier.
"accessType": Type of access (e.g., 3GPP or non-3GPP).
"pei": The UE’s permanent equipment identifier.
"userLoc": User location information.
"timeZone": UE’s time zone.
"servingPlmn": Serving PLMN and optional NID.
"ratType": Radio Access Technology type.
"servAreaRes": Service Area Restrictions mapped to Tracking Area Identities (TAIs).
"rfsp": RFSP index.
"groupIds": Internal group identifiers.
"guami": Globally Unique AMF Identifier.
"serviceName": Service name for receiving notifications.
"altNotifIpv4Addrs", "altNotifIpv6Addrs", "altNotifFqdns": Backup addresses for notifications.
"traceReq": Trace control and configuration.
"ueAmbr": Subscribed UE-AMBR for authorization.
"mappingSnssais": Mapping of allowed S-NSSAI to HPLMN S-NSSAI.

5. Response Handling by PCF
Upon receiving the request:
The PCF assigns a policy association ID and determines the applicable policy.
It modifies or validates attributes like UE-AMBR, Service Area Restrictions, or RFSP index.
In a successful case:
PCF responds with HTTP 201 Created.
It provides the URI for the created resource and includes the PolicyAssociation data structure.
In case of errors:
PCF returns HTTP 400 Bad Request with an appropriate error message.

Use in 5G
The above procedure is used in 5G networks whenever the AMF needs policy decisions for managing UE access and mobility. Typical use cases include:
Initial Registration: When a UE first connects to the network.
Mobility Events: E.g., when a UE moves to a new tracking area.
Slice Selection: To enforce slice-specific policies.

Call Flow in 5G
Step-by-Step Call Flow
UE Registration:
UE sends a registration request to the AMF.
AMF interacts with the UDM to fetch subscription details (e.g., GPSI, RFSP, Service Area Restrictions).
Policy Association Request:
AMF decides to request policies from the PCF.
AMF sends an HTTP POST request to the PCF containing the PolicyAssociationRequest structure.
PCF Processing:
PCF assigns a policy association ID.
Determines the applicable policy based on received data and local rules.
Policy Decision:
PCF responds with HTTP 201 Created.
Includes the PolicyAssociation structure with the applicable policies and triggers (e.g., location changes).
Policy Enforcement:
AMF enforces the received policy (e.g., allowed slices, service area restrictions).
Diagram
[UE] -----> [AMF] -----> [UDM]
       Registration       |
                        Fetch Subscription Info
                            |
               [AMF] -----> [PCF]
                     Policy Association Request
                            |
               [PCF] -----> [AMF]
               Policy Decision Response
                            |
               Policy Enforcement by AMF

Would you like me to expand on specific steps, add more details, or focus on implementation in Go?

