/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UeContextInSmsfData struct {

	SmsfInfo3GppAccess SmsfInfo `json:"smsfInfo3GppAccess,omitempty"`

	SmsfInfoNon3GppAccess SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty"`
}
