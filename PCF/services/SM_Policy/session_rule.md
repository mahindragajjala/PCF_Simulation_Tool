https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.b9xbavpq92at

Session rule
Session rules definition

A session rule consists of policy information elements associated with PDU session. The encoding of the SessionRule data type is defined in subclause 5.6.2.7.
A session rule may include:
- Session Rule ID;
- Authorized Session AMBR;
- Authorized Default QoS;
- Reference to Usage Monitoring Data;
- Reference to Usage Monitoring Data for Non-3GPP access of MA PDU session; and
- Reference to Condition Data. 


In the context of 5G networks, a session rule is a key part of the policy control and charging (PCC) framework. 
These rules govern how data sessions (PDU sessions) are handled by the network, including quality of service (QoS), bandwidth restrictions, and monitoring. 
Each session rule is associated with a PDU session, which is a data connection for user traffic, and it dictates how the session behaves according to network policies.
Let's break down each of the components mentioned in the session rule:
1. Session Rule ID
This is a unique identifier for the session rule. 
It allows the network to refer to the specific rule that applies to a given PDU session.
Example: A session rule with ID rule1234 applies a specific QoS and monitoring policy for a user session.
2. Authorized Session AMBR (Aggregate Maximum Bit Rate)
The AMBR specifies the maximum bit rate that can be used by the user in aggregate across all bearers of the PDU session. This is a cap on the total data rate for the session.
Example: If the AMBR is set to 100 Mbps, the user will not be allowed to exceed this total bit rate across all bearers of the session.
3. Authorized Default QoS (Quality of Service)
The QoS parameters define the characteristics of the data session, including priority, delay tolerance, and packet loss. The Authorized Default QoS refers to the default QoS settings that apply to the session.
Example: A QoS setting might guarantee low latency for voice calls but allow higher latency for background data.
4. Reference to Usage Monitoring Data
This refers to the data that the network uses to track the usage of the PDU session. 
It might be used for billing purposes or to enforce policies like data caps or throttling.
Example: A rule could reference usage data to monitor a user’s data consumption and trigger an alert when the user exceeds a predefined data limit.
5. Reference to Usage Monitoring Data for Non-3GPP Access of MA PDU Session
This refers to usage monitoring for sessions that are established via Non-3GPP (e.g., Wi-Fi) access. Non-3GPP access refers to data connections that are not through the cellular network but still need to be monitored for session control.
Example: A session initiated over a Wi-Fi network (Non-3GPP access) could still be monitored for data usage to ensure that the user's data plan or QoS requirements are respected.
6. Reference to Condition Data
This refers to conditions that might modify the behavior of the session rule based on factors like time of day, location, or network load.
Example: A condition could be applied where the QoS is downgraded during peak network hours, or the session might be paused or restricted during certain conditions (like when the user is roaming).
Example Scenario:
Let's say a 5G user subscribes to a video streaming service. Their session rule might look like this:
Session Rule ID: rule1234
Authorized Session AMBR: 50 Mbps (total aggregate bandwidth across all bearers)
Authorized Default QoS: High priority, low latency (ensures smooth video streaming)
Reference to Usage Monitoring Data: Monitor data usage for billing and data cap enforcement.
Reference to Usage Monitoring Data for Non-3GPP Access: If the user is connected via Wi-Fi, monitor usage for data capping
Reference to Condition Data: If the user is in a high-traffic area, reduce the bandwidth to 10 Mbps during peak hours.
In this example, the session rule ensures that the user’s video stream gets high priority and low latency but is also limited by a total bandwidth cap. Usage is monitored, and data caps are applied even when the user is on Wi-Fi.

Policy Decision types:
A policy decision is a grouping of cohesive information elements describing a specific type of decision, e.g. QoS, Charging data, etc. 
A policy decision can be linked to one or more PCC rules or one or more Session rules. 

A PCC rule or session rule can at most refer to one instance of the policy decision for each type.
The following types of policy decision are defined:
- Traffic control data;
- QoS data;
- Charging data;
- Usage Monitoring data; and
- QoS Monitoring data.
In 5G, a Policy Decision refers to a set of instructions or rules that are made to control or manage certain aspects of the network's operation. 
These decisions are based on different factors like Quality of Service (QoS), charging, traffic control, and monitoring, all of which help to ensure that network resources are used efficiently and in a way that aligns with the operator's business goals or the network’s technical capabilities.
Key Concepts:
Policy Decision: This is a grouping of relevant information that describes the specific decisions made about network management. These decisions could involve different aspects of the network like traffic control, charging data, or usage monitoring. These decisions guide the operation of the network according to pre-established rules.


