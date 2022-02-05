package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Controller(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	getInformationApi(v1)
	commentApi(v1)
	emotionApi(v1)
	fileGetApi(v1)
	uploadApi(v1)
	mainApi(v1)
}
