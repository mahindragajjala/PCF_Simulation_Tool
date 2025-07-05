https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.ygsmmqjy97pg

Policy decisions based on network analytics allow PCF to perform policy decisions taking into account the analytics information provided by the NWDAF. 
The PCF subscribes/unsubscribes to Analytics information as defined in TS 23.288 [24].
The following Analytics IDs are relevant for Policy decisions: "Load level information", "Service Experience","Network Performance" and "Abnormal behaviour". 
The PCF may subscribe to notifications of network analytics related to "Load Level Information" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Load level information", the Analytics Filter "S-NSSAI and NSI ID" and the Analytics Reporting Information set to a load level threshold value. 
The PCF is notified when the load level of the Network Slice Instance reaches the threshold, and then the PCF may verify if the RFSP index value needs to be modified for a SUPI for which an AM Policy Association is created; this is based on operator policies in the PCF, as defined in clause 6.1.2.1.
The PCF may subscribe to notifications of network analytics related to "Service Experience" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Service Experience", the Target of Analytics Reporting "any UE" and the Analytics Filter including one or more "Application ID(s)". 
The PCF is notified on the Service Experience statistics or predictions including, for each Application Id, the list of SUPIs for which Service Experience is provided. 
In addition, both spatial and time validity may be provided as well as the confidence of the prediction. 
The PCF may check the 5QI values assigned to the Application, the number of UEs
affected and may use this as input to calculate and update the authorized QoS for a service data flow template.
The NWDAF service to retrieve the service experience (i.e. the average observed Service MoS) is described in clause 6.4 of TS 23.288 [24].
The PCF may subscribe to notifications of network analytics related to "Network Performance" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Network Performance", the Target of Analytics Reporting "Internal Group Id" and the Analytics Filter including the Area of Interest. 
The PCF is notified on the Network Performance statistics or predictions including the Area of Interest and both the gNB status information and the predicted number of UEs in the Area of Interest. In addition, the confidence of the prediction may be provided. 
The PCF may use this information as input to calculate the background data transfer policies that are negotiated with the ASP, as defined in clause 6.1.2.4.
The NWDAF services to retrieve "Network Performance" as described in clause 6.6 of TS 23.288 [24].
The PCF may subscribe to notifications of network analytics related to "Abnormal behaviour" using the
Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Abnormal behaviour", the
Target of Analytics Reporting "SUPI", "Internal Group Id" or "any UE" and the Analytics Filter including the expected
analytics type or the list of Exceptions IDs and per each Exception Id a possible threshold and other Analytics Filter
Information if needed. The list of Exception IDs is specified in TS 23.288 [24]. The PCF may use "Unexpected UE
location" as input to determine the Service Area Restrictions defined in clause 6.1.2.1, "Suspicion of DDoS attack" or
"Too frequent Service Access" to request the SMF to terminate the PDU session as defined in clause 6.1.3.6, "Wrong
destination address" to perform gating of a service data flow as defined in clause 6.1.3.6 and "Unexpected longlive/large rate flows" to perform QoS related policies such as gating or policing as defined in clause 6.2.1.1.
The NWDAF services to retrieve UE related analytics are described in clause 6.7 of TS 23.288 [24].
The PCF may also use the network analytics as input to its policy decision to apply operator defined actions for
example for the UE context(s) or PDU session(s).


Here’s a simplified explanation of the content and a basic call flow for the PCF (Policy Control Function) leveraging NWDAF (Network Data Analytics Function) analytics to make policy decisions:

Simplified Explanation:
Role of PCF and NWDAF:


PCF: Makes policy decisions for users (UEs), applications, or network services based on analytics information.
NWDAF: Provides insights and predictions about network behavior (e.g., load, performance, user experience).
Analytics IDs Used by PCF:


Load Level Information: Tracks network slice load and adjusts policies when a threshold is reached.
Service Experience: Monitors and predicts user or application service quality (e.g., video quality).
Network Performance: Measures and predicts overall network performance (e.g., congestion).
Abnormal Behavior: Identifies unusual events (e.g., suspicious activity or wrong network behavior).
PCF Subscriptions:


PCF subscribes to NWDAF analytics by specifying:
Analytics ID (e.g., "Load Level Information").
Filters (e.g., targeting specific network slices, UEs, or applications).
Thresholds (e.g., trigger a notification if network load exceeds a certain level).
What Happens After a Notification:


When NWDAF notifies PCF of an event (e.g., high network load):
PCF evaluates its policies (e.g., adjust service quality, restrict areas, or terminate sessions).
PCF updates rules for the affected UEs, applications, or network slices.

Call Flow for PCF Using NWDAF Analytics:
Example: PCF Managing "Load Level Information"
Subscription:


PCF subscribes to NWDAF for "Load Level Information" analytics:
Analytics ID: "Load Level Information".
Filters: Target specific S-NSSAI (Single Network Slice Selection Assistance Information) and NSI ID (Network Slice Instance Identifier).
Threshold: Set a load level limit.
Notification:


NWDAF monitors the network slice.
When the load crosses the threshold, NWDAF sends a Notification to the PCF.
Policy Decision:


PCF receives the notification and checks its policies:
Adjust RFSP Index (used to prioritize services for a user).
Update AM Policy for the affected user (SUPI).
Action:


PCF applies the updated policy and informs the appropriate network function (e.g., Access and Mobility Function or Session Management Function).

Example Call Flow: "Abnormal Behavior" (e.g., DDoS Detection)
Subscription:


PCF subscribes to NWDAF for "Abnormal Behavior" analytics:
Analytics ID: "Abnormal Behavior".
Filters: Target specific UEs, groups, or network areas.
Specify Exception IDs (e.g., DDoS attack or unexpected UE location).
Notification:


NWDAF detects a DDoS attack targeting a group of UEs.
NWDAF sends a Notification to the PCF with details (e.g., affected UEs and network area).
Policy Decision:


PCF evaluates the notification and decides:
Block the data flow for the affected UEs.
Inform the SMF to terminate the PDU session for those UEs.
Action:


PCF communicates the decision to SMF or other functions.
The attack is mitigated, and affected users are secured.

General Call Flow Diagram
1. PCF → NWDAF: Subscribe to analytics (e.g., Load Level Information, Abnormal Behavior).
2. NWDAF → PCF: Monitor analytics and send notifications (e.g., load threshold exceeded, abnormal behavior detected).
3. PCF → Internal Policy Engine: Evaluate operator-defined policies.
4. Policy Engine → PCF: Recommend policy actions (e.g., update RFSP, terminate session, adjust QoS).
5. PCF → Network Functions (e.g., SMF, AMF): Enforce updated policies.
6. Network Functions → UEs/Applications: Implement changes (e.g., restrict area, terminate sessions, adjust service quality).

