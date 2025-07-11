/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// HssInfo - Information of an HSS NF Instance
type HssInfo struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	ImsiRanges []ImsiRange `json:"imsiRanges,omitempty"`

	ImsPrivateIdentityRanges []IdentityRange `json:"imsPrivateIdentityRanges,omitempty"`

	ImsPublicIdentityRanges []IdentityRange `json:"imsPublicIdentityRanges,omitempty"`

	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	HssDiameterAddress NetworkNodeDiameterAddress `json:"hssDiameterAddress,omitempty"`

	AdditionalDiamAddresses []NetworkNodeDiameterAddress `json:"additionalDiamAddresses,omitempty"`
}