Policy Control and Charging (PCC) Rules: In the context of 5G, PCC rules are used to enforce policy decisions. These rules dictate how the network should behave in specific conditions, such as allocating bandwidth, managing traffic, or charging users. A PCC rule can link to a specific Policy Decision, ensuring that the correct decision is applied under the right circumstances.


Session Rules: These are specific rules related to the user’s active session in the network (such as a mobile data session). A session rule can refer to a Policy Decision that affects the current session, like managing QoS or monitoring usage.


Types of Policy Decisions:
The types of policy decisions are based on the particular area of network management:
Traffic Control Data: Refers to decisions that manage how network traffic is handled. This could involve controlling the flow of data based on factors like congestion, priority, or the type of data being transmitted.


QoS (Quality of Service) Data: Defines how network resources should be allocated to ensure a certain level of service quality for users. It may include priorities for different types of data (e.g., voice vs. video streaming) and guarantee certain metrics like latency, bandwidth, or reliability.


Charging Data: Involves decisions related to the charging mechanisms for network usage. This could include how users are billed for data consumption, the pricing model for different services, or determining if additional charges should be applied based on usage patterns.


Usage Monitoring Data: Involves the collection of data about how much network resource a user has consumed. This data is important for charging decisions, but also for managing network resources and detecting abnormal usage patterns.


QoS Monitoring Data: Refers to decisions that involve monitoring the ongoing quality of service on the network. It helps in adjusting resources dynamically to maintain the desired level of service.


How Policy Decisions are Linked to Rules:
PCC Rules or Session Rules can refer to only one instance of a specific Policy Decision per type. This means that each rule can only apply one specific decision for a given category (e.g., charging, QoS) at any time.
A Policy Decision can be linked to multiple PCC or session rules, but these rules must apply the decision in a cohesive and consistent manner.
Example in Practice:
A Traffic Control Policy Decision could dictate that high-priority traffic (e.g., emergency services) gets allocated more bandwidth during peak times. The PCC Rule for this scenario would link to the traffic control decision, ensuring that this policy is applied when the network detects high traffic or congestion.
In summary, Policy Decisions in 5G serve as a centralized mechanism for guiding how network resources are allocated and managed. These decisions are applied through PCC and session rules to ensure the network operates efficiently while meeting the required service levels for users.
Example of the above content 
Let's break down a real-time example using the concepts of Policy Decisions in 5G, focusing on a scenario where multiple network policies are applied to manage traffic, Quality of Service (QoS), charging, and usage monitoring. 
We'll consider a mobile operator managing a 5G network for a smartphone user.
Scenario: Managing Network Traffic During a Concert
Imagine there’s a 5G network operator providing coverage for a large outdoor concert. 
The concert has tens of thousands of attendees, all using their smartphones to access the internet, stream videos, and share content. 
The operator needs to ensure that the network handles the increased traffic smoothly, prioritizes important services, and applies the appropriate billing.
1. Traffic Control Data: Ensuring Network Efficiency
During the concert, the network experiences a large influx of traffic. To avoid congestion and ensure efficient data flow, the network operator needs to manage the traffic efficiently. Here’s where Traffic Control Policy Decisions come into play.
Policy Decision: The operator can create a policy decision that prioritizes high-priority services like emergency communications, voice calls, and video streams for the concert’s live broadcast. All other traffic like social media posts, downloads, and casual web browsing may be deprioritized.
PCC Rule: The operator implements a PCC rule that applies this traffic control decision. If the network experiences congestion, the rule ensures that priority services (e.g., video streams, voice calls) are given higher bandwidth, while non-essential traffic is throttled.
Real-World Example: During the concert, a user trying to stream a live video of the event gets priority bandwidth, while a user trying to browse social media sees slower speeds due to the congestion.
2. Quality of Service (QoS) Data: Maintaining Service Quality
To maintain a high-quality experience for users, especially for video streaming or real-time communications, QoS Data must be managed.
Policy Decision: The operator may create a QoS policy to ensure that the video streaming quality is maintained even in congested network conditions. The policy can guarantee a minimum bandwidth of 10 Mbps for video streaming.
PCC Rule: A PCC rule links this QoS policy to the network traffic. If a user is streaming video, the rule ensures the necessary resources (bandwidth, latency) are available, so the video quality is high (no buffering or lag).
Real-World Example: A user watching the concert live stream experiences smooth video playback without buffering, even though other users are experiencing slower speeds for general browsing.
3. Charging Data: Applying Billing Rules
As the network is heavily used, the operator may want to implement charging policies based on data usage or specific services.
Policy Decision: The operator can set a charging policy where data consumption beyond a certain threshold (e.g., 10 GB) during the concert is subject to additional charges. Alternatively, they could offer a special event-based pricing plan, charging users a premium for accessing the live stream in HD quality.
PCC Rule: A PCC rule links the charging policy to the user’s data usage during the event. If the user exceeds the data limit or accesses high-bandwidth services, the rule triggers additional charges.
Real-World Example: A user who streams the concert in high definition (HD) for several hours may get charged extra due to the high data consumption, as per the operator’s charging policy.
4. Usage Monitoring Data: Tracking Data Consumption
To ensure that users are billed correctly and the network is not being overloaded, the operator needs to track data usage throughout the concert.
Policy Decision: The operator can set a usage monitoring policy that tracks how much data each user consumes during the event.
Session Rule: A session rule could be applied to monitor the user’s real-time data usage and update the operator's billing system accordingly. This can include tracking usage for specific services (e.g., streaming vs. browsing) or monitoring overall consumption.
Real-World Example: A user who consumes 2 GB of data during the concert gets their data usage recorded and monitored for accurate billing at the end of the session.
5. QoS Monitoring Data: Ensuring Consistent Service Quality
To make sure that the QoS metrics are met and adjust resources dynamically, the operator must monitor QoS data during the concert.
Policy Decision: The operator may create a QoS monitoring policy to ensure that the network is consistently delivering the expected performance for high-priority services. If the network’s QoS drops below a threshold (e.g., latency becomes too high), the policy could trigger actions like allocating more bandwidth or rerouting traffic.
Session Rule: A session rule can be applied to monitor real-time QoS metrics (e.g., latency, jitter, throughput) and adjust network resources as necessary.
Real-World Example: During the concert, the operator monitors the latency and packet loss for users streaming the event. If a user's experience begins to degrade, the operator may adjust the network resources to improve their connection quality.

