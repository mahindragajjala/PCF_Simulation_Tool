https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.k6frxydugash

The overall Access and Mobility (AM) policies for a user in a 5G network, governed by the PCF (Policy Control Function), are high-level policies that dictate how a user accesses the network and manages mobility. These policies are designed to ensure seamless connectivity, optimal resource utilization, and compliance with subscription terms.
Key Aspects of Overall AM Policies

1. Access Policies
These determine how and where a user can access the network and under what conditions.
Authorized Access Types:


Specifies whether the user can connect via:
5G NR (New Radio, i.e., cellular access)
LTE
Non-3GPP networks (e.g., Wi-Fi).
Example: A subscription might allow cellular access but restrict Wi-Fi offloading.
Preferred RAT (Radio Access Technology):


Determines which RAT the user should prioritize:
5G preferred over LTE when available.
Wi-Fi preferred for data offloading.
PLMN (Public Land Mobile Network) Selection:


Policies for selecting a specific operator's network (e.g., home network or a roaming partner's network).
Example: "Only connect to the home PLMN unless roaming is allowed."
Network Slice Access:


Determines which network slices the user can access, based on their subscription or service agreement.
Example: Access to eMBB slice for high-speed data but no access to URLLC slice.

2. Mobility Policies
These govern how the user moves between cells, networks, or access types without service interruption.
Handover Policies:


Rules for switching between cells, RATs, or networks during mobility (e.g., seamless handover between 5G NR and LTE).
Example: "Allow inter-RAT handover only if 5G signal strength drops below a certain threshold."
Roaming Policies:


Rules for mobility across geographic regions or operator networks.
Example:
Roaming allowed in specific countries or PLMNs.
Roaming restricted for certain services or slices to control costs.
Session Continuity:


Policies to ensure uninterrupted service for ongoing sessions during mobility.
Example: For a video call, enable IP address preservation during handover.

3. QoS (Quality of Service) Policies
These ensure that users receive the appropriate quality of service based on their subscription and current network conditions.
Priority Levels:


Assigns priority to users or services during congestion.
Example: Emergency services get higher priority than regular users.
QoS Parameters:


Bandwidth allocation, latency, jitter, and packet loss thresholds.
Example: Provide a guaranteed minimum bandwidth of 10 Mbps for streaming services.

4. Subscription-Based Policies
These are tied to the user's subscription plan and profile.
Service Authorization:


Defines which services (e.g., voice, data, messaging) the user can access.
Example: Data-only plan prohibits voice services.
Usage Limits:


Policies for controlling data caps or service restrictions.
Example: "Restrict bandwidth to 1 Mbps after 10 GB of data usage."

5. Security and Authentication Policies
These enforce secure and authenticated access to the network.
Authentication Rules:


Policies for verifying the user’s identity before granting access.
Example: Require SIM-based authentication for 3GPP access and username/password for Wi-Fi.
Device Access Rules:


Policies for allowing or denying access based on the user’s device type or status.
Example: "Block access for devices flagged as stolen or compromised."

6. Location-Based Policies
These depend on the user’s physical location.
Geographic Restrictions:


Limits access to services or networks in specific areas.
Example: "Prohibit access in certain regions for regulatory compliance."
Regional QoS Adjustments:


Dynamic QoS settings based on location.
Example: In urban areas, reduce bandwidth to manage congestion.

7. Slice-Specific Policies
If the user’s session involves network slicing, the AM policies define which slices can be accessed and how.
Slice Admission Rules:


Defines conditions for accessing specific slices.
Example: "Allow access to the URLLC slice only for industrial IoT devices."
Slice Prioritization:


Assigns priorities to slices in case of resource contention.
Example: Prioritize eMBB traffic over mMTC traffic during congestion.




Here’s a practical example and flow of how Access and Mobility (AM) policies are applied in real-time for a 5G user:

