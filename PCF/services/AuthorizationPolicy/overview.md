https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.s672p4lh9omc

 TS 29.514 


AF-NEF-PCF Discovery Flow
AF → NEF: Request for PCF (e.g., for VoLTE policy).
NEF → NRF: Request for available PCF instances that handle VoLTE.
NRF → NEF: Provides list of available PCFs.
NEF → AF: Sends PCF details.
AF → PCF: Sends the Policy Authorization Request to the chosen PCF.


The Npcf_PolicyAuthorization service authorises an AF request and creates policies as requested by the authorised NF service consumer for the PDU session to which the AF session is bound to. 
This service allows the NF service consumer to subscribe/unsubscribe to the notification of events (e.g. access type change, PLMN change, usage report, access network information report). 
In the context of 5G architecture, the Application Function (AF) is a logical entity that provides application services and interacts with the 5G core network (5GC). The AF enables communication between the applications running on user devices or external servers and the 5G core to ensure proper network behavior and Quality of Service (QoS) for specific applications.
Key Roles of the Application Function:
Policy and QoS Control:


The AF interacts with the Policy Control Function (PCF) in the 5GC to request and manage network policies for specific applications or services.
For example, a video streaming service can request high-priority traffic treatment for a video session.
Support for Network Slicing:


The AF can request a specific network slice (a virtualized segment of the network) to meet application requirements, such as ultra-low latency or high bandwidth.
Exposure of Network Capabilities:


The AF can expose 5G core capabilities (like location, QoS, and analytics) to applications.
Session and Data Flow Management:


It ensures that applications receive the appropriate data flow routing and prioritization.

Real-Time Examples of Application Functions in 5G Applications
Smart City Surveillance:


A smart city application uses high-definition cameras for surveillance.
The AF requests low latency and high bandwidth from the 5G core for seamless real-time video streaming to law enforcement agencies.
Autonomous Vehicles:


In vehicle-to-everything (V2X) communication, an autonomous car's application requires ultra-reliable low-latency communication (URLLC).
The AF ensures a dedicated slice of the network is allocated for real-time communication between vehicles, infrastructure, and pedestrians.
Augmented/Virtual Reality (AR/VR):


For AR/VR gaming or industrial training applications, the AF requests high bandwidth and ultra-low latency to provide immersive experiences without lags or interruptions.
Healthcare: Remote Surgery:


In remote surgery, the AF ensures extremely low latency and high reliability for transmitting commands and video feeds between the surgeon and the robotic surgery system.
IoT Applications in Agriculture:


Smart irrigation systems equipped with IoT sensors use AF to ensure priority data transfer from sensors to the control center for timely irrigation based on weather or soil conditions.
Live Sports Streaming:


A video streaming platform like Netflix or Disney+ uses the AF to optimize bandwidth and QoS for delivering high-definition video to end-users during live sports events.
Gaming Platforms (Cloud Gaming):


Cloud gaming platforms like NVIDIA GeForce Now or Microsoft Xbox Cloud Gaming rely on the AF to request network resources that provide low latency and high bandwidth for a seamless gaming experience.

Call Flow Example: Video Streaming (AF Interaction with PCF)
User Requests Video Stream:


The user launches a video streaming app (e.g., Netflix) on a mobile device.
Application Requests QoS:


The AF of the video streaming service sends a request to the PCF via the Network Exposure Function (NEF) to allocate bandwidth and prioritize traffic for the streaming session.
Policy Creation:


The PCF creates the required policies for the streaming session and informs the User Plane Function (UPF) to enforce them.
Traffic Flow Setup:


The 5G core ensures high-priority routing of the video data packets between the streaming server and the user's device.
Stream Playback:


The user enjoys a buffer-free streaming experience with optimal quality.
This shows how the AF is integral to enabling advanced 5G services by ensuring seamless integration between applications and network resources.



The Npcf_PolicyAuthorization_Create service operation authorizes the request from the NF service consumer, and
optionally communicates with Npcf_SMPolicyControl service to determine and install the policy according to the
information provided by the NF service consumer.
The Npcf_PolicyAuthorization_Create service operation creates an application session context in the PCF.
The following procedures using the Npcf_PolicyAuthorization_Create service operation are supported:
- Initial provisioning of service information.
- Gate control.
- Initial Background Data Transfer policy indication.
- Initial provisioning of sponsored connectivity information.
- Subscription to Service Data Flow QoS notification control.
- Subscription to Service Data Flow Deactivation.
- Initial provisioning of traffic routing information.
- Subscription to resources allocation outcome.
- Invocation of Multimedia Priority Services.
- Support of content versioning.
- Request of access network information.
- Initial provisioning of service information status.
- Provisioning of signalling flow information.
- Support of resource sharing.
- Indication of Emergency traffic.
- Invocation of MCPTT. 
- Invocation of MCVideo.
- Priority sharing indication.
- Subscription to out of credit notification.
- Subscription to Service Data Flow QoS Monitoring information.
- Provisioning of TSCAI input information and TSC QoS related data.
- Provisioning of bridge management information and port management information.
- P-CSCF restoration enhancements.
- Support of CHEM feature.
- Support of FLUS feature.
- Subscription to EPS Fallback report.
- Subscription to TSN related events.
- Initial provisioning of required QoS information.
- Support of QoSHint feature.
- Subscription to reallocation of credit notification. 



