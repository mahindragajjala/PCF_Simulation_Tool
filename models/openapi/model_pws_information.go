/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PwsInformation struct {
	MessageIdentifier int32          `json:"messageIdentifier"`
	SerialNumber      int32          `json:"serialNumber"`
	PwsContainer      *N2InfoContent `json:"pwsContainer"`
	SendRanResponse   bool           `json:"sendRanResponse,omitempty"`
	OmcId             string         `json:"omcId,omitempty"`
}
