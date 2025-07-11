/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AfRequestAuthorization - AF Request Authorization
type AfRequestAuthorization struct {

	GptpAllowedInfoList []GptpAllowedInfo `json:"gptpAllowedInfoList,omitempty"`

	AstiAllowedInfo AstiAllowedInfo `json:"astiAllowedInfo,omitempty"`
}
