package server

import (
	"github.com/dracher/fryer/model"
	"github.com/gin-gonic/gin"
)

var authMiddleware = InitJWTAuthware()
var machinePool = NewMachinePool()

// InitApp will return the server instance
func InitApp(prod bool, db *model.Database) *gin.Engine {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(DatabaseWare(db))
	{
		app.POST("/login", authMiddleware.LoginHandler)
		app.GET("/machines", MachinesListHandler)
	}

	api := app.Group("/api/v1").Use(authMiddleware.MiddlewareFunc())
	{
		api.POST("/cobbler", CobblerHandler)
		api.DELETE("/cobbler", CobblerHandler)

		api.POST("/beaker/:bkrname/:action", BeakerHandler)

		api.POST("/provision", ProvisionHandler)
	}
	return app
}
