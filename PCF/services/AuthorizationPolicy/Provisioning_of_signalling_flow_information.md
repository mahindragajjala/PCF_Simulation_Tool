https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.p8cimev1oqlg

This subclause is applicable when IMS restoration is supported according to the supported feature "ProvAFsignalFlow" as described in subclause 5.8.

This procedure allows an AF to provision information about the AF signalling IP flows between the UE and the AF.



The AF shall provide:
- the IP address (IPv4 or IPv6) of the UE in the "ueIPv4" or "ueIPv6" attribute; and
- a media component within the "medComponents" attribute including:
- the "medCompN" attribute set to "0"; and
- one or more media subcomponents within the "medSubComps" attribute representing the AF signalling IP flows, where each media subcomponent shall contain:
- the "flowUsage" attribute set to the value "AF_SIGNALLING";
- the "fNum" attribute set according to the rules described in Annex C;
- the "fDesc" attribute containing the IP flows of the AF signalling flow;
- the "fStatus" set to the value "ENABLED"; and
- the "afSigProtocol" set to the value corresponding to the signalling protocol used between the UE and the AF. 


The PCF shall perform session binding and shall reply to the AF as described in subclause 4.2.2.2.
PCC rules related to the AF signalling IP flows could have been provisioned to SMF using the corresponding procedures specified in 3GPP TS 29.512 [8] at an earlier stage (e.g. typically at the establishment of the QoS flow for AF Signalling IP Flows). 
The PCF shall install the corresponding dynamic PCC rule for the AF signalling IP flows.
The AF may de-provision the information about the AF signalling IP flows at any time. 
To do that, if the "Individual Application Session Context" resource is only used to provide information about the AF Signalling IP flows, the AF shall remove the resource by sending an Npcf_PolicyAuthorization_Delete service operation as service operation towards the PCF as defined in subclause 4.2.4.2. Otherwise, the AF shall remove the IP flows within the media component invoking the Npcf_PolicyAuthorization_Update service operation as defined in subclause 4.2.3.17.
