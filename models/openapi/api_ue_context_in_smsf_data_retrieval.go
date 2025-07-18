/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type UEContextInSMSFDataRetrievalAPI struct {
}

// Get /nudm-sdm/v2/:supi/ue-context-in-smsf-data
// retrieve a UE's UE Context In SMSF Data 
func (api *UEContextInSMSFDataRetrievalAPI) GetUeCtxInSmsfData(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

