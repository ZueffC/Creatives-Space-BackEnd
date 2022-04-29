package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getInformationApi(v1 *gin.RouterGroup) {
	v1.GET("/getVideo/:id", func(context *gin.Context) {
		videoId := context.Param("id")
		id, _ := strconv.ParseUint(videoId, 10, 64)

		var VideoInfo struct {
			Video           Video
			VideoAuthor     User
			CountOfLikes    int64
			CountOfDislikes int64
		}

		db.Find(&VideoInfo.Video, uint(id))
		db.Model(&VideoEmotions{}).Where("video_id = ? AND emotion = ?", id, true).Count(&VideoInfo.CountOfLikes)
		db.Model(&VideoEmotions{}).Where("video_id = ? AND emotion = ?", id, false).Count(&VideoInfo.CountOfDislikes)
		MySQLDB.Find(&VideoInfo.VideoAuthor, VideoInfo.Video.UserId)

		if VideoInfo.Video.ID == uint(id) {
			result, _ := json.Marshal(VideoInfo)
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{404, "Error, video not found!"})
			context.Data(200, "application/json", result)
		}
	})

	v1.GET("/getUser/:id", func(context *gin.Context) {
		userId := context.Param("id")
		id, _ := strconv.ParseUint(userId, 10, 64)

		var AllAboutUser struct {
			User               User
			About              AboutUser
			CountOfVideos      int64
			CountOfSubscribers int64
		}

		MySQLDB.Find(&AllAboutUser.User, uint(id))

		if AllAboutUser.User.ID != 0 && AllAboutUser.User.ID == uint(id) {
			db.Where("user_id = ?", id).Find(&AllAboutUser.About)
			db.Model(&Video{}).Where("user_id = ?", id).Count(&AllAboutUser.CountOfVideos)
			db.Model(&Subscribers{}).Where("to_id = ?", id).Count(&AllAboutUser.CountOfSubscribers)

			result, _ := json.Marshal(AllAboutUser)
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{404, "Nothing found"})
			context.Data(200, "application/json", result)
		}
	})

	v1.GET("/existNick/:nick", func(context *gin.Context) {
		var user User
		nick := context.Param("nick")

		MySQLDB.Where("nick = ?", nick).Find(&user)
		if len(user.Nick) > 1 {
			result, _ := json.Marshal(&Result{200, user.Nick})
			context.Data(200, "application/json", result)
		} else {
			result, _ := json.Marshal(&Result{404, "Nick not exist"})
			context.Data(200, "application/json", result)
		}
	})

	v1.GET("/getAllUserVideos/:id", func(context *gin.Context) {
		var videos []Video
		id := context.Param("id")
		userId, _ := strconv.ParseInt(id, 10, 64)
		db.Where("user_id = ?", userId).Find(&videos)

		result, _ := json.Marshal(videos)
		context.Data(200, "application/json", result)
	})

	v1.GET("/getAvatar/:id", func(context *gin.Context) {
		var aboutUser AboutUser
		id := context.Param("id")
		userId, _ := strconv.ParseInt(id, 10, 64)
		db.Where("user_id = ?", userId).Find(&aboutUser)

		result, _ := json.Marshal(aboutUser.Avatar)
		context.Data(200, "application/json", result)
	})

	v1.POST("/search", func(context *gin.Context) {
		var Videos []Video
		var Users []User

		data := struct {
			Title string `json:"title"`
		}{}

		if err := context.BindJSON(&data); err != nil {
			panic(err)
		}

		var title string = strings.ToLower(data.Title)
		db.Raw("SELECT * FROM videos WHERE LOWER(name) LIKE ?", "%"+title+"%").Scan(&Videos)
		MySQLDB.Raw("SELECT * FROM users WHERE LOWER(nick) LIKE ?", "%"+title+"%").Scan(&Users)

		var SearchResults struct {
			Users  []User
			Videos []Video
		}

		SearchResults.Users = Users
		SearchResults.Videos = Videos

		result, _ := json.Marshal(SearchResults)
		context.Data(200, "application/json", result)
	})
}
