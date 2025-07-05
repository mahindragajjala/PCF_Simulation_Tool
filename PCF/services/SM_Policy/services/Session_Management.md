	https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.tb1rcm9tg91p		

OVERVIEW

The 5GC supports a PDU Connectivity Service i.e. a service that provides exchange of PDUs between a UE and a data network identified by a DNN. 

The PDU Connectivity Service is supported via PDU Sessions that are established upon request from the UE.
The Subscription Information for each S-NSSAI may contain a Subscribed DNN list and one default DNN. 
When the UE does not provide a DNN in a NAS Message containing PDU Session Establishment Request for a given S-NSSAI,the serving AMF determines the DNN for the requested PDU Session by selecting the default DNN for this S-NSSAI if
A default DNN is present in the UE's Subscription Information; otherwise the serving AMF selects a locally configured DNN for this S-NSSAI.
The expectation is that the URSP in the UE is always up to date using the procedure defined in TS 23.502 [3] clause 4.16.12.2 and therefore the UE requested DNN will be up to date.
In order to cover cases that UE operates using local configuration, but also other cases where operator policies can be used in order to replace an "up to date" UE requested DNN with another DNN used only internally in the network, during UE Registration procedure the PCF may indicate, to the AMF, the operator policies to be used at PDU Session Establishment for DNN replacement of a UE requested DNN. 
PCF may indicate a policy for DNN replacement of UE requested DNNs not supported by the network, and/or indicate a list of UE requested DNNs per S-NSSAI valid for the
serving network, that are subject for replacement (details are described in TS 23.503 [45]).
If the DNN provided by the UE is not supported by the network and AMF cannot select an SMF by querying NRF, the AMF shall reject the NAS Message containing PDU Session Establishment Request from the UE with a cause indicating that the DNN is not supported unless the PCF provided the policy to perform a DNN replacement of
unsupported DNNs.
If the DNN requested by the UE is indicated for replacement or the DNN provided by the UE is not supported by the network and the PCF provided the policy to perform DNN replacement of UE requested DNNs not supported by the
network, the AMF shall interact with the PCF to perform a DNN replacement. 
During PDU Session Establishment procedure and as a result of DNN replacement, the PCF provides the selected DNN that is applicable for the S-NSSAI
requested by the UE at the PDU Session Establishment. 
The AMF uses the selected DNN in the query towards the NRF for the SMF selection, as specified in clause 6.3.2, and provides both requested and selected DNN to the selected SMF.
For PDU Session with Home-routed Roaming whether to perform DNN replacement is based on operator agreements.




NOTE 1: 

The selected DNN is determined based on operator preferences and can differ from subscribed DNNs.
The matching of selected DNN to S-NSSAI is assumed to be based on network configuration.
Each PDU Session supports a single PDU Session type i.e. supports the exchange of a single type of PDU requested by the UE at the establishment of the PDU Session. 
The following PDU Session types are defined: IPv4, IPv6, IPv4v6,Ethernet, Unstructured.
PDU Sessions are established (upon UE request), modified (upon UE and 5GC request) and released (upon UE and 5GC request) using NAS SM signalling exchanged over N1 between the UE and the SMF. 
Upon request from an Application Server, the 5GC is able to trigger a specific application in the UE. 
When receiving that trigger message, the UE shall pass it to the identified application in the UE. The identified application in the UE may establish a PDU Session to a specific DNN, see clause 4.4.5.


SMF may support PDU Sessions for LADN where the access to a DN is only available in a specific LADN service area. This is further defined in clause 5.6.5.
SMF may support PDU Sessions for a 5G VN group which offers a virtual data network capable of supporting 5G LAN-type service over the 5G system. 
This is further defined in clause 5.8.2.13.
The SMF is responsible of checking whether the UE requests are compliant with the user subscription. For this purpose, it retrieves and requests to receive update notifications on SMF level subscription data from the UDM. 
Such data may indicate per DNN and per S-NSSAI of the HPLMN:
- The allowed PDU Session Types and the default PDU Session Type.
- The allowed SSC modes and the default SSC mode.
- QoS Information (refer to clause 5.7): the subscribed Session-AMBR, Default 5QI and Default ARP.
- The static IP address/prefix.
- The subscribed User Plane Security Policy.
- the Charging Characteristics to be associated with the PDU Session. 

Whether this information is provided by the UDM to a SMF in another PLMN (for PDU Sessions in LBO mode) is defined by operator policies in the UDM/UDR.


NOTE 2: 

The content of the Charging Characteristics as well as the usage of the Charging Characteristics by the SMF are defined in TS 32.240 [41].

A PDU Session may support:

(a) a single-access PDU Connectivity Service, in which case the PDU Session is associated with a single access type at a given time, i.e. either 3GPP access or non-3GPP access; or
(b) a multi-access PDU Connectivity Service, in which case the PDU Session is simultaneously associated with both 3GPP access and non-3GPP access and simultaneously associated with two independent N3/N9 tunnels between the PSA and RAN/AN.

