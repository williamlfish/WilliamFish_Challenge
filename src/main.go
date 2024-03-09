package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const DefaultTemplatePath = "./templates/*"

func buildRouter(templatePath string) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob(templatePath)
	RegisterRoot(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := buildRouter(DefaultTemplatePath)
	r.Run(":3000")
}
