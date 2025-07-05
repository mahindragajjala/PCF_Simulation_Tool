https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.cit47to1llo

It seems like you're referring to N1 and N2 messages and interfaces, which are part of 5G core network signaling. Let me explain both interfaces and their respective roles with call flows.

N1 Interface
Definition: The N1 interface connects the User Equipment (UE) to the Access and Mobility Management Function (AMF) in the 5G core network.
Purpose: It facilitates signaling for user registration, session management, mobility, and other control plane functions.
Message Examples:
NAS (Non-Access Stratum) messages like Registration Request, Service Request, Deregistration Request, etc.

N2 Interface
Definition: The N2 interface connects the Access Network (e.g., gNB) to the AMF.
Purpose: It handles signaling for mobility management, bearer setup/modification, and session management between the access network and the core network.
Message Examples:
NGAP (Next Generation Application Protocol) messages like Initial UE Message, UE Context Setup Request, and PDU Session Resource Setup.

How N1 and N2 Work Together
The N1 interface is used for direct signaling between the UE and the AMF, carrying NAS messages.
The N2 interface is used for signaling between the gNB and the AMF, carrying NGAP messages.
The gNB serves as a relay for NAS messages over N1 but communicates directly with the AMF over N2 for network-level signaling.

Call Flows
1. UE Registration (Attach Procedure)
Step 1:


N1 (NAS):
UE sends a Registration Request to the AMF via the gNB.
The gNB relays this NAS message over the N2 interface to the AMF in an Initial UE Message.
Step 2:


N2 (NGAP):
The AMF may respond to the gNB with UE Context Setup Request to establish the connection for the UE.
gNB sets up the required resources and responds with UE Context Setup Response.
Step 3:


N1 (NAS):
The AMF sends a Registration Accept message to the UE via the gNB.

2. PDU Session Establishment (Data Session Setup)
Step 1:


N1 (NAS):
UE sends a PDU Session Establishment Request to the AMF via the gNB.
The gNB relays this request to the AMF over the N2 interface.
Step 2:


N2 (NGAP):
The AMF communicates with the SMF (Session Management Function) to create the PDU session.
The AMF sends PDU Session Resource Setup Request to the gNB, providing QoS parameters and tunnel details for the session.
Step 3:


N2 (NGAP):
The gNB establishes the necessary resources and responds with PDU Session Resource Setup Response.
Step 4:


N1 (NAS):
The AMF relays the PDU Session Establishment Accept to the UE, completing the session setup.

3. Handover Procedure
Step 1:


N2 (NGAP):
The source gNB sends a Handover Required message to the AMF when a handover is needed.
Step 2:


N2 (NGAP):
The AMF sends a Handover Request to the target gNB to prepare resources.
Step 3:


N1 (NAS):
After resources are ready, the UE is instructed via NAS messages (Handover Command) to move to the target gNB.
Step 4:


N2 (NGAP):
The target gNB notifies the AMF of a successful handover with Handover Notify.

Summary
The N1 interface deals with UE-specific signaling between the UE and the AMF.
The N2 interface facilitates signaling between the gNB and the AMF for resource and mobility management.
Both interfaces are essential for seamless registration, session management, and mobility handling in 5G networks.
The initial registration procedure in a 5G network is critical as it establishes the UE's first connection with the core network. This process involves signaling over both the N1 (UE ↔ AMF) and N2 (gNB ↔ AMF) interfaces. Below is a step-by-step explanation with call flows.

Initial Registration Call Flow
Step 1: UE Sends a Registration Request
Interface: N1
Message: Registration Request (NAS)
The UE sends a Registration Request NAS message to the gNB. This message contains:
UE identity (e.g., SUPI or GUTI).
Requested network features (e.g., PDU session or mobility).
PLMN (Public Land Mobile Network) information.
The gNB encapsulates the NAS message in an NGAP Initial UE Message and sends it to the AMF over the N2 interface.

Step 2: AMF Selects UE Context
Interface: N2
Message: NGAP Initial Context Setup Request
The AMF processes the Registration Request message.
It selects the appropriate policies, validates the UE's subscription (with the UDM), and assigns a GUTI (Globally Unique Temporary Identifier) to the UE if necessary.
The AMF sends an Initial Context Setup Request to the gNB via the N2 interface to establish resources for the UE.
This message includes:
NAS Authentication Request or Security Mode Command.
Temporary signaling information (e.g., QoS and context).

Step 3: UE Authentication and Security
Interface: N1
Message: Authentication Request/Response and Security Mode Command/Complete
The AMF sends an Authentication Request to the UE via the gNB.
The UE responds with an Authentication Response after validating the AMF.
Security procedures are initiated with the Security Mode Command (e.g., encryption and integrity protection algorithms).
The UE responds with Security Mode Complete.

Step 4: Registration Accept
Interface: N1
Message: Registration Accept (NAS)
After successful authentication and security procedures, the AMF sends a Registration Accept message to the UE.
This message includes:
GUTI assigned to the UE.
Allowed NSSAI (Network Slice Selection Assistance Information).
Any required configuration updates.

Step 5: Initial Context Setup Response
Interface: N2
Message: NGAP Initial Context Setup Response
The gNB sends an Initial Context Setup Response back to the AMF to confirm that resources for the UE have been successfully established.

Step 6: Registration Complete
Interface: N1
Message: Registration Complete (NAS)
The UE acknowledges the successful completion of the registration process by sending a Registration Complete message to the AMF.

Simplified Call Flow Diagram
UE                gNB                AMF                UDM
| Registration Req |                 |                  |
|------------------>                 |                  |
|                    Initial UE Message                |
|------------------------------------>                  |
|                    Auth Req / Security Mode Command   |
|<-----------------------------------                   |
| Auth Resp / Security Mode Complete |                  |
|------------------>                 |                  |
|                    Registration Accept                |
|<-----------------------------------                   |
| Registration Complete              |                  |
|------------------>                 |                  |
|                    Initial Context Setup Req/Resp     |
|------------------------------------>                  |


Key Points to Remember
Role of N1 and N2:


N1 (NAS): Handles direct communication between the UE and the AMF for signaling messages like Registration Request, Authentication, and Registration Complete.
N2 (NGAP): Handles communication between the gNB and the AMF to set up UE-specific context and resources.
Authentication:


The AMF communicates with the UDM (Unified Data Management) to verify the UE's subscription during the Authentication procedure.
Security:


Security Mode Command ensures that all signaling messages after this step are encrypted and integrity-protected.
Finalization:


The process ends with a Registration Complete message from the UE, indicating that it is now registered and ready to access services.

Would you like a deeper explanation of any specific step or message in the call flow?

