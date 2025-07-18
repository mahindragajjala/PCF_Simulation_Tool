/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NfGroupListCond - Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {

	ConditionType string `json:"conditionType"`

	NfType string `json:"nfType"`

	NfGroupIdList []string `json:"nfGroupIdList"`
}
