https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.398ymelqnw0j



Distribution of Policies to UE

The PCF provides UE access selection and PDU Session-related policy information to the UE.
These policies include either Access Network Discovery & Selection Policy (ANDSP), UE Route Selection Policy (URSP), or both, using Npcf and Namf service operations.
The PCF is triggered to provide this information during UE Policy Association Establishment and UE Policy Association Modification procedures, as defined in clauses 4.16.11 and 4.16.12 of TS 23.502.
The PCF can install a PCC Rule and activate the start and stop of application detection in the SMF.
If the same PCF is selected for SM Policy Association Control and UE Policy Association Control, the reporting of start and stop of an application can trigger the installation or update of a URSP rule in the UE.
Operator-defined policies in the PCF may depend on input data such as UE location, time of day, and information provided by other network functions.
The PCF includes the UE access selection and PDU Session-related information in a Policy Section identified by a Policy Section Identifier (PSI).
The PCF ensures that each Policy Section is under a predefined size limit.

Policy Section Details
A list of self-contained UE access selection and PDU Session-related policy information includes:
URSP rules delivered to the UE are provided in the order of precedence, without splitting a URSP rule across Policy Sections.
WLANSP rules are delivered in the order of priority, without splitting a WLANSP rule across Policy Sections.
Non-3GPP access network selection information is provided in one Policy Section.
The PCF decides how to divide the policy information into Policy Sections while meeting the size limit and self-contained content requirements.
The PLMN ID is provided to the UE along with the policies and indicates which PLMN a Policy Section belongs to.

Policy Delivery via AMF
The AMF forwards UE policies transparently to the UE without needing to understand their content.
If the PCF splits policies, it provides multiple Policy Sections separately to the AMF.
The AMF uses the UE Configuration Update procedure to deliver these policies to the UE.

UE Policy Updates
The UE updates stored policies as follows:
If there is no Policy Section with the same PSI, the UE stores the new Policy Section.
If a Policy Section with the same PSI exists, the UE replaces the stored section with the new information.
If the received information contains only the PSI, the UE removes the stored Policy Section.
The UE retains policies even when registering with another PLMN.
For non-HPLMN policies, the number of stored policies is up to the UE implementation.
The UE may remove older stored policies if the maximum number is exceeded.

Policy Synchronization
At Initial Registration or during Registration to 5GS from EPS:
The UE provides a list of stored PSIs for the home and visited PLMNs (if roaming).
If the USIM is changed, no PSIs are provided.
If no policies are stored for the home PLMN, no PSIs are provided for it.
If the UE is roaming and has policies for the home PLMN but none for the visited PLMN, only the PSIs for the home PLMN are included.
The UE may indicate ANDSP support to the PCF.
The PCF takes this into account when determining whether to provide ANDSP rules.
The PCF does not provide ANDSP rules if the UE does not indicate support for ANDSP.
The UE may also provide the OSId.
The UE can trigger an Initial Registration with the stored PSI list to request synchronization, e.g., after powering up without a USIM change.

PSI Handling by the PCF
During Initial Registration, the PCF retrieves the PSI list and content stored in the UDR for the UE's SUPI.
In roaming scenarios, the V-PCF retrieves the PSI list and content from the V-UDR or local configuration.
The PCF compares the PSI list provided by the UE with the list retrieved from the UDR.
The PCF updates the list if there are differences or if updates are required due to operator policies.
The updated list is sent to the AMF, which forwards it to the UE.
The PCF maintains the latest PSI list for each UE until the UE Policy Association termination request is received.
The PCF stores the latest PSI list in the UDR upon termination.

Additional PCF Capabilities
The PCF may use the PEI and OSId to determine the UE operating system.
If the PCF cannot determine the UE operating system but needs to deliver URSP rules containing Application IDs as Traffic Descriptors, it includes multiple Application IDs for supported operating systems.
If the PCF determines the UE operating system, it includes the Application ID specific to that operating system in the URSP rules.
If the PCF does not use the PEI or OSId, it may send URSP rules with traffic descriptors for multiple operating systems.
The above content outlines how the Policy Control Function (PCF) manages and distributes policies to the User Equipment (UE) in the 5G system. It focuses on policies related to Access Network Discovery & Selection Policy (ANDSP) and User Route Selection Policy (URSP). The key points and corresponding call flows are explained in depth below:

