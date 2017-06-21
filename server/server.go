package server

import (
	"net/http"

	"github.com/dracher/rhvhhelpers/model"
	"github.com/gin-gonic/gin"
)

var authMiddleware = InitJWTAuthware()

// InitApp will return the server instance
func InitApp(prod bool, db *model.Database) *gin.Engine {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(DatabaseWare(db))
	app.POST("/login", authMiddleware.LoginHandler)

	api := app.Group("/api/v1").Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"res": "hello"})
		})
	}

	return app
}
