https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.i2vejhgahksr


PCC rules


PCC rules definition

A PCC rule is a set of information elements enabling the detection of a service data flow and providing parameters for policy control and/or charging control. 

The above content describes a Policy and Charging Control (PCC) rule in the context of telecommunications, particularly in mobile networks (e.g., 4G LTE, 5G). Here's a simplified explanation:
Service Data Flow Detection:


A service data flow is a stream of data packets belonging to a specific application or service, such as a video call, web browsing, or online gaming.
The PCC rule contains information elements (criteria) to identify or detect these data flows. For instance:
IP addresses
Port numbers
Protocol types (e.g., HTTP, FTP)
Example: A rule could specify, "Detect all video streaming traffic on port 443 (HTTPS) for a specific user."


Policy Control:


Once the data flow is detected, the PCC rule can enforce policies that define how the network should handle this traffic.
Policies include:
Quality of Service (QoS): Prioritize video streaming traffic to ensure smooth playback.
Bandwidth allocation: Restrict or allow certain amounts of bandwidth for the flow.
Charging Control:


PCC rules can also specify how to charge the user for the detected service data flow.
For example:
Free access to certain apps or services.
Charge based on the amount of data consumed (e.g., ₹10 per GB).
Differentiated rates for different times (e.g., discounted rates at night).
Use Case Example:
A user is streaming video on an app like YouTube.
The PCC rule detects this flow based on IP, port, and protocol.
It assigns high priority to ensure smooth playback (policy control).
It applies charging rules, e.g., deducting from a "video pack" instead of the regular data balance.
In essence, a PCC rule is a framework for recognizing specific data flows, enforcing network behavior (policy).


There are two different types of PCC rules as defined in 3GPP TS 23.503 :

Dynamic PCC rules. PCC rules that are dynamically provisioned by the PCF to the SMF.

These PCC rules may be either predefined or dynamically generated in the PCF. Dynamic PCC rules can be installed, modified and removed at any time.

Predefined PCC rules. PCC rules that are preconfigured in the SMF. Predefined PCC rules can be activated or
deactivated by the PCF at any time. Predefined PCC rules within the PCF may be grouped allowing the PCF to
dynamically activate a set of PCC rules. 


The content describes the two types of Policy and Charging Control (PCC) rules used in 5G networks, as defined in 3GPP TS 23.503. Here's a clear and simplified explanation:

1. Dynamic PCC Rules:
What are they?


These rules are created and managed dynamically(actively or in real time) by the PCF (Policy Control Function) during the network's operation.
They are sent to the SMF (Session Management Function) as needed.
These rules are flexible, meaning they can be:
Created (installed) when a new service or session starts.
Modified based on changes in service requirements or network conditions.
Removed when no longer needed.
How are they created?
They can be:
Predefined: Based on pre-existing(ముందుగా ఉన్న) templates or rules in the PCF.
Dynamically generated: Fully customized by the PCF based on real-time inputs like user requirements or network status.
When are they used?
For specific, dynamic scenarios such as:
Prioritizing video streaming traffic during a session.
Allocating additional bandwidth for gaming.
Enforcing charging policies for premium services.
Example: A user starts a video call. The PCF creates a dynamic PCC rule to prioritize the call traffic with high quality of service (QoS) and sends it to the SMF.



2. Predefined PCC Rules:
What are they?
These rules are preconfigured in the SMF during the initial network setup.
They are fixed in the sense that their content (e.g., traffic filters, QoS settings) is predefined and cannot be changed dynamically during operation.
However, they can be:
Activated (enabled) by the PCF when needed.
Deactivated (disabled) by the PCF when no longer required.
Grouping of Rules:


Predefined rules in the PCF can be grouped into sets.
The PCF can activate or deactivate entire groups of rules dynamically, simplifying the management of multiple related policies.
When are they used?


For standard or commonly recurring scenarios where flexibility is not required, such as:
Default internet browsing policies.
Generic traffic shaping for normal user sessions.
Example: A user connects to the internet. The PCF activates a predefined PCC rule that applies a standard QoS level for general browsing and sets a flat-rate charging policy.



Key Differences:
Aspect
Dynamic PCC Rules
Predefined PCC Rules
Defined by
PCF (dynamically, during runtime)
Preconfigured in the SMF
Flexibility
Can be created, modified, or removed anytime
Can only be activated or deactivated
Customization
Highly customizable for specific scenarios
Fixed content, predefined for common use
Management
Managed dynamically by PCF
Managed via activation/deactivation by PCF
Use Cases
Specific services (e.g., video calls, gaming)
Standard/default sessions (e.g., browsing)


