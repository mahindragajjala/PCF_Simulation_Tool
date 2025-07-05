
func (p *Processor) postAppSessCtxProcedure(appSessCtx *models.AppSessionContext) (*models.AppSessionContext,
	string, *models.ProblemDetails,
) {
// Receive the Application Session Context Request
	ascReqData := appSessCtx.AscReqData
	pcfSelf := p.Context()
//Handle BDT Policy Indication
	// Initial BDT policy indication(the only one which is not related to session)
	if ascReqData.BdtRefId != "" {
		if err := p.handleBDTPolicyInd(pcfSelf, appSessCtx); err != nil {
			problemDetail := util.GetProblemDetail(err.Error(), util.ERROR_REQUEST_PARAMETERS)
			return nil, "", &problemDetail
		}
		appSessID := fmt.Sprintf("BdtRefId-%s", ascReqData.BdtRefId)
		data := pcf_context.AppSessionData{
			AppSessionId:      appSessID,
			AppSessionContext: appSessCtx,
		}
		pcfSelf.AppSessionPool.Store(appSessID, &data)
		locationHeader := util.GetResourceUri(models.ServiceName_NPCF_POLICYAUTHORIZATION, appSessID)
		logger.PolicyAuthLog.Tracef("App Session Id[%s] Create", appSessID)
		return appSessCtx, locationHeader, nil
	}
//Step 3: Validate the Request Parameters
	if ascReqData.UeIpv4 == "" && ascReqData.UeIpv6 == "" && ascReqData.UeMac == "" {
		problemDetail := util.GetProblemDetail("Ue UeIpv4 and UeIpv6 and UeMac are all empty", util.ERROR_REQUEST_PARAMETERS)
		return nil, "", &problemDetail
	}
	if ascReqData.AfRoutReq != nil && ascReqData.Dnn == "" {
		problemDetail := util.GetProblemDetail("DNN shall be present", util.ERROR_REQUEST_PARAMETERS)
		return nil, "", &problemDetail
	}
	var smPolicy *pcf_context.UeSmPolicyData
//Step 4: Perform Session Binding with SMF
	if tempSmPolicy, err := pcfSelf.SessionBinding(ascReqData); err != nil {
		problemDetail := util.GetProblemDetail(fmt.Sprintf("Session Binding failed[%s]",
			err.Error()), util.PDU_SESSION_NOT_AVAILABLE)
		return nil, "", &problemDetail
	} else {
		smPolicy = tempSmPolicy
	}
	logger.PolicyAuthLog.Infof("Session Binding Success - UeIpv4[%s], UeIpv6[%s], UeMac[%s]",
		ascReqData.UeIpv4, ascReqData.UeIpv6, ascReqData.UeMac)
//Negotiate Supported Features
	ue := smPolicy.PcfUe
	updateSMpolicy := false
	var requestSuppFeat openapi.SupportedFeature
	if tempRequestSuppFeat, err := openapi.NewSupportedFeature(ascReqData.SuppFeat); err != nil {
		logger.PolicyAuthLog.Errorf(err.Error())
	} else {
		requestSuppFeat = tempRequestSuppFeat
	}

	nSuppFeat := pcfSelf.PcfSuppFeats[models.ServiceName_NPCF_POLICYAUTHORIZATION].NegotiateWith(requestSuppFeat).String()
	// InfluenceOnTrafficRouting = 1 in 29514 &  Traffic Steering Control support = 1 in 29512
	traffRoutSupp := util.CheckSuppFeat(nSuppFeat, 1) && util.CheckSuppFeat(smPolicy.PolicyDecision.SuppFeat, 1)
	relatedPccRuleIds := make(map[string]string)
//Handle PCC Rules for Media Components
	if ascReqData.MedComponents != nil {
		// Handle Pcc rules
		precedence := getAvailablePrecedence(smPolicy.PolicyDecision.PccRules)
		for _, medComp := range ascReqData.MedComponents {
			var pccRule *models.PccRule
			var appID string
			var routeReq *models.AfRoutingRequirement
			// TODO: use specific algorithm instead of default, details in subsclause 7.3.3 of TS 29513
			var var5qi int32 = 9
			if medComp.MedType != "" {
				var5qi = util.MediaTypeTo5qiMap[medComp.MedType]
			}

			if medComp.MedSubComps != nil {
				for _, medSubComp := range medComp.MedSubComps {
					if tempPccRule, problemDetail := handleMediaSubComponent(smPolicy,
						&medComp, &medSubComp, var5qi); problemDetail != nil {
						return nil, "", problemDetail
					} else {
						pccRule = tempPccRule
					}
					key := fmt.Sprintf("%d-%d", medComp.MedCompN, medSubComp.FNum)
					relatedPccRuleIds[key] = pccRule.PccRuleId
					updateSMpolicy = true
				}
				continue
			} else if medComp.AfAppId != "" {
				appID = medComp.AfAppId
				routeReq = medComp.AfRoutReq
			} else if ascReqData.AfAppId != "" {
				appID = ascReqData.AfAppId
				routeReq = ascReqData.AfRoutReq
			} else {
				problemDetail := util.GetProblemDetail("Media Component needs flows of subComp or afAppId",
					util.REQUESTED_SERVICE_NOT_AUTHORIZED)
				return nil, "", &problemDetail
			}

			// Find pccRule by AfAppId, otherwise create a new pcc rule
			pccRule = util.GetPccRuleByAfAppId(smPolicy.PolicyDecision.PccRules, appID)
			if pccRule == nil {
				pccRule = util.CreatePccRule(smPolicy.PccRuleIdGenerator, precedence, nil, appID)
				// Set QoS Data
				// TODO: use real arp
				qosData := util.CreateQosData(smPolicy.PccRuleIdGenerator, var5qi, 8)
				if var5qi <= 4 {
					// update Qos Data according to request BitRate
					var ul, dl bool
					qosData, ul, dl = updateQosInMedComp(qosData, &medComp)
					if problemDetails := modifyRemainBitRate(smPolicy, &qosData, ul, dl); problemDetails != nil {
						return nil, "", problemDetails
					}
				}
				util.SetPccRuleRelatedData(smPolicy.PolicyDecision, pccRule, nil, &qosData, nil, nil)
				smPolicy.PccRuleIdGenerator++
				if precedence < Precedence_Maximum {
					precedence++
				}
			} else {
				// update pccRule's qos
				var qosData models.QosData
				for _, qosID := range pccRule.RefQosData {
					qosData = *smPolicy.PolicyDecision.QosDecs[qosID]
					if qosData.Var5qi == var5qi && qosData.Var5qi <= 4 {
						var ul, dl bool
						qosData, ul, dl = updateQosInMedComp(*smPolicy.PolicyDecision.QosDecs[qosID], &medComp)
						if problemDetails := modifyRemainBitRate(smPolicy, &qosData, ul, dl); problemDetails != nil {
							return nil, "", problemDetails
						}
						smPolicy.PolicyDecision.QosDecs[qosData.QosId] = &qosData
					}
				}
			}
//Traffic Routing and AF Influence
			// Initial provisioning of traffic routing information
			if traffRoutSupp {
		pccRule = provisioningOfTrafficRoutingInfo(smPolicy, appID, routeReq, medComp.FStatus)
			}
			key := fmt.Sprintf("%d", medComp.MedCompN)
			relatedPccRuleIds[key] = pccRule.PccRuleId
			updateSMpolicy = true
		}
	} else if ascReqData.AfAppId != "" {
		// Initial provisioning of traffic routing information
		if ascReqData.AfRoutReq != nil && traffRoutSupp {
			logger.PolicyAuthLog.Infof("AF influence on Traffic Routing - AppId[%s]", ascReqData.AfAppId)
			pccRule := provisioningOfTrafficRoutingInfo(smPolicy, ascReqData.AfAppId, ascReqData.AfRoutReq, "")
			key := fmt.Sprintf("appID-%s", ascReqData.AfAppId)
			relatedPccRuleIds[key] = pccRule.PccRuleId
			updateSMpolicy = true
		} else {
			problemDetail := util.GetProblemDetail("Traffic routing not supported", util.REQUESTED_SERVICE_NOT_AUTHORIZED)
			return nil, "", &problemDetail
		}
	} else {
		problemDetail := util.GetProblemDetail("AF Request need AfAppId or Media Component to match Service Data Flow",
			util.ERROR_REQUEST_PARAMETERS)
		return nil, "", &problemDetail
	}
//Handle Event Subscriptions
	// Event Subscription
	eventSubs := make(map[models.AfEvent]models.AfNotifMethod)
	if ascReqData.EvSubsc != nil {
		for _, subs := range ascReqData.EvSubsc.Events {
			if subs.NotifMethod == "" {
				// default value "EVENT_DETECTION"
				subs.NotifMethod = models.AfNotifMethod_EVENT_DETECTION
			}
			eventSubs[subs.Event] = subs.NotifMethod
			var trig models.PolicyControlRequestTrigger
			switch subs.Event {
			case models.AfEvent_ACCESS_TYPE_CHANGE:
				trig = models.PolicyControlRequestTrigger_AC_TY_CH
			// case models.AfEvent_FAILED_RESOURCES_ALLOCATION:
			// 	// Subscription to Service Data Flow Deactivation
			// 	trig = models.PolicyControlRequestTrigger_RES_RELEASE
			case models.AfEvent_PLMN_CHG:
				trig = models.PolicyControlRequestTrigger_PLMN_CH
			case models.AfEvent_QOS_NOTIF:
				// Subscriptions to Service Data Flow QoS notification control
				for _, pccRuleID := range relatedPccRuleIds {
					pccRule := smPolicy.PolicyDecision.PccRules[pccRuleID]
					for _, qosID := range pccRule.RefQosData {
						qosData := smPolicy.PolicyDecision.QosDecs[qosID]
						qosData.Qnc = true
						smPolicy.PolicyDecision.QosDecs[qosID] = qosData
					}
				}
				trig = models.PolicyControlRequestTrigger_QOS_NOTIF
			case models.AfEvent_SUCCESSFUL_RESOURCES_ALLOCATION:
				// Subscription to resources allocation outcome
				trig = models.PolicyControlRequestTrigger_SUCC_RES_ALLO
			case models.AfEvent_USAGE_REPORT:
				trig = models.PolicyControlRequestTrigger_US_RE
			default:
				logger.PolicyAuthLog.Warn("AF Event is unknown")
				continue
			}
			if !util.CheckPolicyControlReqTrig(smPolicy.PolicyDecision.PolicyCtrlReqTriggers, trig) {
				smPolicy.PolicyDecision.PolicyCtrlReqTriggers = append(smPolicy.PolicyDecision.PolicyCtrlReqTriggers, trig)
				updateSMpolicy = true
			}
		}
	}

	// Initial provisioning of sponsored connectivity information
	if ascReqData.AspId != "" && ascReqData.SponId != "" {
		// SponsoredConnectivity = 2 in 29514 &  SponsoredConnectivity support = 12 in 29512
		supp := util.CheckSuppFeat(nSuppFeat, 2) && util.CheckSuppFeat(smPolicy.PolicyDecision.SuppFeat, 12)
		if !supp {
			problemDetail := util.GetProblemDetail("Sponsored Connectivity not supported", util.REQUESTED_SERVICE_NOT_AUTHORIZED)
			return nil, "", &problemDetail
		}
		umID := util.GetUmId(ascReqData.AspId, ascReqData.SponId)
		var umData *models.UsageMonitoringData
		if tempUmData, err := extractUmData(umID, eventSubs, ascReqData.EvSubsc.UsgThres); err != nil {
			problemDetail := util.GetProblemDetail(err.Error(), util.REQUESTED_SERVICE_NOT_AUTHORIZED)
			return nil, "", &problemDetail
		} else {
			umData = tempUmData
		}
		if err := handleSponsoredConnectivityInformation(smPolicy, relatedPccRuleIds, ascReqData.AspId,
			ascReqData.SponId, ascReqData.SponStatus, umData, &updateSMpolicy); err != nil {
			problemDetail := util.GetProblemDetail(err.Error(), util.REQUESTED_SERVICE_NOT_AUTHORIZED)
			return nil, "", &problemDetail
		}
	}

//Allocate App Session ID and Store Data
	// Allocate App Session Id
	appSessID := ue.AllocUeAppSessionId(pcfSelf)
	appSessCtx.AscRespData = &models.AppSessionContextRespData{
		SuppFeat: nSuppFeat,
	}
	// Associate App Session to SMPolicy
	smPolicy.AppSessions[appSessID] = true
	data := pcf_context.AppSessionData{
		AppSessionId:      appSessID,
		AppSessionContext: appSessCtx,
		SmPolicyData:      smPolicy,
	}
	if len(relatedPccRuleIds) > 0 {
		data.RelatedPccRuleIds = relatedPccRuleIds
		data.PccRuleIdMapToCompId = reverseStringMap(relatedPccRuleIds)
	}
	appSessCtx.EvsNotif = &models.EventsNotification{}
	// Set Event Subsciption related Data
	if len(eventSubs) > 0 {
		data.Events = eventSubs
		data.EventUri = ascReqData.EvSubsc.NotifUri
		if _, exist := eventSubs[models.AfEvent_PLMN_CHG]; exist {
			afNotif := models.AfEventNotification{
				Event: models.AfEvent_PLMN_CHG,
			}
			appSessCtx.EvsNotif.EvNotifs = append(appSessCtx.EvsNotif.EvNotifs, afNotif)
			plmnID := smPolicy.PolicyContext.ServingNetwork
			if plmnID != nil {
				appSessCtx.EvsNotif.PlmnId = &models.PlmnId{
					Mcc: plmnID.Mcc,
					Mnc: plmnID.Mnc,
				}
			}
		}
		if _, exist := eventSubs[models.AfEvent_ACCESS_TYPE_CHANGE]; exist {
			afNotif := models.AfEventNotification{
				Event: models.AfEvent_ACCESS_TYPE_CHANGE,
			}
			appSessCtx.EvsNotif.EvNotifs = append(appSessCtx.EvsNotif.EvNotifs, afNotif)
			appSessCtx.EvsNotif.AccessType = smPolicy.PolicyContext.AccessType
			appSessCtx.EvsNotif.RatType = smPolicy.PolicyContext.RatType
		}
	}
	if appSessCtx.EvsNotif.EvNotifs == nil {
		appSessCtx.EvsNotif = nil
	}
	pcfSelf.AppSessionPool.Store(appSessID, &data)
	locationHeader := util.GetResourceUri(models.ServiceName_NPCF_POLICYAUTHORIZATION, appSessID)
	logger.PolicyAuthLog.Infof("App Session Id[%s] Create", appSessID)
	// Send Notification to SMF
//Notify SMF of Policy Changes 
	if updateSMpolicy {
		smPolicyID := fmt.Sprintf("%s-%d", ue.Supi, smPolicy.PolicyContext.PduSessionId)
		notification := models.SmPolicyNotification{
			ResourceUri:      util.GetResourceUri(models.ServiceName_NPCF_SMPOLICYCONTROL, smPolicyID),
			SmPolicyDecision: smPolicy.PolicyDecision,
		}
		go p.SendSMPolicyUpdateNotification(smPolicy.PolicyContext.NotificationUri, &notification)
	}
	return appSessCtx, locationHeader, nil
}

