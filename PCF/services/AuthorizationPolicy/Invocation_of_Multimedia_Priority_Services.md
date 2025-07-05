https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.8vmojc11sj8l

This procedure allows an AF, as per 3GPP TS 22.153 [23], to request prioritized access to system resources in situations such as during congestion.
The AF may include the "mpsId" attribute to indicate that the new AF session relates to an MPS session.

The "mpsId" attribute shall contain the national variant for the MPS service name indicating an MPS session. The "resPrio" attribute shall include the priority value of the related priority service.


If the AF supports the SBI Message Priority mechanism for an MPS session, it shall include the "3gpp-Sbi-MessagePriority" custom HTTP header towards the PCF as described in subclause 6.8.2 of 3GPP TS 29.500 [5].

NOTE: If the AF supports the SBI Message Priority mechanism for an MPS session, the AF will include the "3gpp-Sbi-Message-Priority" custom HTTP header with a priority value equivalent to the value of the "resPrio" attribute. 

Highest user priority value is mapped in the corresponding lowest value of the "3gppSbi-Message-Priority" custom HTTP header.

When the PCF receives the "mpsId" attribute indicating an MPS session, the PCF shall take specific actions on the corresponding PDU session to ensure that the MPS session is prioritized as specified in 3GPP TS 29.512 [8].


MPS is applicable in EPS and 5GS systems. 

MPS provides priority treatment to increase the probability of an authorized Service User's Voice, Video, Data and Messaging communication being successful.
