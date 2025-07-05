Event Exposure

This service:
- allows NF service consumers to subscribe, modify and unsubscribe from policy control events; and
- notifies NF service consumers with a corresponding subscription about observed events on the PCF.
The types of observed events include:
- PLMN identifier notification; and
- Access type change.

The target of the event reporting may include a group of UE(s) or any UE (i.e. all UEs). 

When an event occurs, to which the NF service consumer has subscribed, the PCF reports the requested information to the NF service consumer based on the event reporting information definition requested by the NF service consumer (see 3GPP TS 23.502 [3], subclause 4.15.1).

Exposure - the state of having no protection from something harmful.

The network capability exposure comprises
- Exposure of network events externally as well as internally towards core network NFs;
- Exposure of provisioning capability towards external functions;
- Exposure of policy and charging capabilities towards external functions;
- Exposure of core network internal capabilities for analytics.
- Exposure of analytics to external party.
- Retrieval of data from external party by NWDAF.
When subscribing to event reporting the NF consumer(s) provide:

- One or multiple Event ID(s). An Event ID identifies the type of event being subscribed to (e.g. PDU Session Release, UE mobility out of an Area of Interest, etc.).
- Event Filter Information: Provides Event Parameter Types and Event Parameter Value(s) to be matched against, in order to meet the condition for notifying the subscribed Event ID e.g. the Event Parameter Type could be "Area of interest" and Event Parameter Value list could be list of TAs; The Event Filter depends on the Event ID. 

The Event Filter Information is provided per Event ID(s) being subscribed to: within a subscription different Event ID(s) may be associated with different Event Filter Information.

- Event Reporting Information described in the Table 4.15.1-1 below. Within a subscription all Event ID(s) are associated with a unique Event Reporting Information.

- Target of Event Reporting: this may indicate a specific UE or PDU Session, a group of UE(s) or any UE (i.e. all UEs), Within a subscription all Event ID (s) are associated with the same Target of Event Reporting (possibly corresponding to multiple UE or multiple PDU Sessions).

- A Notification Target Address (+ Notification Correlation ID) allowing the Event Receving NF to correlate notifications received from the Event provider with this subscription. A subscription is associated with an unique Notification Target Address (+ Notification Correlation ID). In the case that the NF consumer subscribes to the NF producer on behalf of other NF, the NF consumer includes the Notification Target Address(+Notification Correlation ID) of other NF for the Event ID which is to be notified to other NF directly, and the Notification Target Address(+Notification Correlation ID) of itself for the Subscription change related event notification.
Here’s a detailed breakdown of the explanation in a few key points:
Definition of Notification Target Address and Notification Correlation ID


The Notification Target Address is the endpoint (e.g., URL, IP address, or another identifier) where notifications related to an event subscription should be sent.
The Notification Correlation ID is an identifier used to link received notifications back to the specific subscription request that initiated them.
Uniqueness of Subscription and Notification Address


Each subscription has a unique Notification Target Address (+ Notification Correlation ID).
This ensures that notifications can be properly routed and correlated with the correct subscription.
Event Receiving NF and Event Provider NF


The Event Receiving NF is the entity that receives notifications from the Event Provider NF.
The Notification Correlation ID helps the Event Receiving NF match incoming notifications with the original subscription request.
Handling Subscriptions on Behalf of Another NF


If an NF Consumer (acting as an intermediary) subscribes to an NF Producer on behalf of another NF, it must specify:
The Notification Target Address (+ Correlation ID) of the other NF for event notifications meant for that NF.
Its own Notification Target Address (+ Correlation ID) for subscription-related event notifications (e.g., subscription changes).
Use Case Example


Suppose NF A wants to receive event notifications but NF B (acting as an intermediary) makes the subscription request.
NF B includes:
NF A’s Notification Target Address (+ Correlation ID) for events meant for NF A.
NF B’s Notification Target Address (+ Correlation ID) for subscription change notifications (e.g., modifications or cancellations).
Purpose and Benefit


Ensures that notifications reach the correct NF.
Helps in tracking and managing subscriptions efficiently.
Facilitates inter-NF communication while allowing intermediaries to manage subscriptions effectively.