Scenario
A user, Alice, with a 5G subscription, is streaming a video on her smartphone while traveling by train. The network needs to manage her session dynamically as she moves between cells, RATs (5G and LTE), and even a Wi-Fi hotspot at a train station.

Step-by-Step Flow

1. Initial Registration and Policy Setup
Trigger: Alice’s device powers on and attempts to connect to the 5G network.
Action:
The Access and Mobility Management Function (AMF) sends a request to the Policy Control Function (PCF) to retrieve AM policies for Alice.
The PCF determines Alice’s overall AM policies based on her subscription and sends them back to the AMF.
Policy Enforcement:
Alice is authorized to access the network via:
3GPP access (5G NR as primary, LTE as fallback).
Non-3GPP access (Wi-Fi is allowed for offloading).
The AM policy ensures a preferred connection to 5G NR if available.

2. Seamless Handover Between 5G NR and LTE
Trigger: As the train moves, Alice transitions between areas with weak 5G coverage to areas with stronger LTE coverage.
Action:
The AMF monitors signal strength and initiates a handover from 5G NR to LTE when the 5G signal drops below the configured threshold.
The PCF dynamically updates the mobility policies to allow the inter-RAT (Radio Access Technology) handover.
Policy Enforcement:
The session remains active with continuity of the video stream.
The policy prioritizes LTE for better coverage but monitors for opportunities to reconnect to 5G NR when available.

3. Roaming and Slice Selection
Trigger: The train crosses into a new region where Alice’s home operator doesn’t provide 5G coverage. Roaming agreements come into play.
Action:
The AMF detects the change in PLMN (Public Land Mobile Network) and informs the PCF.
The PCF applies Alice’s roaming policy, allowing access to the visited PLMN but restricting high-priority slices (e.g., URLLC).
Policy Enforcement:
Alice is granted access to the eMBB slice for video streaming, but not the URLLC slice, as her subscription does not cover URLLC services in roaming.

4. QoS Adjustment During Congestion
Trigger: As the train enters a busy urban area, the network experiences congestion.
Action:
The AMF detects congestion and requests a QoS adjustment for active users from the PCF.
The PCF applies QoS policies based on priority:
Premium users are given higher priority.
Alice, as a regular user, receives a reduced bandwidth allocation but maintains a minimum QoS for her video stream.
Policy Enforcement:
Bandwidth for Alice’s session is reduced from 20 Mbps to 10 Mbps, ensuring fair resource distribution across users.

5. Offloading to Non-3GPP (Wi-Fi)
Trigger: At a train station, Alice’s device detects an available Wi-Fi network.
Action:
The device informs the AMF about the detected Wi-Fi.
The PCF checks Alice’s access policy and allows Wi-Fi offloading to reduce cellular network load.
Policy Enforcement:
Alice’s video stream switches to Wi-Fi while maintaining session continuity through a Non-Seamless WLAN Offload (NSWO) mechanism.

6. Security and Subscription Check
Trigger: Alice attempts to use a high-bandwidth video calling app.
Action:
The PCF checks Alice’s subscription, which does not allow high-bandwidth video calls while roaming.
Policy Enforcement:
The network denies the request for the video calling service, notifying Alice of the restriction.

Diagram of the Flow
Initial Access: Device → AMF → PCF (Policies applied)
Mobility Events: Handover: 5G NR ↔ LTE → Policy updates.
Roaming: New PLMN → Policy enforcement (slice restrictions).
Congestion: AMF → PCF (QoS adjustment).
Wi-Fi Offloading: Non-3GPP access → PCF (Policy validation).
Service Restrictions: PCF → Deny unauthorized service.

Key Takeaways
Access Policies: Govern RAT preferences, PLMN selection, and network slicing.
Mobility Policies: Enable seamless handovers and roaming control.
QoS Policies: Dynamically adapt based on network conditions.
Subscription Policies: Enforce limits and restrictions based on user plans.

