/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AuthorizedDefaultQos struct {
	Var5qi          int32 `json:"5qi,omitempty" yaml:"5qi" bson:"5qi" mapstructure:"Var5qi"`
	Arp             *Arp  `json:"arp,omitempty" yaml:"arp" bson:"arp" mapstructure:"Arp"`
	PriorityLevel   int32 `json:"priorityLevel,omitempty" yaml:"priorityLevel" bson:"priorityLevel" mapstructure:"PriorityLevel"`
	AverWindow      int32 `json:"averWindow,omitempty" yaml:"averWindow" bson:"averWindow" mapstructure:"AverWindow"`
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" yaml:"maxDataBurstVol" bson:"maxDataBurstVol" mapstructure:"MaxDataBurstVol"`
}
