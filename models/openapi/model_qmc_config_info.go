/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// QmcConfigInfo - It contains the configuration information for signaling-based activation of the  Quality of Experience (QoE) Measurements Collection (QMC) functionality.  
type QmcConfigInfo struct {

	// String containing MCC (3 digits), MNC (2 or 3 digits)  and QMC ID (3 octets, encoded as 6 hexadecimal digits). 
	QoeReference string `json:"qoeReference" validate:"regexp=^[0-9]{3}-[0-9]{2,3}-[A-Fa-f0-9]{6}$"`

	ServiceType QoeServiceType `json:"serviceType,omitempty"`

	SliceScope []Snssai `json:"sliceScope,omitempty"`

	AreaScope QmcAreaScope `json:"areaScope,omitempty"`

	QoeCollectionEntityAddress *IpAddr `json:"qoeCollectionEntityAddress,omitempty"`

	QoeTarget QoeTarget `json:"qoeTarget,omitempty"`

	// String containing: - Trace Reference: MCC (3 digits), MNC (2 or 3 digits),    Trace ID (3 octets, encoded as 6 hexadecimal digits) - Trace Recording Session Reference (2 octets, encoded as 4 hexadecimal digits) 
	MdtAlignmentInfo *interface{} `json:"mdtAlignmentInfo,omitempty" validate:"regexp=^[0-9]{3}-[0-9]{2,3}-[A-Fa-f0-9]{6}-[A-Fa-f0-9]{4}$"`

	AvailableRanVisibleQoeMetrics []AvailableRanVisibleQoeMetric `json:"availableRanVisibleQoeMetrics,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	ContainerForAppLayerMeasConfig string `json:"containerForAppLayerMeasConfig,omitempty"`

	MbsCommunicationServiceType MbsServiceType `json:"mbsCommunicationServiceType,omitempty"`
}
