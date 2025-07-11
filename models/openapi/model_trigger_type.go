/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TriggerType string

// List of TriggerType
const (
	TriggerType_QUOTA_THRESHOLD                                  TriggerType = "QUOTA_THRESHOLD"
	TriggerType_QHT                                              TriggerType = "QHT"
	TriggerType_FINAL                                            TriggerType = "FINAL"
	TriggerType_QUOTA_EXHAUSTED                                  TriggerType = "QUOTA_EXHAUSTED"
	TriggerType_VALIDITY_TIME                                    TriggerType = "VALIDITY_TIME"
	TriggerType_OTHER_QUOTA_TYPE                                 TriggerType = "OTHER_QUOTA_TYPE"
	TriggerType_FORCED_REAUTHORISATION                           TriggerType = "FORCED_REAUTHORISATION"
	TriggerType_UNUSED_QUOTA_TIMER                               TriggerType = "UNUSED_QUOTA_TIMER"
	TriggerType_ABNORMAL_RELEASE                                 TriggerType = "ABNORMAL_RELEASE"
	TriggerType_QOS_CHANGE                                       TriggerType = "QOS_CHANGE"
	TriggerType_VOLUME_LIMIT                                     TriggerType = "VOLUME_LIMIT"
	TriggerType_TIME_LIMIT                                       TriggerType = "TIME_LIMIT"
	TriggerType_PLMN_CHANGE                                      TriggerType = "PLMN_CHANGE"
	TriggerType_USER_LOCATION_CHANGE                             TriggerType = "USER_LOCATION_CHANGE"
	TriggerType_RAT_CHANGE                                       TriggerType = "RAT_CHANGE"
	TriggerType_UE_TIMEZONE_CHANGE                               TriggerType = "UE_TIMEZONE_CHANGE"
	TriggerType_TARIFF_TIME_CHANGE                               TriggerType = "TARIFF_TIME_CHANGE"
	TriggerType_MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS     TriggerType = "MAX_NUMBER_OF_CHANGES_IN CHARGING_CONDITIONS"
	TriggerType_MANAGEMENT_INTERVENTION                          TriggerType = "MANAGEMENT_INTERVENTION"
	TriggerType_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA TriggerType = "CHANGE_OF_UE_PRESENCE_IN PRESENCE_REPORTING_AREA"
	TriggerType_CHANGE_OF_3_GPP_PS_DATA_OFF_STATUS               TriggerType = "CHANGE_OF_3GPP_PS_DATA_OFF_STATUS"
	TriggerType_SERVING_NODE_CHANGE                              TriggerType = "SERVING_NODE_CHANGE"
	TriggerType_REMOVAL_OF_UPF                                   TriggerType = "REMOVAL_OF_UPF"
	TriggerType_ADDITION_OF_UPF                                  TriggerType = "ADDITION_OF_UPF"
)
