/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UnrelatedClass struct {

	DefaultUnrelatedClass DefaultUnrelatedClass `json:"defaultUnrelatedClass"`

	ExternalUnrelatedClass ExternalUnrelatedClass `json:"externalUnrelatedClass,omitempty"`

	ServiceTypeUnrelatedClasses []ServiceTypeUnrelatedClass `json:"serviceTypeUnrelatedClasses,omitempty"`
}
