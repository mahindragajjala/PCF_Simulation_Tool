https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.8l38i3m7bf8e

The URSP is defined in 3GPP TS 23.503 [2] and is a set of one or more URSP rules, where a URSP rule is composed of:
a) a precedence value of the URSP rule identifying the precedence of the URSP rule among all the existing URSP rules;

b) a traffic descriptor, including either:
1) match-all traffic descriptor; or
2) at least one of the following components:
A) one or more application identifiers;
B) one or more IP 3 tuples as defined in 3GPP TS 23.503 [2] i.e. the destination IP address, the destination port number, and the protocol in use above the IP;
C) one or more non-IP descriptors, i.e. destination information of non-IP traffic;
D) one or more DNNs;
E) one or more connection capabilities; and
F) one or more domain descriptors, i.e. destination FQDN(s) or a regular expression as a domain name matching criteria; and

c) one or more route selection descriptors each consisting of a precedence value of the route selection descriptor and either
1) one PDU session type and, optionally, one or more of the followings:
A) SSC mode;
B) one or more S-NSSAIs. If the URSP rule is a part of a non-subscribed SNPN signalled URSP, the SNSSAI is of the non-subscribed SNPN otherwise the S-NSSAI is of the HPLMN or the subscribed
SNPN. Mapped HPLMN SST and mapped HPLMN SD are not included in the S-NSSAI;
C) one or more DNNs;
D) Void;
E) preferred access type;
F) multi-access preference;
G) a time window;
H) location criteria;
I) PDU session pair ID; and
J) RSN;

2) non-seamless non-3GPP offload indication; or
3) 5G ProSe layer-3 UE-to-network relay offload indication.

Only one URSP rule in the URSP can be a default URSP rule and the default URSP rule shall contain a match all traffic descriptor. 
If a default URSP rule and one or more non-default URSP rules are included in the URSP, any non-default URSP rule shall have lower precedence value than (i.e. shall be prioritised over) the default URSP rule.
If a traffic descriptor lists one or more application identifiers together with one or more connection capabilities, the UE shall consider that the application identifiers identify the applications requesting access to the connection capabilities.

NOTE 1: The connection capabilities requested by the applications are OS dependent. 

The connection capability identifiers defined in table 5.2.1 are OS independent. 

It is based on the UE implementation how the UE matches the connection capabilities requested by the applications to the connection capability identifiers in table 5.2.1.


NOTE 2: If the UE has multiple concurrently active OS, the traffic descriptor can list as many multiple OS Ids.

NOTE 3: It is recommended to avoid the combination of more than two components in the traffic descriptor.



Here’s a simple and clear breakdown of the User Route Selection Policy (URSP) in 5G based on the provided information:
1. What is URSP?
URSP (User Route Selection Policy) is a set of rules that help determine how user traffic is routed in a 5G network.
It is defined in 3GPP TS 23.503.
Each URSP consists of one or more URSP rules.
2. Components of a URSP Rule
Each URSP rule consists of three main parts:
A) Precedence Value
This is a priority number that determines the order of URSP rules.
A rule with a higher precedence is applied before rules with lower precedence.
B) Traffic Descriptor
This defines which type of traffic the rule applies to. It includes:
Match-All Traffic Descriptor


Means the rule applies to all traffic.
Specific Traffic Conditions (At least one of these must be included):


Application Identifiers → Identifies specific apps (e.g., YouTube, WhatsApp).
IP 3-Tuple → Defines destination IP address, port number, and protocol (e.g., TCP, UDP).
Non-IP Descriptors → Defines non-IP traffic (e.g., sensor data, industrial traffic).
DNNs (Data Network Names) → Identifies different networks (e.g., internet, private networks).
Connection Capabilities → Defines network features needed by apps (e.g., low latency).
Domain Descriptors → Matches traffic based on domain names (FQDNs or regex patterns).
C) Route Selection Descriptor
Defines how traffic should be handled and includes:
PDU Session Type (IP, Ethernet, or Unstructured) and optional parameters:


