package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSearchInformation(context *gin.Context) {
	var Videos []db.Video
	var Users []db.User

	data := struct {
		Title string `json:"title"`
	}{}

	if err := context.BindJSON(&data); err != nil {
		panic(err)
	}

	var title string = strings.ToLower(data.Title)
	db.DB.Raw("SELECT * FROM videos WHERE LOWER(name) LIKE ?", "%"+title+"%").Scan(&Videos)
	db.MySQLDB.Raw("SELECT * FROM users WHERE LOWER(nick) LIKE ?", "%"+title+"%").Scan(&Users)

	var SearchResults struct {
		Users  []db.User
		Videos []db.Video
	}

	SearchResults.Users = Users
	SearchResults.Videos = Videos

	result, _ := json.Marshal(SearchResults)
	context.Data(200, "application/json", result)
}
