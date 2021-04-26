package api

import (
	conf "gcapi/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func response(c *gin.Context, r conf.JsonOutput) {
	c.JSON(http.StatusOK, r)
	return
}

func getAuthorization(c *gin.Context) string {
	return strings.ReplaceAll(c.GetHeader("authorization"),"Bearer ","")
}
