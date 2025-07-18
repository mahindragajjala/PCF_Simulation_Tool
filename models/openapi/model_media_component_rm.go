/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This data type is defined in the same way as the MediaComponent data type, but with the OpenAPI nullable property set to true
type MediaComponentRm struct {
	// Contains an AF application identifier.
	AfAppId   string                  `json:"afAppId,omitempty" yaml:"afAppId" bson:"afAppId" mapstructure:"AfAppId"`
	AfRoutReq *AfRoutingRequirementRm `json:"afRoutReq,omitempty" yaml:"afRoutReq" bson:"afRoutReq" mapstructure:"AfRoutReq"`
	// Represents the content version of some content.
	ContVer     int32                          `json:"contVer,omitempty" yaml:"contVer" bson:"contVer" mapstructure:"ContVer"`
	Codecs      []string                       `json:"codecs,omitempty" yaml:"codecs" bson:"codecs" mapstructure:"Codecs"`
	FStatus     FlowStatus                     `json:"fStatus,omitempty" yaml:"fStatus" bson:"fStatus" mapstructure:"FStatus"`
	MarBwDl     string                         `json:"marBwDl,omitempty" yaml:"marBwDl" bson:"marBwDl" mapstructure:"MarBwDl"`
	MarBwUl     string                         `json:"marBwUl,omitempty" yaml:"marBwUl" bson:"marBwUl" mapstructure:"MarBwUl"`
	MedCompN    int32                          `json:"medCompN" yaml:"medCompN" bson:"medCompN" mapstructure:"MedCompN"`
	MedSubComps map[string]MediaSubComponentRm `json:"medSubComps,omitempty" yaml:"medSubComps" bson:"medSubComps" mapstructure:"MedSubComps"`
	MedType     MediaType                      `json:"medType,omitempty" yaml:"medType" bson:"medType" mapstructure:"MedType"`
	MirBwDl     string                         `json:"mirBwDl,omitempty" yaml:"mirBwDl" bson:"mirBwDl" mapstructure:"MirBwDl"`
	MirBwUl     string                         `json:"mirBwUl,omitempty" yaml:"mirBwUl" bson:"mirBwUl" mapstructure:"MirBwUl"`
	ResPrio     ReservPriority                 `json:"resPrio,omitempty" yaml:"resPrio" bson:"resPrio" mapstructure:"ResPrio"`
}
