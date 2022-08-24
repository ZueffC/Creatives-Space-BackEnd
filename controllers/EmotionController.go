package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddEmotionController(context *gin.Context) {
	var count int64
	var emotion bool
	var videoLiked db.VideoEmotions

	data := struct {
		VideoId int `json:"videoId"`
		UserId  int `json:"userId"`
		Emotion int `json:"emotion"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	db.DB.Model(&db.VideoEmotions{}).Where("video_id = ?", data.VideoId).Where("user_id = ?", data.UserId).Count(&count)
	db.DB.Model(&db.VideoEmotions{}).Where("video_id = ?", data.VideoId).Where("user_id = ?", data.UserId).Find(&videoLiked)

	if data.Emotion == 0 {
		emotion = false
	} else {
		emotion = true
	}

	fmt.Println(emotion, data.VideoId)

	if count > 0 {
		if data.UserId > 0 {
			db.DB.Model(&db.VideoEmotions{}).Where("user_id = ? AND video_id = ?", data.UserId, data.VideoId).Updates(map[string]interface{}{"emotion": emotion})

			result, _ := json.Marshal(&db.Result{200, "Succesfully added emotion"})
			context.Data(200, "application/json", result)
		}
	} else {
		if data.UserId > 0 {
			db.DB.Create(&db.VideoEmotions{UserId: uint(data.UserId), VideoId: uint(data.VideoId), Emotion: emotion})

			result, _ := json.Marshal(&db.Result{400, "Unexpected error :("})
			context.Data(200, "application/json", result)
		}
	}

	result, _ := json.Marshal(&db.Result{})
	context.Data(200, "application/json", result)
}

func GetVideosEmotionsByUserId(context *gin.Context) {
	var VideoEmotions []db.VideoEmotions
	data := struct {
		UserId int `json:"userId"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	db.DB.Where("user_id = ? AND emotion = ?", data.UserId, true).Find(&VideoEmotions)

	result, _ := json.Marshal(VideoEmotions)
	context.Data(200, "application/json", result)
}
