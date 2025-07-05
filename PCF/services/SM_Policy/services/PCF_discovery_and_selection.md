https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.30n34p9allp1

PCF discovery and selection


describes the underlying principles for PCF selection and discovery: - 

There may be multiple and separately addressable PCFs in a PLMN. 
The PCF must be able to correlate the AF service session established over N5 or Rx with the associated PDU Session (Session binding) handled over N7. 
It shall be possible to deploy a network so that the PCF may serve only specific DN(s). 
For example, Policy Control may be enabled on a per DNN basis.
Unique identification of a PDU Session in the PCF shall be possible based on the (UE ID, DNN)-tuple, the (UE (IP or MAC) Address(es), DNN)-tuple and the (UE ID, UE (IP or MAC) Address(es), DNN). 
1. Multiple and Separately Addressable PCFs in a PLMN
Explanation:
In a Public Land Mobile Network (PLMN), there can be multiple Policy Control Functions (PCFs) deployed.
These PCFs are independently identifiable and addressable. This means different PCFs can handle different aspects of policy control and decisions within the network.
For example, one PCF might handle policies for certain types of devices (e.g., IoT devices), while another PCF could focus on policies for high-bandwidth applications (e.g., video streaming).
Use Case:
This separation allows for scalability and specialization. For instance, a large network can assign PCFs to specific geographical areas, network slices, or services.
Example:
In a large telecom operator's network:
PCF-1 is deployed for IoT traffic (e.g., connected smart meters or wearables) in Region A.
PCF-2 handles policies for high-speed mobile broadband services (e.g., streaming and gaming) in Region B.
PCF-3 is specialized to manage policy control for enterprise customers (e.g., private 5G networks for factories).
Each PCF operates independently but within the same PLMN, ensuring different use cases are handled efficiently.

Call Flow
User Equipment (UE) initiates a PDU Session request to access a Data Network (DN) through the network.
The request includes the DNN (Data Network Name), e.g., "internet.company.com".
The Access and Mobility Function (AMF) identifies the request and forwards it to the Session Management Function (SMF).
The SMF interacts with the Network Repository Function (NRF) to discover a suitable PCF based on the DNN or UE attributes.
For example:
PCF-1 for IoT devices connecting to iot.network.com.
PCF-2 for consumers connecting to internet.company.com.
PCF-3 for enterprise users connecting to enterprise.company.com.

2. Correlation Between AF Service Session and PDU Session (Session Binding)
Explanation:
The Application Function (AF) interacts with the PCF over the N5 (or Rx) reference point to request policies for an application or service.
The PCF is also responsible for controlling the PDU Session (Packet Data Unit session) over the N7 reference point.
Session binding means the PCF must correlate or link the AF service session with the associated PDU session.
This ensures that the policies requested by the AF are applied to the correct data session.
Use Case:
Imagine a video streaming application requesting specific Quality of Service (QoS) policies (e.g., higher bandwidth). The PCF ensures that these policies are applied to the correct PDU session of the user’s device, enabling smooth playback.
Example:
A user streams a live football match on their smartphone via a video streaming app.
The Application Function (AF), representing the streaming app, requests specific Quality of Service (QoS) requirements from the PCF via N5 (e.g., high bandwidth, low latency).
The PCF correlates this request with the ongoing PDU Session (e.g., the user’s data session to the internet) over N7.
The PCF applies the requested QoS policies (e.g., prioritizing video packets) to the user’s data session, ensuring a smooth streaming experience without buffering.
CALL FLOW

The SMF sends a policy request to the selected PCF over the N7 interface.
This request includes details such as:
UE identity (IMSI).
Requested DNN.
QoS requirements for the session.
UE IP or MAC address (if already assigned).
Simultaneously, an Application Function (AF) (e.g., a video streaming app) interacts with the PCF over the N5 interface to request specific application-level policies.
For example, the AF requests low latency and high bandwidth for video streaming.
The PCF binds the AF service session with the corresponding PDU session using session correlation logic:
The PCF links the AF request (on N5) to the PDU session (on N7) based on identifiers such as the (UE ID, DNN) tuple.

