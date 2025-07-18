/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CommunicationFailure struct {
	NasReleaseCode string     `json:"nasReleaseCode,omitempty" bson:"nasReleaseCode" `
	RanReleaseCode *NgApCause `json:"ranReleaseCode,omitempty" bson:"ranReleaseCode" `
}