SSC Mode → Defines session continuity mode (how sessions persist during movement).
S-NSSAI (Network Slice Selection Assistance Information) → Identifies network slices (services).
DNNs → Specifies which data network to use.
Preferred Access Type → Chooses between 5G, Wi-Fi, or other networks.
Multi-Access Preference → Determines priority for multi-network use.
Time Window → Activates rules at specific times.
Location Criteria → Activates rules in specific locations.
PDU Session Pair ID → Groups related sessions together.
RSN (Route Selection Number) → Used for traffic routing decisions.
Non-Seamless Non-3GPP Offload


Offloads traffic to Wi-Fi or other networks without seamless handover.
5G ProSe Layer-3 UE-to-Network Relay Offload


Allows direct device-to-device communication to reduce network load.
3. Default vs. Non-Default URSP Rules
Only one rule in the URSP can be the default rule.
The default rule must match all traffic.
Non-default rules must have a higher precedence than the default rule.
4. Application Identifiers & Connection Capabilities
If an app requests specific connection capabilities, the UE (User Equipment) decides which capability is needed.
Example: A video streaming app may request high bandwidth and low latency.
5. Important Notes
OS Dependency → Different operating systems (Android, iOS) handle connection capabilities differently.
Multiple OS Handling → If a device runs multiple OSs, the URSP can list multiple OS IDs.
Traffic Descriptor Complexity → It is recommended to avoid combining more than two components in a traffic descriptor for simplicity.
6. Why is URSP Important in 5G?
Helps direct traffic efficiently in multi-network environments (e.g., 5G, Wi-Fi, private networks).
Improves QoS (Quality of Service) by prioritizing traffic based on apps, network slices, and connection needs.
Supports network slicing, allowing different services to use dedicated network resources.



The UE Route Selection Policy is used by the UE to determine how to route outgoing traffic.
The UE Route Selection Policy shall consist of one or several URSP rules.
URSP rules are encoded as defined in 3GPP TS 24.526 [16].
UE Route Selection Policy may only be provided by a H-PCF, but shall not be provided by a V-PCF.
The (H-)PCF shall use the UE subscription stored in UDR as specified in 3GPP TS 29.519 [17] to ensure the values included in the Route Selection Descriptor of the generated URSP rules are always supported by subscription.
The (H-)PCF may obtain the information about the UE's OS from the UE as described in the Annex D of 3GPP TS 24.501 [15] or it may derive the information about the UE's OS from the PEI provided by the AMF.
If the (H-)PCF is required to provide UE policies to the UE that includes application descriptors then:
a) If the (H-)PCF has been provided with one UE's OS Id by the UE, the (H-)PCF shall use either the traffic descriptor "OS App Id type" or the traffic descriptor "OS Id + OS App Id type" as defined in 3GPP TS 24.526 [16].
NOTE 1: The (H-)PCF uses the traffic descriptor "OS Id + OS App Id type" when the (H-)PCF does not take the received UE's OS Id into account.
b) If the (H-)PCF has been provided with more than one UE's OS Id by the UE,
- the (H-)PCF shall use the traffic descriptor "OS Id + OS App Id type" for the UE's OS Id provided by the UE as defined in 3GPP TS 24.526 [16]; and
- the (H-)PCF shall not use the traffic descriptor "OS App Id type" as defined in 3GPP TS 24.526 [16].
c) If the (H-)PCF has not been provided with the UE's OS Id by the UE,
- the (H-)PCF shall use the traffic descriptor "OS Id + OS App Id type" as defined in 3GPP TS 24.526 [16];
and
- the (H-)PCF shall not use the traffic descriptor "OS App Id type" as defined in 3GPP TS 24.526 [16].
d) If the (H-)PCF has been provided with the UE's OS Id by the UE and the (H-)PCF has derived the UE's OS Id from the PEI and if there is an inconsistency between the OS Id provided by the UE and the OS Id derived from the PEI, the (H-)PCF shall use the OS Id provided by the UE for providing UE policies to the UE that include application descriptors.
URSP rules may be used to support end to end redundant user plane paths by establishing two redundant PDU sessions.
NOTE 2: The PCF can provide two distinct URSP rules to support end to end redundant user plane paths using Dual Connectivity for the duplicated traffic of an application. 

Duplicated traffic from the UE application is differentiated by two distinct traffic descriptors (different DNNs, and for IP traffic, different IP descriptors or non-IP descriptors), each one defined in a different URSP rule, so that the two redundant PDU sessions are matched to the specific Route Selection Descriptors of distinct URSP rules.
