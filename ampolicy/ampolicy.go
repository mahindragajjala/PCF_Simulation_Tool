package ampolicy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptrace"
	ampolicy "pcfsimulation/structures"
	"strings"

	"github.com/gin-gonic/gin"
)

func Ampolicy() {
	request := ampolicy.PolicyAssociationRequest{
		NotificationUri:   "https://amf1.nf.example.com/notify",
		AltNotifIpv4Addrs: []string{"192.0.2.1", "192.0.2.2"},
		AltNotifIpv6Addrs: []string{"2001:db8::1", "2001:db8::2"},
		Supi:              "imsi-2089300007487",
		Gpsi:              "msisdn-1234567890",
		AccessType:        ampolicy.AccessType3GPP,
		Pei:               "imei-35678901234567",
		UserLoc: &ampolicy.UserLocation{
			Tai: &ampolicy.Tai{
				PlmnId: &ampolicy.NetworkId{
					Mcc: "208",
					Mnc: "93",
				},
				Tac: "0001",
			},
			Ecgi: &ampolicy.Ecgi{
				PlmnId: &ampolicy.NetworkId{
					Mcc: "208",
					Mnc: "93",
				},
				Ecgi: "00112233",
			},
			Ncgi: &ampolicy.Ncgi{
				PlmnId: &ampolicy.NetworkId{
					Mcc: "208",
					Mnc: "93",
				},
				NrCellId: "44556677",
			},
			UeLocation: "12.9716,77.5946", // GPS-style
		},
		TimeZone: "Asia/Kolkata",
		ServingPlmn: &ampolicy.NetworkId{
			Mcc: "208",
			Mnc: "93",
		},
		RatType:  ampolicy.RatTypeNR,
		GroupIds: []string{"group-1", "group-2"},
		ServAreaRes: &ampolicy.ServiceAreaRestriction{
			RestrictionType: "ALLOWED_AREAS",
			Areas:           []string{"area-101", "area-102"},
		},
		Rfsp: 64,
		Guami: &ampolicy.Guami{
			PlmnId: &ampolicy.NetworkId{
				Mcc: "208",
				Mnc: "93",
			},
			AmfId: "0af1",
		},
		ServiveName: "npcf-am-policy-control",
		TraceReq: &ampolicy.TraceData{
			TraceRef:   "ref123",
			TraceDepth: "MEDIUM",
			NeTypeList: "AMF,SMF",
			EventList:  []string{"LOCATION_UPDATE", "SESSION_ESTABLISHMENT"},
		},
		SuppFeat: "01",
	}

	fmt.Printf("PolicyAssociationRequest: %+v\n", request)

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	req, _ := http.NewRequest("POST", "http://192.168.144.27:8012/npcf-am-policy-control/v1/policies", bytes.NewReader(jsonData))

	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Println("Got Connection:", info.Conn.RemoteAddr())
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			fmt.Println("Wrote request:", info)
		},
		GotFirstResponseByte: func() {
			fmt.Println("Got first response byte")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

// Helper function to split comma-separated string and trim spaces
func splitAndTrim(input string) []string {
	if input == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// For ui
func AmpolicyFromUI(c *gin.Context) {
	var req ampolicy.PolicyAssociationRequest

	// Basic fields
	req.NotificationUri = c.PostForm("notificationUri")
	req.AltNotifIpv4Addrs = splitAndTrim(c.PostForm("altNotifIpv4Addrs"))
	req.AltNotifIpv6Addrs = splitAndTrim(c.PostForm("altNotifIpv6Addrs"))
	req.Supi = c.PostForm("supi")
	req.Gpsi = c.PostForm("gpsi")
	req.AccessType = ampolicy.AccessType(c.PostForm("accessType"))
	req.Pei = c.PostForm("pei")
	req.TimeZone = c.PostForm("timeZone")
	req.RatType = ampolicy.RatType(c.PostForm("ratType"))
	req.GroupIds = splitAndTrim(c.PostForm("groupIds"))
	req.SuppFeat = c.PostForm("suppFeat")

	// Serving PLMN
	req.ServingPlmn = &ampolicy.NetworkId{
		Mcc: c.PostForm("servingPlmnMcc"),
		Mnc: c.PostForm("servingPlmnMnc"),
	}

	// GUAMI
	req.Guami = &ampolicy.Guami{
		PlmnId: &ampolicy.NetworkId{
			Mcc: c.PostForm("guamiMcc"),
			Mnc: c.PostForm("guamiMnc"),
		},
		AmfId: c.PostForm("amfId"),
	}

	// User Location
	req.UserLoc = &ampolicy.UserLocation{
		UeLocation: c.PostForm("ueLocation"),
		Tai: &ampolicy.Tai{
			PlmnId: &ampolicy.NetworkId{
				Mcc: c.PostForm("taiMcc"),
				Mnc: c.PostForm("taiMnc"),
			},
			Tac: c.PostForm("taiTac"),
		},
	}

	// Service Area Restriction
	req.ServAreaRes = &ampolicy.ServiceAreaRestriction{
		RestrictionType: c.PostForm("restrictionType"),
		Areas:           splitAndTrim(c.PostForm("areas")),
	}

	// Optional Trace Data
	if traceRef := c.PostForm("traceRef"); traceRef != "" {
		req.TraceReq = &ampolicy.TraceData{
			TraceRef:   traceRef,
			TraceDepth: c.PostForm("traceDepth"),
			NeTypeList: c.PostForm("neTypeList"),
			EventList:  splitAndTrim(c.PostForm("eventList")),
		}
	}

	// Validate required fields
	if req.Supi == "" || req.NotificationUri == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields: supi and notificationUri",
		})
		return
	}

	// Log request
	fmt.Printf("Received Policy Association Request:\n%+v\n", req)

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Policy Association Request received successfully!",
	})
}
