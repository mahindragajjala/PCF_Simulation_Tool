/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MutingNotificationsSettings - Indicates the Event producer NF settings to the Event consumer NF 
type MutingNotificationsSettings struct {

	MaxNoOfNotif int32 `json:"maxNoOfNotif,omitempty"`

	// indicating a time in seconds.
	DurationBufferedNotif int32 `json:"durationBufferedNotif,omitempty"`
}
