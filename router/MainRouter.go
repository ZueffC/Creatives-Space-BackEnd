package router

import (
	controllers "creativesSpace/controllers"

	"github.com/gin-gonic/gin"
)

func MainRouter(app *gin.Engine) {
	app.GET("/", controllers.IndexPage)

	v1 := app.Group("/api/v1")

	// information routes
	v1.GET("/getVideo/:id", controllers.GetVideoInformation)
	v1.GET("/getUser/:id", controllers.GetUserInformation)
	v1.GET("/history/by/:userId", controllers.GetUserHistoryInformation)
	v1.GET("/existNick/:nick", controllers.CheckNickExist)
	v1.GET("/getAllUserVideos/:id", controllers.GetAllUserVideos)
	v1.GET("/getAvatar/:id", controllers.GetAvatarById)

	v1.POST("/search", controllers.GetSearchInformation)

	// comments routes
	v1.POST("/add-comment", controllers.AddCommentController)
	v1.GET("/getComments/:id", controllers.GetCommentController)

	// emotion routes
	v1.POST("/add-emotion", controllers.AddEmotionController)
	v1.POST("/liked-videos", controllers.GetVideosEmotionsByUserId)

	// get file routes
	v1.Static("/static", "./static")
	v1.GET("/video/:id/:filename", controllers.GetVideoFileByIdAndName)
	v1.GET("uploads/users/:id/videos/:hash/cover/:filename", controllers.GetVideoCoverByHashAndName)
	v1.GET("/uploads/users/:id/cover/:hash/:filename", controllers.GetUserCoverFileByIdAndHash)
	v1.GET("/uploads/users/:id/avatar/:hash/:filename", controllers.GetUserAvatarByIdAndHash)

	// upload routes
	v1.POST("/upload-video/by/:id", controllers.UploadVideoByForm)
	v1.POST("/upload-profile-photo/by/:id", controllers.UploadUserAvararByForm)
	v1.POST("/delete-video", controllers.DeleteVideo)

	// change and adds router
	v1.GET("/", controllers.GetVideosForIndexPage)

	v1.GET("/add-view-to-video/:id", controllers.AddViewToVideo)
	v1.GET("/add-to-history/by/:userId/to/:videoId", controllers.AddToHistory)

	v1.POST("/change-information", controllers.ChangeUserInformation)
	v1.POST("/change-password", controllers.ChangeUserPassword)
}
