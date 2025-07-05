https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.su7pwybk7p9g





The policy control decision and flow based charging control functionalities enable the PCF to provide network control regarding the service data flow detection, gating, QoS and flow based charging (except credit management) towards the SMF/UPF.

The PCF receives session and media related information from the Npcf_PolicyAuthorization service consumers and notifies them of subscribed traffic plane events.

The PCF may receive from the NF service consumers the request to monitor the requested service and media information and notifies them of the UL/DL/round-trip delay of the requested flows.

The PCF may receive service routing requirements and the indication of receiving notifications about user plane path changes from the Npcf_PolicyAuthorization service consumers.

The PCF may receive from the NF service consumers the specific required QoS and a prioritized list of alternative QoS profiles and notifies them about the QoS target the access network guarantees.

The PCF checks that the service information provided by the NF service consumer is consistent with the operator defined policy rules before storing the service information.

The PCF uses the received service information and the subscription information when it applies as basis for the policy and charging control decisions.

The PCF derives PCC rules and provisions them to the SMF via the Npcf_SMPolicyControl service and subscribes to traffic plane events via policy control request triggers as described in 3GPP TS 29.512 [8]. 



The Npcf_PolicyAuthorization_Create service operation authorizes the request from the NF service consumer, and optionally communicates with Npcf_SMPolicyControl service to determine and install the policy according to the information provided by the NF service consumer.

The Npcf_PolicyAuthorization_Create service operation creates an application session context in the PCF. 

The following procedures using the Npcf_PolicyAuthorization_Create service operation are supported:
- Initial provisioning of service information.
- Gate control.
- Initial Background Data Transfer policy indication.
- Initial provisioning of sponsored connectivity information.
- Subscription to Service Data Flow QoS notification control.
- Subscription to Service Data Flow Deactivation.
- Initial provisioning of traffic routing information.
- Subscription to resources allocation outcome.
- Invocation of Multimedia Priority Services.
- Support of content versioning.
- Request of access network information.
- Initial provisioning of service information status.
- Provisioning of signalling flow information.
- Support of resource sharing.
- Indication of Emergency traffic.
- Invocation of MCPTT.
- Invocation of MCVideo.
- Priority sharing indication.
- Subscription to out of credit notification.
- Subscription to Service Data Flow QoS Monitoring information.
- Provisioning of TSCAI input information and TSC QoS related data.
- Provisioning of bridge management information and port management information.
- P-CSCF restoration enhancements.
- Support of CHEM feature.
- Support of FLUS feature.
- Subscription to EPS Fallback report.
- Subscription to TSN related events.
- Initial provisioning of required QoS information.
- Support of QoSHint feature.
- Subscription to reallocation of credit notification.
