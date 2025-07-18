/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProseServiceAuth - Indicates whether the UE is authorized to use related services. 
type ProseServiceAuth struct {

	ProseDirectDiscoveryAuth UeAuth `json:"proseDirectDiscoveryAuth,omitempty"`

	ProseDirectCommunicationAuth UeAuth `json:"proseDirectCommunicationAuth,omitempty"`

	ProseL2RelayAuth UeAuth `json:"proseL2RelayAuth,omitempty"`

	ProseL3RelayAuth UeAuth `json:"proseL3RelayAuth,omitempty"`

	ProseL2RemoteAuth UeAuth `json:"proseL2RemoteAuth,omitempty"`

	ProseL3RemoteAuth UeAuth `json:"proseL3RemoteAuth,omitempty"`

	ProseMultipathComL2RemoteAuth UeAuth `json:"proseMultipathComL2RemoteAuth,omitempty"`

	ProseL2UeRelayAuth UeAuth `json:"proseL2UeRelayAuth,omitempty"`

	ProseL3UeRelayAuth UeAuth `json:"proseL3UeRelayAuth,omitempty"`

	ProseL2EndAuth UeAuth `json:"proseL2EndAuth,omitempty"`

	ProseL3EndAuth UeAuth `json:"proseL3EndAuth,omitempty"`
}
