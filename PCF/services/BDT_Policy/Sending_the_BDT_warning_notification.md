https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.394rv4jl9afj

				Npcf_BDTPolicyControl_Notify

This procedure is used by the PCF to inform the NF service consumer that network performance in the area of interest goes below the criteria set by the operator, as defined in subclause 6.1.2.4 of 3GPP TS 23.503 [4].

Figure 4.2.4.2-1 illustrates a BDT warning notification from the PCF.


When the PCF knows that the network performance in the area of interest goes below the criteria set by the operator
from the NWDAF as described in 3GPP TS 29.520 [22], the PCF retrieves all the background transfer policies from the
UDR. If the PCF determines that the background data traffic is impacted the PCF shall:
- if the BDT warning notification is enabled, invoke the Npcf_BDTPolicyControl_Notify service operation by
sending the HTTP POST request with the BDT warning notification to the NEF so that the NEF can notify the
AF; and
- if the PCF has not locally stored the background transfer policies, invoke the Nudr_DataRepository_Delete
service operation, as described in 3GPP TS 29.504 [11] and 3GPP TS 29.519 [12], to remove the affected
background transfer policy from the UDR.
The PCF shall include a "Notification" data type in a payload body of the HTTP POST request.
The "Notification" data type provided in the request body:
- shall contain the BDT Reference ID of the impacted transfer policy within the "bdtRefId" attribute;
- may contain the time window when the network performance will go below the criteria set by the operator
within the "timeWindow" attribute;
- may contain the network area where the network performance will go below the criteria set by the operator
within the "nwAreaInfo" attribute; and
- may contain the list of candidate transfer policies in the "candPolicies" attribute.
NOTE: The AF might select a new background transfer policy from the offered candidate list when receives the
BDT warning notification.
Upon the reception of the HTTP POST request from the PCF, the NEF shall acknowledge that request by sending an
HTTP response message with the corresponding status code.
If the HTTP POST request from the PCF is accepted, the NEF shall acknowledge the receipt of the notification with a
"204 No Content" response to HTTP POST request, as shown in figure 4.2.4.2-1, step 2.
If the HTTP POST request from the PCF is not accepted, the NEF shall send an HTTP error response as specified in
subclause 5.7.
