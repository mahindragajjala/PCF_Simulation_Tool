/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TrustAfInfo - Information of a trusted AF Instance
type TrustAfInfo struct {

	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`

	AfEvents []AfEvent `json:"afEvents,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	InternalGroupId []string `json:"internalGroupId,omitempty"`

	MappingInd bool `json:"mappingInd,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}
