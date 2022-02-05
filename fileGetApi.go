package main

import (
	"github.com/gin-gonic/gin"
)

func fileGetApi(v1 *gin.RouterGroup) {
	v1.Static("/static", "./static")
	v1.GET("/video/:id/:filename", func(context *gin.Context) {
		context.File("./uploads/users/" + context.Param("id") + "/videos/video/" + context.Param("filename"))
	})

	v1.GET("uploads/users/:id/videos/:hash/cover/:filename", func(context *gin.Context) {
		context.File("./uploads/users/" + context.Param("id") + "/videos/" + context.Param("hash") + "/cover/" + context.Param("filename"))
	})
	v1.GET("/uploads/users/:id/videos/:hash/video/:filename", func(context *gin.Context) {
		context.File("./uploads/users/" + context.Param("id") + "/videos/" + context.Param("hash") + "/video/" + context.Param("filename"))
	})

	v1.GET("/uploads/users/:id/cover/:hash/:filename", func(context *gin.Context) {
		context.File("./uploads/users/" + context.Param("id") + "/cover/" + context.Param("hash") + "/" + context.Param("filename"))
	})
	v1.GET("/uploads/users/:id/avatar/:hash/:filename", func(context *gin.Context) {
		context.File("./uploads/users/" + context.Param("id") + "/avatar/" + context.Param("hash") + "/" + context.Param("filename"))
	})
}
