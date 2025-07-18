/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RequestType string

// List of RequestType
const (
	RequestType_INITIAL_REQUEST                RequestType = "INITIAL_REQUEST"
	RequestType_EXISTING_PDU_SESSION           RequestType = "EXISTING_PDU_SESSION"
	RequestType_INITIAL_EMERGENCY_REQUEST      RequestType = "INITIAL_EMERGENCY_REQUEST"
	RequestType_EXISTING_EMERGENCY_PDU_SESSION RequestType = "EXISTING_EMERGENCY_PDU_SESSION"
)
