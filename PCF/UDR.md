https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.j39sifyyna4y





ETSI TS 129 504

Data present in the data base - ETSI TS 129 519

In 5G architecture, the UDM (Unified Data Management) is a central component responsible for managing subscriber data and providing support for network functions such as authentication, authorization, and mobility management. The UDM is defined by 3GPP standards, and its database holds various types of subscriber data. Here's an overview of the data it manages:

Data in UDM (Unified Data Management):
Subscriber Profile Data:


Subscription data, such as the Subscriber Permanent Identifier (SUPI).
Authentication credentials, including the Authentication Vector (AV).
Service subscription information (e.g., subscribed data plans, QoS profiles, and allowed network slices).
Access and Mobility Data (AMData):


Access control information, such as roaming restrictions, allowed RAT (Radio Access Technology), and PLMN (Public Land Mobile Network).
Parameters needed for mobility management (e.g., location management and tracking area lists).
User location information (e.g., Tracking Area List allowed for the user).
Policy-related Data:


Policy control parameters used by the PCF (Policy Control Function).
Slice-specific data (e.g., Network Slice Selection Assistance Information - NSSAI).
Service Authorization Data:


Rules and policies related to which services a subscriber can access.
Allowed application functions (e.g., application function identifiers).
Session Management Data:


Session data linked to PDU (Packet Data Unit) sessions.
QoS (Quality of Service) rules and default QoS profiles.

Inner Data in AMData (Access and Mobility Data):
The AMData within the UDM contains the following critical components:
Roaming and Access Permissions:
Allowed roaming policies and restrictions (e.g., whether the subscriber can roam internationally or domestically).
Access control data (e.g., allowed RAT types such as 5G, 4G, or 3G).
Tracking Area List (TAL):
The set of allowed Tracking Areas for the user, ensuring the network knows where to locate the user efficiently.
Subscription-Based Data:
IMSI (International Mobile Subscriber Identity) or SUPI (Subscriber Permanent Identifier).
PLMN list (allowed or forbidden Public Land Mobile Networks).
Connection Details:
Parameters related to mobility state and restrictions.
Information for serving cells or access network settings.
Emergency Services Data:
Special permissions for emergency services or Public Safety communications.
3GPP Specification Document References:
To get detailed insights into the UDM and its database, refer to the following 3GPP Technical Specifications:
3GPP TS 23.501 - System Architecture for the 5G System (5GS):


Covers the high-level architecture and roles of UDM and other 5G Core Network (5GC) components.
3GPP TS 23.502 - Procedures for the 5G System (5GS):


Includes detailed signaling flows and interactions involving the UDM, AMF, and other network functions.
3GPP TS 29.503 - Common Data Types and Characteristics for UDM Service Interfaces:


Defines the data structures and protocols used to access and manage UDM data, including AMData.
3GPP TS 29.504 - UDM/NSSF Service-Based Interfaces:


Provides details on the APIs used to access UDM data, including AMData-related APIs.
3GPP TS 29.571 - Data Repository Function (UDR):


Defines the structure of the UDR, which works closely with UDM for storing subscriber data.
AMDATA 
The AMF (Access and Mobility Management Function) interacts with the UDM (Unified Data Management) to fetch AMData (Access and Mobility Subscription Data) as part of the authentication, registration, and mobility management procedures in a 5G network. This process is critical for ensuring that a subscriber is authorized to access the network and for managing their mobility.
Here’s a detailed explanation of the interaction and the key steps involved:
1. High-Level Overview
The AMF communicates with the UDM using the Service-Based Interface (SBI) over the Nausf or Nudm interface.
AMData is part of the subscription data stored in the UDM and includes:
Roaming permissions.
Tracking Area List (TAL).
Access control parameters (e.g., RAT restrictions, PLMN lists).
Allowed NSSAI (Network Slice Selection Assistance Information).
2. Interaction Steps
Step 1: Initial UE Registration
UE sends a Registration Request to the AMF:
The User Equipment (UE) sends a Registration Request message to the AMF, including the SUPI (Subscriber Permanent Identifier) or the encrypted form, SUCI (Subscription Concealed Identifier).
The AMF extracts the SUPI/SUCI to identify the subscriber.
Step 2: AMF Contacts the UDM
AMF resolves the UDM:


