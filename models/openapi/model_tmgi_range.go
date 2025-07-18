/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TmgiRange - Range of TMGIs
type TmgiRange struct {

	MbsServiceIdStart string `json:"mbsServiceIdStart" validate:"regexp=^[A-Fa-f0-9]{6}$"`

	MbsServiceIdEnd string `json:"mbsServiceIdEnd" validate:"regexp=^[A-Fa-f0-9]{6}$"`

	PlmnId PlmnId `json:"plmnId"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty" validate:"regexp=^[A-Fa-f0-9]{11}$"`
}
