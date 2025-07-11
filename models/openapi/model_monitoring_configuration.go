/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MonitoringConfiguration struct {

	EventType EventType `json:"eventType"`

	ImmediateFlag bool `json:"immediateFlag,omitempty"`

	LocationReportingConfiguration LocationReportingConfiguration `json:"locationReportingConfiguration,omitempty"`

	AssociationType AssociationType `json:"associationType,omitempty"`

	DatalinkReportCfg DatalinkReportingConfiguration `json:"datalinkReportCfg,omitempty"`

	LossConnectivityCfg LossConnectivityCfg `json:"lossConnectivityCfg,omitempty"`

	// indicating a time in seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	// indicating a time in seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	SuggestedPacketNumDl int32 `json:"suggestedPacketNumDl,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	SingleNssai Snssai `json:"singleNssai,omitempty"`

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	PduSessionStatusCfg PduSessionStatusCfg `json:"pduSessionStatusCfg,omitempty"`

	ReachabilityForSmsCfg ReachabilityForSmsConfiguration `json:"reachabilityForSmsCfg,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`

	AfId string `json:"afId,omitempty"`

	ReachabilityForDataCfg ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`

	IdleStatusInd bool `json:"idleStatusInd,omitempty"`

	MonitoringSuspension MonitoringSuspension `json:"monitoringSuspension,omitempty"`
}
