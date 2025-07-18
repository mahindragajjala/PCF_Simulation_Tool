/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to retrieve background data transfer policies from the PCF and to update the PCF with the background data transfer policy selected by the NF service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains service requirements for creation a new Individual BDT policy resource.
type BdtReqData struct {
	// Contains an identity of an application service provider.
	AspId      string           `json:"aspId" yaml:"aspId" bson:"aspId" mapstructure:"AspId"`
	DesTimeInt *TimeWindow      `json:"desTimeInt" yaml:"desTimeInt" bson:"desTimeInt" mapstructure:"DesTimeInt"`
	NwAreaInfo *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo" mapstructure:"NwAreaInfo"`
	// Indicates a number of UEs.
	NumOfUes int32           `json:"numOfUes" yaml:"numOfUes" bson:"numOfUes" mapstructure:"NumOfUes"`
	VolPerUe *UsageThreshold `json:"volPerUe" yaml:"volPerUe" bson:"volPerUe" mapstructure:"VolPerUe"`
	SuppFeat string          `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
}
