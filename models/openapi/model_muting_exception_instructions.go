/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MutingExceptionInstructions - Indicates to an Event producer NF instructions for the subscription and stored events when an exception (e.g. full buffer) occurs at the Event producer NF while the event is muted. 
type MutingExceptionInstructions struct {

	BufferedNotifs BufferedNotificationsAction `json:"bufferedNotifs,omitempty"`

	Subscription SubscriptionAction `json:"subscription,omitempty"`
}