The AMF determines the UDM responsible for the subscriber, typically based on the SUPI (e.g., via a routing mechanism using the Home Network Identifier).
AMF sends a Nudm_UEContextManagement Request to UDM:


The AMF calls the Nudm_SubscriberDataRetrieval API provided by the UDM to fetch subscription data, including AMData.
This request contains:
SUPI: To identify the subscriber.
Requested data types (e.g., AMData, SMData, Policy Data).
Step 3: UDM Retrieves Subscriber Data
UDM fetches data from the UDR (Unified Data Repository):


The UDM queries the UDR (if implemented) to retrieve the requested subscriber data.
AMData includes:
Allowed RAT/PLMN access.
Roaming restrictions (e.g., barred services).
Tracking Area List (TAL).
Emergency services permissions.
UDM responds to the AMF:


The UDM sends the requested data (AMData) to the AMF via the Nudm_UEContextManagement response message.
Step 4: AMF Processes the Data
AMF uses AMData:
The AMF processes the received AMData to:
Validate the subscriber's access to the network.
Restrict access to certain PLMNs or RATs if required.
Generate a Tracking Area List for mobility management.

3. API and Message Details
Nudm_SubscriberDataManagement API:


Used by AMF to fetch subscriber-specific data from UDM.
Key endpoints for AMData retrieval:
GET /nudm-sdm/v1/{supi}/am-data: Retrieves Access and Mobility Subscription Data.
Request includes:
supi: The subscriber identifier.
Response includes:
amData: Contains access control information, TAL, roaming restrictions, etc.
Sample Response from UDM:

 {
  "amData": {
    "gpsis": ["msisdn-123456789"],
    "subscribedUeAmbr": {
      "uplink": "100 Mbps",
      "downlink": "1 Gbps"
    },
    "nssai": {
      "defaultSingleNssais": [
        { "sst": 1, "sd": "010203" }
      ]
    },
    "restrictedRatList": ["NR"],
    "restrictedAreaList": ["Area1", "Area2"],
    "plmnAmDataList": [
      {
        "plmnId": { "mcc": "310", "mnc": "410" },
        "servingAreaRestriction": { "trackingAreaList": ["1234", "5678"] }
      }
    ]
  }
}

4. Key 3GPP Specifications
3GPP TS 29.503: Defines the Nudm_SubscriberDataManagement API and the structure of AMData.
3GPP TS 23.501: Explains the role of AMF and UDM in 5G system architecture.
3GPP TS 23.502: Covers the detailed procedures for UE Registration, Mobility Management, and AMF-UDM interactions.
5. Sequence Diagram
Below is the simplified message flow:
[UE] ---- Registration Request ----> [AMF]
          SUPI/SUCI Extraction
               |
               |---- Nudm_SubscriberDataRetrieval ---->
               |      (SUPI, requested data type)
               |
       <---- Nudm_SubscriberDataRetrieval_Response ----
               (AMData)
               |
[AMF] Processes AMData for Access Control, Mobility Mgmt.

Would you like further elaboration on the specific fields in AMData, or how the AMF uses this data during the registration or mobility process?



SUBSCRIBER PROFILE DATA

The Subscriber Profile Data is a core part of the UDM's (Unified Data Management) responsibilities in 5G architecture. It includes all the essential details about a subscriber, such as their identifiers, authentication information, and service subscription details. This data is typically stored in the UDR (Unified Data Repository), which serves as a backend database for UDM.
Here’s how Subscriber Profile Data is stored, structured, and managed in the UDR:
1. Key Elements of Subscriber Profile Data
Subscriber Profile Data includes:
Identifiers:


SUPI (Subscriber Permanent Identifier): Unique identifier for the subscriber, derived from the IMSI (International Mobile Subscriber Identity) in most cases.
SUCI (Subscription Concealed Identifier): Encrypted form of the SUPI used for privacy during signaling.
MSISDN (Mobile Subscriber Integrated Services Digital Network Number): The phone number associated with the subscriber.
Authentication Data:


K (Permanent Key): A unique secret key for the subscriber used for authentication.
OPC (Operator-Provisioned Key): Derived key used for AKA (Authentication and Key Agreement) procedures.
Algorithm Information: E.g., 5G-AKA or EAP-AKA.
Subscription Data:


