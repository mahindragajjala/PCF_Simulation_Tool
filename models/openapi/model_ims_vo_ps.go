/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ImsVoPs - It represents the information indicating homogeneity of IMS Voice over PS Sessions support for the UE
// type ImsVoPs struct {
// }

type ImsVoPs string

// List of ImsVoPs
const (
	ImsVoPs_HOMOGENEOUS_SUPPORT        ImsVoPs = "HOMOGENEOUS_SUPPORT"
	ImsVoPs_HOMOGENEOUS_NON_SUPPORT    ImsVoPs = "HOMOGENEOUS_NON_SUPPORT"
	ImsVoPs_NON_HOMOGENEOUS_OR_UNKNOWN ImsVoPs = "NON_HOMOGENEOUS_OR_UNKNOWN"
)