The provided Go function postAppSessCtxProcedure is part of a 5G Policy Control Function (PCF) implementation. This function processes an Application Session Context (AppSessionContext) request from an Application Function (AF), which plays a crucial role in managing policy authorization and traffic influence in 5G networks. Below is a structured breakdown of how the function works, mapping its flow to 5G call flows and protocol interactions.

1. Function Overview
The function is responsible for:
Handling Application Session Context (ASC) requests from the AF.
Associating the session with a Session Management (SM) Policy.
Creating Policy Control Rules (PCC Rules).
Handling event subscriptions.
Allocating App Session IDs and storing session data.

2. Step-by-Step Execution in 5G Context
The function follows the standard PCF-AF communication flow in 5G System (5GS), aligning with 3GPP TS 29.512 (Policy Authorization) and TS 29.514 (PCF-SMF interactions).
Step 1: Receive the Application Session Context Request
ascReqData := appSessCtx.AscReqData
pcfSelf := p.Context()

ascReqData holds request data from the Application Function (AF).
pcfSelf refers to the PCF instance.
5G Call Flow Equivalent:
AF → PCF: Npcf_PolicyAuthorization_Create Request (HTTP POST)
This request is received by PCF to establish an application session policy.

Step 2: Handle BDT Policy Indication (If Present)
if ascReqData.BdtRefId != "" {
    if err := p.handleBDTPolicyInd(pcfSelf, appSessCtx); err != nil {
        problemDetail := util.GetProblemDetail(err.Error(), util.ERROR_REQUEST_PARAMETERS)
        return nil, "", &problemDetail
    }
}

