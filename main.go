package main

import (
	"pcfsimulation/ampolicy"
)

/*
Project Title: 5G PCF Simulation Tool
Tech Stack: Golang, Gin, REST APIs, HTTP/2, Docker, JSON, 3GPP TS 29.512

Developed a simulation tool that mimics the 5G Policy Control Function (PCF) using Golang and Gin framework. Implemented key Npcf API endpoints to emulate policy decision handling (QoS, slice selection, and SM policy control) as per 3GPP TS 29.512. Integrated HTTP/2 with TLS support and simulated policy behavior for SMF testing. Packaged using Docker for cloud-native deployment.
*/
func main() {
	/*
		router := gin.Default()

		// Serve static files (HTML, CSS)
		router.Static("/static", "./frontend")
		router.LoadHTMLFiles("frontend/index.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		router.POST("/policy", ampolicy.AmpolicyFromUI)
		router.Run(":8080")
	*/
	ampolicy.Ampolicy()
}
