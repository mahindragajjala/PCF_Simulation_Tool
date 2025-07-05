https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.ov0arx9rc5z8

The 5GC shall be able to provide policy information from the PCF to the UE. Such policy information includes:
Access Network Discovery & Selection Policy (ANDSP):
 It is used by the UE for selecting non-3GPP accesses and for selection of the N3IWF in the PLMN. The structure and the content of this policy are specified in clause 6.6.1.


UE Route Selection Policy (URSP):
 This policy is used by the UE to determine if a detected application can be associated to an established PDU Session, can be offloaded to non-3GPP access outside a PDU Session, or can trigger the establishment of a new PDU Session. The structure and the content of this policy are specified in clause 6.6.2. A URSP rule includes one Traffic descriptor that specifies the matching criteria and one or more of the following components:
 2a) SSC Mode Selection Policy (SSCMSP):
 This is used by the UE to associate the matching application with SSC modes.

 2b) Network Slice Selection Policy (NSSP):
 This is used by the UE to associate the matching application with SNSSAI.

 2c) DNN Selection Policy:
 This is used by the UE to associate the matching application with DNN.

 2d) PDU Session Type Policy:
 This is used by the UE to associate the matching application with a PDU Session Type.

 2e) Non-Seamless Offload Policy:
 This is used by the UE to determine that the matching application should be non-seamlessly offloaded to non-3GPP access (i.e. outside of a PDU Session).

 2f) Access Type Preference:
 If the UE needs to establish a PDU Session for the matching application, this indicates the preferred Access Type (3GPP or non-3GPP or Multi-Access).


The ANDSP and URSP may be pre-configured in the UE or may be provisioned to UE from PCF. The pre-configured policy shall be applied by the UE only when it has not received the same type of policy from PCF.
The PCF selects the ANDSP and URSP applicable for each UE based on local configuration and operator policies, taking into consideration the information defined in clause 6.2.1.2.
In the case of a roaming UE, the V-PCF may retrieve ANDSP and URSP from the H-PCF over N24/Npcf. When the UE is roaming and the UE has valid rules from both HPLMN and VPLMN, the UE gives priority to the valid ANDSP rules from the VPLMN.
The ANDSP and URSP shall be provided from the PCF to the AMF via the N15/Namf interface and then from AMF to the UE via the N1 interface as described in TS 23.502 [3] clause 4.2.4.3. The AMF shall not change the ANDSP and the URSP provided by PCF.
The PCF is responsible for the delivery of UE policy. If the PCF is notified about UE Policy delivery failure (e.g., because of UE unreachable), the PCF may provide a new trigger "Connectivity state changes" in Policy Control Request Trigger of UE Policy Association to AMF as defined in TS 23.502 [3] clause 4.16.12.2. After reception of the Notify message indicating that the UE enters the CM-Connected state, the PCF may retry to deliver the UE Policy.
NOTE 1: For backward compatibility, the PCF may subscribe to the "Connectivity state changes (IDLE or CONNECTED)" event in Rel-15 AMF as defined in TS 23.502 [3] clause 5.2.2.3.
If due to UE Local Configurations, a UE application requests a network connection using Non-Seamless Offload, the UE shall use Non-Seamless Offload for this application without evaluating the URSP rules. Otherwise, the UE shall select the PDU Session or Non-Seamless Offload in the following order:
If the UE has a URSP rule (except the URSP rule with the "match all" Traffic descriptor) that matches the application as defined in clause 6.6.2.3, the UE shall perform the association of the application to the corresponding PDU Session or to Non-Seamless Offload according to this rule. Otherwise,


If no URSP rule is applicable for the application (except the URSP rule with the "match all" Traffic descriptor), the UE shall perform the association of the application to a PDU Session according to the applicable UE Local Configurations, if any. If the UE attempts to establish a new PDU Session according to the UE Local Configurations and this PDU Session Establishment request is rejected by the network, then the UE shall perform the association of the application to a PDU Session or to Non-Seamless Offload according to the URSP rule with the "match all" Traffic descriptor. Otherwise,


NOTE 2: It is assumed that the S-NSSAI(s) in the UE Local Configurations are operator-provided S-NSSAI(s). The provision of the S-NSSAI(s) is not specified.
 NOTE 3: The application layer is not allowed to set the S-NSSAI when the UE establishes a PDU Session based on the UE Local Configurations.
 NOTE 4: Any missing information in the UE Local Configurations needed to build the PDU Session Establishment request can be the appropriate corresponding component from the URSP rule with the "match all" Traffic descriptor.
If neither the UE Local Configurations nor the URSP rules are applicable for the application (except the URSP rule with the "match all" Traffic descriptor), the UE shall perform the association of the application to a PDU Session or to Non-Seamless Offload according to the URSP rule with the "match all" Traffic descriptor.
For the existing PDU Session(s), the UE shall examine the URSP rules within the UE Policy in order to determine whether the existing PDU Session(s) (if any) are maintained or not. If not, then the UE may initiate a PDU Session release procedure for the PDU Session(s) that cannot be maintained.
If there are multiple IPv6 prefixes within the PDU Session, then the IPv6 multi-homed routing rules, described in clause 5.8.2.2.2 in TS 23.501 [2], on the UE shall be used to select which IPv6 prefix to route the traffic of the application.
NOTE 5: For the case that an application cannot be associated with any PDU Session.


