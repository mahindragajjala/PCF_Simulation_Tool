/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// NfService - Information of a given NF Service Instance; it is part of the NFProfile of an NF Instance discovered by the NRF
type NfService struct {
	ServiceInstanceId string `json:"serviceInstanceId"`

	ServiceName ServiceName `json:"serviceName"`

	Versions []NfServiceVersion `json:"versions"`

	//Scheme string `json:"scheme"`
	// Scheme UriSchemestring `json:"scheme"`
	Scheme UriScheme `json:"scheme"`

	NfServiceStatus NfServiceStatus `json:"nfServiceStatus"`

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty" validate:"regexp=^([0-9A-Za-z]([-0-9A-Za-z]{0,61}[0-9A-Za-z])?\\\\.)+[A-Za-z]{2,63}\\\\.?$"`

	// Fully Qualified Domain Name
	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty" validate:"regexp=^([0-9A-Za-z]([-0-9A-Za-z]{0,61}[0-9A-Za-z])?\\\\.)+[A-Za-z]{2,63}\\\\.?$"`

	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	ApiPrefix string `json:"apiPrefix,omitempty"`

	CallbackUriPrefixList []CallbackUriPrefixItem `json:"callbackUriPrefixList,omitempty"`

	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`

	AllowedSnpns []PlmnIdNid `json:"allowedSnpns,omitempty"`

	AllowedNfTypes []NfType `json:"allowedNfTypes,omitempty"`

	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`

	AllowedNssais []ExtSnssai `json:"allowedNssais,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	Load int32 `json:"load,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp time.Time `json:"loadTimeStamp,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`

	NfServiceSetIdList []string `json:"nfServiceSetIdList,omitempty"`

	SNssais []ExtSnssai `json:"sNssais,omitempty"`

	PerPlmnSnssaiList []PlmnSnssai `json:"perPlmnSnssaiList,omitempty"`

	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId string `json:"vendorId,omitempty" validate:"regexp=^[0-9]{6}$"`

	// The key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes
	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty"`

	Oauth2Required bool `json:"oauth2Required,omitempty"`

	// A map (list of key-value pairs) where NF Type serves as key
	AllowedOperationsPerNfType map[string][]string `json:"allowedOperationsPerNfType,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	AllowedOperationsPerNfInstance map[string][]string `json:"allowedOperationsPerNfInstance,omitempty"`

	AllowedOperationsPerNfInstanceOverrides bool `json:"allowedOperationsPerNfInstanceOverrides,omitempty"`

	// A map (list of key-value pairs) where a valid JSON pointer Id serves as key
	AllowedScopesRuleSet map[string]RuleSet `json:"allowedScopesRuleSet,omitempty"`

	SelectionConditions SelectionConditions `json:"selectionConditions,omitempty"`

	CanaryRelease bool `json:"canaryRelease,omitempty"`

	ExclusiveCanaryReleaseSelection bool `json:"exclusiveCanaryReleaseSelection,omitempty"`

	SharedServiceDataId string `json:"sharedServiceDataId,omitempty"`
}
