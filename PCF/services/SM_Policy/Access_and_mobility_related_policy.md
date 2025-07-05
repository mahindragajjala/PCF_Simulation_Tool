https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.l0ph2k7idn92

The access and mobility policy control encompasses(surrounded) 
Management of service area restrictions
Management of the RFSP functionalities
UE-AMBR
Management of the SMF selection. 

This clause defines the

Management of service area restrictions and RFSP Index for a UE registered over 3GPP access. 
The management of service area restrictions for a 5G-RG or a FN-CRG using W-5GAN(wi-fi)are specified in TS 23.316 [27].
The management of service area restrictions enables the PCF of the serving PLMN (e.g. V-PCF in roaming case) to modify the service area restrictions used by AMF as described in TS 23.501 [2] clause 5.3.4.
A UE's subscription may contain service area restrictions, which may be further modified by PCF based on operator defined policies at any time, either by expanding a list of allowed TAIs or by reducing non-allowed TAIs or by increasing the maximum number of allowed TAIs. 
Operator defined policies in the PCF may depend on input data such as UE location, time of day, information provided by other NFs, etc.
The AMF may report the subscribed service area restrictions received from UDM during Registration procedure or when the AMF changed, the conditions for reporting are that local policies in the AMF indicate that Access and Mobility Control is enable. 
The AMF reports the subscribed service area restrictions to the PCF also when the policy control request trigger for service area restrictions change, as described in clause 6.1.2.5, is met. 

The AMF receives the modified service area restrictions from the PCF. 

The AMF stores them then use it to determine mobility restriction for a UE. 

The PCF may indicate the AMF that there is an unlimited service area.

The service area restrictions consist of a list of allowed TAI(s) or a list of non-allowed TAI(s) and optionally the maximum number of allowed TAIs.

NOTE 1: The enforcement of the service area restrictions is performed by the UE, when the UE is in CM-IDLE state or in CM-CONNECTED state when in RRC Inactive, and in the RAN/AMF when the UE is in CMCONNECTED state.

The management of the RFSP Index enables the PCF to modify the RFSP Index used by the AMF to perform radio resource management functionality as described in TS 23.501 [2] clause 5.3.4. 

PCF modifies the RFSP Index based on operator policies that take into consideration e.g. accumulated usage, load level information per network slice instance etc. 

The subscribed RFSP Index may be further adjusted by the PCF based on operator policies at any time.

For radio resource management, the AMF may report the subscribed RFSP Index received from UDM during the Registration procedure or when the AMF changed. 

The conditions for reporting are that local policies in the AMF indicate that Access and Mobility Control is enable. 

The AMF reports the subscribed RFSP Index to the PCF when the subscription to RFSP Index change to the PCF is met. 

The AMF receives the modified RFSP Index from the PCF.

NOTE 2: The enforcement of the RFSP Index is performed in the RAN.

Upon change of AMF, the source AMF informs the PCF that the UE context was removed in the AMF in the case of inter-PLMN mobility.

The management of UE-AMBR enables the PCF to provide the UE-AMBR information to AMF based on serving network policy. 

The AMF may report the subscribed UE-AMBR received from UDM. The conditions for reporting are that the PCF provided Policy Control Request Triggers to the AMF to report subscriber UE-AMBR change. 

The AMF receives the modified UE-AMBR from the PCF. The AMF provides a UE-AMBR value of the serving network to RAN as specified in TS 23.501 [2], clause 5.7.2.6.

The management of the SMF selection enables the PCF to instruct the AMF to contact the PCF during the PDU Session Establishment procedure to perform a DNN replacement, as specified in TS 23.501 [2], clause 5.6.1. To indicate the conditions to check whether to contact the PCF at PDU Session establishment (as specified in clause 6.1.2.5), the PCF provides the Policy Control Request Triggers SMF selection management and, if necessary Change of the Allowed NSSAI, together with SMF selection management related policy control information (see clause 6.5) during UE Registration procedure and at establishment of the AM Policy Association.

The PCF may update SMF selection management information based on PCF local decision or upon being informed about a new Allowed NSSAI. 

