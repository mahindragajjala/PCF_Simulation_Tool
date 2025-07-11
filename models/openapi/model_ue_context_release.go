/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UeContextRelease struct {
	Supi                string     `json:"supi,omitempty"`
	UnauthenticatedSupi bool       `json:"unauthenticatedSupi,omitempty"`
	NgapCause           *NgApCause `json:"ngapCause"`
}
