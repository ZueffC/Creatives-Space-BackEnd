package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCommentController(context *gin.Context) {
	data := struct {
		UserID  uint64 `json:"userId"`
		VideoID uint64 `json:"videoId"`
		Comment string `json:"comment"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	db.DB.Create(&db.Comments{UserId: data.UserID, VideoId: data.VideoID, TextComment: data.Comment})

	result, _ := json.Marshal(&db.Result{200, "Comment was added"})
	context.Data(200, "application/json", result)
}

func GetCommentController(context *gin.Context) {
	var comments []db.Comments
	videoId := context.Param("id")
	id, _ := strconv.ParseUint(videoId, 10, 64)

	db.DB.Where("video_id = ?", id).Limit(100).Find(&comments)

	result, _ := json.Marshal(comments)
	context.Data(200, "application/json", result)
}
