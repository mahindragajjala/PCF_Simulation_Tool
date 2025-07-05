This procedure is used by an AF to modify a BDT warning notification request indication when "BdtNotification_5G" feature is supported.

Upon reception of a request from the AF to modify the BDT warning notification request indication, the NEF shall invoke the Npcf_BDTPolicyControl_Update service operation by sending an HTTP PATCH request to the PCF, as described in subclause 4.2.3.2. 

The NEF shall indicate whether a BDT warning notification shall be enabled or disabled by including the "warnNotifReq" attribute in the "BdtPolicyDataPatch" data type.

If the BDT warning notification is not required anymore the NEF shall set the value of the "warnNotifReq" attribute to "false".

If the BDT warning notification is again required the NEF shall set the value of the "warnNotifReq" attribute to "true".

Upon the reception of the HTTP PATCH request from the NEF indicating a modification of the BDT warning notification request indication, the PCF shall acknowledge that request by sending an HTTP response message as described in subclause 4.2.3.2.
