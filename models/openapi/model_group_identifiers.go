/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GroupIdentifiers struct {

	ExtGroupId string `json:"extGroupId,omitempty" validate:"regexp=^extgroupid-[^@]+@[^@]+$"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	IntGroupId string `json:"intGroupId,omitempty" validate:"regexp=^[A-Fa-f0-9]{8}-[0-9]{3}-[0-9]{2,3}-([A-Fa-f0-9][A-Fa-f0-9]){1,10}$"`

	UeIdList []UeId `json:"ueIdList,omitempty"`
}
