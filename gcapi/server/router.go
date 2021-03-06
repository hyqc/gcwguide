package server

import (
	"gcapi/api"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	website(r)
	upload(r)
	auth(r)
}

func website(r *gin.Engine) {
	website := r.Group("/website")
	{
		website.GET("/list", api.WebSiteList)
		website.POST("/add", api.WebSiteAdd)
		website.POST("/update", api.WebSiteUpdate)
		website.POST("/delete", api.WebSiteDelete)
		website.GET("/groups", api.WebsiteGroups)
	}
}

func upload(r *gin.Engine) {
	upload := r.Group("/upload")
	{
		upload.POST("/image", api.Image)
	}
}

func auth(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", api.Login)
	}
}
