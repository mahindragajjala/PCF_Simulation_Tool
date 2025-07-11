/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MbsServiceAreaInfo - MBS Service Area Information for location dependent MBS session
type MbsServiceAreaInfo struct {

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. 
	AreaSessionId int32 `json:"areaSessionId"`

	MbsServiceArea *MbsServiceArea `json:"mbsServiceArea"`
}