3. PCF Serving Specific Data Networks (DNs)
Explanation:
The PCF can be configured to serve specific Data Networks (DNs). A DN refers to external networks, such as the internet or private enterprise networks, that the user accesses via the PLMN.
For example, a network can be set up so that a PCF applies policy control only for certain Data Network Names (DNNs). A DNN is essentially a label or identifier for a data network.
This enables granular control over policy enforcement, tailored to specific DNs.
Use Case:
Suppose there’s a DNN for corporate users and another for general internet access. The PCF may be configured to apply stricter security policies for the corporate DNN while applying QoS policies for the internet DNN.
Example:
A mobile operator provides access to two Data Networks (DNs):
DNN-1: Internet access for regular consumers.
DNN-2: A private corporate network for employees of a specific company.
The network is configured such that:
PCF-1 applies standard QoS and parental control policies for DNN-1.
PCF-2 applies strict security and access control policies for DNN-2, ensuring only authorized devices can access the corporate network.
This separation ensures that the policies for each DN are appropriately applied without interference.

CALL FLOW

The PCF:
Evaluates the policy rules based on the AF request, SMF request, and pre-configured operator policies.
Makes policy decisions, e.g., setting bandwidth limits, prioritizing traffic, or applying security rules.
The PCF sends the policy decisions to the SMF over the N7 interface.
Example policies:
For DNN: internet.company.com: Enable QoS with high bandwidth for video streaming.
For DNN: enterprise.company.com: Apply stricter security and limit access to specific IP ranges.

4. Unique Identification of a PDU Session
Explanation:
A PDU Session represents a user’s connection to a specific data network (DN) through the mobile network.
In the PCF, each PDU session must have a unique identifier to avoid confusion and ensure proper policy enforcement.
The unique identification can be achieved using the following combinations:
(UE ID, DNN): The combination of the User Equipment Identifier (UE ID) (e.g., IMSI or subscription ID) and the Data Network Name (DNN).
(UE Address(es), DNN): The combination of the UE’s IP address or MAC address and the DNN.
(UE ID, UE Address(es), DNN): A more detailed tuple that includes the UE ID, its address(es), and the DNN.
These identifiers provide flexibility in how PDU sessions are tracked and managed.
Use Case:
A device (UE) with multiple active PDU sessions (e.g., one for a corporate network and another for general internet browsing) can be uniquely identified and managed using these tuples.
a) Using (UE ID, DNN):
Example:
A subscriber (UE) with the IMSI (International Mobile Subscriber Identity) 123456789012345 connects to the DNN internet.company.com.
The tuple (123456789012345, internet.company.com) uniquely identifies this PDU session in the PCF.
b) Using (UE Address(es), DNN):
Example:
A smartphone with the IP address 192.168.1.10 connects to the DNN internet.company.com.
The tuple (192.168.1.10, internet.company.com) uniquely identifies this session.
c) Using (UE ID, UE Address(es), DNN):
Example:
A tablet with the IMSI 987654321098765 and IP address 10.0.0.5 connects to the DNN enterprise.company.com.
The tuple (987654321098765, 10.0.0.5, enterprise.company.com) uniquely identifies this session in the PCF.
CALL FLOW
During the session, the PCF uniquely identifies the PDU Session using one of the following tuples:
(UE ID, DNN): For example, IMSI: 123456789012345, DNN: internet.company.com.
(UE Address, DNN): For example, IP Address: 192.168.1.10, DNN: internet.company.com.
(UE ID, UE Address, DNN): For example, IMSI: 987654321098765, IP Address: 10.0.0.5, DNN: enterprise.company.com.
The PCF ensures consistent policy application and handles updates or modifications during the session (e.g., QoS adjustments)

UE → AMF: PDU session establishment request.
AMF → SMF: Forward the request for DNN and session setup.
SMF → PCF (via NRF): Discover and select the appropriate PCF.
AF → PCF: Application policy request (over N5).
PCF: Correlates the AF service session with the PDU session (N7).
PCF → SMF: Policy decisions sent to SMF.
SMF → UPF: Enforce policy rules.
UE ↔ DN: Data exchange begins with applied policies.



PCF discovery and selection for a UE or a PDU Session 

PCF discovery and selection functionality is implemented in AMF, SMF and SCP, and follows the principles in clause 6.3.1. 

The AMF uses the PCF services for a UE and the SMF uses the PCF services for a PDU Session. 

Amf - Pcf services /ue 
Smf - Pcf services /pdu session 


