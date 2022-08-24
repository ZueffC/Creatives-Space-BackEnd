package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInformation(context *gin.Context) {
	userId := context.Param("id")
	id, _ := strconv.ParseUint(userId, 10, 64)

	var AllAboutUser struct {
		User               db.User
		About              db.AboutUser
		CountOfVideos      int64
		CountOfSubscribers int64
	}

	db.MySQLDB.Find(&AllAboutUser.User, uint(id))

	if AllAboutUser.User.ID != 0 && AllAboutUser.User.ID == uint(id) {
		db.DB.Where("user_id = ?", id).Find(&AllAboutUser.About)
		db.DB.Model(&db.Video{}).Where("user_id = ?", id).Count(&AllAboutUser.CountOfVideos)
		db.DB.Model(&db.Subscribers{}).Where("to_id = ?", id).Count(&AllAboutUser.CountOfSubscribers)

		result, _ := json.Marshal(AllAboutUser)
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{404, "Nothing found"})
		context.Data(200, "application/json", result)
	}
}

func CheckNickExist(context *gin.Context) {
	var user db.User
	nick := context.Param("nick")

	db.MySQLDB.Where("nick = ?", nick).Find(&user)

	if len(user.Nick) > 1 {
		result, _ := json.Marshal(&db.Result{200, user.Nick})
		context.Data(200, "application/json", result)
	} else {
		result, _ := json.Marshal(&db.Result{404, "Nick not exist"})
		context.Data(200, "application/json", result)
	}
}

func GetAllUserVideos(context *gin.Context) {
	var videos []db.Video
	id := context.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	db.DB.Where("user_id = ?", userId).Find(&videos)

	result, _ := json.Marshal(videos)
	context.Data(200, "application/json", result)
}

func GetAvatarById(context *gin.Context) {
	var aboutUser db.AboutUser
	id := context.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	db.DB.Where("user_id = ?", userId).Find(&aboutUser)

	result, _ := json.Marshal(aboutUser.Avatar)
	context.Data(200, "application/json", result)
}
