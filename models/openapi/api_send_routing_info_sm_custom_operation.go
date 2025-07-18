/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type SendRoutingInfoSMCustomOperationAPI struct {
}

// Post /nudm-uecm/v1/:ueId/registrations/send-routing-info-sm
// Retreive addressing information for SMS delivery 
func (api *SendRoutingInfoSMCustomOperationAPI) SendRoutingInfoSm(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

