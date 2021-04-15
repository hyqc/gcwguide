package server

import (
	"gcapi/api"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	websites := r.Group("/websites")
	{
		websites.POST("/list", api.WebSiteList)
		websites.POST("/add", api.WebSiteAdd)
		websites.POST("/update", api.WebSiteUpdate)
		websites.POST("/delete", api.WebSiteDelete)
	}
}
