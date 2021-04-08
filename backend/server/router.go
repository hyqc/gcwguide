package server

import (
	"gcwguide/api"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", api.WebSiteList)
}