Why Are These Two Types Important?
Dynamic PCC Rules provide flexibility to handle specific and complex service requirements in real-time.
Predefined PCC Rules offer simplicity and efficiency for standard, recurring scenarios, reducing processing overhead.
Additionally, predefined PCC rules may be grouped within the SMF as predefined PCC rule bases which allow the PCF to dynamically activate these sets of rules. In this case, the PCC rule identifier is used to hold the predefined PCC rule base identifier. NOTE: The operator can define a predefined PCC rule, to be activated by the SMF. Such a predefined rule is not explicitly known in the PCF. 

The content you're referring to explains how Predefined PCC Rules work in 5G networks, particularly how they can be grouped into Predefined PCC Rule Bases and how the PCF (Policy Control Function) interacts with them dynamically. Here's a detailed breakdown of this concept:
1. Predefined PCC Rules and Rule Bases:
Predefined PCC Rules: These are rules that are pre-configured in the SMF (Session Management Function) and are typically used for standard or recurring services (e.g., internet browsing, general data sessions).


Predefined PCC Rule Bases: A collection of predefined PCC rules grouped together. This grouping simplifies the management and activation of related rules.


Example:
A predefined rule base might consist of multiple rules for handling different types of data traffic (e.g., video, audio, and browsing). Instead of activating each rule individually, the operator can activate the entire base of rules as a set.

2. Dynamic Activation by the PCF:
The PCF can dynamically activate or deactivate these predefined PCC rule bases during the session.
The PCC rule identifier (which is associated with the specific set of rules) is used to activate the entire predefined PCC rule base.
How it works:
When the PCF needs to apply certain policies, it sends the PCC rule identifier to the SMF, which refers to the predefined PCC rule base identifier.
The SMF then activates the entire set of predefined rules associated with that identifier.

3. Operator-Controlled Predefined Rules:
The operator (e.g., a telecom service provider) has the ability to define specific predefined PCC rules that can be activated by the SMF during the session.


Important point: These predefined rules, though defined in the SMF, are not explicitly known in the PCF.


Meaning: The predefined rules themselves may not be visible to the PCF, but the PCF can still trigger them via the rule base identifier. 
This allows the operator to configure the network in a way that optimizes performance or charging without requiring the PCF to manage the individual rules.
Example:


An operator could define a predefined rule for handling a new "premium video streaming" service in the SMF.
The PCF doesn’t need to know the exact details of the rule, but it can activate the predefined set of rules using the corresponding PCC rule identifier when the user subscribes to that service.
4. Role of the PCC Rule Identifier:
The PCC rule identifier plays a critical role in the interaction between the PCF and the SMF.
It is used by the PCF to reference a specific predefined PCC rule base in the SMF.
This identifier ensures that the correct set of rules (i.e., rule base) is activated, which may consist of multiple predefined rules that work together.
5. Flow:
The operator defines a set of predefined PCC rules in the SMF.
These rules are grouped into a predefined PCC rule base.
The PCF doesn’t explicitly know the individual predefined rules but is aware of their rule base identifier.
When required (e.g., for a specific user or service), the PCF sends the PCC rule identifier to the SMF.
The SMF uses this identifier to activate the entire predefined PCC rule base for the session.
The rules in the base may control various parameters like QoS, charging, or traffic management, and the SMF applies them accordingly.

Key Points:
Predefined Rules: Pre-configured in the SMF, typically for general services.
Predefined Rule Base: A group of predefined rules that can be activated together.
PCF Role: The PCF triggers the activation of rule bases via the PCC rule identifier.
Operator Control: Operators can define specific rules and control their activation without explicitly involving the PCF in rule management.
This setup helps the network be flexible and efficient by allowing the operator to pre-define traffic policies and apply them dynamically through the system, without needing to update the PCF with every rule change.
Example of the above content 

Let's walk through an example where a user interacts with an application, and this results in the activation of predefined PCC rules in the SMF via the PCF.

