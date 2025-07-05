https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.eng6dknz0zjb

Binding mechanism



The binding mechanism is the procedure that associates a  service data flow (defined in a PCC rule by means of the SDF template), to the QoS Flow deemed to transport the service data flow. 

For service data flows belonging to AF sessions, the binding mechanism shall also associate the AF session information with the QoS Flow that is selected to carry the
service data flow.

NOTE 1: The relation between AF sessions and rules depends only on the operator configuration. 

An AF session can be covered by one or more PCC rules, if applicable (e.g. one rule per media component of an IMS session).

NOTE 2: The PCF may authorize dynamic PCC rules for service data flows without a corresponding AF session.

The binding mechanism includes three steps:

1. Session binding.
2 PCC rule authorization and
3. QoS Flow binding.