The Npcf_PolicyAuthorization_Create service operation is a key function within the Policy Control Function (PCF) in 5G networks, as specified by the 3GPP standards. 

It is responsible for managing and authorizing application session contexts and policies for network functions (NFs) based on their requests.




1. Initial provisioning of service information
Sets up and manages application-specific service policies for an initial application session (e.g., video streaming or data transfer).
Ensures that required service parameters are communicated from the NF to the PCF for accurate policy enforcement.
2. Gate control
Ensures that data traffic is allowed or denied based on rules for specific service flows.
Implements policies such as traffic blocking, bandwidth control, or priority handling.
3. Initial Background Data Transfer (BDT) policy indication
Applies policies for low-priority, delay-tolerant traffic such as software updates or bulk data transfers.
4. Initial provisioning of sponsored connectivity information
Supports sponsored services where a third party (e.g., content provider) pays for the data usage instead of the subscriber.
Manages policies for billing and service differentiation.
5. Subscription to Service Data Flow QoS notification control
Allows NF service consumers to subscribe to notifications for changes in QoS (Quality of Service) parameters for service data flows.
6. Subscription to Service Data Flow Deactivation
Enables notifications when specific service data flows are deactivated or removed.
7. Initial provisioning of traffic routing information
Provides rules to route traffic through specific paths (e.g., low-latency routes) based on service requirements.
8. Subscription to resources allocation outcome
Notifies the NF about the success or failure of resource allocation for service flows.
9. Invocation of Multimedia Priority Services (MPS)
Manages prioritized access for critical multimedia services during emergencies (e.g., video calls for first responders).
10. Support of content versioning
Enables policies that manage different versions of content (e.g., video streams with different resolutions).
11. Request of access network information
Authorizes requests for information about the access network (e.g., signal strength or location).
12. Initial provisioning of service information status
Tracks the status of service-specific policies during application session setup.
13. Provisioning of signaling flow information
Sets up policies for signaling traffic separate from user traffic (e.g., control plane data for call setups).
14. Support of resource sharing
Manages the sharing of network resources across multiple services or applications.
15. Indication of Emergency traffic
Identifies and applies policies for emergency traffic (e.g., 911 calls).
16. Invocation of MCPTT (Mission Critical Push-to-Talk)
Manages policies for Push-to-Talk services used by first responders in mission-critical scenarios.
17. Invocation of MCVideo
Applies policies for Mission Critical Video services, ensuring high-priority access for emergency responders.
18. Priority sharing indication
Manages priority levels across different services or applications for fair resource distribution.
19. Subscription to out-of-credit notification
Notifies the NF if a subscriber runs out of data or credit balance during an application session.
20. Subscription to Service Data Flow QoS Monitoring information
Tracks and reports QoS metrics (e.g., latency, throughput) for specific service data flows.
21. Provisioning of TSCAI input information and TSC QoS related data
Supports policies for Time Sensitive Communications (TSC) applications, ensuring strict latency and synchronization requirements.
22. Provisioning of bridge management information and port management information
Applies policies for managing bridges and ports in network slicing or other specialized use cases.
23. P-CSCF restoration enhancements
Ensures that policies are applied for restoring the Proxy Call Session Control Function (P-CSCF) in IMS-based systems.
24. Support of CHEM feature
Refers to Charging Enhancements for Multimedia, managing charging-related policy for multimedia services.
25. Support of FLUS feature
Refers to Flexible User Subscription, enabling dynamic management of subscription parameters.
26. Subscription to EPS Fallback report
Notifies the NF if a device falls back to EPS (Evolved Packet System) from 5G due to network conditions.
27. Subscription to TSN-related events
Applies policies for Time-Sensitive Networking (TSN) events, ensuring deterministic communication for industrial IoT.
28. Initial provisioning of required QoS information
Sets initial QoS requirements for an application session (e.g., minimum bandwidth, latency, or packet loss).
29. Support of QoSHint feature
Provides hints about QoS requirements, allowing flexible adaptation to network conditions.
30. Subscription to reallocation of credit notification
Notifies the NF when credits (e.g., for data usage or service consumption) are reallocated or adjusted.


