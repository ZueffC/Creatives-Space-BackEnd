package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkAboutTables(id uint) {
	var existAbout AboutUser
	db.Where("user_id = ?", id).Find(&existAbout)

	if existAbout.ID > 0 && len(existAbout.Avatar) > 0 {
		fmt.Println("Ok")
	} else {
		var about = AboutUser{UserId: id, About: "Ну, тут пока пустота."}
		db.Create(&about)
	}
}

func mainApi(v1 *gin.RouterGroup) {
	v1.GET("/", func(context *gin.Context) {
		var videos []Video
		db.Limit(50).Find(&videos)

		result, _ := json.Marshal(videos)
		context.Data(200, "application/json", result)
	})

	v1.GET("/add-view-to-video/:id", func(context *gin.Context) {
		var video Video
		id := context.Param("id")
		videoId, parseError := strconv.ParseUint(id, 10, 64)

		db.Find(&video, videoId)

		if videoId > 0 && parseError == nil {
			db.Model(&Video{}).Where("id = ?", videoId).Update("views", (video.Views + 1))

			result, _ := json.Marshal(&Result{200, "successfully added one view"})
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{500, "ErRoR"})
			context.Data(200, "application/json", result)
		}
	})

	v1.POST("/change-information", func(context *gin.Context) {
		var data UserData

		if err := context.BindJSON(&data); err != nil {
			panic(err)
		}

		checkAboutTables(uint(data.Id))

		if data.Id > 0 {
			if len(data.Nick) > 3 {
				MySQLDB.Model(&User{}).Where("id = ?", data.Id).Update("nick", data.Nick)
			}

			if len(data.Email) > 3 {
				MySQLDB.Model(&User{}).Where("id = ?", data.Id).Update("email", data.Email)
			}

			if len(data.About) > 10 {
				db.Model(&AboutUser{}).Where("user_id = ?", data.Id).Update("about", data.About)
			}

			result, _ := json.Marshal(&Result{200, "Sucessfuly updated inforamtion"})
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{500, "Error while uploading"})
			context.Data(200, "application/json", result)
		}
	})

	v1.POST("/change-password", func(context *gin.Context) {
		data := struct {
			Id      int    `json:"id"`
			OldPass string `json:"oldPass"`
			NewPass string `json:"newPass"`
		}{}

		if err := context.BindJSON(&data); err != nil {
			panic(err)
		}

		if len(data.OldPass) > 63 && len(data.NewPass) > 63 {
			MySQLDB.Model(&User{}).Where("id = ? AND password = ?", data.Id, data.OldPass).Update("password", data.NewPass)
		}

		result, _ := json.Marshal(&Result{200, "Ok"})
		context.Data(200, "application/json", result)
	})
}
