	https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.vb5fl24watk4#heading=h.oznyqku4dhtw	


ONLINE AND OFFLINE CHARGING IN THE NETWORKING
In 5G networks, charging refers to how network usage is tracked and billed. Charging methods are categorized into online charging and offline charging, depending on when and how charging decisions are made.

1. Online Charging (Real-Time Charging)
Definition: Charging happens in real-time, where the network checks the user's balance or available credit before allowing service usage.
How It Works:
The Charging Function (CHF) evaluates user credit during the session.
Services are granted or denied immediately based on available credit.
The network may terminate or modify services if the credit runs out.
Use Cases:
Prepaid plans: The user is billed in real-time, ensuring they can’t exceed their credit balance.
Services requiring strict usage limits (e.g., data caps, streaming subscriptions).
Key Features:
Real-time interaction between SMF and CHF.
Immediate enforcement of credit limits.
Events like credit exhaustion trigger actions (e.g., stopping the service or asking for more credit).

2. Offline Charging (Postpaid Billing)
Definition: Charging happens after the service is used, based on usage records collected by the network.
How It Works:
The network tracks user activity in Charging Data Records (CDRs).
The CDRs are periodically sent to a Billing System (usually via the CHF or directly to the backend).
Users are billed later based on their total usage.
Use Cases:
Postpaid plans: The user is billed at the end of a billing cycle (e.g., monthly bills).
Non-critical services like broadband or enterprise connections.
Key Features:
No real-time interaction with CHF during usage.
Service is not interrupted even if the user exceeds a predefined limit.
Focuses on accurate record-keeping and delayed billing.

Comparison Between Online and Offline Charging
Feature
Online Charging
Offline Charging
Timing
Real-time
After service usage
Use Case
Prepaid plans, strict control
Postpaid plans, delayed billing
Service Control
Enforces credit limits immediately
No real-time control of limits
Interaction with CHF
Real-time interaction for decisions
Usage tracked for later billing
Actions on Exhaustion
Service stopped/modified immediately
No immediate action


Real-Life Examples
Online Charging:
You purchase a prepaid mobile data pack. If your data runs out, the service stops immediately until you recharge.
A video streaming subscription that denies access if your credit expires.
Offline Charging:
You use postpaid mobile data and receive a bill at the end of the month, regardless of how much you use during the month.

Relevance in 5G
5G networks use Online Charging for advanced services that need real-time authorization and enforcement (e.g., IoT applications, gaming, and streaming).
Offline Charging is used for traditional services like voice calls or broadband, where immediate credit control isn’t necessary.
Let me know if you need more details about their architecture or call flows!

			POSTPAID AND PREPAID
In a 5G network, the component responsible for determining whether a user is postpaid or prepaid is the Policy Control Function (PCF) in coordination with the Charging Function (CHF). Here's a detailed explanation:

Key Components Involved
Unified Data Management (UDM):


The UDM stores the subscriber's profile, which includes details about whether the user is on a postpaid or prepaid plan.
During the session establishment (e.g., when the user connects to the 5G network), the UDM provides this information to the PCF or Session Management Function (SMF).
Policy Control Function (PCF):


The PCF retrieves subscription information from the UDM and determines the policies to apply based on whether the user is postpaid or prepaid.
For prepaid users, the PCF ensures that the charging is done in real time (via Online Charging).
For postpaid users, it allows deferred or offline charging.
Charging Function (CHF):


The CHF is responsible for applying the correct charging model based on the user type:
Online Charging (Prepaid): Real-time credit checks and deductions. The CHF communicates with the PCF and SMF to ensure that services are provided only if the user has sufficient balance.
Offline Charging (Postpaid): Charging records are generated (e.g., Call Detail Records - CDRs) and processed later for billing.
Session Management Function (SMF):


The SMF handles session establishment and interacts with the PCF and CHF to enforce charging rules (e.g., limiting data usage for prepaid users with insufficient balance).

How the Decision Works
User Attach Process:


When a user device (UE) connects to the 5G network, it performs an attach or registration process.
The Access and Mobility Management Function (AMF) contacts the UDM to retrieve the user's subscription details, including their prepaid or postpaid status.
Policy Decision:


