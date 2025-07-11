/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IpSmGwRegistration - This data type contains the IP-SW-GW routing information.
type IpSmGwRegistration struct {

	// This data type mentions International E.164 number of the SMSF; shall be present if the SMSF supports MAP. 
	IpSmGwMapAddress string `json:"ipSmGwMapAddress,omitempty"`

	IpSmGwDiameterAddress NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	IpsmgwIpv4 string `json:"ipsmgwIpv4,omitempty"`

	IpsmgwIpv6 Ipv6Addr `json:"ipsmgwIpv6,omitempty"`

	// Fully Qualified Domain Name
	IpsmgwFqdn string `json:"ipsmgwFqdn,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId,omitempty"`

	UnriIndicator bool `json:"unriIndicator,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`

	IpSmGwSbiSupInd bool `json:"ipSmGwSbiSupInd,omitempty"`
}
