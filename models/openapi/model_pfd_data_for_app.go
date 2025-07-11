/*
 * Nnef_PFDmanagement Sevice API
 *
 * Packet Flow Description Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type PfdDataForApp struct {
	ApplicationId string       `json:"applicationId" yaml:"applicationId" bson:"applicationId" mapstructure:"ApplicationId"`
	Pfds          []PfdContent `json:"pfds" yaml:"pfds" bson:"pfds" mapstructure:"Pfds"`
	CachingTime   *time.Time   `json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime" mapstructure:"CachingTime"`
}
