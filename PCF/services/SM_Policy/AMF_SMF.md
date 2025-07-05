https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.bswao72d2i8j

Interaction between AMF and SMF
The AMF and SMF are separate Network Functions.

N1 related interaction with SMF is as follows:

The single N1 termination point is located in AMF. The AMF forwards SM related NAS information to the SMF based on the PDU Session ID in the NAS message. 

Further SM NAS exchanges (e.g. SM NAS message responses) for N1 NAS signalling received by the AMF over an access (e.g. 3GPP access or non-3GPP access)
are transported over the same access.
The serving PLMN ensures that subsequent SM NAS exchanges (e.g. SM NAS message responses) for N1 NAS signalling received by the AMF over an access (e.g. 3GPP access or non-3GPP access) are transported over the same access.
SMF handles the Session management part of NAS signalling exchanged with the UE.
The UE shall only initiate PDU Session Establishment in RM-REGISTERED state.
When a SMF has been selected to serve a specific PDU Session, AMF has to ensure that all NAS signalling related with this PDU Session is handled by the same SMF instance.
Upon successful PDU Session Establishment, the AMF and SMF stores the Access Type that the PDU Session is associated.


N11 related interaction with SMF is as follows:

The AMF reports the reachability of the UE based on a subscription from the SMF, including:
The UE location information with respect to the area of interest indicated by the SMF.
The SMF indicates to AMF when a PDU Session has been released.
Upon successful PDU Session Establishment, AMF stores the identification of serving SMF of UE and SMF stores the identification of serving AMF of UE including the AMF set. 
When trying to reach the AMF serving the UE, the SMF may need to apply the behaviour described for "the other CP NFs" in clause 5.21.


N2 related interaction with SMF is as follows:

Some N2 signalling (such as handover related signalling) may require the action of both AMF and SMF. 
In such case, the AMF is responsible to ensure the coordination between AMF and SMF. 
The AMF may forward the SM N2 signalling towards the corresponding SMF based on the PDU Session ID in N2 signalling.
SMF shall provide PDU Session Type together with PDU Session ID to NG-RAN, in order to facilitate NGRAN to apply suitable header compression mechanism to packet of different PDU type. Details refer to TS 38.413 [34].


N3 related interaction with SMF is as follows:
The N3 interface in 5G is established when the SMF configures the UPF with the necessary tunnel information and the AMF successfully completes the PDU Session Establishment procedure with the UE.

Selective activation and deactivation of UP connection of existing PDU Session is defined in clause 5.6.8.

N4 related interaction with SMF is as follows:

When it is made aware by the UPF that some DL data has arrived for a UE without downlink N3 tunnel information, the SMF interacts with the AMF to initiate Network Triggered Service Request procedure. 

In this case, if the SMF is aware that the UE is unreachable or if the UE is reachable only for regulatory prioritized service and the PDU Session is not for regulatory prioritized service, then the SMF shall not inform DL data notification to the AMF

- When it forwards UL NAS or N2 signalling to a peer NF (e.g. to SMF or to SMSF) or during the UP connection activation of a PDU Session, the AMF provides any User Location Information it has received from the 5G-AN as well as the Access Type (3GPP - Non 3GPP) of the AN over which it has received the UL NAS or N2 signalling. 

The AMF also provides the corresponding UE Time Zone. 
In addition, in order to fulfil regulatory requirement (i.e. providing Network Provided Location Information (NPLI), as defined in TS 23.228 [15]) when the access is non-3GPP, the AMF may also provide the last known 3GPP access User Location Information with its age, if the UE is still attached to the same AMF for 3GPP access (i.e. valid User Location Information).
The User Location Information, the access type and the UE Time Zone may be further provided by SMF to PCF. 
The PCF may get this information from the SMF in order to provide NPLI to applications (such as IMS) that have requested it.
Here’s a simplified explanation of the above interaction and the call flow breakdown:

Key Points:
Downlink Data Notification (DL Data Notification):


When DL data arrives for a UE that has no downlink N3 tunnel information, the SMF informs the AMF to trigger a Network Triggered Service Request procedure.
Exception:
If the SMF knows that the UE is unreachable or can only be reached for regulatory services (and the session is not for such services), the SMF does not notify the AMF about the DL data.
Uplink Signaling with Location Details:


