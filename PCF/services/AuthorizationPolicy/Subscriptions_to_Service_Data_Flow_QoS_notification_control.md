https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.tfkxo3xajiti

Subscriptions to Service Data Flow QoS notification control


The subscription to Service Data Flow QoS notification control is used by an AF to subscribe to receive a notification when the GBR QoS targets for one or more service data flows can no longer (or can again) be guaranteed.

NOTE: It may happen that the GBR QoS targets for one or more PCC rules (i.e. Service Data Flows) cannot be guaranteed, either permanently or temporarily in the radio access network.

The AF shall use the "EventsSubscReqData" data type as described in subclause 4.2.2.2 and shall include in the HTTP POST request message an event within the "events" attribute with the "event" attribute set to "QOS_NOTIF".

The PCF shall reply to the AF as described in subclause 4.2.2.2. As result of this action, the PCF shall set the appropriate subscription to QoS notification control for the corresponding PCC rule(s) as described in in 3GPP TS 29.512 [8].

In 5G, GBR (Guaranteed Bit Rate) QoS targets are defined to ensure a minimum level of service for applications that require consistent data rates, such as voice over NR (VoNR), video streaming, and industrial automation. 
GBR QoS flows are used in scenarios where guaranteed bandwidth and low latency are critical.
Key GBR QoS Targets in 5G
QoS Flow Type:
GBR flows ensure a guaranteed minimum data rate.
Non-GBR flows do not have a guaranteed data rate but can be prioritized using scheduling mechanisms.
QoS Parameters for GBR Flows:
Guaranteed Bit Rate (GBR): The minimum data rate that must always be available for the QoS flow.
Maximum Bit Rate (MBR): The upper limit of the data rate that can be allocated to the flow.
Reflective QoS Indicator (RQI): Used to copy QoS parameters from downlink to uplink for efficiency.
Packet Delay Budget (PDB): Specifies the maximum tolerated delay for packets to ensure real-time application performance.
Packet Error Rate (PER): Defines the allowable packet loss to maintain service quality.
Allocation and Retention Priority (ARP): Determines whether the QoS flow should be maintained during network congestion.
GBR Services in 5G:
Voice over NR (VoNR)
Ultra-Reliable Low Latency Communications (URLLC)
High-quality video streaming
Industrial IoT and mission-critical applications
5G QoS Identifiers (5QI) for GBR:
5QI 1 → Conversational voice (VoNR)
5QI 2 → Conversational video
5QI 3 → Real-time gaming
5QI 4 → Non-conversational video (buffered streaming)
5QI 5 → Mission-critical applications (industrial automation)

Subscription to Service Data Flow Deactivation

This procedure is used by an AF to subscribe to the notification of deactivation of one or more Service Data Flows within the AF application session context.
NOTE: It may happen that one or more PCC rules (i.e. Service Data Flows) are deactivated at the SMF at certain time, either permanently or temporarily, due to e.g. release of resources or out of credit condition.
The AF shall use the "EventsSubscReqData" data type as described in subclause 4.2.2.2 and shall include in the HTTP POST request message an event within the "events" attribute with the "event" attribute set to "FAILED_RESOURCES_ALLOCATION".
The PCF shall reply to the AF as described in subclause 4.2.2.2.
As result of this action, the PCF shall set the appropriate subscription to service data flow deactivation for the corresponding PCC rule(s) as described in in 3GPP TS 29.512 [8].
