/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ChangeType - Indicates the type of change to be performed.
// type ChangeType struct {
// }

type ChangeType string

// List of ChangeType
const (
	ChangeType_ADD     ChangeType = "ADD"
	ChangeType_MOVE    ChangeType = "MOVE"
	ChangeType_REMOVE  ChangeType = "REMOVE"
	ChangeType_REPLACE ChangeType = "REPLACE"
)
