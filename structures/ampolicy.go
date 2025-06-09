package structure

// AccessType represents 3GPP access type.
type AccessType string

const (
	AccessType3GPP    AccessType = "3GPP_ACCESS"
	AccessTypeNon3GPP AccessType = "NON_3GPP_ACCESS"
)

// RatType represents Radio Access Technology type.
type RatType string

const (
	RatTypeNR      RatType = "NR"
	RatTypeEUTRA   RatType = "EUTRA"
	RatTypeUTRA    RatType = "UTRA"
	RatTypeGERAN   RatType = "GERAN"
	RatTypeWLAN    RatType = "WLAN"
	RatTypeVirtual RatType = "VIRTUAL"
)

// UserLocation contains location-related information.
type UserLocation struct {
	Tai        *Tai   `json:"tai,omitempty" yaml:"tai" bson:"tai" mapstructure:"Tai" form:"tai"`
	Ecgi       *Ecgi  `json:"ecgi,omitempty" yaml:"ecgi" bson:"ecgi" mapstructure:"Ecgi" form:"tai"`
	Ncgi       *Ncgi  `json:"ncgi,omitempty" yaml:"ncgi" bson:"ncgi" mapstructure:"Ncgi" form:"tai"`
	UeLocation string `json:"ueLocation,omitempty" yaml:"ueLocation" bson:"ueLocation" mapstructure:"UeLocation"`
}

// Tai - Tracking Area Identity
type Tai struct {
	PlmnId *NetworkId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	Tac    string     `json:"tac" yaml:"tac" bson:"tac"`
}

// Ecgi - E-UTRAN Cell Global Identifier
type Ecgi struct {
	PlmnId *NetworkId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	Ecgi   string     `json:"ecgi" yaml:"ecgi" bson:"ecgi"`
}

// Ncgi - NR Cell Global Identifier
type Ncgi struct {
	PlmnId   *NetworkId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	NrCellId string     `json:"nrCellId" yaml:"nrCellId" bson:"nrCellId"`
}

// NetworkId represents PLMN ID
type NetworkId struct {
	Mcc string `json:"mcc" yaml:"mcc" bson:"mcc"`
	Mnc string `json:"mnc" yaml:"mnc" bson:"mnc"`
}

// ServiceAreaRestriction defines area restrictions
type ServiceAreaRestriction struct {
	RestrictionType string   `json:"restrictionType" yaml:"restrictionType" bson:"restrictionType"`
	Areas           []string `json:"areas" yaml:"areas" bson:"areas"`
}

// Guami - Globally Unique AMF Identifier
type Guami struct {
	PlmnId *NetworkId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	AmfId  string     `json:"amfId" yaml:"amfId" bson:"amfId"`
}

// TraceData contains trace configuration for diagnostics.
type TraceData struct {
	TraceRef         string   `json:"traceRef" yaml:"traceRef" bson:"traceRef"`
	TraceDepth       string   `json:"traceDepth" yaml:"traceDepth" bson:"traceDepth"`
	NeTypeList       string   `json:"neTypeList" yaml:"neTypeList" bson:"neTypeList"`
	EventList        []string `json:"eventList" yaml:"eventList" bson:"eventList"`
	CollectionEntity string   `json:"collectionEntity" yaml:"collectionEntity" bson:"collectionEntity"`
}

type PolicyAssociationRequest struct {
	NotificationUri   string                  `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri" mapstructure:"NotificationUri" form:"notificationUri"`
	AltNotifIpv4Addrs []string                `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs" mapstructure:"AltNotifIpv4Addrs" form:"altNotifIpv4Addrs"`
	AltNotifIpv6Addrs []string                `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs" mapstructure:"AltNotifIpv6Addrs" form:"altNotifIpv6Addrs"`
	Supi              string                  `json:"supi" yaml:"supi" bson:"supi" mapstructure:"Supi" form:"supi"`
	Gpsi              string                  `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi" form:"gpsi"`
	AccessType        AccessType              `json:"accessType,omitempty" yaml:"accessType" bson:"accessType" mapstructure:"AccessType" form:"accessType"`
	Pei               string                  `json:"pei,omitempty" yaml:"pei" bson:"pei" mapstructure:"Pei" form:"pei"`
	UserLoc           *UserLocation           `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc" mapstructure:"UserLoc" form:"userLoc"`
	TimeZone          string                  `json:"timeZone,omitempty" yaml:"timeZone" bson:"timeZone" mapstructure:"TimeZone" form:"timeZone"`
	ServingPlmn       *NetworkId              `json:"servingPlmn,omitempty" yaml:"servingPlmn" bson:"servingPlmn" mapstructure:"ServingPlmn" form:"servingPlmn"`
	RatType           RatType                 `json:"ratType,omitempty" yaml:"ratType" bson:"ratType" mapstructure:"RatType" form:"ratType"`
	GroupIds          []string                `json:"groupIds,omitempty" yaml:"groupIds" bson:"groupIds" mapstructure:"GroupIds" form:"groupIds"`
	ServAreaRes       *ServiceAreaRestriction `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes" mapstructure:"ServAreaRes" form:"servAreaRes"`
	Rfsp              int32                   `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp" mapstructure:"Rfsp" form:"rfsp"`
	Guami             *Guami                  `json:"guami,omitempty" yaml:"guami" bson:"guami" mapstructure:"Guami" form:"guami"`
	ServiveName       string                  `json:"serviveName,omitempty" yaml:"serviveName" bson:"serviveName" mapstructure:"ServiveName" form:"serviveName"`
	TraceReq          *TraceData              `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq" mapstructure:"TraceReq" form:"traceReq"`
	SuppFeat          string                  `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat" form:"suppFeat"`
}
