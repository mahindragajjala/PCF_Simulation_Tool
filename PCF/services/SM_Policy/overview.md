SMPOLICY 
TS 29.512



7a. If dynamic PCC is to be used for the PDU Session, the SMF performs PCF selection as described in TS 23.501 [2], clause 6.3.7.1. If the Request Type indicates "Existing PDU Session" or "Existing Emergency PDU Session", the SMF shall use the PCF already selected for the PDU Session.
 Otherwise, the SMF may apply local policy.

7b. The SMF may perform an SM Policy Association Establishment procedure as defined in clause 4.16.4 to establish an SM Policy Association with the PCF and get the default PCC Rules for the PDU Session. 

The GPSI shall be included if available at SMF. 

If the Request Type in step 3 indicates "Existing PDU Session", the SMF may provide information on the Policy Control Request Trigger condition(s) that have been met by an SMF initiated SM Policy Association Modification procedure as defined in clause 4.16.5.1. The PCF may provide policy information defined in clause 5.2.5.4 (and in TS 23.503 [20]) to SMF. The PCF, based on the Emergency DNN, sets the ARP of the PCC rules to a value that is reserved for Emergency services as described in TS 23.503 [20]. 

The Session Management Policy Control Service performs provisioning, update and removal of session related policies and PCC rules by the Policy Control Function (PCF) to the NF service consumer (i.e. SMF). 
The Session Management Policy Control Service can be used for 
charging control
policy control
application detection 
control and/or access traffic steering
switching 
splitting within a MA PDU Session
Session Management Policy Control Service applies to the cases where the SMF interacts with the PCF in the non-roaming scenario, the V-SMF interacts with the V-PCF in the local breakout roaming scenario and the H-SMF interacts with the H-PCF in the home-routed scenario. 
The PCF is responsible for policy control decisions and flow based charging control functionalities. The PCF provides
the following:
Policies for 
Application
service data flow detection
Gating
QoS
flow based charging
traffic steering control
usage monitoring control
access traffic steering
switching and steering within a MA PDU Session
access network information report
TSN BMIC
TSN port management information container 
TSN TSCAI input container 
RAN support information to the SMF.


The policy decisions made by the PCF may be based on one or more of the following:

- Information taking from the AF, e.g. the session, media and subscriber related information;
- Information taking from the UDR;
- Information taking from the AMF, e.g. UE related and access related information;
- Information taking from the SMF;
- Information taking from the NWDAF;
- Information taking from the NEF;
- information from CHF; and
- PCF pre-configured policy context.


NF Service Consumers

The SMF is responsible for the enforcement(అమలు) of session management related policy decisions from the PCF, related to service flow detection, QoS, charging, gating, traffic usage reporting, traffic steering and access traffic steering,
switching and splitting within a MA PDU Session.
The SMF shall support:
- sending the PDU session related attributes to the PCF;
- requesting and receiving the PCC rule(s) from the PCF;
- binding of service data flows to QoS flow as defined in 3GPP TS 29.513 [7];
- deriving rule(s) from the PCC rule(s) and then providing those rules to the user plane function or remove the rule(s) from the user plane as defined in 3GPP TS 29.244 [13];
- deriving the QoS rules towards the UE;
- deriving the QoS profile towards the access network;
- deriving the ATSSS rules towards the UE if applicable;
- handling the policy control request trigger; and
- handling the PDU session related policy information.
NOTE: SMF functionality related to event exposure is defined in 3GPP TS 29.508 [12].

A rule is a set of policy information elements associated with a PDU session, or with service data flows (i.e., with a PCC rule).
Two types of rules are defined:
- Session rule; and
- PCC rule.
Both Session rules and PCC rules are composed of embedded information elements as well as information elements that are part of the referenced objects (e.g. condition data, or usage monitoring policy data type) by the rule.

PCC rule is defined in subclause 4.1.4.2. Session rule is defined in subclause 4.1.4.3.