Allowed PLMNs (Public Land Mobile Networks).
Allowed RATs (Radio Access Technologies) such as 5G, LTE, or NR.
Roaming restrictions or permissions.
Service Profiles:


Quality of Service (QoS) parameters.
NSSAI (Network Slice Selection Assistance Information): Defines the network slices the subscriber can access.
Subscribed AMBR (Aggregate Maximum Bit Rate) for uplink and downlink.
Policy Control Information:


Information related to data usage caps, charging profiles, and service restrictions.

2. Storage Structure in the UDR
The Subscriber Profile Data can be stored in a document-oriented database (e.g., MongoDB) or a relational database (e.g., MySQL, PostgreSQL), depending on the implementation. Here’s how the data might be structured:
Document-Oriented Storage Example (e.g., MongoDB):
In a NoSQL database, data is stored in JSON-like documents. Each document corresponds to a subscriber and contains their profile data.
{
  "subscriberId": "imsi-310410123456789",    // Unique identifier (SUPI)
  "authenticationData": {
    "permanentKey": "8A9D61...",            // Authentication key (K)
    "opc": "00112233445566778899AABBCCDDEE", // Operator-provided key
    "authMethod": "5G_AKA"                  // Authentication method
  },
  "accessAndMobilitySubscriptionData": {
    "gpsis": ["tel:+123456789"],            // GPSI (e.g., MSISDN)
    "allowedPlmnList": [
      { "mcc": "310", "mnc": "410" },       // Allowed PLMNs
      { "mcc": "311", "mnc": "480" }
    ],
    "trackingAreaList": ["1234", "5678"],   // Allowed tracking areas
    "subscribedUeAmbr": {                  // Subscribed AMBR
      "uplink": "100 Mbps",
      "downlink": "1 Gbps"
    }
  },
  "sessionManagementSubscriptionData": {
    "dnnConfigurations": {
      "internet": {
        "sessionAmbr": {
          "uplink": "500 Mbps",
          "downlink": "1 Gbps"
        },
        "pduSessionTypes": {
          "defaultSessionType": "IPv4",
          "allowedSessionTypes": ["IPv4", "IPv6"]
        }
      }
    }
  },
  "policyControlData": {
    "defaultPolicy": "basic_plan",         // Default policy
    "dataUsageCap": "50GB"                // Data cap for the subscriber
  }
}

Relational Database Example (e.g., MySQL/PostgreSQL):
In a relational database, subscriber profile data is normalized into multiple tables for efficient storage and retrieval.
Table: SubscriberInfo
Column
Data Type
Description
subscriber_id
VARCHAR(50)
SUPI or IMSI
msisdn
VARCHAR(15)
Phone number
auth_method
VARCHAR(20)
Authentication method (e.g., 5G-AKA)
permanent_key
VARCHAR(64)
Authentication key (encrypted)
opc
VARCHAR(64)
Operator-provided key

Table: SubscriptionData
Column
Data Type
Description
subscriber_id
VARCHAR(50)
Foreign key to SubscriberInfo
plmn_id
VARCHAR(10)
Allowed PLMN (e.g., MCC+MNC)
tracking_areas
TEXT
JSON array of tracking areas
ambr_uplink
VARCHAR(20)
Subscribed uplink AMBR (e.g., "100 Mbps")
ambr_downlink
VARCHAR(20)
Subscribed downlink AMBR (e.g., "1 Gbps")

Table: PolicyControlData
Column
Data Type
Description
subscriber_id
VARCHAR(50)
Foreign key to SubscriberInfo
policy_id
VARCHAR(20)
Policy name (e.g., "basic_plan")
data_usage_cap
VARCHAR(10)
Data usage cap (e.g., "50GB")


3. Storage Workflow
When storing Subscriber Profile Data in the UDR, the following workflow is followed:
Data Ingestion:


Data is provisioned by the operator’s OSS/BSS (Operations Support System / Business Support System).
UDM interacts with UDR via the Nudr interface using standardized APIs.
APIs for Storing Data:


PUT /nudr-dr/v1/subscription-data/{supi}/am-data:
Stores or updates Access and Mobility Subscription Data.
PUT /nudr-dr/v1/subscription-data/{supi}/sm-data:
Stores or updates Session Management Data.
PUT /nudr-dr/v1/subscription-data/{supi}/auth-data:
Stores or updates Authentication Data.
Data Validation:


