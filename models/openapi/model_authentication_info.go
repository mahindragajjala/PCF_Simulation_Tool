/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AuthenticationInfo struct {
	SupiOrSuci            string                 `json:"supiOrSuci" yaml:"supiOrSuci" bson:"supiOrSuci"`
	ServingNetworkName    string                 `json:"servingNetworkName" yaml:"servingNetworkName" bson:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty" yaml:"resynchronizationInfo" bson:"resynchronizationInfo"`
	TraceData             *TraceData             `json:"traceData,omitempty" yaml:"traceData" bson:"traceData"`
}