Purpose of UE Policies
UE Policies define rules for:


Access network discovery and selection (ANDSP): Determines which access networks the UE can connect to (e.g., 3GPP or non-3GPP networks like Wi-Fi).
User Route Selection Policy (URSP): Determines how traffic is routed to various PDU Sessions based on applications or services.
These policies are delivered from the PCF to the UE via Access and Mobility Management Function (AMF) and NAS (Non-Access Stratum) signaling.



Trigger Events for Policy Delivery
UE Policy Association Establishment:


Triggered when the UE registers with the network.
PCF establishes a policy association with the UE to deliver policies.
UE Policy Association Modification:


Triggered when conditions change (e.g., UE moves to a new location or updates its capabilities).
PCF modifies the policies and updates the UE.

Policy Information Structure
PCF organizes the UE policies into Policy Sections, each identified by a Policy Section Identifier (PSI):
Each Policy Section is self-contained and includes all necessary rules (e.g., URSP rules, ANDSP rules, or WLANSP rules).
Policy Sections are sized to fit within the predefined limits for NAS transport.

UE Policy Management
Key Scenarios:
Storing and Updating Policies in the UE:


The UE stores the received policies using the PSI as a reference.
If the UE receives a new PSI, it stores it.
If the UE receives an updated PSI, it replaces the existing one.
If the PSI is empty, the UE deletes the corresponding stored policies.
Persistent Storage:


The UE retains policies even when switching to a new PLMN.
The number of policies stored for visited PLMNs is implementation-specific, and old policies can be deleted if necessary.
ANDSP Rules for VPLMN:


When the UE roams, ANDSP rules may apply to Visited PLMN (VPLMN) if specified in the policy.
Initial Registration:


The UE sends a list of stored PSIs to the network during registration (if the USIM is not changed).
If no policies are stored, the UE does not provide any PSIs.

Role of AMF
The AMF acts as a transparent relay:
It forwards policies from the PCF to the UE without interpreting their content.
It uses the UE Configuration Update Procedure (specified in TS 23.502) to send the policies.

Call Flow Scenarios
1. Policy Delivery During Initial Registration
Steps:
UE → AMF: The UE sends a Registration Request, optionally including a list of stored PSIs.
AMF → PCF: The AMF forwards the registration request to the Home-PCF (H-PCF).
H-PCF → UDR:
The PCF retrieves stored PSI and associated policies from the Unified Data Repository (UDR).
If the UE is roaming, the Visited-PCF (V-PCF) retrieves policies from the Visited-UDR (V-UDR).
PCF Policy Decision:
The PCF compares the UE's provided PSIs with those in the UDR.
If updates are needed, the PCF prepares updated Policy Sections.
H-PCF → AMF: The updated policies are sent to the AMF.
AMF → UE: The AMF uses the NAS transport to deliver the updated policies to the UE.

2. Policy Update During Roaming
Steps:
UE → V-AMF: The UE provides a list of stored PSIs associated with its Home PLMN (HPLMN) during registration.
V-AMF → V-PCF:
The V-AMF forwards the registration request to the Visited PCF (V-PCF).
The V-PCF retrieves local PSI data from the Visited UDR.
V-PCF → H-PCF:
The V-PCF forwards the home PSIs to the H-PCF.
H-PCF Decision:
The H-PCF updates the home policies based on operator rules, such as location and time.
H-PCF → V-PCF → V-AMF → UE:
The updated policies are sent back to the UE via the AMF.

3. Policy Modification Based on Trigger (e.g., Location Change)
Steps:
UE → AMF: The UE provides location updates to the AMF.
AMF → PCF: The AMF notifies the PCF of the location change.
PCF Decision:
The PCF evaluates the need for policy updates (e.g., ANDSP changes due to location).
PCF → AMF → UE:
The updated policies are delivered to the UE via NAS signaling.

Handling Application Detection and Traffic Rules
When the same PCF manages SM Policy Association and UE Policy Association, application events can trigger updates:
Example: Start/stop of an application may lead to:
Installation or update of a URSP rule.
Selection of a specific PDU Session for application traffic.

Operator-Specific Considerations
Policies are tailored to operator configurations, including:
UE location.
Time of day.
Information from other network functions (e.g., SMF, AMF).
Policies may vary between users or groups of users.

