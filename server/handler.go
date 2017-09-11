package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"fmt"

	"github.com/dracher/fryer/utils/beaker"
	"github.com/dracher/fryer/utils/cobbler"
	"github.com/gin-gonic/gin"
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
	c.String(http.StatusOK, fmt.Sprintf("%s %s is complete", name, action))
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
			c.String(http.StatusOK, "create system ok")
		}

	} else if c.Request.Method == "DELETE" {
		if err := c.BindJSON(&sp); err != nil {
			log.Error(err.Error())
			c.String(http.StatusPreconditionFailed, "remove system failed")
		} else {
			cb.RemoveSystem(sp.Name)
			c.String(http.StatusOK, "remove system ok")
		}
	}
}

// ProvisionHandler is
func ProvisionHandler(c *gin.Context) {
	var sp SystemParams
	if err := c.BindJSON(&sp); err != nil {
		log.Error(err.Error())
		c.String(http.StatusPreconditionFailed, "provision system failed")
	} else {
		if machinePool.CheckInUse(sp.Name) {
			c.String(http.StatusForbidden, fmt.Sprintf("%s in used", sp.Name))
		} else {
			machinePool.MarkMachine(sp.Name)
			// cobbler.NewCobbler().NewSystem(sp.Name, sp.Profile, sp.Comment, sp.Status, sp.Kargs, sp.Nic)
			// beaker.Beaker{SystemName: sp.Name}.Reboot()
			c.String(http.StatusOK, fmt.Sprintf("%s provision complete", sp.Name))
		}
	}
}

// MachinesListHandler is
func MachinesListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, machinePool.MachinesList())
}
