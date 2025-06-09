# ChangeOfSupiPeiAssociationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPei** | **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | 

## Methods

### NewChangeOfSupiPeiAssociationReport

`func NewChangeOfSupiPeiAssociationReport(newPei string, ) *ChangeOfSupiPeiAssociationReport`

NewChangeOfSupiPeiAssociationReport instantiates a new ChangeOfSupiPeiAssociationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeOfSupiPeiAssociationReportWithDefaults

`func NewChangeOfSupiPeiAssociationReportWithDefaults() *ChangeOfSupiPeiAssociationReport`

NewChangeOfSupiPeiAssociationReportWithDefaults instantiates a new ChangeOfSupiPeiAssociationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPei

`func (o *ChangeOfSupiPeiAssociationReport) GetNewPei() string`

GetNewPei returns the NewPei field if non-nil, zero value otherwise.

### GetNewPeiOk

`func (o *ChangeOfSupiPeiAssociationReport) GetNewPeiOk() (*string, bool)`

GetNewPeiOk returns a tuple with the NewPei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPei

`func (o *ChangeOfSupiPeiAssociationReport) SetNewPei(v string)`

SetNewPei sets NewPei field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