Summary of Real-Time Example:
During a large 5G network deployment at a concert, the operator uses different Policy Decisions to control the network’s performance:
Traffic Control to prioritize high-bandwidth, important traffic (live streaming and emergency services).
QoS to guarantee video quality for users streaming the concert.
Charging to apply extra fees for high data consumption or premium services.
Usage Monitoring to track data consumption for accurate billing.
QoS Monitoring to ensure service quality remains optimal for streaming and real-time communications.
In each case, the operator uses PCC Rules and Session Rules to ensure these policies are applied and enforced across the network in real time, providing a smooth user experience while managing resources efficiently.
Traffic control data definition

In computer networking, traffic control data is a method for managing, controlling, or reducing network traffic, especially internet bandwidth. Network administrators use traffic control to reduce congestion, packet loss, and latency.

Traffic control data defines how traffic data flows associated with a rule are treated (e.g. blocked, redirected). 
The traffic control data encoding table is defined in subclause 5.6.2.10.
Traffic control data shall include:
- Traffic Control Data ID.
Traffic control data may include:
- Flow status;
- Redirect Information;
- Mute Notification;
- Traffic Steering Policy ID UL;
- Traffic Steering Policy ID DL;
- Routing requirements;
- UP path change event subscription from the AF;
- Indication of traffic correlation;
- Access Traffic Steering Functionality;
- Access Traffic Steering Mode DL;
- Access Traffic Steering Mode UL; and
- Multicast Access Control. 


The Traffic Control Data is a set of parameters and attributes that determine how specific traffic flows are treated within a network based on policies or rules. Below is a detailed explanation of the components mentioned in your query:

Mandatory Element
Traffic Control Data ID
This is a unique identifier for a particular traffic control policy or data set. It is used to associate specific traffic flows with the corresponding rules or actions.