The UDM validates incoming data for correctness (e.g., SUPI format, PLMN codes) before storing it in the UDR.
Data Security:


Sensitive fields like permanentKey and opc are encrypted before storage.
Access to the data is protected using role-based access control (RBAC).

4. Security and Integrity
Encryption:
Keys and sensitive subscriber data are encrypted (e.g., using AES-256).
Access Control:
Only authorized network functions like AMF, SMF, and PCF can query the subscriber data via the UDM APIs.
Auditing:
All data changes are logged to ensure traceability.




OSS/BSS(OPERATION SUPPORT/BUSSINESS SUPPORT SYSTEM)

Yes, data provisioning in the 5G ecosystem is primarily managed through the OSS/BSS (Operations Support System/Business Support System). Here's a detailed explanation of how the OSS/BSS provisions subscriber data in the UDM and UDR:

1. Role of OSS/BSS in Data Provisioning
OSS (Operations Support System):


Manages the technical and operational aspects of the network.
Responsible for configuring the subscriber's technical data, such as:
Allowed PLMNs.
Tracking areas.
Quality of Service (QoS) profiles.
Network slice configurations (NSSAI).
BSS (Business Support System):


Manages the business-related aspects of the subscriber, such as:
Billing and charging profiles.
Subscription plans and data caps.
Roaming agreements and restrictions.
Policy rules for service usage.
Both systems communicate with the UDM (and indirectly with the UDR) to store subscriber-related data in a centralized repository.

2. Workflow of Data Provisioning
Step 1: Subscriber Registration
When a new subscriber is added to the network:
The BSS system registers the subscriber and assigns them a SUPI (Subscriber Permanent Identifier), typically derived from their IMSI.
Other details like MSISDN (phone number), subscription plan, and roaming permissions are also provisioned.
Step 2: Subscriber Profile Creation
The OSS configures the technical profile of the subscriber:
Allowed RAT types (e.g., 5G, LTE).
PLMNs where the subscriber is permitted to operate.
Tracking areas (TAL) and location-related data.
QoS parameters like uplink and downlink AMBR (Aggregate Maximum Bit Rate).
Step 3: Authentication Data Provisioning
The BSS provisions the subscriber's authentication data into the UDR:
Permanent key (K).
Operator-provided key (OPC).
Selected authentication method (e.g., 5G-AKA or EAP-AKA).
Step 4: Policy and Charging Configuration
The BSS provisions policy and charging data, such as:
Allowed network slices (NSSAI).
Policy rules for data usage, throttling, or service restrictions.
Billing profiles for prepaid or postpaid accounts.
Step 5: Storage in the UDR
The data from OSS/BSS is sent to the UDM via standard provisioning interfaces.
The UDM validates the data and stores it in the UDR, which acts as the centralized database.

3. Provisioning Interfaces and Protocols
Provisioning APIs
Operators use RESTful APIs exposed by the UDM (or UDR) to provision subscriber data. Key interfaces include:
Nudr Interface:


Used by UDM to interact with the UDR.
Example APIs:
PUT /nudr-dr/v1/subscription-data/{supi}/am-data: Store Access and Mobility data.
PUT /nudr-dr/v1/subscription-data/{supi}/sm-data: Store Session Management data.
PUT /nudr-dr/v1/subscription-data/{supi}/auth-data: Store Authentication data.
Service Management Interfaces:


Proprietary APIs provided by the OSS/BSS systems to push data to UDM.
Data Formats
JSON: Commonly used for REST API payloads.
XML: Sometimes used in legacy OSS/BSS systems.
CSV or Bulk Files: For batch provisioning of subscriber data.

4. Example of Data Provisioning via API
Here’s an example of how an API call from OSS/BSS to UDM might look:
Request: Provisioning Access and Mobility Subscription Data
PUT /nudr-dr/v1/subscription-data/imsi-310410123456789/am-data
Content-Type: application/json

