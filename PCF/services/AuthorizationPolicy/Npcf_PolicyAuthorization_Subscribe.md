https://docs.google.com/document/d/1Zetks3YhpaQtLLzyElvU3efc-3FUTTuiJVA7TiOnUGY/edit?tab=t.rmg5c2d311i3

Npcf_PolicyAuthorization_Subscribe service operation

The Npcf_PolicyAuthorization_Subscribe service operation enables NF service consumers handling of subscription to events for the existing application session context. Subscription to events shall be created:


- within the application session context establishment procedure by invoking the
Npcf_PolicyAuthorization_Create service operation, as described in subclause 4.2.2; or
- within the application session context modification procedure by invoking the
Npcf_PolicyAuthorization_Update service operation, as described in subclause 4.2.3; or
- by invoking the Npcf_PolicyAuthorization_Subscribe service operation for the existing application session context, as described in subclause 4.2.6.2.

The following procedures using the Npcf_PolicyAuthorization_Subscribe service operation is supported:
- Handling of subscription to events for the existing application session context.
- Initial subscription to events without provisioning of service information.
- Subscription to usage monitoring of sponsored data connectivity.
- Request of access network information.
- Subscription to notification of signalling path status.
- Modification of Subscription to Service Data Flow QoS Monitoring Information.
