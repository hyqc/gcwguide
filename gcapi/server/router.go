package server

import (
	"gcapi/api"
	conf "gcapi/config"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	websites := r.Group("/websites")
	{
		websites.POST("/list", api.WebSiteList)
		websites.POST("/add", api.WebSiteAdd)
		websites.POST("/update", api.WebSiteUpdate)
		websites.POST("/delete", api.WebSiteDelete)
		websites.GET("/groups", api.WebsiteGroups)
	}

	upload := r.Group("/upload")
	{
		upload.POST("/image", api.Image)
	}
}

func Static(r *gin.Engine)  {
	r.Static("/" + conf.App.Static.Static, "./" + conf.App.Static.Static )
}
