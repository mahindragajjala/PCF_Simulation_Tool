/*
 * NudmUEAU
 *
 * UDM UE Authentication Service. � 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AuthenticationVector struct {
	AvType AvType `json:"avType"`

	Rand string `json:"rand"`

	Xres string `json:"xres"`

	Autn string `json:"autn"`

	CkPrime string `json:"ckPrime"`

	IkPrime string `json:"ikPrime"`

	XresStar string `json:"xresStar"`

	Kausf string `json:"kausf"`
}
