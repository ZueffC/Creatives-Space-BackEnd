package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func uploadApi(v1 *gin.RouterGroup) {
	v1.POST("/upload-video/by/:id", func(context *gin.Context) {
		userId := context.Param("id")
		video, _ := context.FormFile("video")
		cover, _ := context.FormFile("cover")

		name, _ := context.GetPostForm("name")
		about, _ := context.GetPostForm("about")

		if video != nil {
			os.MkdirAll("./uploads/users/"+userId+"/videos/"+MD5(video.Filename)+"/video", 0777)
			os.MkdirAll("./uploads/users/"+userId+"/videos/"+MD5(video.Filename)+"/cover", 0777)

			pathToVideo := "uploads/users/" + userId + "/videos/" + MD5(video.Filename) + "/video/" + video.Filename
			pathToCover := "uploads/users/" + userId + "/videos/" + MD5(video.Filename) + "/cover/" + cover.Filename

			context.SaveUploadedFile(video, fmt.Sprintf("./%s", pathToVideo))
			context.SaveUploadedFile(cover, fmt.Sprintf("./%s", pathToCover))

			id, _ := strconv.ParseUint(userId, 10, 16)

			videos := Video{
				UserId:      id,
				PathToVideo: pathToVideo,
				Cover:       pathToCover,
				Name:        name,
				About:       about,
			}
			db.Create(&videos)

			result, _ := json.Marshal(&Result{200, "Video was succesfully uploaded"})
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{500, "Error while uploading"})
			context.Data(200, "application/json", result)
		}
	})

	v1.POST("/upload-profile-photo/by/:id", func(context *gin.Context) {
		userId := context.Param("id")
		id, _ := strconv.ParseUint(userId, 10, 64)
		avatar, _ := context.FormFile("avatar")
		cover, _ := context.FormFile("cover")

		checkAboutTables(uint(id))

		if avatar != nil {
			os.RemoveAll("./uploads/users/" + userId + "/avatar")
			os.MkdirAll("./uploads/users/"+userId+"/avatar/"+MD5(avatar.Filename), 0777)
			pathToAvatar := "uploads/users/" + userId + "/avatar/" + MD5(avatar.Filename) + "/" + avatar.Filename
			context.SaveUploadedFile(avatar, fmt.Sprintf("./%s", pathToAvatar))

			db.Model(&AboutUser{}).Where("user_id = ?", id).Update("avatar", pathToAvatar)
		}

		if cover != nil {
			os.RemoveAll("./uploads/users/" + userId + "/cover")
			os.MkdirAll("./uploads/users/"+userId+"/cover/"+MD5(cover.Filename), 0777)
			pathToCover := "uploads/users/" + userId + "/cover/" + MD5(cover.Filename) + "/" + cover.Filename
			context.SaveUploadedFile(cover, fmt.Sprintf("./%s", pathToCover))

			db.Model(&AboutUser{}).Where("user_id = ?", id).Update("cover", pathToCover)
		}

		result, _ := json.Marshal(&Result{200, "Success"})
		context.Data(200, "application/json", result)
	})
}
