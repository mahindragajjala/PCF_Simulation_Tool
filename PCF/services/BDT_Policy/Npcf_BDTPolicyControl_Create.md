https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.x99bkpe8rajy

The Npcf_BDTPolicyControl_Create service operation is used by an NF service consumer to retrieve BDT policies from the PCF.

The following procedure using the Npcf_BDTPolicyControl_Create service operation is supported:
- retrieval of BDT policies


Retrieval of BDT policies
This procedure is used by the NEF to request BDT policies from the PCF, as defined in 3GPP TS 23.501 [2], 3GPP TS 23.502 [3] and 3GPP TS 23.503 [4].

Figure 4.2.2.2-1 illustrates a retrieval of BDT policies


Upon reception of a Background Data Transfer request from the AF indicating a transfer policy request, the NEF shall invoke the Npcf_BDTPolicyControl_Create service operation by sending an HTTP POST request to the URI representing a "BDT policies" collection resource of the PCF (as shown in figure 4.2.2.2-1, step 1).

The NEF shall include a "BdtReqData" data type in a payload body of the HTTP POST request. The "BdtReqData" data type shall contain:

- an ASP identifier in the "aspId" attribute;
- a volume of data per UE in the "volPerUe" attribute;
- an expected number of UEs in the "numOfUes" attribute;
- a desired time window in the "desTimeInt" attribute; and
- if "BdtNotification_5G" feature is supported a notification URI in the "notifUri" attribute,
and may include:
- a network area information (e.g. list of TAIs and/or NG-RAN nodes and/or cells identifiers) in the "nwAreaInfo" attribute;
- an identification of a group of UE(s) via an "interGroupId" attribute;
- a traffic descriptor of background data within the "trafficDes" attribute;
- if "BdtNotification_5G" feature is supported an indication that BDT warning notification is requested in the "warnNotifReq" attribute; and
- a DNN and an S-NSSAI, corresponding to the ASP identifier, in the "dnn" attribute and the "snssai" attribute respectively.
If the PCF cannot successfully fulfil the received HTTP POST request due to the internal PCF error or due to the error in the HTTP POST request, the PCF shall send the HTTP error response as specified in subclause 5.7.
Otherwise, upon the reception of the HTTP POST request from the NEF indicating a BDT policies request, the PCF:
- may invoke the Nudr_DataRepository_Query service operation, as described in 3GPP TS 29.504 [11] and 3GPP TS 29.519 [12], to request from the UDR all stored transfer policies; 

NOTE 1: In case only one PCF is deployed in the network, transfer policies can be locally stored in the PCF and the interaction with the UDR is not required.

- shall determine one or more acceptable transfer policies based on:
a) information provided by the NEF; and
b) other available information (e.g. the existing transfer policies, network policy, load status estimation for the desired time window); and

shall create a BDT Reference ID.

The PCF shall send to the NEF a "201 Created" response to the HTTP POST request, as shown in figure 4.2.2.2-1, step 2. The PCF shall include in the "201 Created" response:
a Location header field; and
a "BdtPolicy" data type in the payload body containing the BDT Reference ID in the "bdtRefId" attribute and acceptable transfer policy/ies in the "transfPolicies" attribute.
The Location header field shall contain the URI of the created individual BDT policy resource i.e. "{apiRoot}/npcfbdtpolicycontrol/v1/bdtpolicies/{bdtPolicyId}".

For each included transfer policy, the PCF shall provide:
- a transfer policy ID in the "transPolicyId" attribute;
- a recommended time window in the "recTimeInt" attribute; and
- a reference to charging rate for the recommended time window in the "ratingGroup" attribute,
and may provide a maximum aggregated bitrate for the uplink direction in the "maxBitRateUl" attribute and/or a maximum aggregated bitrate for the downlink direction in the "maxBitRateDl" attribute.

The PCF may map the ASP identifier into a target DNN and S-NSSAI based on local configuration, if the NEF did not provide the DNN and S-NSAAI to the PCF.

If the PCF included in the "BdtPolicy" data type:

- more than one transfer policy, the PCF shall wait for the transfer policy selected by the NEF as described in subclause 4.2.3; or

- only one transfer policy, the PCF may invoke the Nudr_DataRepository_Update service operation, as described in 3GPP TS 29.504 [11] and 3GPP TS 29.519 [12], to update the UDR with the selected transfer policy, the corresponding BDT Reference ID, the volume of data per UE, the expected number of UEs and, if available, a network area information, the associated DNN and S-NSSAI for the provided ASP identifier.

NOTE 2: In case only one PCF is deployed in the network, transfer policies can be locally stored in the PCF and the interaction with the UDR is not required.

