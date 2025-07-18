/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MbsServiceArea - MBS Service Area
type MbsServiceArea struct {

	// List of NR cell Ids
	NcgiList []NcgiTai `json:"ncgiList,omitempty"`

	// List of tracking area Ids
	TaiList []Tai `json:"taiList,omitempty"`
}