Optional Elements
The following elements may be included in the Traffic Control Data based on the use case:
1. Flow Status
Describes the current status of the traffic flow, such as active, paused, or terminated. This status can help the network make decisions about resource allocation or session handling.
2. Redirect Information
Provides details on where specific traffic should be redirected. This is useful for scenarios like redirecting to a content delivery network (CDN), proxy servers, or another server for load balancing or optimization.
3. Mute Notification
A parameter that indicates whether notifications or alerts about this traffic flow should be suppressed or muted for the user or network functions.
4. Traffic Steering Policy ID UL (Uplink)
Refers to the policy or rule governing how uplink traffic (data sent from the user device to the network) should be handled, such as selecting specific paths or prioritizing certain traffic.
5. Traffic Steering Policy ID DL (Downlink)
Refers to the policy or rule for downlink traffic (data sent from the network to the user device), ensuring that traffic is routed or prioritized as required.
6. Routing Requirements
Specifies the constraints or preferences for routing the traffic, such as avoiding certain nodes, preferring low-latency paths, or using secure tunnels.
7. UP Path Change Event Subscription from the AF
The Application Function (AF) can subscribe to events indicating changes in the User Plane (UP) path, such as when traffic is rerouted due to network conditions.
8. Information on User Plane Latency Requirements
Indicates latency requirements for the traffic flow. For instance, applications like online gaming or video conferencing might specify low latency requirements to ensure smooth performance.
9. EAS IP Replacement Information
Provides instructions for replacing the IP address of the Edge Application Server (EAS) during certain scenarios, such as edge relocation or network optimization.
10. Indication of Traffic Correlation
Signals that multiple traffic flows are correlated and should be treated together. For example, audio and video streams in a call may need synchronized handling.
11. Indication of Simultaneous Connectivity Temporarily Maintained
Describes situations where simultaneous connectivity is maintained to both source and target PSA (Packet Session Anchors) during edge relocation. It also provides guidance on when connectivity to the source PSA can be safely removed.
12. Access Traffic Steering Functionality
Indicates the capability to steer traffic between different access networks, such as between Wi-Fi and 5G, to optimize performance or balance loads.
13. Access Traffic Steering Mode DL (Downlink)
Specifies the mode for steering downlink traffic. It may include rules or thresholds for determining which access path is used.
14. Access Traffic Steering Mode UL (Uplink)
Similar to downlink, this specifies the mode for steering uplink traffic. It ensures that user devices send data using the most suitable access network.
15. Optionally, Access Traffic Steering Mode Indicator or Threshold
These indicators or thresholds can define when to switch access traffic steering modes. For example, a network may switch to a different mode when a certain bandwidth threshold is reached.
16. Multicast Access Control
Governs how multicast traffic (data sent to multiple recipients simultaneously) is controlled, such as allowing or blocking certain multicast streams based on policies.

Use Cases of Traffic Control Data
Quality of Service (QoS): Ensures high-priority traffic, such as emergency calls or video conferencing, receives preferential treatment.
Load Balancing: Redirects or steers traffic to less congested routes or servers.
Edge Computing: Manages traffic in scenarios where data is processed closer to the user at the network edge.
Mobility Management: Facilitates seamless connectivity during user movement across network boundaries.
Multicast Optimization: Ensures efficient delivery of live events or updates to multiple users.
Each of these attributes is part of a flexible and customizable framework that allows operators to implement diverse traffic management strategies tailored to network needs and user demands.

Policy control request trigger
Policy control request trigger is a condition when the SMF shall interact again with PCF for further policy decision of a PDU session.
The policy control request trigger is designed as an Enumeration type defined in the subclause 5.6.3.6.
The PCF can provide an array of policy control request triggers in policy decision to subscribe the triggers in SMF.
When SMF interacts with PCF due to the triggering of the policy control request triggers, the SMF shall send the related attributes that have changed together with the corresponding triggers.
The SMF (Session Management Function) sends the Policy Control Request to the PCF (Policy Control Function) when a Policy Control Request Trigger is activated. This trigger occurs when there is a significant change in the PDU session (e.g., data usage, QoS requirements). The request is sent whenever the conditions defined in the PCF's policy triggers are met, and the SMF needs updated policy decisions from the PCF based on those changes.
Requested rule data
Requested rule data consists of requested information by the PCF associated with one or more PCC rules.
The requested rule data is designed as a subresource of the policy decision within an attribute called "lastReqRuleData".
The PCF only records the last requested rule data.
When requesting rule data, the PCF shall include the types of data requested for the rules within the "reqData" array of
the "lastReqRuleData" and shall also provide the corresponding policy control request triggers if the triggers are not yet
set.
The encoding of the requested rule data is further specified in subclause 5.6.2.24.
When the SMF receives the requested rule data, the SMF shall report the corresponding information to the PCF for the
associated PCC rule(s). 

					Requested rule data