Each Notification Target Address(+ Notification Correlation ID) is associated with related (set of) Event ID(s).
- An Expiry time represents the time upto which the subscription is desired to be kept as active. The NF service consumer may suggest an Expiry time and provide to the NF service producer.  Based on the operator's policy,
the NF service producer decides whether the subscription can be expired. If the subscripton can be expired, the NF service producer determines the Expiry time and provide it in the response to the NF service consumer. If the event subscription is about to expire based on the received Expiry time and the NF service consumer wants to keep receiving notifications, the NF service consumer update the subscription with the NF service producer in order to extend the Expiry time. Once the Expiry time associated with the subscription is reached, the subscription becomes invalid at the NF service producer. If the NF service consumer wants to keep receiving
notifications, it shall create a new subscription with the NF service producer.

Here’s a detailed explanation of the additional points:
1. Association of Notification Target Address (+ Notification Correlation ID) with Event ID(s)
Each Notification Target Address (+ Notification Correlation ID) is linked to a specific set of Event IDs.
This means that different events may have different Notification Target Addresses, ensuring that notifications are routed correctly.
Example: If an NF subscribes to multiple events (e.g., "Event A" and "Event B"), it can specify different Notification Target Addresses for each.

2. Definition and Role of Expiry Time in Subscriptions
The Expiry Time defines how long a subscription remains active.
The NF service consumer (subscriber) can suggest an Expiry Time while requesting a subscription.
The NF service producer (event provider) determines whether the requested Expiry Time is acceptable based on operator policies.

3. Determination of Expiry Time by NF Service Producer
The NF service producer evaluates the subscription request and:
Accepts the suggested Expiry Time or
Modifies it based on network policies and includes the final Expiry Time in its response to the NF service consumer.

4. Subscription Renewal Before Expiry
If the subscription is about to expire, and the NF service consumer still wants to receive event notifications, it must:
Update the subscription with the NF service producer to extend the Expiry Time.
The update ensures continuous event notifications without interruption.

5. Handling of Subscription Expiry
Once the Expiry Time is reached, the subscription automatically becomes invalid at the NF service producer.
If the NF service consumer still requires notifications after expiry, it must create a new subscription with the NF service producer.

6. Key Benefits of Expiry Time Mechanism
Efficient resource management: Expired subscriptions are removed automatically, preventing unnecessary load.
Flexibility: NF service consumers can extend or create new subscriptions based on their requirements.
Control for service providers: NF service producers can enforce operator policies on subscription durations.

When the subscription is accepted by the Event provider NF, the consumer NF receives from the event provider NF an identifier (Subscription Correlation ID) allowing to further manage (modify, delete) this subscription.
NOTE 1: The Notification Correlation ID is allocated by the consumer NF that subscribes to event reporting and the Subscription Correlation ID is allocated by the NF that notifies when the event is met. Both correlation identifiers can be assigned the same value, although in principle they are supposed to be different, as they are optimized for finding the subscription related context within each NF.
The consumer NF may use an operation dedicated to subscription modification to add or remove Event ID(s) to this subscription or to modify Event Filter Information.
Events are subscribed by the consumer NF(s) by providing Event Filters. The contents of the Event Reporting Information along with the presence requirement of each information element is described in Table 4.15.1-1. 

Maximum number of reports is applicable to the subscription to one UE or a group of UE(s). When the subscription is
applied to a group of UE(s), the parameter is applied to each individual member UE. The count of number of reports is
per Event Type granularity.
Maximum duration of reporting is applicable to the subscription to one UE, a group of UE(s) or any UE. When the
subscription is applied to a group of UE(s), this parameter applies to each group member UE. When the subscription is
applied to any UE, this parameter applies to all the impacted UEs.
If for a given subscription both Maximum Number of reports and Maximum duration of reporting are included then the
subscription is considered to expire as soon as one of the conditions is met.
Sampling ratio is applicable to subscription targeting a group of UEs or any UE. When a sampling ratio is provided, a
random subset is selected among the target UEs according to the sampling ratio and only the events related to this
subset are reported. A UE remains selected until it is not managed by the event provider NF any more. A UE newly
managed by the NF may become selected.
Group Reporting Guard Time is an optional parameter for group-based monitoring configuration to indicate the time for
which the Monitoring Event Reporting(s) related with the UEs in a group can be aggregated before sending them to the
consumer NF. The value of the Group Reporting Guard time should be set less than the Maximum duration of reporting.
For the continuous monitoring reporting, unless the Maximum duration of reporting has been reached, the Group
Reporting Guard timer is restarted when it expires. If the time left until the Maximum duration of reporting is less than
the Group Reporting Guard Time, then the Group Reporting Guard timer shall be set to expire when the Maximum
duration of reporting expires. If the Maximum duration of reporting is expired, the Group Reporting Guard Time, if
running, shall be considered to expire and aggregated Monitoring Event Reporting(s) is sent to destination immediately.
Table 4.15.1-1 indicates the presence requirements for the Event Reporting Information.
Corresponding notifications contain at least the Notification Correlation ID together with the Event ID and the
individual target (e.g. UE or PDU Session ID) associated with the notification.
If the NF service consumer decides to terminate the event subscription, it unsubscribes the event subscription by
sending unsubscription request to the event provider NF. After receiving unsubscription request from the NF service
consumer, the event provider NF terminates the event subscription.
The following clauses describe the external exposure of network capabilities and core network internal event and
capability exposure.
When the immediate reporting flag is set, the first corresponding event report is included in the output message, if corresponding information is available at the reception of the subscription request of the event. The optional parameter MTC Provider Information as used e.g. in clause 4.25.3, is a reference parameter that may be provided by AF or determined by NEF based on which AF it communicates with. The MTC Provider Information identifies the MTC Service Provider and/or MTC Application. NOTE 5: The MTC Provider Information can be used by Service Providers for, e.g. to distinguish their different customers. 

