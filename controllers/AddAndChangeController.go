package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddViewToVideo(context *gin.Context) {
	var video db.Video
	id := context.Param("id")
	videoId, parseError := strconv.ParseUint(id, 10, 64)

	db.DB.Find(&video, videoId)

	if videoId > 0 && parseError == nil {
		db.DB.Model(&db.Video{}).Where("id = ?", videoId).Update("views", (video.Views + 1))

		result, _ := json.Marshal(&db.Result{200, "successfully added one view"})
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{500, "ErRoR"})
		context.Data(200, "application/json", result)
	}
}

func AddToHistory(context *gin.Context) {
	var history db.History

	userId := context.Param("userId")
	videoId := context.Param("videoId")

	userIdInt, _ := strconv.ParseUint(userId, 10, 64)
	videoIdInt, _ := strconv.ParseUint(videoId, 10, 64)

	db.DB.Where("user_id = ? AND video_id = ?", userIdInt, videoIdInt).Find(&history)

	if history.ID > 0 {
		fmt.Printf("")
	} else {
		db.DB.Create(&db.History{UserId: userIdInt, VideoId: videoIdInt})
	}
}

func ChangeUserInformation(context *gin.Context) {
	var data db.UserData

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	checkAboutTables(uint(data.Id))

	if data.Id > 0 {
		if len(data.Nick) > 3 {
			db.MySQLDB.Model(&db.User{}).Where("id = ?", data.Id).Update("nick", data.Nick)
		}

		if len(data.Email) > 3 {
			db.MySQLDB.Model(&db.User{}).Where("id = ?", data.Id).Update("email", data.Email)
		}

		if len(data.About) > 10 {
			db.DB.Model(&db.AboutUser{}).Where("user_id = ?", data.Id).Update("about", data.About)
		}

		result, _ := json.Marshal(&db.Result{200, "Sucessfuly updated inforamtion"})
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{500, "Error while uploading"})
		context.Data(200, "application/json", result)
	}
}

func ChangeUserPassword(context *gin.Context) {
	data := struct {
		Id      int    `json:"id"`
		OldPass string `json:"oldPass"`
		NewPass string `json:"newPass"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	if len(data.OldPass) > 63 && len(data.NewPass) > 63 {
		db.MySQLDB.Model(&db.User{}).Where("id = ? AND password = ?", data.Id, data.OldPass).Update("password", data.NewPass)
	}

	result, _ := json.Marshal(&db.Result{200, "Ok"})
	context.Data(200, "application/json", result)
}