Requested rule data consists of requested information by the PCF associated with one or more PCC rules.
The requested rule data is designed as a subresource of the policy decision within an attribute called "lastReqRuleData".
The PCF only records the last requested rule data.
When requesting rule data, the PCF shall include the types of data requested for the rules within the "reqData" array of the "lastReqRuleData" and shall also provide the corresponding policy control request triggers if the triggers are not yet set.
The encoding of the requested rule data is further specified in subclause 5.6.2.24.
When the SMF receives the requested rule data, the SMF shall report the corresponding information to the PCF for the associated PCC rule(s). 
The content you provided relates to a specific part of the Policy Control Function (PCF) and the communication between the PCF and the Session Management Function (SMF) in a 5G network. Here's an explanation of the concepts mentioned and when this is used in a call flow:
Explanation of the Content
Requested Rule Data:


This is the information that the PCF requests related to one or more Policy and Charging Control (PCC) rules. PCC rules define how the network handles data flow, including traffic control and charging.
Subresource of the Policy Decision:


The requested rule data is part of a broader policy decision. The PCF sends a policy decision to different network functions, and the requested rule data is recorded under the attribute called lastReqRuleData.
Recording of Rule Data:


The PCF only keeps track of the last requested rule data. This means that, at any given time, it only records the most recent request for rule data.
"reqData" Array in "lastReqRuleData":


When the PCF makes a request, it includes the types of data it needs in the reqData array. This array specifies the exact data needed for the associated PCC rules.
Additionally, if there are any policy control request triggers that are not yet set, the PCF will also provide these triggers. Triggers are used to initiate certain actions based on network conditions or events.
Encoding of Rule Data (Subclause 5.6.2.24):


The content of the requested rule data follows a specific encoding format, which is defined in subclause 5.6.2.24. This ensures that the data is structured correctly for processing by the network functions.
SMF's Role:


After receiving the requested rule data, the Session Management Function (SMF) processes it and reports the corresponding information back to the PCF for the associated PCC rules.
This means that the SMF acts as the intermediary that reports the requested data from the network.
Call Flow of When This is Used
Here’s when this process is typically used in a call flow, focusing on the interactions between the PCF and SMF:
Step 1 - Policy Decision:


The PCF makes a policy decision for a user or service flow. This policy decision may involve several PCC rules.
Step 2 - Rule Data Request:


The PCF needs specific data related to the PCC rules (e.g., traffic management data, QoS parameters, etc.). It sends a request for rule data to the SMF.
This request includes the reqData array, specifying the types of data needed for the PCC rules.
Step 3 - Handling of Triggers:


If certain policy control triggers are not yet set, the PCF includes those in the request as well. These triggers help ensure that specific actions are taken once certain network conditions are met.
Step 4 - Encoding of Data:


The requested rule data is encoded as per the standard defined (subclause 5.6.2.24). This ensures that the data is in the proper format for SMF to process.
Step 5 - SMF Processes the Request:


The SMF receives the request for rule data. It processes the request and gathers the necessary information related to the PCC rules.
Step 6 - SMF Reports to PCF:


Once the SMF has collected the requested data, it sends the information back to the PCF.
The PCF now has the data it requested for the associated PCC rules and can make further decisions accordingly.
Step 7 - Final Policy Decision Update:


Based on the received rule data and triggers, the PCF can update its policy decisions or apply further control over the user or service flow.
Use Cases:
This process is commonly used when the PCF needs to gather specific information to make an informed policy decision for a user’s traffic or charging behavior, especially when dynamic data such as QoS or usage data is required.
Example in Real-World Call Flow:
In a situation where a user is consuming high-bandwidth content (e.g., streaming video), the PCF may need real-time traffic data (like bandwidth usage) to ensure that the correct PCC rules for QoS are applied. It would request this data from the SMF, which would return it, and the PCF could then adjust the QoS rules accordingly to optimize the user’s experience or manage network resources efficiently.


