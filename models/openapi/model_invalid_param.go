/*
 * NRF Bootstrapping
 *
 * NRF Bootstrapping.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InvalidParam - It contains an invalid parameter and a related description.
type InvalidParam struct {

	// If the invalid parameter is an attribute in a JSON body, this IE shall contain the  attribute's name and shall be encoded as a JSON Pointer. If the invalid parameter is  an HTTP header, this IE shall be formatted as the concatenation of the string \"header \"  plus the name of such header. If the invalid parameter is a query parameter, this IE  shall be formatted as the concatenation of the string \"query \" plus the name of such  query parameter. If the invalid parameter is a variable part in the path of a resource  URI, this IE shall contain the name of the variable, including the symbols \"{\" and \"}\"  used in OpenAPI specification as the notation to represent variable path segments. 
	Param string `json:"param"`

	// A human-readable reason, e.g. \"must be a positive integer\". In cases involving failed  operations in a PATCH request, the reason string should identify the operation that  failed using the operation's array index to assist in correlation of the invalid  parameter with the failed operation, e.g.\" Replacement value invalid for attribute  (failed operation index= 4)\" 
	Reason string `json:"reason,omitempty"`
}
