package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type applicationJSON struct {
	ContentType string `header:"Content-type" binding:"required,eq=application/json"`
}

func validateMiddleware(c *gin.Context) {
	var body map[string]interface{}
	rawData := getBody(c.Request)
	bodyError := json.Unmarshal(rawData, &body)

	headerError := c.ShouldBindHeader(&applicationJSON{})

	if headerError == nil && bodyError == nil && len(body) > 0 {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
}
