/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EdrxParameters struct {

	RatType RatType `json:"ratType"`

	EdrxValue string `json:"edrxValue" validate:"regexp=^([0-1]{4})$"`
}
