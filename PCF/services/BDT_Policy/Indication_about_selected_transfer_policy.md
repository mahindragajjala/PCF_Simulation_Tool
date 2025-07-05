				Npcf_BDTPolicyControl_Update

This procedure is used by the NEF to inform the PCF about selected transfer policy, as defined in 3GPP TS 23.501 [2], 3GPP TS 23.502 [3] and 3GPP TS 23.503 [4], if the AF selected the transfer policy from the received transfer policy list after:

- retrieval of the BDT policies as described in subclause 4.2.2; or

- reception of the BDT warning notification as described in subclause 4.2.4.

Figure 4.2.3.2-1 illustrates an indication about selected transfer policy.



Upon reception of a Background Data Transfer request from the AF indicating transfer policy selection, the NEF shall invoke the Npcf_BDTPolicyControl_Update service operation by sending an HTTP PATCH request to the PCF, as shown in figure 4.2.3.2-1, step 1. 

The NEF shall set the request URI to "{apiRoot}/npcfbdtpolicycontrol/v1/bdtpolicies/{bdtPolicyId}".

The NEF shall include a "BdtPolicyDataPatch" data type in a payload body of the HTTP PATCH request. The "BdtPolicyDataPatch" data type shall contain a transfer policy ID of the selected transfer policy in the "selTransPolicyId" attribute.

If the PCF cannot successfully fulfil the received HTTP PATCH request due to the internal PCF error or due to the error in the HTTP PATCH request, the PCF shall send the HTTP error response as specified in subclause 5.7.

Otherwise, upon the reception of the HTTP PATCH request from the NEF indicating a selected transfer policy, the PCF:

- may invoke the Nudr_DataRepository_Update service operation, as described in 3GPP TS 29.504 [11] and 3GPP TS 29.519 [12], to update the UDR with the selected transfer policy, the corresponding BDT Reference ID, the volume of data per UE, the expected number of UEs and, if available, a network area information, the associated DNN and S-NSSAI for the provided ASP identifier;
