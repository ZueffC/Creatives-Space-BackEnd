package controllers

import "github.com/gin-gonic/gin"

func GetVideoFileByIdAndName(context *gin.Context) {
	context.File("./uploads/users/" + context.Param("id") + "/videos/video/" + context.Param("filename"))
}

func GetVideoCoverByHashAndName(context *gin.Context) {
	context.File("./uploads/users/" + context.Param("id") + "/videos/" + context.Param("hash") + "/cover/" + context.Param("filename"))
}

func GetUserCoverFileByIdAndHash(context *gin.Context) {
	context.File("./uploads/users/" + context.Param("id") + "/cover/" + context.Param("hash") + "/" + context.Param("filename"))
}

func GetUserAvatarByIdAndHash(context *gin.Context) {
	context.File("./uploads/users/" + context.Param("id") + "/avatar/" + context.Param("hash") + "/" + context.Param("filename"))
}