{
  "gpsis": ["tel:+123456789"],
  "subscribedUeAmbr": {
    "uplink": "100 Mbps",
    "downlink": "1 Gbps"
  },
  "nssai": {
    "defaultSingleNssais": [
      { "sst": 1, "sd": "010203" }
    ],
    "singleNssais": [
      { "sst": 1, "sd": "010203" },
      { "sst": 2, "sd": "040506" }
    ]
  },
  "restrictedRatList": ["EUTRA"],
  "plmnAmDataList": [
    {
      "plmnId": { "mcc": "310", "mnc": "410" },
      "servingAreaRestriction": {
        "trackingAreaList": ["1234", "5678"]
      }
    }
  ]
}

Response:
200 OK
{
  "status": "success",
  "message": "Subscription data provisioned successfully."
}


5. Database Design for Provisioned Data
The UDR stores subscriber data in a structured format. Below is an example schema:
Field
Description
SUPI
Unique identifier (e.g., imsi-310410123456789).
GPSIs
Global Public Subscriber Identities (e.g., MSISDN).
AMBR
Uplink and downlink bit rates (e.g., 100 Mbps).
NSSAI
Network Slice information (e.g., SST: 1, SD: 010203).
PLMN IDs
Allowed PLMNs (e.g., MCC: 310, MNC: 410).
Authentication Key (K)
Permanent key used for authentication.
QoS Profiles
Quality of Service details (e.g., latency, throughput).
Policy Rules
Usage policies, restrictions, and billing profiles.


6. Bulk Provisioning
For bulk subscriber provisioning, OSS/BSS systems may:
Export subscriber data in a file format (e.g., CSV, JSON).
Use batch APIs to provision multiple subscribers at once:
Example API: POST /nudr-dr/v1/subscription-data/batch.
Automate data validation and synchronization with the UDM and UDR.

7. Security in Data Provisioning
Encryption: Subscriber data is encrypted during transmission (e.g., using TLS).
Authentication: OSS/BSS systems authenticate with the UDM using OAuth2 tokens or similar mechanisms.
Access Control: Role-based access control (RBAC) ensures only authorized entities can provision data.

Would you like an example of bulk provisioning or further details about how OSS/BSS integrates with the 5G core network?


POLICY RULES EXPLANATION

Policy Rules in 5G networks define how a subscriber's usage is controlled, restricted, and billed. These rules are crucial for enabling Quality of Service (QoS), enforcing service-level agreements (SLAs), and ensuring fair resource usage. They are primarily managed by the Policy Control Function (PCF) and stored in the UDR (Unified Data Repository) as part of the subscriber's profile. The data is provisioned by the operator’s OSS/BSS systems.
Here's a detailed breakdown of Policy Rules, their components, and how they work:

1. Components of Policy Rules
Policy Rules include the following key elements:
a) Data Usage Policies
Define data limits, such as:
Data caps (e.g., 10GB/month).
Speed throttling after data cap exhaustion (e.g., 64 kbps).
Usage rules for specific applications or services (e.g., video streaming, VoIP).
b) QoS (Quality of Service) Policies
Define the subscriber's QoS profile, including:
Latency: Required delay tolerance (e.g., ultra-low for gaming).
Throughput: Guaranteed uplink and downlink speeds.
Packet loss tolerance.
Example: A subscriber may have priority for mission-critical applications.
c) Charging Profiles
Define how a subscriber is billed for their usage:
Prepaid or Postpaid billing mechanisms.
Rates for different services (e.g., standard rates for browsing vs. higher rates for streaming).
Time-based billing (e.g., cheaper at night).
d) Access and Service Restrictions
Restrictions on accessing certain services:
Blocking specific websites or applications.
Allowing roaming only in specified regions or networks.
Enforcing parental controls or corporate usage policies.
e) Network Slice Policies
Define the allowed Network Slices (NSSAI) a subscriber can access:
A subscriber may have access to an enhanced Mobile Broadband (eMBB) slice but not a massive IoT (mIoT) slice.
Each slice can have its own QoS and policy settings.
f) Time-Based Rules
Policies that apply during specific times:
Peak vs. off-peak hours.
Time-limited promotions (e.g., unlimited data for one week).
g) Application-Specific Policies
Policies for specific services, such as:
Prioritizing VoLTE/VoNR (Voice over LTE/5G NR) traffic.
Throttling video streaming to standard definition (SD).

