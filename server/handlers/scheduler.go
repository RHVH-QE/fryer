package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/dracher/fryer/autocore"
	"github.com/gin-gonic/gin"
)

var runType = [3]string{"manual", "autoinstall", "upgrade"}

// CurrentSchedulerStatusHandler is
func CurrentSchedulerStatusHandler(c *gin.Context) {
	autocore.Scheduler.Lock.RLock()
	defer autocore.Scheduler.Lock.RUnlock()
	c.JSON(OK, autocore.Scheduler.Pool)
}

// DebugHandler is
func DebugHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	fmt.Println(claims)
	c.JSON(OK, "OK")
}

// ProvisonHandler is
func ProvisonHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uname := claims["id"].(string)
	bkrName := c.Param("bkrname")
	runType := c.Param("runtype")
	err := autocore.CheckHostAvailable(bkrName, uname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ret": err.Error()})
		return
	}
	switch runType {
	case "manual":
		c.JSON(OK, "Manual like provsion started")
	case "autoinstall":
		log.Info("autoinstall provison started")
	case "upgrade":
		log.Info("upgrade provison started")
	default:
		c.JSON(http.StatusBadRequest, gin.H{"ret": fmt.Sprintf("%s is not supported", runType)})
	}
}
