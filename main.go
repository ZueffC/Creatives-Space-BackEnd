package main

import (
	admin "creativesSpace/admin"
	db "creativesSpace/database"
	"creativesSpace/router"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	app := gin.Default()
	app.Use(cors.Default())
	app.LoadHTMLGlob("views/*.html")

	admin.CreateAdminPanels(app)
	router.MainRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	app.Run(":" + port)
}