2. Policy Storage in the UDR
The Policy Control Rules for each subscriber are stored in the UDR. The data structure can look like the following:
Example JSON Representation
{
  "policyId": "basic_plan",
  "subscriberId": "imsi-310410123456789",
  "dataUsagePolicy": {
    "monthlyDataCap": "50GB",
    "throttleAfterCap": "128 kbps",
    "unlimitedApplications": ["whatsapp", "youtube"]
  },
  "qosPolicy": {
    "priorityLevel": 8,
    "maxUplink": "10 Mbps",
    "maxDownlink": "50 Mbps",
    "latency": "50ms"
  },
  "chargingPolicy": {
    "chargingMethod": "prepaid",
    "rate": {
      "defaultRate": "₹0.10/MB",
      "streamingRate": "₹0.20/MB"
    }
  },
  "serviceAccessRestrictions": {
    "restrictedApplications": ["torrent"],
    "allowedRoamingAreas": [
      { "mcc": "310", "mnc": "410" },
      { "mcc": "311", "mnc": "480" }
    ]
  },
  "timeBasedPolicy": {
    "offPeakHours": {
      "start": "00:00",
      "end": "06:00",
      "rateDiscount": "50%"
    }
  },
  "networkSlicePolicy": {
    "defaultSlice": { "sst": 1, "sd": "010203" },
    "allowedSlices": [
      { "sst": 1, "sd": "010203" },
      { "sst": 2, "sd": "040506" }
    ]
  }
}


3. Workflow: How Policy Rules Are Applied
Step 1: Policy Configuration by OSS/BSS
Operators configure policy rules for a subscriber in the OSS/BSS.
The policy is provisioned into the UDR via the UDM using standard Nudr APIs.
Step 2: Policy Query by PCF
The PCF (Policy Control Function) retrieves the subscriber's policies from the UDR when needed.
Example API: GET /nudr-dr/v1/policy-data/subscriber/{supi}.
Step 3: Policy Enforcement
Policies are enforced in real-time by various 5G core functions:
SMF (Session Management Function): Enforces QoS and data usage policies for PDU sessions.
UPF (User Plane Function): Implements packet filtering, throttling, and charging.
AMF (Access and Mobility Function): Ensures access and service restrictions.
Step 4: Real-Time Adjustments
Policies can be dynamically modified by the operator via the OSS/BSS or triggered by network events (e.g., reaching a data cap).
The PCF updates the SMF and UPF with new rules in real-time.

4. Example Scenarios
Scenario 1: Data Throttling After Cap
A subscriber with a 10GB monthly data cap consumes 10GB.
The PCF notifies the SMF to reduce the subscriber’s speed to 128 kbps.
The UPF enforces the throttling for all data sessions.
Scenario 2: Roaming Restrictions
A subscriber tries to attach to a network in an unallowed roaming area.
The AMF queries the PCF for roaming policies.
The PCF denies the attach request based on the roaming restrictions in the policy.
Scenario 3: Charging for Video Streaming
A subscriber starts streaming a video on YouTube.
The SMF identifies the traffic and notifies the PCF.
The PCF applies a special charging rate (e.g., ₹0.20/MB for video).
The UPF collects the usage and reports it to the CHF (Charging Function).

5. Provisioning Policy Rules
Using Nudr API to Create Policy Rules
PUT /nudr-dr/v1/policy-data/subscriber/imsi-310410123456789
Content-Type: application/json

{
  "policyId": "basic_plan",
  "dataUsagePolicy": {
    "monthlyDataCap": "10GB",
    "throttleAfterCap": "128 kbps"
  },
  "qosPolicy": {
    "priorityLevel": 5,
    "maxUplink": "5 Mbps",
    "maxDownlink": "20 Mbps"
  },
  "chargingPolicy": {
    "chargingMethod": "postpaid",
    "rate": {
      "defaultRate": "₹0.05/MB",
      "streamingRate": "₹0.10/MB"
    }
  }
}


6. Security and Integrity
Encryption: Policy data is encrypted during transmission (e.g., TLS) and at rest.
Access Control: Only authorized OSS/BSS systems and network functions (e.g., PCF) can modify or access policies.
Auditing: All policy changes are logged to ensure accountability.


Subscriber data stored in the udr 

In a 5G network, the Subscriber Database is primarily managed by the UDM (Unified Data Management). It stores critical information required for identifying, authenticating, and managing subscribers. Below is a detailed breakdown of the types of data stored in the Subscriber Database (UDM):

