/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MappingOfSnssai struct {
	ServingSnssai *Snssai `json:"servingSnssai" bson:"servingSnssai" yaml:"servingSnssai"`

	HomeSnssai *Snssai `json:"homeSnssai" bson:"homeSnssai" yaml:"homeSnssai"`
}
