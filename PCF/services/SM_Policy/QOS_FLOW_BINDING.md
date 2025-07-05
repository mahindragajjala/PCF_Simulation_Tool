
QoS Flow binding

QoS Flow binding is the association of a PCC rule to a QoS Flow within a PDU Session. The binding is performed
using the following binding parameters:
- 5QI;
- ARP;
- QNC (if available in the PCC rule);
- Priority Level (if available in the PCC rule);
- Averaging Window (if available in the PCC rule);
- Maximum Data Burst Volume (if available in the PCC rule);
- Alternative QoS Parameter Set(s) (if available in the PCC rule).

1. 5QI (5G QoS Identifier):
Definition: A numerical identifier (ranging from 1 to 255) that maps to a set of QoS characteristics like latency, packet loss rate, and priority.
Purpose: It simplifies QoS enforcement by avoiding detailed signaling of each QoS attribute.
Example:
5QI 1: Reserved for ultra-reliable and low-latency communication (URLLC), with stringent latency and reliability requirements (e.g., autonomous driving or industrial automation).
5QI 9: Typically used for default internet traffic like web browsing and video streaming.
Key Attributes Defined by 5QI:
Resource type (e.g., Guaranteed Bit Rate (GBR) or Non-GBR)
Priority level
Packet delay budget
Packet error rate
2. ARP (Allocation and Retention Priority):
Definition: Determines the priority for allocating resources to a session and retaining resources during congestion.
Purpose: Ensures that critical services (e.g., emergency calls) are given higher priority over less critical ones.
Components:
Priority Level: Indicates how important the session is compared to others.
Pre-emption Capability: Indicates whether the session can preempt lower-priority sessions.
Pre-emption Vulnerability: Indicates whether the session can be preempted by higher-priority ones.
Example: Emergency services (e.g., ambulances) will have a high ARP to ensure they get resources even during heavy network traffic.

3. QNC (QoS Notification Control) (if available in the PCC rule):
Definition: A mechanism where the network notifies the User Equipment (UE) of changes in QoS parameters, such as bandwidth reduction or increased latency.
Purpose: Enables the UE to adapt to changes in network conditions or resources dynamically.
Example: If network congestion leads to a reduced data rate for a video stream, the UE may be informed to adjust the video quality.

4. Priority Level (if available in the PCC rule):
Definition: Indicates the relative importance of traffic flows within the network.
Purpose: Ensures critical traffic (e.g., voice calls, emergency services) is prioritized over less critical traffic (e.g., file downloads or bulk data transfers).
Range: Priority levels are usually defined from 1 (highest priority) to 256 (lowest priority).
Example: A VoIP call might have a priority level of 5, whereas a bulk data transfer might have a priority level of 100.

5. Averaging Window (if available in the PCC rule):
Definition: A time window used to calculate and enforce QoS parameters like Guaranteed Bit Rate (GBR) and Maximum Bit Rate (MBR).
Purpose: Ensures that traffic flows are averaged over a specific period to meet their QoS requirements without short-term spikes or dips impacting performance.
Example: For a video stream with a GBR of 2 Mbps, the network will average the data flow over a 1-second window to ensure the rate is consistently maintained.

6. Maximum Data Burst Volume (if available in the PCC rule):
Definition: The maximum amount of data that can be transmitted in a single burst without violating QoS constraints.
Purpose: Helps manage traffic bursts that could overload the network or affect other users' QoS.
Example: In video streaming, a buffer fill might cause a sudden spike in data flow. The Maximum Data Burst Volume parameter ensures this does not exceed a pre-defined limit.

7. Alternative QoS Parameter Set(s) (if available in the PCC rule):
Definition: Specifies backup QoS parameter sets that can be used if the primary QoS set cannot be fulfilled (e.g., due to congestion).
Purpose: Enhances flexibility and ensures service continuity even when network conditions degrade.
Example:
Primary QoS Set: 5QI = 1, Guaranteed Bit Rate (GBR) = 1 Mbps
Alternative QoS Set: 5QI = 2, GBR = 512 Kbps
If the network cannot meet the primary set, it falls back to the alternative set.


