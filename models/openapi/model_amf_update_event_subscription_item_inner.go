/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AmfUpdateEventSubscriptionItemInner struct {
	Op    string    `json:"op" bson:"op" `
	Path  string    `json:"path" bson:"path" `
	Value *AmfEvent `json:"value,omitempty" bson:"value" `
}
