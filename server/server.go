package server

import (
	"github.com/dracher/fryer/model"
	"github.com/gin-gonic/gin"
)

var authMiddleware = InitJWTAuthware()

// InitApp will return the server instance
func InitApp(prod bool, q *model.Query) *gin.Engine {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(DatabaseWare(q))
	{

		app.POST("/login", authMiddleware.LoginHandler)
		app.POST("/cobbler", CobblerHandler)
		app.DELETE("/cobbler", CobblerHandler)
		app.POST("/beaker/:bkrname/:action", BeakerHandler)

		app.GET("/config/:name", ConfigParamsHandler)
		app.GET("/current/scheduler", CurrentSchedulerStatusHandler)
	}

	auth := app.Group("/auth").Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/debug", DebugHandler)
		auth.GET("/refresh", authMiddleware.LoginHandler)

		auth.GET("/provision/:bkrname", ProvisonHandler)
	}
	return app
}