The PCF, using the subscription data from the UDM, identifies whether the user is prepaid or postpaid.
Based on this, the PCF defines specific policy rules, such as:
Prepaid: Enforcing credit limits, balance checks, and data usage thresholds.
Postpaid: Allowing unrestricted usage within the plan limits.
Charging Enforcement:


The CHF ensures that:
For prepaid users, charges are applied in real time, and service is denied if the balance is insufficient.
For postpaid users, the usage is recorded and billed periodically.

Summary of Decision Flow
UDM: Stores user type (prepaid or postpaid) and shares it with the PCF and SMF.
PCF: Applies policies specific to prepaid or postpaid users.
CHF: Handles charging (online for prepaid, offline for postpaid).
SMF: Enforces session and charging rules in the data path.

Examples of Policy Enforcement
Prepaid:
User has 1 GB of data balance. The CHF deducts usage in real time, and the PCF ensures data stops once the balance is exhausted.
Postpaid:
User has unlimited data in their plan. The PCF allows unrestricted usage, and the CHF generates a monthly usage report for billing.
—-----------------------------------------------------------------------------------------------------------------------
In 5G, the Unified Data Management (UDM) stores a subscriber's profile when the subscriber's data is provisioned into the 5G Core Network. 
The subscriber profile contains all necessary information required for managing the subscriber's access, policies, and services. Here’s when and how the profile is stored in the UDM:

When is the Subscriber Profile Stored in the UDM?
Subscriber Provisioning (Initial Subscription Setup):


When a new subscriber is registered by the network operator, their profile is created and provisioned into the UDM.
This typically happens during the onboarding of the subscriber, such as when a SIM card (or eSIM) is issued or a new account is created.
The provisioning process involves storing all subscription-related information like IMSI (International Mobile Subscriber Identity), QoS profiles, policy details, and subscription type (e.g., prepaid/postpaid).
Profile Updates:


If there are any changes to the subscriber's profile (e.g., upgrading their plan, changing their service preferences, or modifying policies), the UDM is updated accordingly.
Updates can be triggered manually by customer service or automatically through backend systems.
During Network Integration or Migration:


When a network transitions to 5G or integrates new systems, subscriber profiles from legacy systems like Home Subscriber Server (HSS) in 4G may be migrated to the UDM.

What Information is Stored in the UDM?
The UDM stores comprehensive subscriber information, including:
Subscriber Identity:


IMSI (International Mobile Subscriber Identity)
MSISDN (Mobile Station International Subscriber Directory Number, i.e., phone number)
SUPI (Subscription Permanent Identifier)
Subscription Type and Plan:


Prepaid or Postpaid status
Allowed data, voice, and SMS quotas
Plan validity and service entitlements
Policy and QoS Information:


Allowed Quality of Service (QoS) levels
Application-specific policies (e.g., streaming, gaming)
Authentication and Security:


Authentication vectors for secure access
Keys for encryption and integrity protection
Location and Mobility Information:


Information to help locate and route services to the subscriber.
Network Slicing Information:


Allowed network slices and corresponding policies.
Roaming Permissions:


Information on allowed roaming partners and policies.

How is the Subscriber Profile Stored?
Data Structure:


The UDM stores profiles in a structured database format to ensure efficient access and scalability.
The profiles are organized in a way that allows retrieval based on unique identifiers like IMSI or SUPI.
Interaction with Other 5G Core Components:


The UDM shares subscriber information with other components in the 5G Core (e.g., AMF, PCF, CHF, and SMF) to manage authentication, policy enforcement, charging, and session establishment.
Security and Encryption:


Subscriber data in the UDM is encrypted to ensure privacy and prevent unauthorized access.

Workflow Example
Provisioning:


A new user purchases a 5G plan. The operator's backend system creates the subscriber profile and stores it in the UDM.
Network Registration:


When the subscriber's device (UE) powers on, the Access and Mobility Management Function (AMF) queries the UDM for the subscriber's profile using the SUPI or IMSI.
Policy Retrieval:


The UDM sends the relevant data (e.g., authentication vectors, QoS policies) to the AMF, PCF, and SMF to enable the user to access the network and services.

Conclusion
The subscriber's profile is stored in the UDM during the provisioning or migration phase. It is updated as necessary and plays a vital role in enabling 5G services by providing authentication, policy, and subscription-related information to other 5G Core Network functions. Let me know if you'd like more details about the UDM's interaction with other network components!