Here's a detailed explanation of the key points in simple English:

1. Maximum Number of Reports
This defines the maximum number of event reports that can be generated for a subscription.
The rule applies whether the subscription is for one UE (User Equipment) or a group of UEs.
If applied to a group, each UE in the group has the same individual report limit.
The count is based on each Event Type separately, meaning different event types have their own limits.

2. Maximum Duration of Reporting
This sets the time limit for how long reports can be generated for an event subscription.
It can apply to:
One specific UE
A group of UEs (each UE follows the same time limit)
Any UE in the network (applies to all affected UEs)
If both "Maximum Number of Reports" and "Maximum Duration of Reporting" are set, the subscription will end as soon as one of these conditions is met.

3. Sampling Ratio
This is used when a subscription applies to a group of UEs or any UE.
It means that not all UEs in the group will be monitored, only a random subset.
The subset is selected based on the given sampling ratio.
Once a UE is selected, it stays in the monitored group until it is no longer managed by the event provider NF.
If a new UE joins, it may become part of the monitored group.

4. Group Reporting Guard Time
This is an optional setting for group-based monitoring.
It delays sending reports immediately and instead aggregates reports over a short period.
The guard time must always be shorter than the Maximum Duration of Reporting.
How it works:
If the guard timer expires, it resets and starts again unless the Maximum Duration of Reporting is reached.
If the time left until Maximum Duration of Reporting is less than the Guard Time, the guard timer is adjusted to match the maximum duration.
When Maximum Duration of Reporting expires, all pending reports are sent immediately.

5. Event Reporting Notifications
Each notification (event report) must include:
Notification Correlation ID (to link the notification to the subscription request)
Event ID (identifies which event triggered the report)
Target UE or PDU Session ID (which UE or session the event is related to)

6. Unsubscribing from an Event
If an NF service consumer (such as an application or another network function) wants to stop receiving event reports, it must send an unsubscribe request.
Once the event provider NF receives this request, it terminates the event subscription and stops sending reports.

7. External Exposure of Network Capabilities
The system allows external entities (such as application providers) to access and use certain network capabilities.
It also supports exposing core network events and capabilities to external applications.
This helps in better service management and integration with external applications.

8. Immediate Reporting Flag
If the Immediate Reporting Flag is set:
The first event report is included immediately in the response message when the subscription request is received.
This happens only if relevant data is available at that time.

9. MTC Provider Information
This is an optional parameter used in Machine-Type Communications (MTC).
It helps identify:
The MTC Service Provider (who is providing the service).
The MTC Application (which application is requesting the data).
This parameter is useful for:
Differentiating multiple customers who use the same service provider.
Helping the network function (NF) decide how to handle requests from different providers.

Final Summary
Maximum Number of Reports = Limits the number of reports per event type.
Maximum Duration of Reporting = Limits how long reports will be generated.
Both limits together = Subscription ends when any one condition is met.
Sampling Ratio = Selects a subset of UEs randomly for reporting.
Group Reporting Guard Time = Delays and aggregates reports before sending.
Event Reporting Notifications = Must include correlation ID, event ID, and target UE/session.
Unsubscribing = Requires a request from the NF service consumer.
Immediate Reporting = Sends first event report immediately if available.
MTC Provider Information = Helps identify different service providers and applications.
