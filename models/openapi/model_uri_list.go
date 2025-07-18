/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UriList - Represents a set of URIs following the 3GPP hypermedia format (containing a \"_links\" attribute). 
type UriList struct {

	// List of the URI of NF instances. It has two members whose names are item and self. The item attribute contains an array of URIs. 
	Links map[string]LinksValueSchema `json:"_links,omitempty"`

	TotalItemCount int32 `json:"totalItemCount,omitempty"`
}
