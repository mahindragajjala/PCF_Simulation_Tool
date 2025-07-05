				TS 29.554 


The BDT Policy Control Service, as defined in 3GPP TS 23.502 [3] and 3GPP TS 23.503 [4], is provided by the Policy Control Function (PCF).
This service enables the NF service consumer to negotiate policy for a future background data transfer and offers the following functionalities:


- GET background data transfer policies based on the request from the NEF;

- UPDATE background data transfer policies based on the selection provided by the NEF; and

- Provide background data transfer warning notification to trigger renegotiation of background data transfer policy.

The 5G System Architecture is defined in 3GPP TS 23.501 [2]. The Policy and Charging related 5G architecture is also described in 3GPP TS 29.513 [5].
The BDT Policy Control Service (Npcf_BDTPolicyControl) is part of the Npcf service-based interface exhibited by the Policy Control Function (PCF).
The only known NF service consumer of the Npcf_BDTPolicyControl service is the Network Exposure Function (NEF).
The NEF accesses the BDT Policy Control Service at the PCF via the N30 Reference point. In the roaming scenario, the N30 reference point is located between the PCF and the NEF in the home network only. 




NF Service Consumers
The Network Exposure Function (NEF):
- requests the PCF to provide background data transfer policies;
- provides the selected background data transfer policy to the PCF; and
- indicates to the PCF whether to provide a BDT warning notification.