Scenario: A User Streaming a Premium Video Service
1. The User Starts Using the Application:
A user opens a premium video streaming app on their smartphone (e.g., a new service that provides high-definition video streaming, requiring higher bandwidth and lower latency).
The app sends a data request to the network to start the video stream.
2. SMF Pre-configured with Predefined PCC Rules:
In this scenario, the SMF has predefined PCC rules set up for handling different types of traffic. One of these predefined rules is designed specifically for premium video streaming.
These rules might include:
High bandwidth allocation.
Prioritization of video streaming traffic to minimize latency.
Charging policies based on the volume of HD streaming data.
These predefined rules are grouped into a predefined PCC rule base (let's call this Rule Base 1).
3. PCF Involvement:
The PCF doesn’t need to know all the individual rules in Rule Base 1, but it has the identifier for this set of predefined rules.
When the user starts the premium video service, the PCF is notified that the user is consuming high-priority traffic (e.g., video streaming).
4. PCF Sends Rule Activation to SMF:
Based on the user's activity, the PCF sends the PCC rule identifier (which corresponds to Rule Base 1) to the SMF. The identifier instructs the SMF to apply the entire set of predefined PCC rules for video streaming to the session.
5. SMF Activates the Predefined PCC Rule Base:
Upon receiving the PCC rule identifier from the PCF, the SMF looks up Rule Base 1 and activates the relevant predefined rules.
This ensures that the user’s video traffic is prioritized, given high bandwidth, and that the session is charged according to the premium streaming plan.
6. User’s Session with Activated PCC Rules:
The SMF enforces the active predefined rules during the session:
QoS: High video resolution and low latency are ensured for a smooth streaming experience.
Charging: The user is charged based on the predefined video streaming plan (e.g., per GB of HD data).
7. Session Continues with Dynamic Changes:
If the user switches to a different type of traffic (e.g., browsing or audio), the PCF can deactivate the video streaming rules and activate a different rule base for normal internet browsing.
Similarly, if the user switches from HD to SD streaming, the PCF can modify the existing rules to lower the bandwidth allocation.

Flow of Events:
User starts using the premium video streaming app.
SMF has predefined rules for video streaming stored in Rule Base 1.
PCF identifies the user’s service (video streaming) and sends the PCC rule identifier for Rule Base 1 to the SMF.
SMF activates the entire Rule Base 1 with predefined policies (e.g., high bandwidth, charging).
SMF enforces these policies on the user’s data flow, ensuring high-quality streaming and charging based on usage.

Summary:
In this example, the PCF acts as a central controller for activating predefined PCC rules that are already configured in the SMF. The user’s action (starting the video stream) triggers the activation of these rules through the PCF, and the SMF applies them to the user’s session. This allows the network to handle specific traffic types efficiently, ensuring quality of service and appropriate charging.


A PCC rule consists of: 
PCC rule identifier -> smf to apply the pcc rules in the pcf


PCC rules operation

For dynamic PCC rules, the following applies:
- Installation: to provision the PCC rules.
- Modification: to modify the PCC rules.
- Removal: to remove the PCC rules.

For predefined PCC rules, the following operations are available:
- Activation: to activate the PCC rules.
- Deactivation: to deactivate the PCC rules.



Rule identifier : 

Uniquely identifies the PCC rule, within a PDU Session. 
It is used between PCF and SMF for referencing PCC rules. 

The Rule Identifier (Rule ID) in the 5G context, particularly in the Policy and Charging Control (PCC) rule, is a unique identifier assigned to a PCC rule that is used for enforcing policies related to Quality of Service (QoS), charging, and traffic handling. 
This concept is managed within the 5G Core Network by the Policy Control Function (PCF) and enforced by the Session Management Function (SMF) and the User Plane Function (UPF).

What is the Rule Identifier (Rule ID)?
Definition:


It is a unique identifier assigned to each PCC rule.
Helps the SMF and UPF differentiate between various rules during traffic handling.
Purpose:


Ensure appropriate QoS for different types of user data flows.
Enable proper charging mechanisms (e.g., by differentiating between free or billable services).
Define traffic treatment policies like shaping, policing, or marking packets.
Key Attributes in the Rule:


Rule ID: A unique identifier for the rule.
QoS parameters: Includes Guaranteed Bit Rate (GBR), Maximum Bit Rate (MBR), and 5QI (5G QoS Identifier).
Charging parameters: Specifies how traffic is charged, like online or offline charging.
Traffic filters: Match traffic flows based on IP headers, ports, and protocols.
Activation/deactivation triggers: Conditions for when the rule should be active.

Real-Time Example of Rule ID in Action
Scenario: Video Streaming Service with Multiple QoS Levels
Context:


A user subscribes to a 5G network and streams video from an OTT platform like Netflix.
Video streaming requires two different levels of QoS:
High-quality video stream (e.g., 4K resolution): Needs high bandwidth and low latency.
Low-quality stream for advertisements: Can tolerate lower QoS.
PCC Rule Setup:


The PCF generates two PCC rules and assigns unique Rule IDs:
Rule ID 1: For 4K video streaming (high priority, high bandwidth, low latency).
Rule ID 2: For advertisements (lower priority, reduced bandwidth).
The rules are pushed to the SMF, which translates them into policies for the UPF.
Rule ID in Operation:


Traffic Classification:
When traffic arrives at the UPF, it matches the packets with the Rule IDs using predefined traffic filters (e.g., IP address of Netflix servers, video stream port).
Enforcement:
For traffic matching Rule ID 1:
Allocates high bandwidth (e.g., 50 Mbps), applies low latency QoS parameters (e.g., 5QI = 1).
Ensures a seamless 4K video experience.
For traffic matching Rule ID 2:
Allocates lower bandwidth (e.g., 5 Mbps), applies higher latency parameters (e.g., 5QI = 8).
Displays ads without disrupting the main streaming experience.
Charging:
Traffic matching Rule ID 1 is charged at a premium rate due to higher resource usage.
Traffic matching Rule ID 2 may not be charged or could have lower rates.
Dynamic Behavior:


When the user switches from 4K streaming to standard resolution, the PCF dynamically updates the PCC rule:
Rule ID 1 is modified for standard QoS.
When the streaming session ends, both Rule ID 1 and Rule ID 2 are deactivated.

Rule ID Call Flow in 5G
Session Establishment:


User Equipment (UE) sends a PDU Session Establishment Request to the SMF via the Access and Mobility Management Function (AMF).
SMF interacts with the PCF to fetch PCC rules.
Rule Installation:


PCF sends PCC rules (including Rule IDs) to the SMF.
SMF translates PCC rules into QoS Flow Descriptions and installs them in the UPF.
Traffic Handling:


UPF enforces the rules by matching incoming traffic against the Rule IDs and applying the corresponding policies.
Session Modification/Termination:


When QoS needs change or the session ends, the SMF updates or removes the rules in the UPF.

Benefits of Rule ID
Traffic Differentiation:
Enables granular control of traffic based on application requirements.
Dynamic QoS Management:
Rules can be updated in real-time based on user behavior or network conditions.
Efficient Charging:
Different Rule IDs help implement tiered charging models.
Scalability:
Supports multiple services and users simultaneously with unique rule identifiers.

Conclusion
The Rule Identifier (Rule ID) in the PCC rule is a critical component for enabling the 5G network to manage diverse traffic flows effectively. 
It ensures seamless service delivery, optimized resource utilization, and accurate charging, making it a backbone for QoS and monetization in 5G networks.

Service data flow detection 

				Precedence 


Determines the order, in which the service data flow templates are applied at service data flow detection, enforcement and charging. 
Each PCC rule has a built-in precedence value that specifies its priority when compared to other rules.




				Service data flow 
In 5G, a Service Data Flow (SDF) is a group of data packets with the same Quality of Service (QoS) requirements, enabling precise traffic management within a PDU session.


Mute for notification
Defines whether applications start’s or stop’s notification is to muted.
				Charging key 

PDU Session Establishment:


The User Equipment (UE) sends a request to the network to establish a PDU session.
The SMF receives this request and initiates policy and charging determination.
Determine Charging Key:
The SMF identifies the Charging Key based on:
Service type (e.g., video streaming, voice call).
SDF (Service Data Flow) rules.
Subscription profiles from the Unified Data Management (UDM).
Policy and Charging Rules:
The PCF defines the QoS and charging rules (including the Charging Key) and sends them to the SMF.
Send Charging Key to CHF:
The SMF forwards the Charging Key and associated QoS information to the CHF.
The CHF uses this to apply charging policies (e.g., rate-based, volume-based).
Session Established:
The PDU session is successfully established, and the UE can start data transmission.
Data Traffic with Charging Key:
As the UE sends/receives traffic, the SMF tags it with the corresponding Charging Key for billing purposes.

Service identifier
In 5G, the Service Identifier (e.g., S-NSSAI) is used to differentiate and allocate resources for specific services via network slicing. For example, in a smart city:
URLLC slice: For autonomous vehicles needing ultra-low latency.
eMBB slice: For high-speed internet and video streaming.
mMTC slice: For IoT sensors managing smart grids or waste.
This ensures each service gets tailored performance based on its needs.
Sponsor identifier

In 5G, the Sponsor Identifier is used to associate specific network usage with a third-party entity (the sponsor) that pays for or manages the service on behalf of the user. It enables business models where sponsors (e.g., enterprises or app providers) fund the data usage for specific services.
Example: Streaming App Partnership
A video streaming app partners with a 5G operator to offer free data for its users. The Sponsor Identifier ensures that:
App users stream content without being charged for data.
The app provider (sponsor) is billed for the data usage.
This identifier links the traffic to the sponsor for charging, policy control, or service differentiation.
	

		Application Service Provider Identifier
In 5G, the Application Service Provider Identifier (ASPID) is used to identify the specific service provider offering an application or service over the network. It helps the network differentiate and apply policies, charging, or QoS (Quality of Service) rules for traffic related to that application.

Example: Gaming Platform
A cloud gaming service (e.g., "GameX") partners with a 5G operator to offer low-latency gaming. The ASPID identifies all traffic from "GameX" and ensures:
Low latency is prioritized for gaming sessions.
Billing is applied separately for "GameX" traffic.
User differentiation, e.g., premium users may get higher resolution streams.
The ASPID ensures tailored treatment for the service based on provider agreements.