PCF Discovery and Selection by AMF
The AMF (Access and Mobility Function) is responsible for selecting the appropriate PCF for both Access and Mobility (AM) policy association and UE policy association.
Key Factors for PCF Discovery and Selection
Use of NRF:
The AMF uses the NRF to discover available PCF instances.
Alternatively, PCF information can be locally configured in the AMF based on operator policies.
Selection Criteria:
SUPI (Subscription Permanent Identifier):
The AMF selects a PCF instance based on the SUPI range the UE's SUPI belongs to or through discovery using the UE’s SUPI.
S-NSSAI (Single-Network Slice Selection Assistance Information):
In roaming scenarios, the selection is based on:
S-NSSAI of the visited PLMN (VPLMN) for V-PCF.
S-NSSAI of the home PLMN (HPLMN) for H-PCF.
PCF Group ID and Set ID:
The AMF infers the PCF Group ID through the NRF discovery process.
DNN (Data Network Name) replacement capability:
This ensures the PCF instance supports the necessary DNN configurations.
Behavior in Roaming and Non-Roaming Cases:
Non-Roaming Case:
The same PCF instance is used for AM policy and UE policy associations.
Roaming Case:
The AMF selects a V-PCF (in the visited network) for both AM and UE policy associations.
Call Flow: PCF Discovery and Selection by AMF
Trigger:
AMF receives a request to establish a policy association for a UE.
Discovery:
AMF queries the NRF with discovery parameters such as SUPI, S-NSSAI, or PCF Group ID.
Response:
NRF returns a list of PCF instances that match the discovery criteria.
Selection:
AMF selects one PCF instance based on operator policies and local configurations.
Association:
AMF establishes an AM policy association with the selected PCF.

PCF Discovery and Selection by SMF
The SMF (Session Management Function) is responsible for selecting a PCF for PDU Session policy control.
Key Factors for PCF Discovery and Selection
Use of NRF:
SMF uses the NRF to discover PCF instances, or it relies on locally configured PCF information.
Selection Criteria:
Operator Policies:
These dictate how PCF instances are selected.
DNN:
The selected DNN determines which PCF instance is used.
S-NSSAI:
Similar to AMF, the S-NSSAI influences PCF selection in roaming cases.
SUPI:
The PCF is selected based on the SUPI range or through NRF discovery.
PCF Group ID and Set ID:
Provided by the AMF, this aids in the selection process.
Multi-Access (MA) PDU Sessions:
Specific PCFs supporting MA PDU sessions are chosen.
Delegated Discovery and Selection by SCP
The SCP (Service Communication Proxy) can perform PCF discovery and selection on behalf of the SMF.
SCP uses discovery parameters provided by the SMF (e.g., S-NSSAI, PCF Group ID).
Call Flow: PCF Discovery and Selection by SMF
Trigger:
SMF receives a PDU session establishment request from the AMF or UE.
Discovery:
SMF queries the NRF with discovery parameters such as DNN, SUPI, or S-NSSAI.
Response:
NRF returns a list of candidate PCF instances.
Selection:
SMF selects one PCF instance based on local configurations or operator policies.
Association:
SMF establishes a session policy association with the selected PCF.
AMF Relocation and PCF Re-selection
During AMF relocation, the target AMF may receive the selected PCF ID and PCF Group ID from the source AMF.
The target AMF can decide to:
Use the same PCF.
Select a new PCF based on operator policies.
Call Flow: AMF Relocation and PCF Re-selection
Trigger:
AMF relocation is initiated due to UE mobility.
Source AMF Behavior:
Source AMF transfers the current PCF ID and Group ID to the target AMF.
Target AMF Behavior:
Target AMF evaluates operator policies and either:
Retains the current PCF.
Initiates a new PCF discovery and selection procedure.
Association Update:
If a new PCF is selected, the AMF establishes a new policy association.

Redirection of SM Policy Control
If the PCF redirects the SMF to another PCF instance:
SMF terminates the current policy association.
SMF performs a new PCF discovery and selection based on the redirected PCF ID.
Call Flow: PCF Redirection by SMF
Trigger:
SMF receives a redirection indication from the current PCF.
Termination:
SMF terminates the existing policy association.
Discovery and Selection:
SMF discovers and selects a new PCF instance.
Reassociation:
SMF establishes a new policy control association with the redirected PCF.

Delegated Discovery and Selection in SCP
SCP can be delegated to perform PCF discovery and selection on behalf of the AMF or SMF.
SCP uses available parameters (e.g., PCF ID, Group ID, DNN, S-NSSAI) to identify the appropriate PCF.

Summary of Call Flows
The described scenarios involve multiple call flows:
Initial PCF Discovery:
NRF is queried by the AMF or SMF for available PCFs.
PCF Association Establishment:
The selected PCF instance is used to establish policy associations.
AMF Relocation:
PCF information is transferred between source and target AMFs.
SM Policy Control Redirection:
SMF reselects a new PCF upon receiving a redirection indication.
SCP Delegated Discovery:
SCP performs PCF discovery and provides the PCF ID to the AMF or SMF.

