package api

import (
	conf "gcwguide/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(c *gin.Context, r conf.JsonOutput) {
	c.JSON(http.StatusOK, r)
	return
}
