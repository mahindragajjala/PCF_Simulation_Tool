
https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.srzt65q4imcm




mobile was switched on 
new sim has been inserted 
new sim has activated before only 
how mobile data will come to mobile to search the www.youtube.com 
what is the call flow to get that mobile data.
explain the call flow from the initial registration to PDU session in the 5g

To access mobile data on a 5G network, several steps are involved from the initial registration to establishing a PDU (Packet Data Unit) session. Here’s a detailed explanation of the call flow:
1. Mobile Registration
a. Power On and Attach Request
Mobile Device (UE) powers on and sends an Attach Request to the nearest 5G base station, known as gNB (gNodeB).
This request includes the UE’s identity, such as the IMSI (International Mobile Subscriber Identity).
b. gNB Processing
The gNB forwards this request to the AMF (Access and Mobility Management Function).
c. Authentication
The AMF sends an authentication request to the AUSF (Authentication Server Function).
The AUSF verifies the UE’s identity by communicating with the UDM (Unified Data Management), where subscriber information is stored.
Once authenticated, the AUSF sends an authentication response to the AMF.
d. Security Mode Command
The AMF sends a Security Mode Command to the UE to establish encryption and integrity protection for the communication.
The UE responds with a Security Mode Complete message.
e. Registration Accept
The AMF sends a Registration Accept message to the UE, confirming that it is registered to the network.
The UE acknowledges this with a Registration Complete message.
2. Establishing a PDU Session
a. PDU Session Establishment Request
The UE sends a PDU Session Establishment Request to the gNB.
This request includes information about the desired data session, such as the Data Network Name (DNN, often referred to as APN in 4G), and requested Quality of Service (QoS).
b. gNB Processing
The gNB forwards this request to the AMF.
The AMF then communicates with the SMF (Session Management Function), which manages the session details.
c. SMF Interaction with UPF
The SMF communicates with the UPF (User Plane Function) to set up the data path for the PDU session.
The UPF allocates an IP address to the UE and establishes the necessary tunnels for data transfer.
d. PDU Session Accept
The SMF sends a PDU Session Establishment Accept message to the AMF, which then forwards it to the gNB, and finally to the UE.
The message includes the assigned IP address, QoS parameters, and other session information.
e. PDU Session Establishment Complete
The UE sends a PDU Session Establishment Complete message to the gNB, which forwards it to the AMF and SMF.
The PDU session is now fully established, allowing the UE to send and receive data.
3. Data Transfer
With the PDU session established, the UE can now access the internet.
The UE initiates a DNS query to resolve the IP address of www.youtube.com.
Once the IP address is resolved, the UE can start sending HTTP requests to YouTube’s servers, and data will flow through the established PDU session.
Summary
Initial Registration: UE → gNB → AMF → AUSF/UDM → AMF → UE
PDU Session Establishment: UE → gNB → AMF → SMF → UPF → SMF → AMF → gNB → UE
Data Transfer: UE → gNB → UPF → Internet (YouTube)
This call flow ensures that the mobile device is authenticated, registered, and has an active data session for internet access.


