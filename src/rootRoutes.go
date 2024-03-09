package main

import (
	"comacast_challenge/pkg/CcValidate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func serveHelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

type CCValidateRequest struct {
	CCNumber string `json:"ccNumber"`
}

func validateCCNumbers(c *gin.Context) {
	var cc CCValidateRequest
	if err := c.BindJSON(&cc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ccNumbers := CcValidate.NewCCNumber(cc.CCNumber)
	ccNumbers.Validate()
	if ccNumbers.ErrorMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "error": ccNumbers.ErrorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valid": true})
}

func RegisterRoot(r *gin.Engine) {
	r.GET("/", serveHelloWorld)
	r.POST("/", validateCCNumbers)
}
