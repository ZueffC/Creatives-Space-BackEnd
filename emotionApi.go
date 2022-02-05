package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func emotionApi(v1 *gin.RouterGroup) {
	v1.POST("/add-emotion", func(context *gin.Context) {
		var count int64
		var emotion bool
		var videoLiked VideoEmotions

		data := struct {
			VideoId int `json:"videoId"`
			UserId  int `json:"userId"`
			Emotion int `json:"emotion"`
		}{}

		if err := context.BindJSON(&data); err != nil {
			panic(err)
		}

		db.Model(&VideoEmotions{}).Where("video_id = ?", data.VideoId).Where("user_id = ?", data.UserId).Count(&count)
		db.Model(&VideoEmotions{}).Where("video_id = ?", data.VideoId).Where("user_id = ?", data.UserId).Find(&videoLiked)

		if data.Emotion == 0 {
			emotion = false
		} else {
			emotion = true
		}

		fmt.Println(emotion, data.VideoId)

		if count > 0 {
			if data.UserId > 0 {
				db.Model(&VideoEmotions{}).Where("user_id = ? AND video_id = ?", data.UserId, data.VideoId).Updates(map[string]interface{}{"emotion": emotion})

				result, _ := json.Marshal(&Result{200, "Succesfully added emotion"})
				context.Data(200, "application/json", result)
			}
		} else {
			if data.UserId > 0 {
				db.Create(&VideoEmotions{UserId: uint(data.UserId), VideoId: uint(data.VideoId), Emotion: emotion})

				result, _ := json.Marshal(&Result{400, "Unexpected error :("})
				context.Data(200, "application/json", result)
			}
		}

		result, _ := json.Marshal(&Result{})
		context.Data(200, "application/json", result)
	})
}