If the request contains a Bulk Data Transfer (BDT) Policy, it is processed separately.
5G Call Flow Equivalent:
AF → PCF: BDT Policy Request (N5)
The PCF stores and processes BDT policies for large-scale data transfers.

Step 3: Validate the Request Parameters
if ascReqData.UeIpv4 == "" && ascReqData.UeIpv6 == "" && ascReqData.UeMac == "" {
    problemDetail := util.GetProblemDetail("UeIpv4, UeIpv6, and UeMac are all empty", util.ERROR_REQUEST_PARAMETERS)
    return nil, "", &problemDetail
}

The PCF ensures that at least one of the UE identifiers (IPv4, IPv6, MAC) is present.
5G Call Flow Equivalent:
AF → PCF: Application Session Establishment (N5)
The PCF checks whether the UE is reachable via available IP or MAC.

Step 4: Perform Session Binding with SMF
var smPolicy *pcf_context.UeSmPolicyData
if tempSmPolicy, err := pcfSelf.SessionBinding(ascReqData); err != nil {
    problemDetail := util.GetProblemDetail(fmt.Sprintf("Session Binding failed[%s]", err.Error()), util.PDU_SESSION_NOT_AVAILABLE)
    return nil, "", &problemDetail
} else {
    smPolicy = tempSmPolicy
}

