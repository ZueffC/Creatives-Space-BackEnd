package admin

import (
	"creativesSpace/config/bindatafs"
	db "creativesSpace/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

func CreateAdminPanels(app *gin.Engine) {
	Admin := admin.New(&admin.AdminConfig{
		SiteName: "TiAdmin",
		DB:       db.DB,
	})

	UserAdmin := admin.New(&admin.AdminConfig{
		SiteName: "UserAdmin",
		DB:       db.MySQLDB,
	})

	Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	UserAdmin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))

	adminServer := http.NewServeMux()
	Admin.MountTo("/admin", adminServer)

	userServer := http.NewServeMux()
	UserAdmin.MountTo("/kuklovody", userServer)
	UserAdmin.AddResource(&db.User{})

	Admin.AddResource(&db.VideoEmotions{})
	Admin.AddResource(&db.Subscribers{})
	Admin.AddResource(&db.AboutUser{})
	Admin.AddResource(&db.Comments{})
	Admin.AddResource(&db.History{})
	Admin.AddResource(&db.Video{})

	app.Any("/admin/*resources", gin.WrapH(adminServer))
	app.Any("/kuklovody/*resources", gin.WrapH(userServer))
}
