package openapi

type AmbrRm struct {
    Uplink   string `json:"uplink" yaml:"uplink" bson:"uplink" mapstructure:"Uplink"`
    Downlink string `json:"downlink" yaml:"downlink" bson:"downlink" mapstructure:"Downlink"`
}
