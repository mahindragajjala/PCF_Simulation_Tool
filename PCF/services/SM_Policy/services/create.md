https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.3yfemn2xdbbv

Request to create an SM Policy Association with
the PCF to receive the policy for a PDU session.

Npcf_SMPolicyControl_Create
In Npcf_SMPolicyControl_Create:
Sender: SMF
Receiver: PCF

The Npcf_SMPolicyControl_Create service operation provides means for the SMF to request the creation of a corresponding SM Policy Association with PCF.

The Session Management procedures of the SMF and related to policies are defined in 3GPP TS 23.501 [2],3GPP TS 23.502 [3] and 3GPP TS 23.503 [6].

The following procedures using the Npcf_SMPolicyControl_Create service operation are supported:

- Request for creation of a corresponding SM Policy Association with PCF.
- Provisioning of PCC rules.
- Provisioning of policy control request triggers.
- Provisioning of charging related information for PDU sessions.
- Provisioning of revalidation time.
- Policy provisioning and enforcement of authorized AMBR per PDU session.
- Policy provisioning and enforcement of authorized default QoS.
- Provisioning of PCC rule for Application Detection and Control.
- 3GPP PS Data Off Support.
- IMS Emergency Session Support.
- Request Usage Monitoring Control.
- Access Network Charging Identifier report.
- Request for the successful resource allocation notification.
- Provisioning of IP Index Information.
- Negotiation of the QoS flow for IMS signalling.
- PCF resource cleanup.
- Access traffic steering, switching and splitting support.
- DNN Selection Mode Support.
- Detection of TSN related SM Policy Association. 

When the EMDBV feature defined in sub clause 5.8 is supported by both the PCF and the SMF, the PCF shall use the extMaxDataBurstVol attribute instead of the maxDataBurstVol attribute to signal maximum data burst volume values higher than 4095 Bytes.

In 5G, Signal Maximum Data Burst Volume (MDBV) refers to a parameter that defines the maximum volume of data (measured in bytes) that a device (UE, User Equipment) is allowed to transmit or receive during a single burst of activity over a specific QoS (Quality of Service) flow.
This parameter is part of the QoS control framework in 5G, which ensures that different types of traffic (e.g., video streaming, voice calls, or IoT device telemetry) are handled with the appropriate priority, latency, and reliability. MDBV is typically associated with Guaranteed Bit Rate (GBR) flows, where specific data rates and volumes are assured for the user or application.
When the EMDBV feature is supported by the PCF but not supported by SMF and the PCF needs to signal maximum data burst volume values higher than 4095 Bytes, the PCF shall use the maxDataBurstVol attribute set to 4095 Bytes.

For values lower than or equal to 4095 Bytes, the PCF shall use the maxDataBurstVol attribute.

NOTE: Maximum data burst volume values are sent by the PCF in responses to the SMF or in an SM Policy Association Update request i.e. after feature negotiation, so the PCF knows whether the SMF supports the EMDBV feature. 


When the SMF receives the Nsmf_PDUSession_CreateSMContext Request as defined in subclause 5.2.2.2 of 3GPP TS 29.502 [22], if the SMF was requested not to interact with the PCF, the SMF shall not interact with the PCF; otherwise, the SMF shall send the POST method as step 1 of the figure 4.2.2.2-1 to request to create an "Individual SM Policy". 
In 5G networks, the SM (Session Management) Policy Association Establishment Request is sent by the SMF (Session Management Function) to the PCF (Policy Control Function). This request is part of the interaction between the SMF and PCF to establish the policy rules associated with a particular session, such as the QoS (Quality of Service) parameters, charging rules, and traffic management rules.
Call Flow for SM Policy Association Establishment
SMF to PCF: SM Policy Association Establishment Request


The SMF initiates the request to the PCF to establish a policy association for a specific session. This request includes relevant session-related information such as:
Session ID
UE (User Equipment) context or subscriber info
Bearer context
Requested QoS parameters
APN (Access Point Name) context or specific service information
PCF: Policy Decision


The PCF evaluates the request based on the subscriber's policy profile, service plan, network conditions, and the specific parameters included in the request.
The PCF then generates the appropriate policy decisions, such as QoS parameters, traffic treatment rules, charging rules, and any other service-specific rules.
PCF to SMF: SM Policy Association Establishment Answer


The PCF responds with the SM Policy Association Establishment Answer, which includes the policy decisions for the session.
This response can include:
QoS rules
Traffic steering and filtering rules
Charging rules
Application-specific rules (e.g., DPI or traffic prioritization)
Flow descriptors (for packet forwarding treatment)
SMF Configures the Session


Upon receiving the policy rules from the PCF, the SMF applies these rules to the session. This may involve setting up bearer resources, enforcing traffic management, and applying QoS policies.
SMF to AMF (Access and Mobility Management Function): Session Establishment Request
If the session is being established, the SMF might forward the appropriate session information to the AMF for UE context setup, including the policy parameters received from the PCF.
UE Configuration


The UE (User Equipment) might also be configured for the session, with specific parameters based on the policy decisions.
When is it Used?
The SM Policy Association Establishment Request is used during the process of establishing a new session or modifying an existing one, particularly in scenarios such as:
Initial session establishment: When a UE first attaches to the network and requests a data session, the SMF needs to establish a policy association with the PCF to apply the necessary policies.
Session modification: If the session parameters change (e.g., a change in QoS requirements, service-level agreements, or traffic handling), the SMF will request an update to the policy association.
Interworking with other services: In cases where a user’s traffic needs to be treated differently (e.g., prioritization or differential handling for voice vs. video data), the SMF uses the SM Policy Association to configure the necessary rules.
Example Use Cases
Quality of Service (QoS) Management: The PCF can enforce QoS rules like maximum bandwidth, latency limits, or traffic prioritization.
Charging: Charging rules such as data volume limits, time-based usage, or service-based billing can be enforced.
Traffic Steering: In case of a need to steer traffic through specific network slices or apply special treatment for certain types of traffic (e.g., video, VoLTE).
This call flow helps ensure that the network provides the right treatment to the user’s traffic based on the policy decisions that align with their subscription and service-level agreement.

