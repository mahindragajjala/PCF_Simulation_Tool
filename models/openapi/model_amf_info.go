/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AmfInfo - Information of an AMF NF Instance
type AmfInfo struct {

	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits). 
	AmfSetId string `json:"amfSetId" validate:"regexp=^[0-3][A-Fa-f0-9]{2}$"`

	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits) 
	AmfRegionId string `json:"amfRegionId" validate:"regexp=^[A-Fa-f0-9]{2}$"`

	GuamiList []Guami `json:"guamiList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	BackupInfoAmfFailure []Guami `json:"backupInfoAmfFailure,omitempty"`

	BackupInfoAmfRemoval []Guami `json:"backupInfoAmfRemoval,omitempty"`

	N2InterfaceAmfInfo *N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty"`

	AmfOnboardingCapability bool `json:"amfOnboardingCapability,omitempty"`

	HighLatencyCom bool `json:"highLatencyCom,omitempty"`
}
