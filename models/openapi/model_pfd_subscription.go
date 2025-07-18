/*
 * Nnef_PFDmanagement Sevice API
 *
 * Packet Flow Description Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PfdSubscription struct {
	ApplicationIds    []string `json:"applicationIds,omitempty" yaml:"applicationIds" bson:"applicationIds" mapstructure:"ApplicationIds"`
	NotifyUri         string   `json:"notifyUri" yaml:"notifyUri" bson:"notifyUri" mapstructure:"NotifyUri"`
	SupportedFeatures string   `json:"supportedFeatures" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
