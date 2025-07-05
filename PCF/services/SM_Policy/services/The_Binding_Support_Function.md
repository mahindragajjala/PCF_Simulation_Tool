https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.f1s9js548ttj

The BSF has the following characteristics:
- For a certain PDU session, the BSF stores internally information about the user identity, the DNN, the UE (IP or Ethernet) address(es), the S-NSSAI, the selected PCF address and  if available the associated PCF instance ID, PCF set ID and the level of binding (see clause 6.3.1.0 of TS 23.501 [2]).

NOTE 1: Only NF instance or NF set Level of Binding indication are supported at the BSF.

The PCF registers, updates and removes the stored information in the BSF using the Nbsf management service operations defined in TS 23.502 [3].


The PCF ensures that it is updated each time an IP address is allocated or de-allocated to the PDU Session or, for Ethernet PDU Sessions supporting binding of AF request based on MAC address, each time it has been detected that a MAC address is used or no more used by the UE in the PDU Session.


Based on operator's policies and configuration, the PCF determines whether the same PCF shall be selected for the SM Policy associations to the same UE ID, S-NSSAI and DNN combination in the non-roaming or home-routed scenario.

NOTE 2: This applies to usage monitoring.

The selected PCF (if needed) downloads the user profile from the UDR as described in TS 23.502 [3] 4.16.4 step 2. If usage monitoring is enabled for the user, and based on operator's policies, the PCF checks if the BSF
has already existing PCF serving the combination of SUPI, S-NSSAI, DNN.

If no such PCF is found the PCF shall register itself to the BSF as described above in this clause.
Else if an existing PCF is found for the above combination, the PCF shall return to the SMF the available information about the existing PCF and a redirection indication.

NOTE 3: The assumption is that for DNN, S-NSSAI combinations where usage monitoring be applied, the same BSF instance or the same BSF SET is selected for all UE PDU Sessions to the same DNN, S-NNSAI.


For retrieval binding information, any NF, such as NEF or AF, that needs to discover the selected PCF address(es), and if available, the associated PCF instance ID, PCF set ID and level of binding (see clause 6.3.1.0 of TS 23.501 [2]) for the tuple (UE address, DNN, S-NSSAI, SUPI, GPSI) (or for a subset of this Tuple) uses the Nbsf management service discovery service operation defined in TS 23.502 [3].
The NF may discover the BSF via NRF or based on local configuration. 
In the case of via NRF the BSF registers the NF profile in NRF. 
The Range(s) of UE IPv4 addresses, Range(s) of UE IPv6 prefixes supported by the BSF may be provided to NRF.
If the NF received a PCF set ID or a PCF instance ID with an indication of level of binding as result of the Nbsf management service discovery service operation, it should use that information as NF set level or NF instance level Binding Indication to route requests to the PCF as defined in clause 6.3.1.0 of TS 23.501 [2] and according to the following provisions:

For the NF set level of binding, the NF will receive a PCF set ID but no PCF instance ID. 
If an NF is not able to reach the received PCF address(es) and applies direct discovery, it should query the NRF for PCF instances within the PCF set and select another instance.
For the NF instance level of binding, the NF will receive a PCF set ID and a PCF instance ID. 
If an NF is not able to reach the received PCF address(es) and applies direct discovery, it should query the NRF for PCF service instances within the PCF and select another instance.
The NF should provide a Routing Binding Indication based on the received PCF set ID, level of binding and possible PCF instance ID in requests it sends to the PCF.
For an ongoing NF service session, the PCF may provide Binding indication to the NF (see clause 6.3.1.0 of TS 23.501 [2]). 
This Binding indication shall then be used instead of any PCF information received from the BSF.
If a new PCF instance is selected, the new PCF should invoke Nbsf_Management_Update service operation to update the binding information in BSF.

The BSF may be deployed standalone or may be collocated with other network functions, such as PCF, UDR, NRF, SMF. 

NOTE 4: Collocation allows combined implementation. 
