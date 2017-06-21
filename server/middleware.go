package server

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dracher/rhvhhelpers/model"
	"github.com/gin-gonic/gin"
	jwt "gopkg.in/appleboy/gin-jwt.v2"
)

const (
	jwtKey = "KKbJi&w6oFMYeoimhnYuyyJM8fftt3$r"
)

// InitJWTAuthware is
func InitJWTAuthware() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:   "RHVH Helper",
		Key:     []byte(jwtKey),
		Timeout: 24 * time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			ret, _ := c.Get("db")
			db := ret.(*model.Database)
			user := db.CheckUser(userID, password)
			if user.Username == userID {
				return user.Username, true
			}
			return "", false

		},
		Authorizator: func(userID string, c *gin.Context) bool {
			ret, _ := c.Get("db")
			db := ret.(*model.Database)
			user := db.FindUser(userID)
			log.Debug(user)
			return user.Regular
		},
		PayloadFunc: func(userID string) map[string]interface{} {
			return make(map[string]interface{})
		},
	}
	return authMiddleware
}

// DatabaseWare is
func DatabaseWare(db *model.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
