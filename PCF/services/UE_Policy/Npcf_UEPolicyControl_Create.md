https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.94zuwug2q53z

The procedure in the present subclause is applicable when the NF service consumer creates a UE policy association in the following cases:

- UE initial registers to the network as defined in subclause 5.5.1.2.2 of 3GPP TS 24.501 [15];

- UE performs the mobility registration if the UE operating in the single-registration mode performs inter-system change from S1 mode to N1 mode as defined in subclause 5.5.1.3.2 of 3GPP TS 24.501 [15] and there is no existing UE Policy Association between AMF and PCF for this UE;

- the AMF is relocated (between the different AMF sets) and the new AMF selects a new PCF. 
The procedure for the case where the AMF is relocated and the new AMF selects the old PCF is defined in subclause 4.2.3.1.

The creation of an UE policy association only applies for normally registered UEs, i.e., it does not apply for emergency registered UEs.

Figure 4.2.2.1-1 illustrates the creation of a policy association.
NOTE 1: For the roaming case, the PCF represents the V-PCF if the NF service consumer is an AMF and the PCF represents the H-PCF if the NF service consumer is a V-PCF.

When a UE registers and a UE context is being established, if the AMF obtains from the UE an UE policy delivery protocol message as defined in Annex D of 3GPP TS 24.501 [15] the AMF shall establish a UE policy association with the (V-)PCF in case that there is no existing UE policy association for the UE; otherwise the AMF may establish UE Policy Association with the (V-)PCF based on AMF local configuration.

NOTE 2: In roaming scenario, the AMF local configuration can indicate whether UE Policy delivery is needed based on the roaming agreement with home PLMN of the UE.

To establish a UE policy association with the PCF, the NF service consumer (e.g. AMF) shall send an HTTP POST request with: "{apiRoot}/npcf-ue-policy-control/v1/policies/" as Resource URI and the PolicyAssociationRequest data structure as request body that shall include:

- Notification URI encoded as "notificationUri" attribute; and
- SUPI encoded as "supi" attribute,
and that shall include when available:
- GPSI encoded as "gpsi" attribute;
- Access type encoded as "accessType" attribute;
- Permanent Equipment Identifier (PEI) encoded as "pei" attribute;
- User Location Information encoded as "userLoc" attribute;
- UE Time Zone encoded as "timeZone" attribute;
- Serving PLMN Identifier and for SNPN the NID encoded as "servingPlmn" attribute;
- RAT type encoded as "ratType" attribute;
- the received UE policy delivery protocol message defined in Annex D of 3GPP TS 24.501 [15] or defined in subclause 7.2.1.1 of 3GPP TS 24.587 [24] encoded as "uePolReq" attribute;
- if the NF service consumer is an AMF, H-PCF ID (if the consumer is V-PCF, when receiving the H-PCF ID from AMF) encoded as "hPcfId" attribute;
- Internal Group Identifier(s) encoded as "groupIds" attribute;
- the PC5 capability for V2X encoded as "pc5Capab" attribute if the "V2X" feature defined in subclause 5.8 is suppoted;
- if the NF service consumer is an AMF, the GUAMI encoded as "guami" attribute;
- if the NF service consumer is an AMF, the name of a service produced by the AMF that expects to receive information within Npcf_UEPolicyControl_UpdateNotify service operation encoded as "serviceName" attribute;
- if the NF service consumer is an AMF, alternate or backup IPv4 Address(es) where to send Notifications encoded as "altNotifIpv4Addrs" attribute;
- if the NF service consumer is an AMF, alternate or backup IPv6 Address(es) where to send Notifications encoded as "altNotifIpv6Addrs" attribute;
- if the NF service consumer is an AMF, alternate or backup FQDN(s) where to send Notifications encoded as "altNotifFqdns" attribute; and
- if the NF service consumer is an AMF, serving AMF Id encoded in the "servingNfId" attribute.

Upon the reception of the HTTP POST request,
- the (V-)(H-)PCF shall assign a UE policy association ID;
- based on operator policy the V-PCF should send as the NF service consumer towards the H-PCF a request for the Creation of a UE policy association as described in the present clause;
- the (V-)(H-)PCF shall determine the applicable UE policy as detailed in subclause 4.2.2.2, for the V-PCF taking into consideration any policy received from the H-PCF in the reply to the possible request for the Creation of a policy association;
- if the (V-)PCF determines that UE policy needs to be provisioned, it shall use the Namf_Communication service
specified in 3GPP TS 29.518 [14] to provision the UE policy according to subclause 4.2.2.2 and as follows:
(i) the V-PCF shall subscribe at the AMF to notifications of N1 messages for UE Policy Delivery Results using
the Namf_Communication_N1N2MessageSubscribe service operation;
(ii) the V-PCF shall send the determined UE policy using Namf_Communication_N1N2MessageTransfer service
operation(s); and
(iii) the V-PCF shall be prepared to receive UE Policy Delivery Results from the AMF within the
Namf_Communication_N1MessageNotify service operation and for the V-PCF if the received UE Policy
Delivery results relate to UE policy sections provided by the H-PCF shall use the
Npcf_UEPolicyControl_Update Service Operation to send those UE Policy Delivery results to the H-PCF;
- If the UE indicates the support of V2X communications over PC5 reference point and "V2X" feature is
supported, the (H-)PCF shall determine the applicable N2 PC5 policy as detailed in subclause 4.2.2.3 based on
the subscription data and operator’s policy;
- if the (V-)PCF determines that N2 PC5 policy needs to be provisioned, it shall use the Namf_Communication
service specified in 3GPP TS 29.518 [14] to provision the N2 PC5 policy according to subclause 4.2.2.3.
- for the succesfull case the (V-)(H-)PCF shall send a HTTP "201 Created" response with the URI for the created
resource in the "Location" header field
NOTE 3: The assigned policy association ID is part of the URI for the created resource and is thus associated with
the SUPI.
and the the PolicyAssociation data type as body including:
- optionally for the H-PCF as service producer communicating with the V-PCF, UE policy (see
subclause 4.2.2.2) encoded as "uePolicy" attribute;
- optionally for the H-PCF as service producer communicating with the V-PCF, N2 PC5 policy (see
subclause 4.2.2.3) encoded as "n2Pc5Pol" attribute;
- optionally one or several of the following Policy Control Request Trigger(s) encoded as "triggers" attribute
(see subclause 4.2.3.2):
a) Location change (tracking area); and
b) Change of UE presence in PRA; and
