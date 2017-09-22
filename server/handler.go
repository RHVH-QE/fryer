package server

import (
	"fmt"
	"net/http"

	"github.com/dracher/fryer/autocore"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/dracher/fryer/utils/beaker"
	"github.com/dracher/fryer/utils/cobbler"
)

const (
	OK = http.StatusOK
)

// BeakerHandler is
func BeakerHandler(c *gin.Context) {
	name := c.Param("bkrname")
	action := c.Param("action")
	bkr := beaker.Beaker{SystemName: name}

	switch action {
	case "reboot":
		log.Infof("%s %s", action, name)
		bkr.Reboot()
	case "on":
		log.Infof("%s %s", action, name)
		bkr.PowerOn()
	case "off":
		log.Infof("%s %s", action, name)
		bkr.PowerOff()
	default:
		log.Infof("no action performed on %s", name)
	}
	c.String(OK, fmt.Sprintf("%s %s is complete", name, action))
}

// SystemParams is
type SystemParams struct {
	Name    string
	Profile string
	Comment string
	Status  string
	Kargs   string
	Nic     []string
}

// CobblerHandler is
func CobblerHandler(c *gin.Context) {
	cb := cobbler.NewCobbler()
	var sp SystemParams
	if c.Request.Method == "POST" {
		if err := c.BindJSON(&sp); err != nil {
			log.Error(err.Error())
			c.String(http.StatusPreconditionFailed, "create system failed")
		} else {
			cb.NewSystem(sp.Name, sp.Profile, sp.Comment, sp.Status, sp.Kargs, sp.Nic)
			c.String(OK, "create system ok")
		}

	} else if c.Request.Method == "DELETE" {
		if err := c.BindJSON(&sp); err != nil {
			log.Error(err.Error())
			c.String(http.StatusPreconditionFailed, "remove system failed")
		} else {
			cb.RemoveSystem(sp.Name)
			c.String(OK, "remove system ok")
		}
	}
}

// ConfigParamsHandler is
func ConfigParamsHandler(c *gin.Context) {
	name := c.Param("name")
	q := getQuery(c)
	switch name {
	case "common":
		ret, _ := q.CommonParams()
		c.JSON(OK, ret)
	case "hosts":
		ret, _ := q.Hosts()
		c.JSON(OK, ret)
	case "tiers":
		ret, _ := q.AutoTestTiers()
		c.JSON(OK, ret)
	case "casemap":
		ret, _ := q.AutoTestCaseMap()
		c.JSON(OK, ret)
	default:
		c.JSON(OK, nil)
	}
}

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
	err := autocore.CheckHostAvailable(bkrName, uname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ret": err.Error()})
	} else {
		c.JSON(OK, gin.H{"ret": "Provison Started"})
	}
}
