/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Cause string

// List of Cause
const (
	Cause_REL_DUE_TO_HO                   Cause = "REL_DUE_TO_HO"
	Cause_EPS_FALLBACK                    Cause = "EPS_FALLBACK"
	Cause_REL_DUE_TO_UP_SEC               Cause = "REL_DUE_TO_UP_SEC"
	Cause_DNN_CONGESTION                  Cause = "DNN_CONGESTION"
	Cause_S_NSSAI_CONGESTION              Cause = "S-NSSAI_CONGESTION"
	Cause_REL_DUE_TO_REACTIVATION         Cause = "REL_DUE_TO_REACTIVATION"
	Cause__5_G_AN_NOT_RESPONDING          Cause = "5G_AN_NOT_RESPONDING"
	Cause_REL_DUE_TO_SLICE_NOT_AVAILABLE  Cause = "REL_DUE_TO_SLICE_NOT_AVAILABLE"
	Cause_REL_DUE_TO_DUPLICATE_SESSION_ID Cause = "REL_DUE_TO_DUPLICATE_SESSION_ID"
	Cause_PDU_SESSION_STATUS_MISMATCH     Cause = "PDU_SESSION_STATUS_MISMATCH"
	Cause_HO_FAILURE                      Cause = "HO_FAILURE"
)
