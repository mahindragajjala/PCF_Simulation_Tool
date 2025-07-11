/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IpSmGwGuidance - Contains guidance information (e.g. minimum and recommended delivery times) of the IP-SM-GW 
type IpSmGwGuidance struct {

	MinDeliveryTime int32 `json:"minDeliveryTime"`

	RecommDeliveryTime int32 `json:"recommDeliveryTime"`
}