When the AMF forwards Uplink NAS or N2 signaling to another Network Function (e.g., SMF or SMSF):
The AMF includes:
User Location Information (ULI): The location data received from the 5G-AN (Access Network).
Access Type: Whether the signal is over 3GPP (e.g., 5G NR) or Non-3GPP (e.g., Wi-Fi).
UE Time Zone: Time zone information of the UE.
Non-3GPP Access with Regulatory Requirements:


For Non-3GPP access, if there’s a regulatory requirement (e.g., for Network Provided Location Information (NPLI)):
The AMF may provide the last known 3GPP User Location Information along with its age (if the UE is still attached to the same AMF for 3GPP access).
SMF to PCF Interaction:


The SMF can pass the User Location Information, Access Type, and UE Time Zone to the PCF.
The PCF may need this data to provide NPLI to applications like IMS that require location-based services.

Call Flow with Nodes:
Here’s how the interactions happen:
Call Flow Diagram:
UE            5G-AN         AMF          SMF          PCF          Applications
 |              |            |            |            |                 |
 |--- UL NAS/N2 Signaling --->|            |            |                 |
 |              |--- Forward Signaling --->|            |                 |
 |              |            |--- Notify Downlink Data-->|                 |
 |              |            |            |--- User Location, Access Info --->|
 |              |            |            |            |--- Pass Info --->|
 |              |            |            |            |             Location Service

Explanation of the Flow:
Uplink Signaling:


The UE sends an uplink NAS/N2 signaling message to the 5G-AN.
The 5G-AN forwards this to the AMF, including any User Location Information it has.
Forwarding by AMF:


The AMF forwards the signaling message to the appropriate Network Function (e.g., SMF or SMSF).
The AMF also includes:
User Location Information (from the 5G-AN).
Access Type (3GPP/Non-3GPP).
UE Time Zone.
Downlink Data Notification:


If the SMF receives DL data for a UE without an active N3 tunnel:
It informs the AMF to initiate a Network Triggered Service Request.
If the UE is unreachable or restricted to regulatory services, the SMF does not notify the AMF.
Regulatory Compliance:


For Non-3GPP Access, the AMF may include the last known 3GPP User Location Information (with its age) if required by regulations.
SMF to PCF:


The SMF sends User Location Information, Access Type, and UE Time Zone to the PCF.
The PCF uses this information to provide location-based services (e.g., to IMS applications).


The User Location Information may correspond to:
In the case of 3GPP access: Cell-Id. The AMF includes only the Primary Cell-Id even if it had received also the Cell-Id of the Primary cell in the Secondary RAN node from NG-RAN.
In the case of Untrusted non-3GPP access: a UE local IP address (used to reach the N3IWF) and optionally UDP or TCP source port number (if NAT is detected).
In the case of Trusted non-3GPP access: TNAP/TWAP Identifier, a UE/N5CW device local IP address (used to reach the TNGF/TWIF) and optionally UDP or TCP source port number (if NAT is detected).
When the UE uses WLAN based on IEEE 802.11 technology to reach the TNGF, the TNAP Identifier shall include the SSID of the access point to which the UE is attached. The TNAP Identifier shall include at least one of the following elements, unless otherwise determined by the TWAN operator's policies:

the BSSID (see IEEE Std 802.11-2012 [106]);
civic address information of the TNAP to which the UE is attached.

The TWAP Identifier shall include the SSID of the access point to which the NC5W is attached. The TWAP Identifier shall also include at least one of the following elements, unless otherwise determined by the TWAN
operator's policies:
- the BSSID (see IEEE Std 802.11-2012 [106]);
- civic address information of the TWAP to which the UE is attached.

NOTE 1: The SSID can be the same for several TNAPs/TWAPs and SSID only cannot provide a location, but it might be sufficient for charging purposes.

NOTE 2: the BSSID associated with a TNAP/TWAP is assumed to be static.

In the case of W-5GAN access: The User Location Information for W-5GAN is defined in TS 23.316 [84].
When the SMF receives a request to provide Access Network Information reporting while there is no action to carry out towards the 5G-AN or the UE (e.g. no QoS flow to create / Update / modify), the SMF may request User Location Information from the AMF.
The interaction between AMF and SMF(s) for the case of a I-SMF insertion, relocation or removal for a PDU session is described in clause 5.34.
