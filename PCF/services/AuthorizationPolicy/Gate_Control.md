https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.xxtoazbl5bte

Gate control
This procedure is used by an AF to instruct the PCF about when the service data flow(s) are to be enabled or disabled for a PDU session.
The AF shall include in the HTTP POST request message described in subclause 4.2.2.2 the "fStatus" attribute for the flows to be enabled or disabled within the "medComponents" or "medSubComponents" attributes.
If a "medSubComponents" attribute contains a "flowUsage" attribute with the value "RTCP", then the IP Flows described by that media subcomponent shall be enabled in both directions irrespective of the value of the "fStatus" attribute of the corresponding media component.
As result of this action, the PCF shall set the appropriate gate status for the corresponding active PCC rule(s).
The PCF shall reply to the AF as described in subclause 4.2.2.2.
