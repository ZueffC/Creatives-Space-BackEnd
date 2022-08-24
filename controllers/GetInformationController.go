package controllers

import (
	"encoding/json"
	"strconv"

	db "creativesSpace/database"

	"github.com/gin-gonic/gin"
)

func GetVideoInformation(context *gin.Context) {
	videoId := context.Param("id")
	id, _ := strconv.ParseUint(videoId, 10, 64)

	var VideoInfo struct {
		Video           db.Video
		VideoAuthor     db.User
		CountOfLikes    int64
		CountOfDislikes int64
	}

	db.DB.Find(&VideoInfo.Video, uint(id))
	db.DB.Model(&db.VideoEmotions{}).Where("video_id = ? AND emotion = ?", id, true).Count(&VideoInfo.CountOfLikes)
	db.DB.Model(&db.VideoEmotions{}).Where("video_id = ? AND emotion = ?", id, false).Count(&VideoInfo.CountOfDislikes)
	db.MySQLDB.Find(&VideoInfo.VideoAuthor, VideoInfo.Video.UserId)

	if VideoInfo.Video.ID == uint(id) {
		result, _ := json.Marshal(VideoInfo)
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{404, "Error, video not found!"})
		context.Data(200, "application/json", result)
	}
}

func GetVideosForIndexPage(context *gin.Context) {
	var videos []db.Video
	db.DB.Limit(50).Find(&videos)

	result, _ := json.Marshal(videos)
	context.Data(200, "application/json", result)
}
