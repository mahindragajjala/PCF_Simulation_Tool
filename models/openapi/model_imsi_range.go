/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ImsiRange - A range of IMSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type ImsiRange struct {

	Start string `json:"start,omitempty"`

	End string `json:"end,omitempty"`

	Pattern string `json:"pattern,omitempty"`
}