1. Subscriber Identity Information
These are identifiers used to uniquely recognize a subscriber in the network:
IMSI (International Mobile Subscriber Identity):
A globally unique identifier for the subscriber.
Used for authentication and network access.
SUPI (Subscription Permanent Identifier):
A unique identifier in 5G that replaces IMSI for enhanced security.
Typically stored in the format: IMSI-based or Network-Specific Identifier.
GUTI (Globally Unique Temporary Identifier):
A temporary identifier allocated by the AMF to avoid transmitting IMSI/SUPI over the air.

2. Authentication Information
Used to securely authenticate the subscriber with the network:
Authentication Vectors (AVs):
Generated by the UDM for use in the 5G-AKA (Authentication and Key Agreement) protocol.
Includes:
RAND (Random Number): For challenge-response authentication.
AUTN (Authentication Token): To verify the network.
CK/IK (Cipher and Integrity Keys): Used for encryption and integrity protection.
Subscription Key (K):
A secret key stored on the SIM and UDM for authentication.
Used for generating authentication vectors.

3. Subscriber Profile and Service Data
Defines the services and capabilities available to the subscriber:
Subscribed Services:
Voice (e.g., VoNR - Voice over New Radio).
Data (internet access).
SMS/MMS.
Value-added services (e.g., roaming, content delivery).
QoS (Quality of Service) Profile:
Specifies the subscriber's data priority and speed, including:
5QI (5G QoS Identifier): Defines QoS for specific flows.
ARP (Allocation and Retention Priority).
GBR (Guaranteed Bit Rate) or Non-GBR.
AMBR (Aggregate Maximum Bit Rate):
Maximum data rate allocated for uplink and downlink.
Data Network (DN) Access Permissions:
Defines which data networks (e.g., internet, private networks) the subscriber can access.

4. Mobility and Roaming Data
Supports subscriber mobility and roaming across networks:
PLMN (Public Land Mobile Network) List:
Home PLMN.
Preferred and forbidden PLMNs for roaming.
Roaming Restrictions:
Allowed or restricted countries/operators.
Current Location Information:
Temporarily stored when needed for handover, registration, or paging.
Includes TAC (Tracking Area Code) and Cell ID.

5. Subscriber Policies
Policies that govern the subscriber's network behavior:
Policy Control Rules:
Stored in PCF (Policy Control Function), linked with UDM.
Define rules for data usage, QoS, throttling, etc.
Charging Policies:
Prepaid or postpaid account type.
Charging thresholds for data, calls, or messages.

6. IP and PDU Session Information
Used to support data services for the subscriber:
PDU Session Configuration:
Associated IP address(es) for active sessions.
Data networks (DNN - Data Network Name, similar to APN in 4G).
Session QoS parameters.
Session State:
Active or inactive sessions.

7. Security Data
Enhances secure communication in the 5G network:
Encryption Algorithms:
Supported encryption and integrity protection algorithms for NAS and AS layers.
E.g., AES, SNOW 3G, or ZUC.
Key Hierarchies:
NAS Keys (KNAS), AS Keys (KAS), and RRC keys derived during authentication.

8. Device Information
Details about the subscriber's device:
IMEI (International Mobile Equipment Identity):
A unique identifier for the device.
Device Capabilities:
Radio capabilities (e.g., support for specific frequency bands).
5G features (e.g., NSA or SA mode).

9. Billing and Usage Data
Stored for account management and charging:
Credit Balance (for prepaid subscribers).
Usage Counters:
Data consumed (MB/GB).
Minutes used for voice calls.
Number of SMS sent.
CDRs (Call Detail Records):
Temporarily stored before being sent to the Billing System.

10. Temporary Session Data
Data temporarily stored for active sessions:
NAS Context:
Temporary security and mobility-related data for an ongoing session.
Tracking Area Updates (TAU):
Tracks the subscriber's location for efficient paging.

Summary of UDM Role in Subscriber Data Management
The UDM plays a central role in:
Authenticating subscribers (with authentication data like SUPI and keys).
Providing subscriber profiles and QoS policies to other network functions like AMF, SMF, and PCF.
Supporting mobility and session management by providing real-time updates on subscriber state and permissions.
If you need further details on any specific aspect, such as how this data flows between network functions or detailed protocols, let me know!