Detailed Explanation of Background Data Transfer (BDT) Policy Control in 5G
1. Overview of Background Data Transfer (BDT)
Background Data Transfer (BDT) is a feature in 5G networks that allows the transfer of non-time-sensitive data in an optimized way. This mechanism helps to reduce network congestion by scheduling data transfers when network resources are less utilized.
The Network Exposure Function (NEF) requests a BDT policy from the Policy Control Function (PCF) to manage the transfer of background data.
The PCF evaluates the request and provides an acceptable transfer policy.
If applicable, the Unified Data Repository (UDR) stores or updates policy data.

2. Step-by-Step Process of BDT Policy Control
The process follows the steps below:
Step 1: NEF Sends a Background Data Transfer Request
When an Application Function (AF) wants to transfer background data, it sends a request to the NEF.


The NEF formulates a BDT policy request and sends an HTTP POST request to the PCF.


The request URI follows this format:

 {apiRoot}/npcf-bdtpolicycontrol/v1/bdtpolicies


The request includes a "BdtReqData" object containing essential attributes:


Attribute
Description
aspId
ASP identifier of the requesting entity
volPerUe
Volume of data per User Equipment (UE)
numOfUes
Expected number of UEs
desTimeInt
Desired time window for transfer
notifUri (if supported)
Notification URI for BDT status
nwAreaInfo (optional)
Network area information (TAI, NG-RAN node, cell ID)
interGroupId (optional)
UE group identifier
trafficDes (optional)
Traffic descriptor for background data
warnNotifReq (optional)
Request for warning notifications
dnn & snssai (optional)
Data Network Name (DNN) and Slicing identifier (S-NSSAI)

Step 2: PCF Processes the Request
Upon receiving the HTTP POST request from NEF, the PCF performs the following actions:
Validate the Request:


Checks for errors in the received request.
If any error is found, the PCF responds with an HTTP error response.
Retrieve Stored Policies (Optional)


If needed, the PCF queries the UDR using Npcf_BDTPolicyControl_Create service to retrieve stored BDT transfer policies.
If the network has only one PCF, policies may be stored locally, eliminating the need for UDR interaction.
Determine Acceptable Transfer Policies


Based on:
Information from the NEF request.
Existing transfer policies stored in UDR or PCF.
Network policy and resource load estimation for the desired time window.
PCF creates a BDT Reference ID for the request.
Step 3: PCF Sends a Response to NEF
If the request is valid and processed successfully:
PCF sends a "201 Created" HTTP response.


The response contains:


A Location header with the URI of the created BDT policy:
 {apiRoot}/npcf-bdtpolicycontrol/v1/bdtpolicies/{bdtPolicyId}


A "BdtPolicy" object that includes:
Attribute
Description
bdtRefId
BDT Reference ID
transfPolicies
List of acceptable transfer policies


Each Transfer Policy in the response contains:


Attribute
Description
transPolicyId
Unique identifier for the transfer policy
recTimeInt
Recommended time window for data transfer
ratingGroup
Reference to charging rate for the time window
maxBitRateUl (optional)
Max uplink bitrate
maxBitRateDl (optional)
Max downlink bitrate

Step 4: Handling Multiple Transfer Policies
If multiple transfer policies are provided:


NEF must select one and inform PCF.
PCF waits for this selection before proceeding.
If only one policy is provided:


PCF may update the UDR with:
Selected transfer policy.
BDT Reference ID.
Volume of data per UE.
Expected number of UEs.
Network area information (if available).
Associated DNN and S-NSSAI.
Step 5: NEF Uses the BDT Policy
NEF receives the BDT policy and informs the AF about the accepted transfer policy.
AF follows the policy to schedule and execute background data transfers.

3. Error Handling
If the PCF cannot process the request due to:
Internal PCF error → Sends an HTTP error response.
Invalid request parameters → Responds with an error specifying the issue.

4. Summary of the Interaction
Step
Entity
Action
1
AF → NEF
Sends Background Data Transfer request
2
NEF → PCF
Sends HTTP POST request with "BdtReqData"
3
PCF
Validates and processes the request
4
PCF → UDR
(Optional) Queries stored policies
5
PCF
Determines acceptable transfer policies
6
PCF → NEF
Sends "201 Created" response with BDT policy
7
NEF → PCF
(If multiple policies) Selects one policy
8
PCF → UDR
(If needed) Updates UDR with selected policy
9
NEF → AF
Informs AF of the BDT policy


5. Key Takeaways
BDT is used to optimize non-time-sensitive data transfers.
NEF sends a request to PCF, which evaluates and provides a policy.
PCF determines acceptable transfer policies, possibly querying UDR.
PCF responds with a 201 Created message, containing a BDT Reference ID and acceptable policies.
NEF selects a policy (if multiple options exist) and informs the AF.
PCF may store the selected policy in the UDR.
If errors occur, PCF returns an HTTP error response.