When the PCF provisions a PCC Rule, the SMF shall evaluate whether a QoS Flow with QoS parameters identical to
the binding parameters exists unless the PCF requests to bind the PCC rule to the QoS Flow associated with the default
QoS rule. If no such QoS Flow exists, the SMF derives the QoS parameters, using the parameters in the PCC Rule, for a
new QoS Flow, binds the PCC Rule to the QoS Flow and then proceeds as described TS 23.501 [2] clause 5.7.1.5 to
establish the new QoS Flow. If a QoS Flow with QoS parameters identical to the binding parameters exists, the SMF
binds the PCC Rule to this QoS Flow and proceeds as described TS 23.501 [2] clause 5.7.1.5 to modify the QoS Flow
unless local policies, e.g. to ensure that a PCC Rule with Alternative QoS parameter Set(s) is bound to a separate QoS
flow, or the below mentioned conditions (which QoS Flow binding shall ensure), require the establishment of a new
QoS Flow following the actions described above.
NOTE 1: For PCC rules containing a delay critical GBR 5QI value, the SMF can bind PCC Rules with the same
binding parameters to different QoS Flows to ensure that the GFBR of the QoS Flow can be achieved
with the Maximum Data Burst Volume of the QoS Flow.
The SMF shall identify the QoS Flow associated with the default QoS rule based on the fact that the PCC rule(s) bound
to this QoS Flow contain:
- 5QI and ARP values that are identical to the PDU Session related information Authorized default 5QI/ARP; or
- a Bind to QoS Flow associated with the default QoS rule and apply PCC rule parameters Indication.
NOTE 2: The Bind to QoS Flow associated with the default QoS rule and apply PCC rule parameters Indication has
to be used whenever the PDU Session related information Authorized default 5QI/ARP (as described in
clause 6.3.1) cannot be directly used as the QoS parameters of the QoS Flow associated with the default
QoS rule, for example when a GBR 5QI is used or the 5QI priority level has to be changed.
When a QoS Flow associated with the default QoS rule exists, the PCF can request that a PCC rule is bound to this QoS
Flow by including the Bind to QoS Flow associated with the default QoS rule Indication in a dynamic PCC rule. In this
case, the SMF shall bind the dynamic PCC rule to the QoS Flow associated with the default QoS rule (i.e. ignoring the
binding parameters) and keep the binding as long as this indication remains set. When the PCF removes the association
of a PCC rule to the QoS Flow associated with the default QoS rule, a new binding may need to be created between this
PCC rule and a QoS Flow based on the binding mechanism described above.
The binding created between a PCC Rule and a QoS Flow causes the downlink part of the service data flow to be
directed to the associated QoS Flow at the UPF (as described in TS 23.501 [2] clause 5.7.1). In the UE, the QoS rule
associated with the QoS Flow (which is generated by the SMF and explicitly signalled to the UE as described in
TS 23.501 [2] clause 5.7.1) instructs the UE to direct the uplink part of the service data flow to the QoS Flow in the
binding.
Whenever the binding parameters of a PCC rule changes, the binding of this PCC rule shall be re-evaluated, i.e. the
binding mechanism described above is performed again. The re-evaluation may, for a PCC rule, result in a new binding
with another QoS Flow. If the PCF requests the same change of the binding parameter value(s) for all PCC rules that
are bound to the same QoS Flow, the SMF should not re-evaluate the binding of these PCC rules and instead, modify
the QoS parameter value(s) of the QoS Flow accordingly.


This explanation breaks down the technical details into simpler terms for better understanding:

Context:
The Policy and Charging Function (PCF) in 5G creates rules (PCC Rules) that specify how network traffic should be handled for different services. The Session Management Function (SMF) enforces these rules by managing QoS Flows—dedicated paths in the network with specific Quality of Service (QoS) guarantees.
Here’s a step-by-step explanation of the process:

1. Checking for Existing QoS Flows:
When the PCF sends a new PCC Rule to the SMF, the SMF checks if a QoS Flow already exists with the same QoS parameters (e.g., 5QI, ARP, etc.).
If such a QoS Flow exists:
The SMF binds the new PCC Rule to this existing QoS Flow.
If needed, it updates (modifies) the QoS Flow to meet the new requirements.
If no matching QoS Flow exists:
The SMF creates a new QoS Flow based on the parameters in the PCC Rule.
The SMF then establishes this new QoS Flow and binds the PCC Rule to it.

2. Special Cases for QoS Flow Binding:
Delay-Critical Traffic:
 For services that need extremely low delay (e.g., real-time video or industrial automation), the SMF may create separate QoS Flows even if some parameters are the same. This ensures the service gets enough bandwidth and meets its performance guarantees.
Default QoS Rule Flow:
Every PDU session (data session) has a default QoS rule.
If the PCF specifies that a PCC Rule should use the default QoS Flow, the SMF will bind the rule to it.
Example: Regular internet browsing might use this default flow.

3. Parameters for Default QoS Rule Binding:
The SMF recognizes a QoS Flow as the default one if:
The 5QI and ARP values match the session’s default settings, OR
The PCF specifically requests the PCC Rule to bind to the default flow.

4. Re-Evaluation of QoS Flow Binding:
If the QoS parameters of a PCC Rule change (e.g., priority or bandwidth requirements), the SMF checks if the existing binding is still valid:
If it’s not valid: The SMF creates or uses a different QoS Flow.
If it’s valid for all bound rules: The SMF updates the QoS Flow’s parameters without creating a new one.

5. How Binding Works in Practice:
Downlink (Network to Device):
 The SMF tells the User Plane Function (UPF) to direct traffic for a specific service to the correct QoS Flow.
Uplink (Device to Network):
 The SMF generates a QoS rule for the device (UE), instructing it to send the service data to the appropriate QoS Flow.

6. Why Re-Evaluation Matters:
Whenever a PCC Rule’s parameters change or the PCF requests updates, the SMF revisits the bindings to ensure they still align with the current requirements.
For example:
If the PCF changes the 5QI for all rules bound to a QoS Flow, the SMF might directly update the QoS Flow instead of creating new ones.

Real-Life Analogy:
Imagine a shipping company managing delivery trucks (QoS Flows) for parcels (PCC Rules):
If a new parcel fits perfectly in an existing truck (matching QoS Flow), it’s added to that truck.
If no truck can handle the new parcel (e.g., weight or size limits), the company sends out a new truck designed specifically for that parcel.
If a parcel’s delivery priority changes (e.g., expedited shipping), the company rechecks which truck it belongs to or sends a new one.

Summary:
The SMF ensures that each service (defined by PCC Rules) gets the right network resources (QoS Flows) to meet its performance needs. It:
Reuses existing QoS Flows when possible,
Creates new ones when necessary,
Continuously monitors and adjusts bindings to maintain efficient and high-quality service delivery.