In a real-time 5G network scenario, the Presence Reporting Area (PRA) is used to monitor whether a User Equipment (UE) is within a specific geographical or logical area, defined by cells, tracking areas, or a combination. This is especially relevant for applications such as location-based services, IoT device monitoring, and optimized paging for specific regions.
Here's an example to clarify the praInfos map (PresenceInfo) and how it is used in practice.

Scenario: Smart Warehouse Monitoring
Use Case:
A logistics company manages a smart warehouse with several sections, each equipped with IoT devices that communicate with the 5G network. 
They use Presence Reporting Areas (PRAs) to monitor the presence of devices in specific zones of the warehouse for operational purposes, such as asset tracking and inventory management.
The warehouse is divided into sections:
Zone A: Goods receiving area.
Zone B: Storage shelves.
Zone C: Packing area.
Zone D: Dispatch area.
The company uses PRA Identifiers (praId) to represent these zones:
praId_1: Zone A.
praId_2: Zone B.
praId_3: Zone C.
praId_4: Zone D.
Each PRA contains a set of cells (gNBs) that correspond to the geographical coverage of the zone. For example:
praId_1 (Zone A) = Cell 101, Cell 102.
praId_2 (Zone B) = Cell 103, Cell 104, Cell 105.

Real-Time Example of praInfos Map:
Suppose a smart pallet (IoT device with a 5G module) enters the warehouse. The 5G core network tracks its presence within the PRAs using the praInfos map.
Here’s how the praInfos data structure could look for this scenario:
{
  "praInfos": {
    "praId_1": {
      "presenceReportingArea": ["Cell_101", "Cell_102"],
      "isDedicated": true,
      "ueList": ["UE_1234"]
    },
    "praId_2": {
      "presenceReportingArea": ["Cell_103", "Cell_104", "Cell_105"],
      "isDedicated": false,
      "ueList": []
    },
    "praId_3": {
      "presenceReportingArea": ["Cell_106", "Cell_107"],
      "isDedicated": true,
      "ueList": []
    }
  }
}


Explanation of the Fields:
Key (praId):


Each entry in the praInfos map represents a unique Presence Reporting Area Identifier (e.g., praId_1, praId_2).
Presence Reporting Area Composition (presenceReportingArea):


Lists the cells or regions (e.g., Cell_101, Cell_102) that define the coverage area for that PRA.
Dedicated vs. Shared PRAs (isDedicated):


Dedicated: PRAs assigned for specific UEs (e.g., a device for Zone A).
Shared: PRAs available for multiple UEs.
UE List (ueList):


Identifies the UEs currently within the PRA.
Example: UE_1234 (the smart pallet) is currently in Zone A.

Network Operations in Real-Time:
PRA Update (UE Movement):


When the smart pallet (UE_1234) moves from Zone A to Zone B, the UE triggers a Location Reporting message.
The 5G core network updates the praInfos map:
Removes UE_1234 from praId_1 (Zone A).
Adds UE_1234 to praId_2 (Zone B).
Updated praInfos:

 {
  "praInfos": {
    "praId_1": {
      "presenceReportingArea": ["Cell_101", "Cell_102"],
      "isDedicated": true,
      "ueList": []
    },
    "praId_2": {
      "presenceReportingArea": ["Cell_103", "Cell_104", "Cell_105"],
      "isDedicated": false,
      "ueList": ["UE_1234"]
    },
    "praId_3": {
      "presenceReportingArea": ["Cell_106", "Cell_107"],
      "isDedicated": true,
      "ueList": []
    }
  }
}


Event Notifications:


Based on the praInfos updates, the network can send event notifications to applications subscribed to PRA events (e.g., the warehouse management system).

Practical Applications:
Warehouse Example:


Automated tracking of IoT devices or vehicles as they move between zones.
Optimized inventory and logistics operations.
Smart Cities:


Monitoring public transport vehicles within specific urban zones.
Enhancing public safety by tracking emergency vehicles.
Industrial IoT:


Supervising machines or robots operating within a factory’s defined areas.

Conclusion:
The praInfos map allows the 5G network to dynamically track the presence of devices in specific areas by using Presence Reporting Areas (PRAs). This feature is highly useful for real-time monitoring in IoT, logistics, and mobility scenarios. Let me know if you'd like a call flow or further technical details!

