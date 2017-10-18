package handlers

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/dracher/fryer/model"
	"github.com/gin-gonic/gin"

	"github.com/dracher/fryer/helper"
)

const (
	jwtKey = "KKbJi&w6oFMYeoimhnYuyyJM8fftt3$r"
)

// InitJWTAuthware is
func InitJWTAuthware() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "fryer",
		Key:        []byte(jwtKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			q := helper.GetQuery(c)
			return q.CheckUser(userId, password)
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			q := helper.GetQuery(c)
			u, _ := q.User("KrbID", userId)
			return u.Admin
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	return authMiddleware
}

// DatabaseWare is
func DatabaseWare(q *model.Query) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("q", q)
		c.Next()
	}
}
