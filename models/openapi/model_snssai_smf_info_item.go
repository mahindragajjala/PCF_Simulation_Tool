/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SnssaiSmfInfoItem - Set of parameters supported by SMF for a given S-NSSAI
type SnssaiSmfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnSmfInfoList []DnnSmfInfoItem `json:"dnnSmfInfoList,omitempty"`

	DnnSmfInfoListId int32 `json:"dnnSmfInfoListId,omitempty"`
}
