/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ExtSnssai - The sdRanges and wildcardSd attributes shall be exclusive from each other. If one of these attributes is present,  the sd attribute shall also be present and it shall contain one Slice Differentiator value within the range of SD  (if the sdRanges attribute is present) or with any value (if the wildcardSd attribute is present). 
type ExtSnssai struct {

	// Unsigned integer, within the range 0 to 255, representing the Slice/Service Type.  It indicates the expected Network Slice behaviour in terms of features and services. Values 0 to 127 correspond to the standardized SST range. Values 128 to 255 correspond  to the Operator-specific range. See clause 28.4.2 of 3GPP TS 23.003. Standardized values are defined in clause 5.15.2.2 of 3GPP TS 23.501.  
	Sst int32 `json:"sst"`

	// 3-octet string, representing the Slice Differentiator, in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the SD shall appear first in the string, and the character representing the 4 least significant bit of the SD shall appear last in the string. This is an optional parameter that complements the Slice/Service type(s) to allow to  differentiate amongst multiple Network Slices of the same Slice/Service type. This IE shall be absent if no SD value is associated with the SST. 
	Sd string `json:"sd,omitempty" validate:"regexp=^[A-Fa-f0-9]{6}$"`

	// When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type 
	SdRanges []SdRange `json:"sdRanges,omitempty"`

	// When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type. 
	WildcardSd bool `json:"wildcardSd,omitempty"`
}
