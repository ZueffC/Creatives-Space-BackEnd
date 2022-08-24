package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	db "creativesSpace/database"

	"github.com/gin-gonic/gin"
)

func checkAboutTables(id uint) {
	var existAbout db.AboutUser
	db.DB.Where("user_id = ?", id).Find(&existAbout)

	if existAbout.ID > 0 && len(existAbout.Avatar) > 0 {
		fmt.Println("Ok")
	} else {
		var about = db.AboutUser{UserId: id, About: "Ну, тут пока пустота."}
		db.DB.Create(&about)
	}
}

func UploadVideoByForm(context *gin.Context) {
	userId := context.Param("id")
	video, _ := context.FormFile("video")
	cover, _ := context.FormFile("cover")

	name, _ := context.GetPostForm("name")
	about, _ := context.GetPostForm("about")

	nameOfVideo := strings.ToLower(name)

	if video != nil {
		os.MkdirAll("./uploads/users/"+userId+"/videos/"+db.MD5(nameOfVideo)+"/video", 0777)
		os.MkdirAll("./uploads/users/"+userId+"/videos/"+db.MD5(nameOfVideo)+"/cover", 0777)

		pathToVideo := "uploads/users/" + userId + "/videos/" + db.MD5(nameOfVideo) + "/video/" + video.Filename
		pathToCover := "uploads/users/" + userId + "/videos/" + db.MD5(nameOfVideo) + "/cover/" + cover.Filename

		context.SaveUploadedFile(video, fmt.Sprintf("./%s", pathToVideo))
		context.SaveUploadedFile(cover, fmt.Sprintf("./%s", pathToCover))

		id, _ := strconv.ParseUint(userId, 10, 16)

		videos := db.Video{
			UserId:      id,
			PathToVideo: pathToVideo,
			Cover:       pathToCover,
			Name:        nameOfVideo,
			About:       about,
		}
		db.DB.Create(&videos)

		result, _ := json.Marshal(&db.Result{200, "Video was succesfully uploaded"})
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{500, "Error while uploading"})
		context.Data(200, "application/json", result)
	}
}

func UploadUserAvararByForm(context *gin.Context) {
	userId := context.Param("id")
	id, _ := strconv.ParseUint(userId, 10, 64)
	avatar, _ := context.FormFile("avatar")
	cover, _ := context.FormFile("cover")

	checkAboutTables(uint(id))

	if avatar != nil {
		os.RemoveAll("./uploads/users/" + userId + "/avatar")
		os.MkdirAll("./uploads/users/"+userId+"/avatar/"+db.MD5(avatar.Filename), 0777)
		pathToAvatar := "uploads/users/" + userId + "/avatar/" + db.MD5(avatar.Filename) + "/" + avatar.Filename
		context.SaveUploadedFile(avatar, fmt.Sprintf("./%s", pathToAvatar))

		db.DB.Model(&db.AboutUser{}).Where("user_id = ?", id).Update("avatar", pathToAvatar)
	}

	if cover != nil {
		os.RemoveAll("./uploads/users/" + userId + "/cover")
		os.MkdirAll("./uploads/users/"+userId+"/cover/"+db.MD5(cover.Filename), 0777)
		pathToCover := "uploads/users/" + userId + "/cover/" + db.MD5(cover.Filename) + "/" + cover.Filename
		context.SaveUploadedFile(cover, fmt.Sprintf("./%s", pathToCover))

		db.DB.Model(&db.AboutUser{}).Where("user_id = ?", id).Update("cover", pathToCover)
	}

	result, _ := json.Marshal(&db.Result{200, "Success"})
	context.Data(200, "application/json", result)
}

func DeleteVideo(context *gin.Context) {
	var video db.Video

	data := struct {
		VideoId int `json:"video_id"`
		UserId  int `json:"user_id"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	db.DB.Model(&db.Video{}).Where("id = ?", data.VideoId).Find(&video)

	pathToVideoFolder := "./uploads/users/" + fmt.Sprint(data.UserId) + "/videos/" + db.MD5(video.Name)
	os.RemoveAll(pathToVideoFolder)

	db.DB.Where("id = ? AND user_id = ?", data.VideoId, data.UserId).Delete(&db.Video{})

	result, _ := json.Marshal(&db.Result{200, "Success"})
	context.Data(200, "application/json", result)
}