UE Policy Storage and Synchronization
If the UE powers on without a USIM change, it may trigger Initial Registration to synchronize policies.
During synchronization:
The PCF ensures that stored policies align with the operator's latest rules.
Outdated or mismatched policies are replaced.

Special Cases
Operating System Identification:


The PCF can use PEI (Permanent Equipment Identifier) or OSId to determine the UE's OS.
If OS-specific rules (e.g., for application traffic) are required, the PCF includes the appropriate Application ID.
Multi-OS Support:


If the PCF cannot determine the UE's OS, it includes traffic descriptors for multiple operating systems.




Here are the key takeaways and important points from the above data:

1. Purpose of UE Policies
Access Network Discovery & Selection Policy (ANDSP): Guides the UE on how to select the best access network (e.g., 5G, Wi-Fi).
User Route Selection Policy (URSP): Specifies how UE traffic is routed to specific PDU Sessions based on application or service needs.

2. Policy Delivery by PCF
The Policy Control Function (PCF) delivers policies to the UE using Npcf and Namf service operations.
Policies are sent during:
Policy Association Establishment (when UE registers to the network).
Policy Association Modification (when conditions like location or operator rules change).

3. Policy Sections and PSI
Policy Sections: Policies are divided into small, self-contained segments to meet size limits for NAS transport.
Policy Section Identifier (PSI): Each Policy Section has a unique ID for identification.
Policies (URSP, ANDSP) are grouped logically into sections for efficient delivery.
Each section is kept self-contained (e.g., no splitting of rules across sections).

4. Role of AMF
The Access and Mobility Management Function (AMF) is a transparent relay:
It forwards policies from the PCF to the UE without understanding the content.
It uses the UE Configuration Update Procedure to send policies via NAS signaling.

5. How UE Stores and Updates Policies
The UE maintains a local policy database:


Stores New Policies: If a received PSI is new, it is stored.
Updates Existing Policies: If the received PSI matches an existing one, the stored policy is updated.
Deletes Policies: If a PSI is received with no content, the corresponding policy is deleted.
The UE keeps policies even when switching to another Public Land Mobile Network (PLMN). However, the number of stored policies for visited PLMNs is UE-implementation specific.



6. Triggers for Policy Updates
Initial Registration: When the UE registers with the network, it provides a list of stored PSIs to synchronize policies.
Location Change: A change in UE location may trigger updates to policies, such as ANDSP.
Application Events: Starting or stopping an application may update URSP rules to ensure proper traffic routing.

7. Policy Customization Based on Operator Rules
Policies can be tailored based on:
UE Location: Different policies for different regions.
Time of Day: Dynamic policies based on usage patterns.
Other Network Functions: Information from the SMF, AMF, or UDR can influence policy decisions.

8. Roaming Scenarios
When the UE roams:
Visited-PCF (V-PCF) retrieves policies for the visited PLMN.
Home-PCF (H-PCF) updates or synchronizes policies based on operator rules.
The V-PCF forwards any home policies received from the UE to the H-PCF.

9. Application and Traffic Management
If the same PCF manages SM Policy Association and UE Policy Association:
Starting or stopping an application can trigger a URSP rule update.
The URSP ensures application traffic is routed to the appropriate PDU Session.

10. Operating System Identification
The PCF can determine the UE's operating system using:
Permanent Equipment Identifier (PEI).
OS Identifier (OSId).
OS-specific URSP rules can then be applied.

11. AMF Transparency
The AMF does not process the policy content; it simply delivers them to the UE.

12. Synchronization During Power-Up
If the UE powers on (without a USIM change), it may trigger Initial Registration to synchronize its policies with the network.

13. Size Limit for NAS Transport
Policies are divided into sections to comply with size limits specified in TS 29.507. This ensures efficient delivery to the UE.

14. Multi-OS Support
If the PCF cannot identify the UE’s OS, it includes traffic descriptors for multiple OS types to ensure compatibility.

15. Storing Policies in UDR
The PCF stores the latest policies in the Unified Data Repository (UDR) for:
Retrieval during future sessions.
Synchronization during roaming.

These points summarize the critical aspects of how PCF delivers and manages UE policies in a 5G system.
