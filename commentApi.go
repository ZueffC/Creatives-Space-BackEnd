package main

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func commentApi(v1 *gin.RouterGroup) {
	v1.POST("/add-comment", func(context *gin.Context) {
		data := struct {
			UserID  uint64 `json:"userId"`
			VideoID uint64 `json:"videoId"`
			Comment string `json:"comment"`
		}{}

		if err := context.BindJSON(&data); err != nil {
			panic(err)
		}

		db.Create(&Comments{UserId: data.UserID, VideoId: data.VideoID, TextComment: data.Comment})

		result, _ := json.Marshal(&Result{200, "Comment was added"})
		context.Data(200, "application/json", result)
	})

	v1.GET("/getComments/:id", func(context *gin.Context) {
		var comments []Comments
		videoId := context.Param("id")
		id, _ := strconv.ParseUint(videoId, 10, 64)

		db.Where("video_id = ?", id).Limit(100).Find(&comments)

		result, _ := json.Marshal(comments)
		context.Data(200, "application/json", result)
	})
}
