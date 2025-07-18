/*
Nudm_UECM

Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the SmfRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfRegistration{}

// SmfRegistration This datatype contains a complete set of mandatory information relevant to an SMF  serving the UE. 
type SmfRegistration struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmfInstanceId string `json:"smfInstanceId"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	SmfSetId *string `json:"smfSetId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId"`
	SingleNssai Snssai `json:"singleNssai"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	EmergencyServices *bool `json:"emergencyServices,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PcscfRestorationCallbackUri *string `json:"pcscfRestorationCallbackUri,omitempty"`
	PlmnId PlmnId `json:"plmnId"`
	// Fully Qualified Domain Name
	PgwFqdn *string `json:"pgwFqdn,omitempty" validate:"regexp=^([0-9A-Za-z]([-0-9A-Za-z]{0,61}[0-9A-Za-z])?\\\\.)+[A-Za-z]{2,63}\\\\.?$"`
	PgwIpAddr NullableIpAddress `json:"pgwIpAddr,omitempty"`
	EpdgInd *bool `json:"epdgInd,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DeregCallbackUri *string `json:"deregCallbackUri,omitempty"`
	RegistrationReason *RegistrationReason `json:"registrationReason,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId *string `json:"pcfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri *string `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	UdrRestartInd *bool `json:"udrRestartInd,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty"`
	PduSessionReActivationRequired *bool `json:"pduSessionReActivationRequired,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	StaleCheckCallbackUri *string `json:"staleCheckCallbackUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	UdmStaleCheckCallbackUri *string `json:"udmStaleCheckCallbackUri,omitempty"`
	WildcardInd *bool `json:"wildcardInd,omitempty"`
}

type _SmfRegistration SmfRegistration

// NewSmfRegistration instantiates a new SmfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfRegistration(smfInstanceId string, pduSessionId int32, singleNssai Snssai, plmnId PlmnId) *SmfRegistration {
	this := SmfRegistration{}
	this.SmfInstanceId = smfInstanceId
	this.PduSessionId = pduSessionId
	this.SingleNssai = singleNssai
	this.PlmnId = plmnId
	var epdgInd bool = false
	this.EpdgInd = &epdgInd
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	var pduSessionReActivationRequired bool = false
	this.PduSessionReActivationRequired = &pduSessionReActivationRequired
	return &this
}

// NewSmfRegistrationWithDefaults instantiates a new SmfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfRegistrationWithDefaults() *SmfRegistration {
	this := SmfRegistration{}
	var epdgInd bool = false
	this.EpdgInd = &epdgInd
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	var pduSessionReActivationRequired bool = false
	this.PduSessionReActivationRequired = &pduSessionReActivationRequired
	return &this
}

// GetSmfInstanceId returns the SmfInstanceId field value
func (o *SmfRegistration) GetSmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmfInstanceId
}

// GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetSmfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmfInstanceId, true
}

// SetSmfInstanceId sets field value
func (o *SmfRegistration) SetSmfInstanceId(v string) {
	o.SmfInstanceId = v
}

// GetSmfSetId returns the SmfSetId field value if set, zero value otherwise.
func (o *SmfRegistration) GetSmfSetId() string {
	if o == nil || IsNil(o.SmfSetId) {
		var ret string
		return ret
	}
	return *o.SmfSetId
}

// GetSmfSetIdOk returns a tuple with the SmfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetSmfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmfSetId) {
		return nil, false
	}
	return o.SmfSetId, true
}

// HasSmfSetId returns a boolean if a field has been set.
func (o *SmfRegistration) HasSmfSetId() bool {
	if o != nil && !IsNil(o.SmfSetId) {
		return true
	}

	return false
}

// SetSmfSetId gets a reference to the given string and assigns it to the SmfSetId field.
func (o *SmfRegistration) SetSmfSetId(v string) {
	o.SmfSetId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmfRegistration) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmfRegistration) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmfRegistration) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetPduSessionId returns the PduSessionId field value
func (o *SmfRegistration) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *SmfRegistration) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetSingleNssai returns the SingleNssai field value
func (o *SmfRegistration) GetSingleNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleNssai, true
}

// SetSingleNssai sets field value
func (o *SmfRegistration) SetSingleNssai(v Snssai) {
	o.SingleNssai = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SmfRegistration) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SmfRegistration) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SmfRegistration) SetDnn(v string) {
	o.Dnn = &v
}

// GetEmergencyServices returns the EmergencyServices field value if set, zero value otherwise.
func (o *SmfRegistration) GetEmergencyServices() bool {
	if o == nil || IsNil(o.EmergencyServices) {
		var ret bool
		return ret
	}
	return *o.EmergencyServices
}

// GetEmergencyServicesOk returns a tuple with the EmergencyServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetEmergencyServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.EmergencyServices) {
		return nil, false
	}
	return o.EmergencyServices, true
}

// HasEmergencyServices returns a boolean if a field has been set.
func (o *SmfRegistration) HasEmergencyServices() bool {
	if o != nil && !IsNil(o.EmergencyServices) {
		return true
	}

	return false
}

// SetEmergencyServices gets a reference to the given bool and assigns it to the EmergencyServices field.
func (o *SmfRegistration) SetEmergencyServices(v bool) {
	o.EmergencyServices = &v
}

// GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field value if set, zero value otherwise.
func (o *SmfRegistration) GetPcscfRestorationCallbackUri() string {
	if o == nil || IsNil(o.PcscfRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.PcscfRestorationCallbackUri
}

// GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.PcscfRestorationCallbackUri) {
		return nil, false
	}
	return o.PcscfRestorationCallbackUri, true
}

// HasPcscfRestorationCallbackUri returns a boolean if a field has been set.
func (o *SmfRegistration) HasPcscfRestorationCallbackUri() bool {
	if o != nil && !IsNil(o.PcscfRestorationCallbackUri) {
		return true
	}

	return false
}

// SetPcscfRestorationCallbackUri gets a reference to the given string and assigns it to the PcscfRestorationCallbackUri field.
func (o *SmfRegistration) SetPcscfRestorationCallbackUri(v string) {
	o.PcscfRestorationCallbackUri = &v
}

// GetPlmnId returns the PlmnId field value
func (o *SmfRegistration) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *SmfRegistration) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetPgwFqdn returns the PgwFqdn field value if set, zero value otherwise.
func (o *SmfRegistration) GetPgwFqdn() string {
	if o == nil || IsNil(o.PgwFqdn) {
		var ret string
		return ret
	}
	return *o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPgwFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PgwFqdn) {
		return nil, false
	}
	return o.PgwFqdn, true
}

// HasPgwFqdn returns a boolean if a field has been set.
func (o *SmfRegistration) HasPgwFqdn() bool {
	if o != nil && !IsNil(o.PgwFqdn) {
		return true
	}

	return false
}

// SetPgwFqdn gets a reference to the given string and assigns it to the PgwFqdn field.
func (o *SmfRegistration) SetPgwFqdn(v string) {
	o.PgwFqdn = &v
}

// GetPgwIpAddr returns the PgwIpAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmfRegistration) GetPgwIpAddr() IpAddress {
	if o == nil || IsNil(o.PgwIpAddr.Get()) {
		var ret IpAddress
		return ret
	}
	return *o.PgwIpAddr.Get()
}

// GetPgwIpAddrOk returns a tuple with the PgwIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmfRegistration) GetPgwIpAddrOk() (*IpAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PgwIpAddr.Get(), o.PgwIpAddr.IsSet()
}

// HasPgwIpAddr returns a boolean if a field has been set.
func (o *SmfRegistration) HasPgwIpAddr() bool {
	if o != nil && o.PgwIpAddr.IsSet() {
		return true
	}

	return false
}

// SetPgwIpAddr gets a reference to the given NullableIpAddress and assigns it to the PgwIpAddr field.
func (o *SmfRegistration) SetPgwIpAddr(v IpAddress) {
	o.PgwIpAddr.Set(&v)
}
// SetPgwIpAddrNil sets the value for PgwIpAddr to be an explicit nil
func (o *SmfRegistration) SetPgwIpAddrNil() {
	o.PgwIpAddr.Set(nil)
}

// UnsetPgwIpAddr ensures that no value is present for PgwIpAddr, not even an explicit nil
func (o *SmfRegistration) UnsetPgwIpAddr() {
	o.PgwIpAddr.Unset()
}

// GetEpdgInd returns the EpdgInd field value if set, zero value otherwise.
func (o *SmfRegistration) GetEpdgInd() bool {
	if o == nil || IsNil(o.EpdgInd) {
		var ret bool
		return ret
	}
	return *o.EpdgInd
}

// GetEpdgIndOk returns a tuple with the EpdgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetEpdgIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EpdgInd) {
		return nil, false
	}
	return o.EpdgInd, true
}

// HasEpdgInd returns a boolean if a field has been set.
func (o *SmfRegistration) HasEpdgInd() bool {
	if o != nil && !IsNil(o.EpdgInd) {
		return true
	}

	return false
}

// SetEpdgInd gets a reference to the given bool and assigns it to the EpdgInd field.
func (o *SmfRegistration) SetEpdgInd(v bool) {
	o.EpdgInd = &v
}

// GetDeregCallbackUri returns the DeregCallbackUri field value if set, zero value otherwise.
func (o *SmfRegistration) GetDeregCallbackUri() string {
	if o == nil || IsNil(o.DeregCallbackUri) {
		var ret string
		return ret
	}
	return *o.DeregCallbackUri
}

// GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetDeregCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.DeregCallbackUri) {
		return nil, false
	}
	return o.DeregCallbackUri, true
}

// HasDeregCallbackUri returns a boolean if a field has been set.
func (o *SmfRegistration) HasDeregCallbackUri() bool {
	if o != nil && !IsNil(o.DeregCallbackUri) {
		return true
	}

	return false
}

// SetDeregCallbackUri gets a reference to the given string and assigns it to the DeregCallbackUri field.
func (o *SmfRegistration) SetDeregCallbackUri(v string) {
	o.DeregCallbackUri = &v
}

// GetRegistrationReason returns the RegistrationReason field value if set, zero value otherwise.
func (o *SmfRegistration) GetRegistrationReason() RegistrationReason {
	if o == nil || IsNil(o.RegistrationReason) {
		var ret RegistrationReason
		return ret
	}
	return *o.RegistrationReason
}

// GetRegistrationReasonOk returns a tuple with the RegistrationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetRegistrationReasonOk() (*RegistrationReason, bool) {
	if o == nil || IsNil(o.RegistrationReason) {
		return nil, false
	}
	return o.RegistrationReason, true
}

// HasRegistrationReason returns a boolean if a field has been set.
func (o *SmfRegistration) HasRegistrationReason() bool {
	if o != nil && !IsNil(o.RegistrationReason) {
		return true
	}

	return false
}

// SetRegistrationReason gets a reference to the given RegistrationReason and assigns it to the RegistrationReason field.
func (o *SmfRegistration) SetRegistrationReason(v RegistrationReason) {
	o.RegistrationReason = &v
}

// GetRegistrationTime returns the RegistrationTime field value if set, zero value otherwise.
func (o *SmfRegistration) GetRegistrationTime() time.Time {
	if o == nil || IsNil(o.RegistrationTime) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationTime
}

// GetRegistrationTimeOk returns a tuple with the RegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetRegistrationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RegistrationTime) {
		return nil, false
	}
	return o.RegistrationTime, true
}

// HasRegistrationTime returns a boolean if a field has been set.
func (o *SmfRegistration) HasRegistrationTime() bool {
	if o != nil && !IsNil(o.RegistrationTime) {
		return true
	}

	return false
}

// SetRegistrationTime gets a reference to the given time.Time and assigns it to the RegistrationTime field.
func (o *SmfRegistration) SetRegistrationTime(v time.Time) {
	o.RegistrationTime = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *SmfRegistration) GetContextInfo() ContextInfo {
	if o == nil || IsNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || IsNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *SmfRegistration) HasContextInfo() bool {
	if o != nil && !IsNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *SmfRegistration) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *SmfRegistration) GetPcfId() string {
	if o == nil || IsNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *SmfRegistration) HasPcfId() bool {
	if o != nil && !IsNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *SmfRegistration) SetPcfId(v string) {
	o.PcfId = &v
}

// GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field value if set, zero value otherwise.
func (o *SmfRegistration) GetDataRestorationCallbackUri() string {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.DataRestorationCallbackUri
}

// GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetDataRestorationCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		return nil, false
	}
	return o.DataRestorationCallbackUri, true
}

// HasDataRestorationCallbackUri returns a boolean if a field has been set.
func (o *SmfRegistration) HasDataRestorationCallbackUri() bool {
	if o != nil && !IsNil(o.DataRestorationCallbackUri) {
		return true
	}

	return false
}

// SetDataRestorationCallbackUri gets a reference to the given string and assigns it to the DataRestorationCallbackUri field.
func (o *SmfRegistration) SetDataRestorationCallbackUri(v string) {
	o.DataRestorationCallbackUri = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *SmfRegistration) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *SmfRegistration) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *SmfRegistration) SetResetIds(v []string) {
	o.ResetIds = v
}

// GetUdrRestartInd returns the UdrRestartInd field value if set, zero value otherwise.
func (o *SmfRegistration) GetUdrRestartInd() bool {
	if o == nil || IsNil(o.UdrRestartInd) {
		var ret bool
		return ret
	}
	return *o.UdrRestartInd
}

// GetUdrRestartIndOk returns a tuple with the UdrRestartInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetUdrRestartIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UdrRestartInd) {
		return nil, false
	}
	return o.UdrRestartInd, true
}

// HasUdrRestartInd returns a boolean if a field has been set.
func (o *SmfRegistration) HasUdrRestartInd() bool {
	if o != nil && !IsNil(o.UdrRestartInd) {
		return true
	}

	return false
}

// SetUdrRestartInd gets a reference to the given bool and assigns it to the UdrRestartInd field.
func (o *SmfRegistration) SetUdrRestartInd(v bool) {
	o.UdrRestartInd = &v
}

// GetLastSynchronizationTime returns the LastSynchronizationTime field value if set, zero value otherwise.
func (o *SmfRegistration) GetLastSynchronizationTime() time.Time {
	if o == nil || IsNil(o.LastSynchronizationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastSynchronizationTime
}

// GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSynchronizationTime) {
		return nil, false
	}
	return o.LastSynchronizationTime, true
}

// HasLastSynchronizationTime returns a boolean if a field has been set.
func (o *SmfRegistration) HasLastSynchronizationTime() bool {
	if o != nil && !IsNil(o.LastSynchronizationTime) {
		return true
	}

	return false
}

// SetLastSynchronizationTime gets a reference to the given time.Time and assigns it to the LastSynchronizationTime field.
func (o *SmfRegistration) SetLastSynchronizationTime(v time.Time) {
	o.LastSynchronizationTime = &v
}

// GetPduSessionReActivationRequired returns the PduSessionReActivationRequired field value if set, zero value otherwise.
func (o *SmfRegistration) GetPduSessionReActivationRequired() bool {
	if o == nil || IsNil(o.PduSessionReActivationRequired) {
		var ret bool
		return ret
	}
	return *o.PduSessionReActivationRequired
}

// GetPduSessionReActivationRequiredOk returns a tuple with the PduSessionReActivationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetPduSessionReActivationRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PduSessionReActivationRequired) {
		return nil, false
	}
	return o.PduSessionReActivationRequired, true
}

// HasPduSessionReActivationRequired returns a boolean if a field has been set.
func (o *SmfRegistration) HasPduSessionReActivationRequired() bool {
	if o != nil && !IsNil(o.PduSessionReActivationRequired) {
		return true
	}

	return false
}

// SetPduSessionReActivationRequired gets a reference to the given bool and assigns it to the PduSessionReActivationRequired field.
func (o *SmfRegistration) SetPduSessionReActivationRequired(v bool) {
	o.PduSessionReActivationRequired = &v
}

// GetStaleCheckCallbackUri returns the StaleCheckCallbackUri field value if set, zero value otherwise.
func (o *SmfRegistration) GetStaleCheckCallbackUri() string {
	if o == nil || IsNil(o.StaleCheckCallbackUri) {
		var ret string
		return ret
	}
	return *o.StaleCheckCallbackUri
}

// GetStaleCheckCallbackUriOk returns a tuple with the StaleCheckCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetStaleCheckCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.StaleCheckCallbackUri) {
		return nil, false
	}
	return o.StaleCheckCallbackUri, true
}

// HasStaleCheckCallbackUri returns a boolean if a field has been set.
func (o *SmfRegistration) HasStaleCheckCallbackUri() bool {
	if o != nil && !IsNil(o.StaleCheckCallbackUri) {
		return true
	}

	return false
}

// SetStaleCheckCallbackUri gets a reference to the given string and assigns it to the StaleCheckCallbackUri field.
func (o *SmfRegistration) SetStaleCheckCallbackUri(v string) {
	o.StaleCheckCallbackUri = &v
}

// GetUdmStaleCheckCallbackUri returns the UdmStaleCheckCallbackUri field value if set, zero value otherwise.
func (o *SmfRegistration) GetUdmStaleCheckCallbackUri() string {
	if o == nil || IsNil(o.UdmStaleCheckCallbackUri) {
		var ret string
		return ret
	}
	return *o.UdmStaleCheckCallbackUri
}

// GetUdmStaleCheckCallbackUriOk returns a tuple with the UdmStaleCheckCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetUdmStaleCheckCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.UdmStaleCheckCallbackUri) {
		return nil, false
	}
	return o.UdmStaleCheckCallbackUri, true
}

// HasUdmStaleCheckCallbackUri returns a boolean if a field has been set.
func (o *SmfRegistration) HasUdmStaleCheckCallbackUri() bool {
	if o != nil && !IsNil(o.UdmStaleCheckCallbackUri) {
		return true
	}

	return false
}

// SetUdmStaleCheckCallbackUri gets a reference to the given string and assigns it to the UdmStaleCheckCallbackUri field.
func (o *SmfRegistration) SetUdmStaleCheckCallbackUri(v string) {
	o.UdmStaleCheckCallbackUri = &v
}

// GetWildcardInd returns the WildcardInd field value if set, zero value otherwise.
func (o *SmfRegistration) GetWildcardInd() bool {
	if o == nil || IsNil(o.WildcardInd) {
		var ret bool
		return ret
	}
	return *o.WildcardInd
}

// GetWildcardIndOk returns a tuple with the WildcardInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfRegistration) GetWildcardIndOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardInd) {
		return nil, false
	}
	return o.WildcardInd, true
}

// HasWildcardInd returns a boolean if a field has been set.
func (o *SmfRegistration) HasWildcardInd() bool {
	if o != nil && !IsNil(o.WildcardInd) {
		return true
	}

	return false
}

// SetWildcardInd gets a reference to the given bool and assigns it to the WildcardInd field.
func (o *SmfRegistration) SetWildcardInd(v bool) {
	o.WildcardInd = &v
}

func (o SmfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smfInstanceId"] = o.SmfInstanceId
	if !IsNil(o.SmfSetId) {
		toSerialize["smfSetId"] = o.SmfSetId
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	toSerialize["pduSessionId"] = o.PduSessionId
	toSerialize["singleNssai"] = o.SingleNssai
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.EmergencyServices) {
		toSerialize["emergencyServices"] = o.EmergencyServices
	}
	if !IsNil(o.PcscfRestorationCallbackUri) {
		toSerialize["pcscfRestorationCallbackUri"] = o.PcscfRestorationCallbackUri
	}
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.PgwFqdn) {
		toSerialize["pgwFqdn"] = o.PgwFqdn
	}
	if o.PgwIpAddr.IsSet() {
		toSerialize["pgwIpAddr"] = o.PgwIpAddr.Get()
	}
	if !IsNil(o.EpdgInd) {
		toSerialize["epdgInd"] = o.EpdgInd
	}
	if !IsNil(o.DeregCallbackUri) {
		toSerialize["deregCallbackUri"] = o.DeregCallbackUri
	}
	if !IsNil(o.RegistrationReason) {
		toSerialize["registrationReason"] = o.RegistrationReason
	}
	if !IsNil(o.RegistrationTime) {
		toSerialize["registrationTime"] = o.RegistrationTime
	}
	if !IsNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !IsNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !IsNil(o.DataRestorationCallbackUri) {
		toSerialize["dataRestorationCallbackUri"] = o.DataRestorationCallbackUri
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !IsNil(o.UdrRestartInd) {
		toSerialize["udrRestartInd"] = o.UdrRestartInd
	}
	if !IsNil(o.LastSynchronizationTime) {
		toSerialize["lastSynchronizationTime"] = o.LastSynchronizationTime
	}
	if !IsNil(o.PduSessionReActivationRequired) {
		toSerialize["pduSessionReActivationRequired"] = o.PduSessionReActivationRequired
	}
	if !IsNil(o.StaleCheckCallbackUri) {
		toSerialize["staleCheckCallbackUri"] = o.StaleCheckCallbackUri
	}
	if !IsNil(o.UdmStaleCheckCallbackUri) {
		toSerialize["udmStaleCheckCallbackUri"] = o.UdmStaleCheckCallbackUri
	}
	if !IsNil(o.WildcardInd) {
		toSerialize["wildcardInd"] = o.WildcardInd
	}
	return toSerialize, nil
}

func (o *SmfRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"smfInstanceId",
		"pduSessionId",
		"singleNssai",
		"plmnId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSmfRegistration := _SmfRegistration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmfRegistration)

	if err != nil {
		return err
	}

	*o = SmfRegistration(varSmfRegistration)

	return err
}

type NullableSmfRegistration struct {
	value *SmfRegistration
	isSet bool
}

func (v NullableSmfRegistration) Get() *SmfRegistration {
	return v.value
}

func (v *NullableSmfRegistration) Set(val *SmfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfRegistration(val *SmfRegistration) *NullableSmfRegistration {
	return &NullableSmfRegistration{value: val, isSet: true}
}

func (v NullableSmfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


