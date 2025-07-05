https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.chwb8mdyu7s4



When multiple and separately addressable PCFs have been deployed, a network functionality is required in order to ensure that an AF needing to send policies about UE traffic identified by an UE address can reach over N5 the PCF holding the corresponding PDU Session information.

This network functionality has the following characteristics:

- It has information about the user identity, the DNN, the UE (IP or Ethernet) address(es), the S-NSSAI and the selected PCF address for a certain PDU Session.

- For IP PDU Session type, it shall receive information when an IP address is allocated or released for a PDU Session.

- For Ethernet PDU Sessions supporting binding of AF request based on MAC address, it shall receive information when a MAC address is detected as being used by the UE over the PDU Session; this detection takes place at the UPF under control of SMF; This is defined in TS 23.501 [2] clause 5.8.2.

- The functionality determines the PCF address and if available the associated PCF instance ID and PCF set ID,selected by the PCF discovery and selection function described in TS 23.501 [2], according to the information provided by the AF or the NEF.

A private IPv4 address may be allocated to different PDU Sessions, e.g.

- The same UE IPv4 address is allocated to different PDU Sessions to the same DNN and different S-NSSAI;

- The same UE IPv4 address is allocated to different PDU Sessions to the same S-NSSAI and different DNN.

Same IPv4, Different S-NSSAI: A video streaming session (eMBB slice) and an IoT data session (MIoT slice) both use the same IPv4 address while connecting to the same internet (DNN: "internet").


Same IPv4, Different DNN: A gaming session (DNN: "gaming") and a corporate VPN session (DNN: "vpn") both use the same IPv4 address while operating on the same network slice (e.g., eMBB).



In the case of private IPv4 address being used for the UE, the AF or the NEF may send DNN S-NSSAI, in addition, in Npcf_PolicyAuthorization_Create request and Nbsf_Management_Discovery request. 

The DNN and S-NSSAI can be used by the PCF for session binding, and they can be also used to help selecting the correct PCF.