The AMF applies the updated SMF selection management information to new PDUSessions only, i.e. already established PDU Sessions are not affected.


The above explanation discusses the Access and Mobility Policy Control (AMPC), focusing on how the Policy Control Function (PCF) interacts with other network elements like AMF (Access and Mobility Management Function), UDM (Unified Data Management), and RAN (Radio Access Network) to manage the following aspects:
1. Service Area Restrictions
This ensures that a user's device (UE) can access only specific geographical areas in the network based on policies.
Steps in Call Flow:
During registration or AMF change, AMF gets service area restrictions (allowed/non-allowed Tracking Area Identifiers or TAIs) from UDM.
AMF sends this restriction information to the PCF.
PCF modifies the restrictions based on operator-defined policies (e.g., UE location, time of day, etc.).
PCF sends the modified restrictions back to AMF.
AMF enforces these restrictions on the UE when determining mobility rules.
UE implements these restrictions when idle or in RRC Inactive state; RAN/AMF enforces them when the UE is active.

2. Radio Frequency Selection Priority (RFSP) Index
This manages network resources for the UE, ensuring efficient radio resource allocation.
Steps in Call Flow:
During registration or AMF change, AMF gets the RFSP Index from UDM.
AMF shares this RFSP Index with the PCF.
PCF adjusts the RFSP Index based on policies (e.g., accumulated usage, load levels, etc.).
PCF sends the modified RFSP Index back to AMF.
AMF ensures the RAN enforces the updated RFSP Index for radio resource management.

3. UE-AMBR (Aggregate Maximum Bit Rate)
This manages the total bandwidth available to a UE for all its active data sessions.
Steps in Call Flow:
During registration, AMF receives UE-AMBR from UDM.
AMF shares this information with the PCF.
PCF may modify the UE-AMBR based on policies defined by the serving network.
PCF sends the updated UE-AMBR to AMF.
AMF provides this value to the RAN for enforcing bandwidth restrictions on the UE.

4. SMF (Session Management Function) Selection
This enables the PCF to guide the AMF in selecting or modifying the SMF during a PDU (Packet Data Unit) session establishment.
Steps in Call Flow:
During registration or PDU session establishment, AMF checks if it needs to contact the PCF for SMF selection.
PCF evaluates the policies and may suggest a DNN replacement (Data Network Name, e.g., internet, IMS) or provide SMF selection information.
AMF applies this guidance only to new PDU sessions (existing sessions are unaffected).

Simplified Call Flow Diagram
Service Area Restrictions


UE → AMF (Registration) → UDM (Fetch Restrictions) → AMF → PCF (Policies Applied) → AMF (Modified Restrictions) → UE/RAN (Enforcement).
RFSP Index Management


UE → AMF (Registration) → UDM (Fetch RFSP Index) → AMF → PCF (Policy Adjustments) → AMF → RAN (Enforcement).
UE-AMBR Management


UE → AMF (Registration) → UDM (Fetch UE-AMBR) → AMF → PCF (Policy Adjustments) → AMF → RAN (Enforcement).
SMF Selection


UE → AMF (PDU Session Establishment) → PCF (Policies Applied) → AMF (DNN/SMF Guidance) → UE.

Key Notes:
PCF's Role: PCF acts as a policy engine that modifies parameters based on operator-defined rules and sends them back to the AMF for enforcement.
AMF's Role: AMF communicates with both the UE and the PCF and acts as the enforcement point for policies (e.g., mobility restrictions, resource management).
Enforcement Points:
Service Area Restrictions: Enforced by UE, RAN, and AMF.
RFSP Index: Enforced by RAN.
UE-AMBR: Enforced by RAN.
SMF Selection: Enforced during session establishment.

Here's a graphical representation of the Access and Mobility Policy Control Call Flow:

UE initiates the request (e.g., Registration).
AMF communicates with UDM to fetch subscription details.
UDM provides data (e.g., restrictions, RFSP, AMBR).
AMF sends the subscription info to the PCF.
PCF applies operator policies and sends adjustments to the AMF.
AMF forwards enforcement data to the RAN.
RAN implements policies and enforces them for the UE.

