Initial Background Data Transfer policy indication

This procedure is used by an AF to indicate a transfer policy negotiated for background data transfer using the Npcf_BDTPolicyControl service as described in 3GPP TS 29.554 [14].


The AF may include in the HTTP POST request message described in subclause 4.2.2.2 a reference identifier related to a transfer policy negotiated for background data transfer in the "bdtRefId" attribute.


NOTE 1: 

The PCF will retrieve the corresponding transfer policy from the UDR based on the reference identifier within the "bdtRefId" attribute. 

In case only one PCF is deployed in the network, transfer policies can be locally stored in the PCF and the interaction with the UDR is not required.


If the PCF cannot retrieve the transfer policy, the PCF shall set to TP_NOT_KNOWN the "servAuthInfo" attribute in the HTTP response message to the AF to indicate that the transfer policy is unknown.

If the time window of the received transfer policy has expired, the PCF shall set to TP_EXPIRED the "servAuthInfo" attribute in the HTTP response message to indicate to the AF that the transfer policy has expired. 

Otherwise, if the time window of the received transfer policy has not yet occurred, the PCF shall set to TP_NOT_YET_OCCURRED the "servAuthInfo" attribute in the HTTP response message to the AF to indicate that the time window of the transfer policy has not yet occurred.

NOTE 2: In the case that the PCF cannot retrieve the transfer policy, the transfer policy time window has not yet occurred or the transfer policy expired, the PCF makes the decision without considering the transfer policy.

The PCF shall reply to the AF as described in subclause 4.2.2.2.