Understanding PCF Policy Information Delivery in 5G:
In a 5G network, the Policy Control Function (PCF) provides policy rules to the User Equipment (UE). 
These rules guide the UE on how to behave in specific scenarios, such as selecting networks, creating sessions, and routing traffic.
The policies discussed here are:
Access Network Discovery and Selection Policy (ANDSP)
UE Route Selection Policy (URSP)
Let's break these down.
1. Access Network Discovery and Selection Policy (ANDSP)
Purpose:
 This policy helps the UE decide which non-3GPP network to connect to (e.g., Wi-Fi) and which N3IWF (Non-3GPP Interworking Function) to use in the network.


Example: If a UE detects multiple Wi-Fi networks, ANDSP guides it to choose the one preferred by the operator.
How it works:


The ANDSP can either be:
Pre-configured in the UE (factory settings or manual configuration).
Provisioned by the PCF.
If the PCF provides ANDSP, the UE uses it. Pre-configured policies are a fallback if the PCF policies aren't received.
2. UE Route Selection Policy (URSP)
Purpose:
 URSP helps the UE decide how to handle application traffic:


Which established PDU Session (Packet Data Unit Session) to use.
Whether to create a new PDU Session for the traffic.
Whether to offload the traffic to a non-3GPP network.
Components of URSP:
Traffic Descriptor: Defines the types of traffic (e.g., video streaming, email) the rule applies to.
SSC Mode Selection Policy (SSCMSP): Guides the UE on which SSC (Session and Service Continuity) mode to use (e.g., seamless session continuity or connection rebuild on demand).
Network Slice Selection Policy (NSSP): Maps traffic to specific slices (e.g., enhanced mobile broadband, ultra-reliable low-latency communication).
DNN Selection Policy: Determines which Data Network Name (DNN) (e.g., internet or private network) to use for traffic.
PDU Session Type Policy: Defines the type of session (IPv4, IPv6, or Ethernet) for the traffic.
Non-Seamless Offload Policy: Routes traffic outside the PDU session directly to non-3GPP access.
Access Type Preference: Indicates whether traffic should use 3GPP, non-3GPP, or both types of networks.
How URSP Works:
URSP rules are matched against application traffic.
If a rule matches, the UE follows it (e.g., use an existing PDU session, create a new one, or offload traffic).
If no URSP rule matches:
The UE falls back to local configurations (operator-provided settings).
If those fail, a default rule (called the "match all" rule) is applied.
Call Flows
Provisioning ANDSP and URSP to the UE
Policy Selection by PCF:


The PCF determines the applicable ANDSP and URSP rules based on:
Local configuration.
Operator policies.
Information such as subscriber profile and network conditions.
ANDSP/URSP Delivery:


The PCF sends the ANDSP/URSP rules to the AMF (Access and Mobility Management Function) over the N15/Namf interface.
The AMF relays these rules to the UE over the N1 interface (control plane signaling).
Roaming Scenario:


If the UE is roaming:
The visited PCF (V-PCF) may retrieve policies from the home PCF (H-PCF) over the N24/Npcf interface.
The UE prioritizes ANDSP rules from the visited network (VPLMN) over the home network (HPLMN).
Handling UE Application Traffic
Application Request:
When a UE application (e.g., video streaming app) requests a network connection:
The UE checks the URSP rules for matching criteria.
Matching Traffic with URSP:
If a URSP rule matches:
The UE associates the application traffic with an existing PDU session or establishes a new one based on the rule.
If no URSP rule matches:
The UE checks local configurations.
If those fail, the "match all" rule is applied as a fallback.
Non-Seamless Offload:
If the UE local configurations indicate non-seamless offload for the application, the traffic is routed to non-3GPP access directly.
Handling Policy Delivery Failures
Initial Delivery:
The PCF sends policies to the AMF, which forwards them to the UE.
Delivery Failure:
If the UE is unreachable (e.g., in IDLE state):
The AMF notifies the PCF of the failure.
The PCF sets a trigger: "Connectivity state changes."
Retry on Reconnection:
When the UE transitions to CONNECTED state, the AMF notifies the PCF.
The PCF retries sending the policies.

PDU Session Maintenance
The UE periodically checks whether existing PDU sessions comply with updated URSP rules:
If not, the UE may release the non-compliant PDU sessions.

Summary of Call Flow Interfaces
N15/Namf: PCF to AMF communication.
N1: AMF to UE signaling for policy delivery.
N24/Npcf: V-PCF to H-PCF communication for retrieving policies in roaming scenarios.

Key Notes
Policy Hierarchy:
PCF-provided policies take precedence over pre-configured policies in the UE.
Roaming Priority:
ANDSP rules from the visited network (VPLMN) are prioritized over the home network (HPLMN).
Fallback Mechanism:
Local configurations and the "match all" rule ensure connectivity even if specific rules are unavailable.
