Npcf_PolicyAuthorization_Unsubscribe service operation

The Npcf_PolicyAuthorization_Unsubscribe service operation enables NF service consumers to remove subscription to all subscribed events for the existing application session context. Subscription to events shall be removed:

- by invoking the Npcf_PolicyAuthorization_Unsubscribe service operation for the existing application session context, as described in subclause 4.2.7.2; or
- within the application session context modification procedure by invoking the
Npcf_PolicyAuthorization_Update service operation, as described in subclause 4.2.3; or
- within the application session context termination procedure by invoking the Npcf_PolicyAuthorization_Delete service operation, as described in subclause 4.2.4.

The following procedure using the Npcf_PolicyAuthorization_Unsubscribe service operation is supported:
- Unsubscription to events.
