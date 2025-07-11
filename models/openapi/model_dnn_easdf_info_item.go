/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DnnEasdfInfoItem - Set of parameters supported by EASDF for a given DNN
type DnnEasdfInfoItem struct {

	Dnn DnnSmfInfoItemDnn `json:"dnn"`

	DnaiList []string `json:"dnaiList,omitempty"`
}
