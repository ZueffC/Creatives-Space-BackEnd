package main

import (
	"net/http"
	"os"
	"trubaApi/config/bindatafs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

func main() {
	connect()

	app := gin.Default()
	app.Use(cors.Default())
	app.LoadHTMLGlob("views/*.html")

	Admin := admin.New(&admin.AdminConfig{
		SiteName: "TiAdmin",
		DB:       db,
	})

	UserAdmin := admin.New(&admin.AdminConfig{
		SiteName: "UserAdmin",
		DB:       MySQLDB,
	})

	Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	UserAdmin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))

	adminServer := http.NewServeMux()
	Admin.MountTo("/admin", adminServer)

	userServer := http.NewServeMux()
	UserAdmin.MountTo("/kuklovody", userServer)
	UserAdmin.AddResource(&User{})

	Admin.AddResource(&VideoEmotions{})
	Admin.AddResource(&Subscribers{})
	Admin.AddResource(&AboutUser{})
	Admin.AddResource(&Comments{})
	Admin.AddResource(&History{})
	Admin.AddResource(&Video{})

	app.Any("/admin/*resources", gin.WrapH(adminServer))
	app.Any("/kuklovody/*resources", gin.WrapH(userServer))

	Controller(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	app.Run(":" + port)
}
