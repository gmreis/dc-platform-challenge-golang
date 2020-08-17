package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type applicationJSON struct {
	ContentType string `header:"Content-type" binding:"required,eq=application/json"`
}

func isArrayJSON(body []byte) bool {
	var array []interface{}
	parserError := json.Unmarshal(body, &array)
	if parserError == nil && len(array) > 0 {
		return true
	}

	return false
}

func isObjectJSON(body []byte) bool {
	var object map[string]interface{}
	parserError := json.Unmarshal(body, &object)
	if parserError == nil && len(object) > 0 {
		return true
	}

	return false
}

func validateMiddleware(c *gin.Context) {
	rawData := getBody(c.Request)
	headerError := c.ShouldBindHeader(&applicationJSON{})

	if headerError == nil && (isArrayJSON(rawData) || isObjectJSON(rawData)) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
}
