/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ServiceName - Service names known to NRF
type ServiceName struct {
	ServiceName string
}

// type ServiceName string

// // List of ServiceName
// const (
// 	ServiceName_NNRF_NFM                  ServiceName = "nnrf-nfm"
// 	ServiceName_NNRF_DISC                 ServiceName = "nnrf-disc"
// 	ServiceName_NUDM_SDM                  ServiceName = "nudm-sdm"
// 	ServiceName_NUDM_UECM                 ServiceName = "nudm-uecm"
// 	ServiceName_NUDM_UEAU                 ServiceName = "nudm-ueau"
// 	ServiceName_NUDM_EE                   ServiceName = "nudm-ee"
// 	ServiceName_NUDM_PP                   ServiceName = "nudm-pp"
// 	ServiceName_NAMF_COMM                 ServiceName = "namf-comm"
// 	ServiceName_NAMF_EVTS                 ServiceName = "namf-evts"
// 	ServiceName_NAMF_MT                   ServiceName = "namf-mt"
// 	ServiceName_NAMF_LOC                  ServiceName = "namf-loc"
// 	ServiceName_NSMF_PDUSESSION           ServiceName = "nsmf-pdusession"
// 	ServiceName_NSMF_EVENT_EXPOSURE       ServiceName = "nsmf-event-exposure"
// 	ServiceName_NAUSF_AUTH                ServiceName = "nausf-auth"
// 	ServiceName_NAUSF_SORPROTECTION       ServiceName = "nausf-sorprotection"
// 	ServiceName_NAUSF_UPUPROTECTION       ServiceName = "nausf-upuprotection"
// 	ServiceName_NNEF_PFDMANAGEMENT        ServiceName = "nnef-pfdmanagement"
// 	ServiceName_NPCF_AM_POLICY_CONTROL    ServiceName = "npcf-am-policy-control"
// 	ServiceName_NPCF_SMPOLICYCONTROL      ServiceName = "npcf-smpolicycontrol"
// 	ServiceName_NPCF_POLICYAUTHORIZATION  ServiceName = "npcf-policyauthorization"
// 	ServiceName_NPCF_BDTPOLICYCONTROL     ServiceName = "npcf-bdtpolicycontrol"
// 	ServiceName_NPCF_EVENTEXPOSURE        ServiceName = "npcf-eventexposure"
// 	ServiceName_NPCF_UE_POLICY_CONTROL    ServiceName = "npcf-ue-policy-control"
// 	ServiceName_NSMSF_SMS                 ServiceName = "nsmsf-sms"
// 	ServiceName_NNSSF_NSSELECTION         ServiceName = "nnssf-nsselection"
// 	ServiceName_NNSSF_NSSAIAVAILABILITY   ServiceName = "nnssf-nssaiavailability"
// 	ServiceName_NUDR_DR                   ServiceName = "nudr-dr"
// 	ServiceName_NLMF_LOC                  ServiceName = "nlmf-loc"
// 	ServiceName_N5G_EIR_EIC               ServiceName = "n5g-eir-eic"
// 	ServiceName_NBSF_MANAGEMENT           ServiceName = "nbsf-management"
// 	ServiceName_NCHF_SPENDINGLIMITCONTROL ServiceName = "nchf-spendinglimitcontrol"
// 	ServiceName_NCHF_CONVERGEDCHARGING    ServiceName = "nchf-convergedcharging"
// 	ServiceName_NNWDAF_EVENTSSUBSCRIPTION ServiceName = "nnwdaf-eventssubscription"
// 	ServiceName_NNWDAF_ANALYTICSINFO      ServiceName = "nnwdaf-analyticsinfo"
// )
