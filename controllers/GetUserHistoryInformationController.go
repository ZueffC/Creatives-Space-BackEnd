package controllers

import (
	db "creativesSpace/database"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserHistoryInformation(context *gin.Context) {
	var History []db.History

	userId := context.Param("userId")
	id, _ := strconv.ParseUint(userId, 10, 64)

	db.DB.Where("user_id = ?", id).Find(&History)

	result, _ := json.Marshal(History)
	context.Data(200, "application/json", result)
}
