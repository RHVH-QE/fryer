package server

import (
	"github.com/dracher/fryer/model"
	hf "github.com/dracher/fryer/server/handlers"
	"github.com/gin-gonic/gin"
)

var authMiddleware = hf.InitJWTAuthware()

// InitApp will return the server instance
func InitApp(prod bool, q *model.Query) *gin.Engine {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(hf.DatabaseWare(q))

	{
		app.POST("/login", authMiddleware.LoginHandler)

		app.POST("/cobbler", hf.CobblerHandler)
		app.DELETE("/cobbler", hf.CobblerHandler)

		app.POST("/beaker/:bkrname/:action", hf.BeakerHandler)

		app.GET("/config/:name", hf.ConfigParamsHandler)
		app.GET("/current/scheduler", hf.CurrentSchedulerStatusHandler)
	}

	auth := app.Group("/auth").Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/debug", hf.DebugHandler)
		auth.GET("/refresh", authMiddleware.LoginHandler)

		auth.POST("/provision/:bkrname/:runtype", hf.ProvisonHandler)
	}

	{
		auth.POST("/host", hf.AddHostHandler)
		auth.PUT("/host", hf.UpdateHandler)
	}
	return app
}
