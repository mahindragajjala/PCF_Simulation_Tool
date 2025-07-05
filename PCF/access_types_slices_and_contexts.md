https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.qn267yhyw1jt


In 5G networks, access types, slices, and contexts are key dimensions used to differentiate and apply policies to user sessions. Here's a detailed explanation of each:

1. Access Types
Access types refer to the different ways a user connects to the network. In a 5G environment, multiple access technologies are supported, and policies can be tailored based on the type of access.
Examples of Access Types:


3GPP Access: This is the primary access type provided by 5G NR (New Radio) or LTE (4G). It involves cellular connections through base stations.
Non-3GPP Access: Refers to non-cellular access networks such as Wi-Fi or fixed-line connections (e.g., DSL or fiber).
Interworking Access: Seamless transition between 3GPP and non-3GPP networks (e.g., switching between mobile and Wi-Fi).
Policy Relevance: Policies may vary depending on whether the user is connected via 5G NR or a Wi-Fi network. For instance:


5G NR might offer higher priority for ultra-reliable low-latency communication (URLLC).
Wi-Fi may have lower priority or different QoS requirements.

2. Slices
Network slicing is one of the defining features of 5G, where the physical network is divided into multiple virtual networks (slices) to serve different use cases.
What is a Slice?
 A network slice is a logically isolated part of the network, tailored to meet specific service requirements. Each slice can have its own QoS, security, and policy settings.


Examples of Network Slices:


eMBB (Enhanced Mobile Broadband): High-speed internet access for data-intensive applications like video streaming.
URLLC (Ultra-Reliable Low-Latency Communication): Critical for applications like autonomous driving or industrial automation.
mMTC (Massive Machine-Type Communications): Designed for IoT applications requiring large-scale device connectivity with low bandwidth.
Policy Relevance: Different slices can have different AM policies based on their use case:


eMBB slice: Allows unrestricted internet access with high bandwidth.
URLLC slice: Implements stricter policies for latency and reliability.
mMTC slice: Enforces low-power and low-priority access policies.

3. Contexts
Contexts refer to the specific conditions or situations of a user’s session, which can influence the applied policies.
Examples of Contexts:


Location Context: Policies based on the user's geographic location (e.g., policies can differ for urban, rural, or roaming scenarios).
RAT Type Context: The specific Radio Access Technology in use (e.g., LTE vs. 5G NR).
Subscription Context: The user’s subscription plan or profile (e.g., premium users vs. regular users).
Application Context: The type of application or service being accessed (e.g., streaming video, online gaming, or IoT telemetry).
Mobility Context: Whether the user is stationary or moving (e.g., traveling between cells or across networks).
Policy Relevance:


A user in a high-density area might have restricted access to certain services to avoid network congestion.
Policies for users on a high-speed train (mobility context) might ensure seamless handovers between cells.

How They Work Together
The AM Policy Association applies general policies across all access types, slices, and contexts. The Individual AM Policy Association, however, fine-tunes these policies for specific conditions:
Example:
A user connects to the 5G network using eMBB (slice), via 3GPP access (access type), and is streaming a video while roaming (context).
The AM Policy Association sets general mobility and QoS rules.
The Individual AM Policy Association ensures the user has high bandwidth, minimal packet loss, and unrestricted streaming while managing roaming conditions.

Let me know if you'd like further examples or a visual representation!