A PDU Session supporting a single-access PDU Connectivity Service is also referred to as single-access PDU Session, while a PDU Session supporting a multi-access PDU Connectivity Service is referred to as Multi-Access PDU (MAPDU) Session and it is used to support the ATSSS feature (see clause 5.32 for details).


A UE that is registered over multiple accesses chooses over which access to establish a PDU Session. As defined in TS 23.503 [45], the HPLMN may send policies to the UE to guide the UE selection of the access over which to establish a PDU Session.
The key differences between Multi-Access PDU Session and Single-Access PDU Session in 5G are:
Feature
Multi-Access PDU Session
Single-Access PDU Session
Access Network Usage
Uses multiple access networks simultaneously (e.g., 5G + Wi-Fi).
Uses a single access network (e.g., only 5G).
Data Flow
Data can flow through multiple interfaces concurrently.
Data flows through a single interface.
Redundancy
Offers redundancy by utilizing multiple paths.
No redundancy; depends on a single connection.
Bandwidth Aggregation
Aggregates bandwidth from multiple access networks.
Bandwidth is limited to one access network.
Latency Optimization
Can reduce latency by offloading traffic to the optimal access.
Latency depends on the single network's conditions.
Use Cases
Ideal for applications requiring high reliability, bandwidth, or seamless handover.
Suitable for simpler or less demanding applications.



NOTE 3: 

In this Release of the specification, at any given time, a PDU Session is routed over only a single access network, unless it is an MA PDU Session in which case it can be routed over one 3GPP access network and one Non 3GPP access network concurrently.
A UE may request to move a single-access PDU Session between 3GPP and Non 3GPP accesses. 
The decision to move single-access PDU Sessions between 3GPP access and Non 3GPP access is made on a per PDU Session basis, i.e. the UE may, at a given time, have some PDU Sessions using 3GPP access while other PDU Sessions are using Non 3GPP access.
If the UE is attempting to move a single-access PDU session from 3GPP access to non-3GPP access and the PDU session is associated with control plane only indication, then the AMF shall reject the PDU Session Establishment request as related CIoT 5GS optimisation features are not supported over non-3GPP access as described in clause 5.4.5.2.5 of TS 24.501 [47]. 
If the UE is attempting to move a single-access PDU session from non-3GPP access
to NB-N1 mode of 3GPP access, the PDU Session Establishment request would also be rejected by AMF when the UP resources for the UE exceed the maximum number of supported UP resources as described in clause 5.4.5.2.4 of TS 24.501 [47].



In a PDU Session Establishment Request message sent to the network, the UE shall provide a PDU Session ID. 
The PDU Session ID is unique per UE and is the identifier used to uniquely identify one of a UE's PDU Sessions. 
The PDU Session ID shall be stored in the UDM to support handover between 3GPP and non-3GPP access when different PLMNs are used for the two accesses. 
The UE also provides as described in TS 24.501 [47]:


(a) PDU Session Type.

(b) S-NSSAI of the HPLMN that matches the application (that is triggering the PDU Session Request) within the NSSP in the URSP rules or within the UE Local Configuration as defined in clause 6.1.2.2.1 of TS 23.503 [45].

NOTE 4: 

If the UE cannot determine any S-NSSAI after performing the association of the application to a PDU Session, then it does not indicate any S-NSSAI in the PDU Session Establishment procedure as defined in clause 5.15.5.3.

(c) S-NSSAI of the Serving PLMN from the Allowed NSSAI, corresponding to the S-NSSAI of the HPLMN (b).

NOTE 5: 
Generally, in non-roaming scenario the mapping of the Allowed NSSAI to HPLMN S-NSSAIs is not provided to the UE (because the S-NSSAI of the Serving PLMN (c) has the same value of the S-NSSAI of the HPLMN (b)), therefore the UE provides in the PDU Session Request only the S-NSSAI of the Serving PLMN (c). 

However, if the UE is provided with the mapping of the Allowed NSSAI to HPLMN S-NSSAIs even in non-roaming scenario, then the UE provides in the PDU Session Request both the SNSSAI of the HPLMN (b) and the S-NSSAI from the Allowed NSSAI (c) that maps to the S-NSSAI of the HPLMN.


NOTE 6: In roaming scenarios the UE provides in the PDU Session Request both the S-NSSAI of the HPLMN (b) and the S-NSSAI of the VPLMN from the Allowed NSSAI (c) that maps to the S-NSSAI of the HPLMN.

(d) DNN (Data Network Name).
(e) SSC mode (Service and Session Continuity mode defined in clause 5.6.9.2).

Additionally, if the UE supports ATSSS and wants to activate a MA PDU Session, the UE shall provide Request Type as "MA PDU Request" and shall indicate the supported ATSSS capabilities (see clause 5.32 for details).