The PCF retrieves the session policy by binding to an existing SMF session.
If no session is found, the request fails.
5G Call Flow Equivalent:
PCF ↔ SMF: Npcf_SMPolicyControl_Create (N7 Interface)
PCF binds the AF request to the Session Management (SM) Policy.

Step 5: Negotiate Supported Features
requestSuppFeat := openapi.SupportedFeature{}
if tempRequestSuppFeat, err := openapi.NewSupportedFeature(ascReqData.SuppFeat); err != nil {
    logger.PolicyAuthLog.Errorf(err.Error())
} else {
    requestSuppFeat = tempRequestSuppFeat
}

The PCF negotiates supported features between itself and the AF.
5G Call Flow Equivalent:
AF ↔ PCF: Feature Negotiation
The PCF checks what capabilities (e.g., traffic routing, QoS support) are available.

Step 6: Handle PCC Rules for Media Components
if ascReqData.MedComponents != nil {
    precedence := getAvailablePrecedence(smPolicy.PolicyDecision.PccRules)
    for _, medComp := range ascReqData.MedComponents {
        var5qi := 9
        if medComp.MedType != "" {
            var5qi = util.MediaTypeTo5qiMap[medComp.MedType]
        }

The function creates and assigns PCC rules for each media component.
The QoS (5QI) for each media flow is determined.
5G Call Flow Equivalent:
PCF → SMF: Npcf_SMPolicyControl_Update
PCF instructs SMF to enforce QoS and traffic rules.

Step 7: Traffic Routing and AF Influence
if traffRoutSupp {
    pccRule = provisioningOfTrafficRoutingInfo(smPolicy, appID, routeReq, medComp.FStatus)
}

If Traffic Routing Control is supported, the PCF provisions it.
5G Call Flow Equivalent:
PCF ↔ SMF: Traffic Routing Request (N7)
PCF sends rules to SMF to enforce traffic steering.

Step 8: Handle Event Subscriptions
eventSubs := make(map[models.AfEvent]models.AfNotifMethod)
if ascReqData.EvSubsc != nil {
    for _, subs := range ascReqData.EvSubsc.Events {
        eventSubs[subs.Event] = subs.NotifMethod
    }
}

The PCF subscribes to UE events like:
Access Type Change
PLMN Change
QoS Notification
5G Call Flow Equivalent:
PCF → AF: Npcf_PolicyAuthorization_Update
PCF sends event notifications to AF.

Step 9: Allocate App Session ID and Store Data
appSessID := ue.AllocUeAppSessionId(pcfSelf)
data := pcf_context.AppSessionData{
    AppSessionId:      appSessID,
    AppSessionContext: appSessCtx,
    SmPolicyData:      smPolicy,
}
pcfSelf.AppSessionPool.Store(appSessID, &data)

A unique session ID is allocated.
The session data is stored in PCF's session pool.
5G Call Flow Equivalent:
PCF → AF: Response with AppSessionID
The PCF acknowledges the session creation.

Step 10: Notify SMF of Policy Changes (If Required)
if updateSMpolicy {
    smPolicyID := fmt.Sprintf("%s-%d", ue.Supi, smPolicy.PolicyContext.PduSessionId)
    notification := models.SmPolicyNotification{
        ResourceUri:      util.GetResourceUri(models.ServiceName_NPCF_SMPOLICYCONTROL, smPolicyID),
        SmPolicyDecision: smPolicy.PolicyDecision,
    }
    go p.SendSMPolicyUpdateNotification(smPolicy.PolicyContext.NotificationUri, &notification)
}

If policies have changed, PCF sends a notification to SMF.
5G Call Flow Equivalent:
PCF → SMF: Npcf_SMPolicyControl_Update
PCF triggers an update to enforce new policy rules.

Conclusion
This function is a core part of the 5G PCF module, implementing Application Session Policy Control by:
Handling AF trequests.
Validating session parameters.
Binding sessions to SMF policies.
Enforcing QoS, PCC rules, and traffic steering.
Managing event subscriptions.
Allocating session IDs and notifying AF/SMF.
This aligns with 5G 3GPP specifications for PCF interactions over N5 and N7 interfaces. 🚀

