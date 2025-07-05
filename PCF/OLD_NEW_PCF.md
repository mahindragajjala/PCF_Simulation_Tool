https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.r1tr9i8jfchm


In the context of 5G networks, the terms NewPcf and OldPcf generally refer to Policy Control Functions (PCFs) that are involved during the handover or transition of policy control in the 5G Core network. These terms typically come into play when discussing:
Session Continuity: When a UE (User Equipment) moves from one network area to another, such as between slices, regions, or nodes, there might be a need to transfer the policy control from one PCF instance to another.


Policy Context Relocation: When the policy context (rules, QoS parameters, session data, etc.) is transferred from the OldPcf (the PCF currently handling the session) to the NewPcf (the PCF that will take over control), this ensures continuity of services and adherence to network policies.


Here’s what they represent:
1. OldPcf
Definition: The PCF instance currently managing the policy control for a particular session or UE.
Role: It holds all the current session-related policy contexts and might transfer this information to the NewPcf during a handover or network optimization.
Context: Typically active before the handover or relocation process begins.
2. NewPcf
Definition: The PCF instance that is expected to take over policy control for the session or UE after a handover or relocation.
Role: It initializes, updates, or reconfigures policy contexts based on the information received from the OldPcf.
Context: Becomes the primary PCF post-handover.

Use Cases in 5G:
Handover Between Network Slices:


When a UE moves from one network slice to another, the OldPcf associated with the first slice hands over the policy control to the NewPcf in the new slice.
Geographical Movement:


When the UE moves geographically, requiring it to change the controlling PCF due to the deployment of regional PCFs.
Load Balancing:


The network may decide to shift the policy control from an overloaded OldPcf to a less loaded NewPcf.
Service-Specific Optimization:


For certain services or QoS levels, the control may shift to a NewPcf optimized for those specific policies.

Call Flow (Simplified):
Trigger: A mobility or slice change event (e.g., handover or slice switching) occurs, and the PCF relocation is triggered.
Policy Transfer:
The OldPcf communicates with the NewPcf, transferring the policy context via N7 signaling.
The policy context includes UE-specific policies, session rules, QoS configurations, and any active PCC (Policy and Charging Control) rules.
Validation:
The NewPcf validates and adapts the policy context to the new network conditions, if necessary.
Activation:
The NewPcf becomes the active PCF, and the session proceeds under its control.
Acknowledgment:
The OldPcf is notified, and resources or contexts associated with the session may be released.

Example (Protocol Perspective):
The interaction typically happens over N7 (PCF-SMF) and other 5G Core interfaces. Here's a Go-like pseudocode to conceptualize the context transfer:
type PolicyContext struct {
    UEID        string
    QoSRules    []QoSRule
    ChargingRules []ChargingRule
}

// OldPcf sends policy context to NewPcf
func transferPolicyContext(oldPcf *PcfInstance, newPcf *PcfInstance, ueID string) error {
    context := oldPcf.GetPolicyContext(ueID)
    if context == nil {
        return errors.New("Policy context not found")
    }
    
    err := newPcf.ApplyPolicyContext(context)
    if err != nil {
        return fmt.Errorf("Failed to transfer policy context: %v", err)
    }

    oldPcf.ReleasePolicyContext(ueID)
    return nil
}


Key Interfaces and Protocols:
N7 Interface: Between the PCF and SMF, used for policy control signaling.
HTTP/2-based API: PCFs use RESTful HTTP/2 APIs for communication in 5G.
Let me know if you'd like more details about the architecture or real-world implementation!

